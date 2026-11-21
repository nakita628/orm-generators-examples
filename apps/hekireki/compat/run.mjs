#!/usr/bin/env node
// Holds what hekireki generated to the database Prisma creates from the same schema.
//
//   node run.mjs [--orm drizzle,sqlalchemy,gorm] [schema ...]
//
// For every schema and ORM:
//
// 1. crud: `prisma db push` creates the database, exactly as the project's own Prisma would, and the
//    ORM's harness runs the plan (lib/plan.mjs) through the generated code: one row per model,
//    inserted with the ORM filling every default, read back, updated, joined through every
//    relation from both sides, and deleted.
// 2. ddl: the ORM creates the tables from the generated code on its own, into an empty database,
//    and `prisma migrate diff` prints what it would take to turn that database into the one the
//    Prisma schema describes. Empty means the ORM's tables are Prisma's.
//
// The databases are PostgreSQL 17 and MySQL 8.4 in Docker (started here when they are not up) and
// SQLite files. The findings go to results/<schema>.md, and the matrix to results/README.md.
import { spawnSync } from 'node:child_process'
import { cpSync, mkdirSync, readdirSync, readFileSync, rmSync, writeFileSync } from 'node:fs'
import { createRequire } from 'node:module'
import { dirname, join, resolve } from 'node:path'
import { fileURLToPath } from 'node:url'
import { makePlan } from './lib/plan.mjs'
import { compare } from './lib/compare.mjs'
import { renderSchema, renderIndex } from './lib/report.mjs'

const here = dirname(fileURLToPath(import.meta.url))
const require = createRequire(import.meta.url)
const prismaCli = require.resolve('prisma/build/index.js')
const schemasDir = resolve(here, '../schemas')
const workRoot = join(here, '.work')
const resultsDir = join(here, 'results')

const ORMS = ['drizzle', 'sqlalchemy', 'gorm']
const PG = { port: 55432, image: 'postgres:17-alpine', name: 'hekireki-compat-pg' }
const MYSQL = { port: 53306, image: 'mysql:8.4', name: 'hekireki-compat-mysql' }

const args = process.argv.slice(2)
let orms = ORMS
const schemas = []
for (let i = 0; i < args.length; i++) {
  if (args[i] === '--orm') orms = args[++i].split(',')
  else schemas.push(args[i])
}
if (schemas.length === 0) throw new Error('usage: node run.mjs [--orm a,b] <schema> [schema ...]')

async function main() {
mkdirSync(resultsDir, { recursive: true })
for (const name of schemas) {
  const source = readFileSync(join(schemasDir, name, 'schema.prisma'), 'utf8')
  const plan = await makePlan(source)
  const work = join(workRoot, name)
  rmSync(work, { recursive: true, force: true })
  mkdirSync(work, { recursive: true })
  // The schema without its generator blocks: `db push` and `migrate diff` need only the models.
  const bare = source.replace(/^generator\s+[\w-]+\s*\{[^}]*\}\s*$/gms, '')
  writeFileSync(join(work, 'schema.prisma'), bare)
  writeFileSync(join(work, 'plan.json'), JSON.stringify(plan, null, 2))
  if (plan.dialect === 'postgresql') await up(PG)
  if (plan.dialect === 'mysql') await up(MYSQL)

  const results = []
  for (const orm of orms) {
    console.log(`${name}: ${orm}`)
    const result = { orm }
    results.push(result)
    try {
      const crudUrl = await freshDatabase(plan.dialect, work, `${name}_${orm}`)
      prisma(work, crudUrl, ['db', 'push'])
      const harness = HARNESSES[orm]
      const prepared = harness.prepare(name, work)
      const run = sh(harness.crud(prepared, join(work, 'plan.json'), crudUrl))
      if (run.status !== 0 || !run.stdout.trim().startsWith('{')) {
        result.fatal = (run.stderr || run.stdout).trim().slice(-3000)
      } else {
        result.crud = compare(plan, JSON.parse(run.stdout))
      }
      // ddl
      const ddlUrl = await freshDatabase(plan.dialect, work, `${name}_${orm}_ddl`)
      const ddl = harness.ddl(prepared, join(work, 'plan.json'), ddlUrl)
      if (ddl.error) result.ddl = { error: ddl.error }
      else {
        const diff = prisma(work, ddlUrl, ['migrate', 'diff', '--from-config-datasource', '--to-schema', join(work, 'schema.prisma'), '--script'], true)
        result.ddl = diff.status === 0 ? { diff: cleanDiff(diff.stdout) } : { error: (diff.stderr || diff.stdout).trim().slice(-2000) }
      }
    } catch (e) {
      result.fatal = String(e?.stack ?? e).slice(0, 3000)
    }
  }
  writeFileSync(join(resultsDir, `${name}.md`), renderSchema(name, plan, results))
  writeFileSync(join(resultsDir, `${name}.json`), JSON.stringify(summary(results), null, 2) + '\n')
}
const all = readdirSync(resultsDir)
  .filter((f) => f.endsWith('.json'))
  .map((f) => ({ name: f.replace(/\.json$/, ''), ...JSON.parse(readFileSync(join(resultsDir, f), 'utf8')) }))
writeFileSync(join(resultsDir, 'README.md'), renderIndex(all))
}

// --- harnesses -------------------------------------------------------------------------------

const HARNESSES = {
  drizzle: {
    prepare(name, work) {
      const dir = join(here, '.work', name, 'drizzle')
      mkdirSync(dir, { recursive: true })
      cpSync(join(schemasDir, name, 'generated/drizzle/schema.ts'), join(dir, 'schema.ts'))
      return { schema: join(dir, 'schema.ts') }
    },
    crud: (p, plan, url) => ['node', join(here, 'drizzle/harness.mjs'), 'crud', plan, p.schema, url],
    ddl(p, plan, url) {
      const sql = sh(['node', join(here, 'drizzle/harness.mjs'), 'ddl', plan, p.schema])
      if (sql.status !== 0) return { error: sql.stderr.trim().slice(-2000) }
      const file = join(dirname(p.schema), 'ddl.sql')
      writeFileSync(file, sql.stdout)
      const applied = prisma(dirname(dirname(p.schema)), url, ['db', 'execute', '--file', file], true)
      return applied.status === 0 ? {} : { error: (applied.stderr || applied.stdout).trim().slice(-2000) }
    },
  },
  sqlalchemy: {
    prepare(name, work) {
      return { models: join(schemasDir, name, 'generated/sqlalchemy/models.py') }
    },
    crud: (p, plan, url) => [...uvRun('sqlalchemy'), 'crud', plan, p.models, url],
    ddl(p, plan, url) {
      const run = sh([...uvRun('sqlalchemy'), 'ddl', plan, p.models, url])
      return run.status === 0 ? {} : { error: (run.stderr || run.stdout).trim().slice(-2000) }
    },
  },
  gorm: {
    prepare(name, work) {
      // The generated package, copied beside the harness's main package with a registry of every
      // struct that has a TableName, which Go cannot list by reflection.
      const dir = join(here, '.work', name, 'gorm')
      rmSync(dir, { recursive: true, force: true })
      cpSync(join(here, 'gorm'), dir, { recursive: true })
      cpSync(join(schemasDir, name, 'generated/gorm'), join(dir, 'model'), { recursive: true })
      const source = readdirSync(join(dir, 'model'))
        .filter((f) => f.endsWith('.go'))
        .map((f) => readFileSync(join(dir, 'model', f), 'utf8'))
        .join('\n')
      const types = [...source.matchAll(/^func \((?:\w+ )?\*?(\w+)\) TableName\(\) string/gm)].map((m) => m[1])
      writeFileSync(
        join(dir, 'model', 'zz_registry.go'),
        `package model\n\n// All is every model, for the compat harness.\nvar All = []any{${types.map((t) => `&${t}{}`).join(', ')}}\n`,
      )
      const build = sh(['go', 'build', '-o', join(dir, 'harness'), '.'], { cwd: dir, env: goEnv() })
      if (build.status !== 0) {
        const tidy = sh(['go', 'mod', 'tidy'], { cwd: dir, env: goEnv() })
        const again = sh(['go', 'build', '-o', join(dir, 'harness'), '.'], { cwd: dir, env: goEnv() })
        if (again.status !== 0) throw new Error(`go build: ${(tidy.stderr + again.stderr).trim().slice(-2000)}`)
      }
      return { bin: join(dir, 'harness') }
    },
    crud: (p, plan, url) => [p.bin, 'crud', plan, url],
    ddl(p, plan, url) {
      const run = sh([p.bin, 'ddl', plan, url])
      return run.status === 0 ? {} : { error: (run.stderr || run.stdout).trim().slice(-2000) }
    },
  },
}

function uvRun(harness) {
  return ['uv', 'run', '--quiet', '--project', join(here, harness), 'python', join(here, harness, 'harness.py')]
}

function goEnv() {
  return { ...process.env, GOFLAGS: '-mod=mod', GOTOOLCHAIN: 'auto' }
}

// --- databases --------------------------------------------------------------------------------

async function up({ name, image, port }) {
  const running = sh(['docker', 'ps', '-q', '-f', `name=^${name}$`]).stdout.trim()
  if (!running) {
    sh(['docker', 'rm', '-f', name])
    const env = image.startsWith('postgres') ? ['-e', 'POSTGRES_PASSWORD=compat'] : ['-e', 'MYSQL_ROOT_PASSWORD=compat']
    const started = sh(['docker', 'run', '-d', '--rm', '--name', name, ...env, '-p', `${port}:${image.startsWith('postgres') ? 5432 : 3306}`, image])
    if (started.status !== 0) throw new Error(`docker run ${image}: ${started.stderr}`)
  }
  for (let i = 0; i < 90; i++) {
    try {
      await admin(image.startsWith('postgres') ? 'postgresql' : 'mysql', 'SELECT 1')
      return
    } catch {
      await new Promise((r) => setTimeout(r, 1000))
    }
  }
  throw new Error(`${image} did not come up on port ${port}`)
}

async function admin(dialect, sql) {
  if (dialect === 'postgresql') {
    const { default: postgres } = await import('postgres')
    const client = postgres(`postgresql://postgres:compat@127.0.0.1:${PG.port}/postgres`, { onnotice: () => {}, max: 1 })
    try {
      await client.unsafe(sql)
    } finally {
      await client.end()
    }
  } else {
    const mysql = await import('mysql2/promise')
    const conn = await mysql.createConnection(`mysql://root:compat@127.0.0.1:${MYSQL.port}`)
    try {
      await conn.query(sql)
    } finally {
      await conn.end()
    }
  }
}

async function freshDatabase(dialect, work, name) {
  const db = name.toLowerCase().replace(/[^a-z0-9_]/g, '_')
  if (dialect === 'postgresql') {
    await admin(dialect, `DROP DATABASE IF EXISTS "${db}" WITH (FORCE)`)
    await admin(dialect, `CREATE DATABASE "${db}"`)
    return `postgresql://postgres:compat@127.0.0.1:${PG.port}/${db}`
  }
  if (dialect === 'mysql') {
    await admin(dialect, `DROP DATABASE IF EXISTS \`${db}\``)
    await admin(dialect, `CREATE DATABASE \`${db}\``)
    return `mysql://root:compat@127.0.0.1:${MYSQL.port}/${db}`
  }
  const file = join(work, `${db}.db`)
  rmSync(file, { force: true })
  return `file:${file}`
}

function prisma(work, url, command, allowFail = false) {
  const config = join(work, 'prisma.config.mjs')
  writeFileSync(config, `export default { schema: './schema.prisma', datasource: { url: process.env.DATABASE_URL } }\n`)
  const run = sh(['node', prismaCli, ...command, '--config', config], {
    cwd: work,
    env: { ...process.env, DATABASE_URL: url, CHECKPOINT_DISABLE: '1', PRISMA_HIDE_UPDATE_MESSAGE: '1' },
  })
  if (run.status !== 0 && !allowFail) throw new Error(`prisma ${command.join(' ')}: ${(run.stderr || run.stdout).trim().slice(-2000)}`)
  return run
}

function cleanDiff(sql) {
  const text = sql.trim()
  return text === '' || /^-- This is an empty migration\.?$/.test(text) ? '' : text
}

function summary(results) {
  return {
    orms: results.map((r) => ({
      orm: r.orm,
      fatal: Boolean(r.fatal),
      checks: r.crud?.counts ?? null,
      findings: r.crud?.findings.length ?? null,
      ddl: r.ddl?.error ? 'error' : r.ddl?.diff === '' ? 'same' : r.ddl ? 'differs' : null,
    })),
  }
}

function sh([cmd, ...rest], options = {}) {
  const run = spawnSync(cmd, rest, { encoding: 'utf8', maxBuffer: 64 * 1024 * 1024, ...options })
  if (run.error) return { status: 1, stdout: '', stderr: String(run.error) }
  return run
}

await main()
