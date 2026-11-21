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
import { cpSync, existsSync, mkdirSync, symlinkSync, readdirSync, readFileSync, rmSync, writeFileSync } from 'node:fs'
import { createRequire } from 'node:module'
import { homedir } from 'node:os'
import { dirname, join, resolve } from 'node:path'
import { fileURLToPath } from 'node:url'
import { makePlan } from './lib/plan.mjs'
import { compare } from './lib/compare.mjs'
import { renderSchema, renderIndex } from './lib/report.mjs'

const here = dirname(fileURLToPath(import.meta.url))
const require = createRequire(import.meta.url)
const prismaCli = require.resolve('prisma/build/index.js')
const schemasDir = resolve(here, '../schemas')
// What a run writes (plans, SQLite files, copies of the generated code) goes outside the source
// tree, where a bind-mounted workspace cannot leave stale files; its node_modules links back here,
// so the copied drizzle schema resolves drizzle-orm from this package.
const workRoot = process.env.COMPAT_WORK ?? join(homedir(), '.cache', 'hekireki-compat')
const resultsDir = join(here, 'results')

const ORMS = ['drizzle', 'kysely', 'sqlalchemy', 'django', 'gorm', 'sea-orm', 'ecto', 'activerecord', 'eloquent', 'efcore', 'exposed']
// pgvector's image is PostgreSQL 17 with the vector extension, which formbricks asks for.
const PG = { port: 55432, image: 'pgvector/pgvector:pg17', name: 'hekireki-compat-pg' }
const MYSQL = { port: 53306, image: 'mysql:8.4', name: 'hekireki-compat-mysql' }
// The Docker network the databases join, for the harnesses that run in a container of their own
// (Active Record, Eloquent): they reach the databases by container name.
const NETWORK = 'hekireki-compat'

const args = process.argv.slice(2)
let orms = ORMS
const schemas = []
for (let i = 0; i < args.length; i++) {
  if (args[i] === '--orm') orms = args[++i].split(',')
  else schemas.push(args[i])
}
if (schemas.length === 0) throw new Error('usage: node run.mjs [--orm a,b] <schema> [schema ...]')

async function main() {
for (const orm of ['efcore', 'exposed']) if (orms.includes(orm)) sh(['docker', 'volume', 'rm', '-f', `hekireki-compat-${orm}`])
// The SeaORM volume keeps its target directory (the crates' build) and drops the binaries.
if (orms.includes('sea-orm')) sh(['docker', 'run', '--rm', '-v', 'hekireki-compat-sea-orm:/cache', 'busybox', 'rm', '-rf', '/cache/bin'])
mkdirSync(resultsDir, { recursive: true })
mkdirSync(workRoot, { recursive: true })
if (!existsSync(join(workRoot, 'node_modules'))) symlinkSync(join(here, 'node_modules'), join(workRoot, 'node_modules'), 'dir')
for (const name of schemas) {
  const source = readFileSync(join(schemasDir, name, 'schema.prisma'), 'utf8')
  const plan = await makePlan(source)
  const work = join(workRoot, name)
  rmSync(work, { recursive: true, force: true })
  mkdirSync(work, { recursive: true })
  // The schema without its generator blocks: `db push` and `migrate diff` need only the models.
  // The preview features the generators asked for (views, partialIndexes, postgresqlExtensions,
  // ...) are kept on one generator block of their own, which `db push` never runs.
  const features = [...new Set([...source.matchAll(/^\s*previewFeatures\s*=\s*\[([^\]]*)\]/gm).flatMap((m) => m[1].split(',').map((x) => x.trim()).filter(Boolean))])]
  const bare =
    source.replace(/^generator\s+[\w-]+\s*\{[^}]*\}\s*$/gms, '') +
    (features.length ? `\ngenerator compat {\n  provider        = "prisma-client"\n  output          = "./client"\n  previewFeatures = [${features.join(', ')}]\n}\n` : '')
  writeFileSync(join(work, 'schema.prisma'), bare)
  writeFileSync(join(work, 'plan.json'), JSON.stringify(plan, null, 2))
  if (plan.dialect === 'postgresql') await up(PG)
  if (plan.dialect === 'mysql') await up(MYSQL)

  const results = []
  for (const orm of orms) {
    console.log(`${name}: ${orm}`)
    const result = { orm }
    results.push(result)
    // hekireki leaves a generator out of a schema it refuses (schemas.bzl, SKIPPED).
    const generatedDir = { drizzle: 'drizzle', kysely: 'kysely', sqlalchemy: 'sqlalchemy', django: 'django', gorm: 'gorm', ecto: 'ecto', activerecord: 'activerecord', eloquent: 'eloquent', efcore: 'efcore', exposed: 'exposed', 'sea-orm': 'sea-orm' }[orm]
    if (!existsSync(join(schemasDir, name, 'generated', generatedDir))) {
      result.notGenerated = true
      continue
    }
    try {
      const crudUrl = await freshDatabase(plan.dialect, work, `${name}_${orm}`)
      prisma(work, crudUrl, ['db', 'push'])
      const harness = HARNESSES[orm]
      const prepared = harness.prepare(name, work)
      const command = harness.crud(prepared, join(work, 'plan.json'), crudUrl)
      const run = Array.isArray(command) ? sh(command) : command.run ? command.run() : sh(command.cmd, { cwd: command.cwd, env: command.env })
      if (run.status !== 0 || !run.stdout.trim().startsWith('{')) {
        result.fatal = cleanError(run.stderr || run.stdout)
      } else {
        result.crud = compare(plan, JSON.parse(run.stdout))
      }
      // ddl
      if (!harness.ddl) {
        result.ddl = { na: harness.noDdl }
        continue
      }
      const ddlUrl = await freshDatabase(plan.dialect, work, `${name}_${orm}_ddl`)
      const ddl = harness.ddl(prepared, join(work, 'plan.json'), ddlUrl)
      if (ddl.error) result.ddl = { error: ddl.error }
      else {
        const diff = prisma(work, ddlUrl, ['migrate', 'diff', '--from-config-datasource', '--to-schema', join(work, 'schema.prisma'), '--script'], true)
        result.ddl = diff.status === 0 ? { diff: cleanDiff(diff.stdout) } : { error: cleanError(diff.stderr || diff.stdout) }
      }
    } catch (e) {
      result.fatal = String(e?.stack ?? e).slice(0, 3000)
    }
  }
  // The schema's databases go once it is done: thirty schemas times eleven ORMs fill a disk.
  for (const db of created.splice(0)) await admin(plan.dialect, plan.dialect === 'postgresql' ? `DROP DATABASE IF EXISTS "${db}" WITH (FORCE)` : `DROP DATABASE IF EXISTS \`${db}\``).catch(() => {})
  if (plan.dialect === 'sqlite') for (const f of readdirSync(work).filter((f) => f.endsWith('.db'))) rmSync(join(work, f), { force: true })
  // A run of some ORMs keeps what earlier runs found for the others.
  const file = join(resultsDir, `${name}.json`)
  const earlier = existsSync(file) ? (JSON.parse(readFileSync(file, 'utf8')).results ?? []) : []
  const merged = ORMS.map((orm) => results.find((r) => r.orm === orm) ?? earlier.find((r) => r.orm === orm)).filter(Boolean)
  writeFileSync(join(resultsDir, `${name}.md`), renderSchema(name, plan, merged))
  writeFileSync(file, JSON.stringify({ ...summary(merged), results: merged }, null, 2) + '\n')
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
      const dir = join(workRoot, name, 'drizzle')
      mkdirSync(dir, { recursive: true })
      cpSync(join(schemasDir, name, 'generated/drizzle/schema.ts'), join(dir, 'schema.ts'))
      return { schema: join(dir, 'schema.ts') }
    },
    crud: (p, plan, url) => ['node', join(here, 'drizzle/harness.mjs'), 'crud', plan, p.schema, url],
    ddl(p, plan, url) {
      const sql = sh(['node', join(here, 'drizzle/harness.mjs'), 'ddl', plan, p.schema])
      if (sql.status !== 0) return { error: cleanError(sql.stderr) }
      const file = join(dirname(p.schema), 'ddl.sql')
      writeFileSync(file, sql.stdout)
      const applied = prisma(dirname(dirname(p.schema)), url, ['db', 'execute', '--file', file], true)
      return applied.status === 0 ? {} : { error: cleanError(applied.stderr || applied.stdout) }
    },
  },
  kysely: {
    noDdl: 'hekireki-kysely writes types only; the tables come from Prisma.',
    prepare(name) {
      return { types: join(schemasDir, name, 'generated/kysely/types.ts') }
    },
    crud: (p, plan, url) => ['node', join(here, 'kysely/harness.mjs'), 'crud', plan, p.types, url],
  },
  django: {
    prepare(name) {
      // The generated models.py as the app `app` of a throwaway project, with an empty migrations
      // package for `makemigrations`.
      const dir = join(workRoot, name, 'django')
      rmSync(dir, { recursive: true, force: true })
      mkdirSync(join(dir, 'app', 'migrations'), { recursive: true })
      writeFileSync(join(dir, 'app', '__init__.py'), '')
      writeFileSync(join(dir, 'app', 'migrations', '__init__.py'), '')
      cpSync(join(schemasDir, name, 'generated/django/models.py'), join(dir, 'app', 'models.py'))
      return { project: dir }
    },
    crud: (p, plan, url) => [...uvRun('django'), 'crud', plan, p.project, url],
    ddl(p, plan, url) {
      const run = sh([...uvRun('django'), 'ddl', plan, p.project, url])
      return run.status === 0 ? {} : { error: cleanError(run.stderr || run.stdout) }
    },
  },
  ecto: {
    noDdl: 'Ecto schemas do not create tables (migrations are written by hand); the tables come from Prisma.',
    prepare(name) {
      return { image: dockerBuild('ecto'), generated: join(schemasDir, name, 'generated/ecto') }
    },
    // The image holds the mix project with its deps compiled; the generated schemas go to
    // lib/generated and the harness from this tree replaces the image's.
    crud: (p, plan, url) => ({
      run: () =>
        dockerHarness(
          p.image,
          { 'plan.json': plan, generated: p.generated, 'harness.exs': join(here, 'ecto/harness.exs') },
          ['sh', '-c', 'cp -r /work/generated /harness/lib/generated && cp /work/harness.exs /harness/ && cd /harness && mix compile >/dev/null 2>/tmp/compile.log || { cat /tmp/compile.log >&2; exit 1; }; mix run harness.exs crud /work/plan.json "$0"'],
          url,
        ),
    }),
  },
  efcore: {
    prepare(name) {
      return { image: dockerBuild('efcore'), models: join(schemasDir, name, 'generated/efcore'), name }
    },
    // The build goes to a Docker volume, per schema, so that the ddl run reuses the crud run's.
    crud: (p, plan, url) => ({ run: () => efcoreRun(p, plan, url, 'crud') }),
    ddl(p, plan, url) {
      const run = efcoreRun(p, plan, url, 'ddl')
      return run.status === 0 ? {} : { error: cleanError(run.stderr || run.stdout) }
    },
  },
  exposed: {
    prepare(name) {
      const models = join(schemasDir, name, 'generated/exposed')
      const pkg = /^package\s+([\w.]+)/m.exec(readFileSync(join(models, 'PrismaSchema.kt'), 'utf8'))[1]
      return { image: dockerBuild('exposed'), models, name, schemaClass: `${pkg}.PrismaSchema` }
    },
    crud: (p, plan, url) => ({ run: () => exposedRun(p, plan, url, 'crud') }),
    ddl(p, plan, url) {
      const run = exposedRun(p, plan, url, 'ddl')
      return run.status === 0 ? {} : { error: cleanError(run.stderr || run.stdout) }
    },
  },
  'sea-orm': {
    prepare(name) {
      const models = join(schemasDir, name, 'generated/sea-orm')
      // Each entity module by its table, for the registry that calls the generic harness.
      const entities = readdirSync(models)
        .filter((f) => f.endsWith('.rs') && !['mod.rs', 'prelude.rs'].includes(f))
        .map((f) => ({ module: f.slice(0, -3), table: /table_name\s*=\s*"([^"]+)"/.exec(readFileSync(join(models, f), 'utf8'))?.[1] }))
        .filter((e) => e.table)
      const arms = (call, wrap) => entities.map((e) => `        ${JSON.stringify(e.table)} => ${wrap(`crate::${call}::<crate::models::r#${e.module}::Entity>`)},`).join('\n')
      const registry = `// Written by run.mjs: the generic harness of main.rs for each entity, by its table.
#![allow(unused)]
use crate::State;
use sea_orm::{DatabaseConnection, DbErr};
use serde_json::Value as Json;

pub async fn insert(db: &DatabaseConnection, m: &Json, st: &mut State) -> Option<Json> {
    match m["table"].as_str()? {
${arms('insert', (f) => `Some(${f}(db, m, st).await)`)}
        _ => None,
    }
}

pub async fn relations(db: &DatabaseConnection, m: &Json, st: &State) -> Option<Json> {
    match m["table"].as_str()? {
${arms('relations', (f) => `${f}(db, m, st).await`)}
        _ => None,
    }
}

pub async fn delete(db: &DatabaseConnection, m: &Json, st: &State) -> Option<Json> {
    match m["table"].as_str()? {
${arms('delete', (f) => `${f}(db, m, st).await`)}
        _ => None,
    }
}

pub async fn create(db: &DatabaseConnection, table: &str) -> Option<Result<(), DbErr>> {
    match table {
${entities.map((e) => `        ${JSON.stringify(e.table)} => Some(crate::create_table(db, crate::models::r#${e.module}::Entity).await),`).join('\n')}
        _ => None,
    }
}
`
      const stage = join(workRoot, name, 'sea-orm')
      rmSync(stage, { recursive: true, force: true })
      mkdirSync(stage, { recursive: true })
      writeFileSync(join(stage, 'registry.rs'), registry)
      return { image: dockerBuild('sea-orm'), models, registry: join(stage, 'registry.rs'), name }
    },
    crud: (p, plan, url) => ({ run: () => seaOrmRun(p, plan, url, 'crud') }),
    ddl(p, plan, url) {
      const work = dirname(plan)
      const types = enumTypes(work, JSON.parse(readFileSync(plan, 'utf8')).dialect)
      if (types) {
        const applied = prisma(work, url, ['db', 'execute', '--file', types], true)
        if (applied.status !== 0) return { error: cleanError(applied.stderr || applied.stdout) }
      }
      const run = seaOrmRun(p, plan, url, 'ddl')
      return run.status === 0 ? {} : { error: cleanError(run.stderr || run.stdout) }
    },
  },
  eloquent: {
    noDdl: 'hekireki-eloquent writes models only (no migrations); the tables come from Prisma.',
    prepare(name) {
      return { image: dockerBuild('eloquent'), models: join(schemasDir, name, 'generated/eloquent') }
    },
    crud: (p, plan, url) => ({
      run: () => dockerHarness(p.image, { 'plan.json': plan, models: p.models, 'harness.php': join(here, 'eloquent/harness.php') }, ['php', '/work/harness.php', 'crud', '/work/plan.json', '/work/models'], url),
    }),
  },
  activerecord: {
    noDdl: 'hekireki-activerecord writes models only (no migrations or schema.rb); the tables come from Prisma.',
    prepare(name, work) {
      return { image: dockerBuild('activerecord'), models: join(schemasDir, name, 'generated/activerecord'), work }
    },
    crud: (p, plan, url) => ({
      run: () =>
        dockerHarness(p.image, { 'plan.json': plan, models: p.models, 'harness.rb': join(here, 'activerecord/harness.rb') }, ['bundle', 'exec', '--gemfile=/harness/Gemfile', 'ruby', '/work/harness.rb', 'crud', '/work/plan.json', '/work/models'], url),
    }),
  },
  sqlalchemy: {
    prepare(name, work) {
      return { models: join(schemasDir, name, 'generated/sqlalchemy/models.py') }
    },
    crud: (p, plan, url) => [...uvRun('sqlalchemy'), 'crud', plan, p.models, url],
    ddl(p, plan, url) {
      const run = sh([...uvRun('sqlalchemy'), 'ddl', plan, p.models, url])
      return run.status === 0 ? {} : { error: cleanError(run.stderr || run.stdout) }
    },
  },
  gorm: {
    prepare(name, work) {
      // The generated package, copied beside the harness's main package with a registry of every
      // struct that has a TableName, which Go cannot list by reflection.
      const dir = join(workRoot, name, 'gorm')
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
      return run.status === 0 ? {} : { error: cleanError(run.stderr || run.stdout) }
    },
  },
}

// Runs a harness in its image: the files go in on stdin as a tar (the Docker daemon may not see
// this machine's paths), a SQLite database among them, and the URL points at the databases by
// container name.
// A tool's error without its stack trace: the message is what tells what went wrong.
function cleanError(text) {
  const lines = text.split('\n').filter((l) => !/^\s+at /.test(l) && !/^\s*--- End of/.test(l))
  return lines.join('\n').trim().slice(0, 2000)
}

function dockerHarness(image, files, args, url, volumes = []) {
  const stage = join(workRoot, `.stage-${process.pid}`)
  rmSync(stage, { recursive: true, force: true })
  mkdirSync(stage, { recursive: true })
  for (const [dest, src] of Object.entries(files)) cpSync(src, join(stage, dest), { recursive: true })
  let inner = url
  const sqlite = url.startsWith('file:') ? url.slice('file:'.length) : null
  if (sqlite) {
    if (existsSync(sqlite)) cpSync(sqlite, join(stage, 'db.sqlite'))
    inner = 'file:/work/db.sqlite'
  }
  inner = inner.replace(`127.0.0.1:${PG.port}`, `${PG.name}:5432`).replace(`127.0.0.1:${MYSQL.port}`, `${MYSQL.name}:3306`)
  const script = `mkdir -p /work && tar -C /work -xf - && cd /work && exec ${args.map((a) => `'${a}'`).join(' ')} '${inner}'`
  const mounts = volumes.map((v) => `-v ${v}`).join(' ')
  // The container is kept until a SQLite database it wrote (the ddl run's) is copied back.
  const name = `hekireki-compat-run-${process.pid}`
  sh(['docker', 'rm', '-f', '-v', name]) // -v: an image's declared VOLUME (Gradle's cache) would be left behind
  const run = sh(['bash', '-c', `tar -C '${stage}' -cf - . | docker run -i --name ${name} --network ${NETWORK} ${mounts} ${image} sh -c "$0"`, script])
  if (sqlite) sh(['docker', 'cp', `${name}:/work/db.sqlite`, sqlite])
  sh(['docker', 'rm', '-f', '-v', name]) // -v: an image's declared VOLUME (Gradle's cache) would be left behind
  if (run.status !== 0) run.stderr = cleanError(run.stderr || '')
  return run
}

function exposedRun(p, plan, url, mode) {
  const out = `/cache/${p.name}`
  const build = `if [ ! -x ${out}/bin/compat ]; then mkdir -p /harness/src/main/kotlin/compat && cp /work/Main.kt /harness/src/main/kotlin/compat/ && cp -r /work/models /harness/src/main/kotlin/models && cd /harness && gradle installDist --offline --no-daemon -q >/tmp/build.log 2>&1 || { tail -40 /tmp/build.log >&2; exit 1; }; rm -rf ${out} && cp -r build/install/compat ${out}; fi`
  return dockerHarness(
    p.image,
    { 'plan.json': plan, models: p.models, 'Main.kt': join(here, 'exposed/Main.kt') },
    ['sh', '-c', `${build}; exec ${out}/bin/compat ${mode} /work/plan.json ${p.schemaClass} "$0"`],
    url,
    ['hekireki-compat-exposed:/cache'],
  )
}

// The enum types of a PostgreSQL schema, from Prisma's own DDL: SeaORM's create_table_from_entity
// creates tables only (create_enum_from_active_enum takes each enum by its type).
function enumTypes(work, provider) {
  if (provider !== 'postgresql') return null
  const diff = prisma(work, 'postgresql://unused', ['migrate', 'diff', '--from-empty', '--to-schema', join(work, 'schema.prisma'), '--script'], true)
  const types = diff.stdout.split(/;\s*\n/).filter((st) => /^\s*(--[^\n]*\n\s*)*CREATE TYPE/i.test(st))
  if (types.length === 0) return null
  const file = join(work, 'enum-types.sql')
  writeFileSync(file, types.map((t) => t.trim() + ';').join('\n') + '\n')
  return file
}

function seaOrmRun(p, plan, url, mode) {
  // The target directory lives in the volume, so the crates build once for every schema; the
  // binary is copied out per schema.
  const out = `/cache/bin/${p.name}`
  const build = `if [ ! -x ${out} ]; then mkdir -p /harness/src /cache/bin && cp /work/main.rs /work/registry.rs /harness/src/ && cp -r /work/models /harness/src/models && cd /harness && CARGO_INCREMENTAL=0 CARGO_PROFILE_DEV_CODEGEN_UNITS=256 CARGO_TARGET_DIR=/cache/target cargo build -j 1 >/tmp/build.log 2>&1 || { grep -E -A12 "^error" /tmp/build.log | head -60 >&2; tail -3 /tmp/build.log >&2; exit 1; }; cp /cache/target/debug/compat ${out}; fi`
  return dockerHarness(
    p.image,
    { 'plan.json': plan, models: p.models, 'main.rs': join(here, 'sea-orm/main.rs'), 'registry.rs': p.registry },
    ['sh', '-c', `${build}; exec ${out} ${mode} /work/plan.json "$0"`],
    url,
    ['hekireki-compat-sea-orm:/cache'],
  )
}

function efcoreRun(p, plan, url, mode) {
  const provider = JSON.parse(readFileSync(plan, 'utf8')).dialect
  const out = `/cache/${p.name}`
  const build = `if [ ! -f ${out}/Compat.dll ]; then cp /work/Program.cs /harness/ && cp -r /work/models /harness/Models && cd /harness && dotnet build -c Release -p:HekirekiProvider=${provider} -o ${out} >/tmp/build.log 2>&1 || { tail -40 /tmp/build.log >&2; exit 1; }; fi`
  return dockerHarness(
    p.image,
    { 'plan.json': plan, models: p.models, 'Program.cs': join(here, 'efcore/Program.cs') },
    ['sh', '-c', `${build}; exec dotnet ${out}/Compat.dll ${mode} /work/plan.json "$0"`],
    url,
    ['hekireki-compat-efcore:/cache'],
  )
}

function dockerBuild(orm) {
  const image = `hekireki-compat-${orm}`
  const built = sh(['docker', 'build', '-q', '-t', image, join(here, orm)])
  if (built.status !== 0) throw new Error(`docker build ${orm}: ${built.stderr.trim().slice(-2000)}`)
  return image
}

function uvRun(harness) {
  return ['uv', 'run', '--quiet', '--project', join(here, harness), 'python', join(here, harness, 'harness.py')]
}

function goEnv() {
  return { ...process.env, GOFLAGS: '-mod=mod', GOTOOLCHAIN: 'auto' }
}

// --- databases --------------------------------------------------------------------------------

async function up({ name, image, port }) {
  sh(['docker', 'network', 'create', NETWORK])
  const running = sh(['docker', 'ps', '--format', '{{.Image}}', '-f', `name=^${name}$`]).stdout.trim()
  if (running !== image) {
    sh(['docker', 'rm', '-f', '-v', name]) // -v: an image's declared VOLUME (Gradle's cache) would be left behind
    const env = isPg(image) ? ['-e', 'POSTGRES_PASSWORD=compat'] : ['-e', 'MYSQL_ROOT_PASSWORD=compat']
    const started = sh(['docker', 'run', '-d', '--rm', '--name', name, '--network', NETWORK, ...env, '-p', `${port}:${isPg(image) ? 5432 : 3306}`, image])
    if (started.status !== 0) throw new Error(`docker run ${image}: ${started.stderr}`)
  }
  sh(['docker', 'network', 'connect', NETWORK, name]) // a container started before the network existed
  for (let i = 0; i < 90; i++) {
    try {
      await admin(isPg(image) ? 'postgresql' : 'mysql', 'SELECT 1')
      return
    } catch {
      await new Promise((r) => setTimeout(r, 1000))
    }
  }
  throw new Error(`${image} did not come up on port ${port}`)
}

function isPg(image) {
  return image.startsWith('postgres') || image.startsWith('pgvector')
}

async function admin(dialect, sql, database = 'postgres') {
  if (dialect === 'postgresql') {
    const { default: postgres } = await import('postgres')
    const client = postgres(`postgresql://postgres:compat@127.0.0.1:${PG.port}/${database}`, { onnotice: () => {}, max: 1 })
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

const created = []

async function freshDatabase(dialect, work, name) {
  const db = name.toLowerCase().replace(/[^a-z0-9_]/g, '_')
  if (dialect !== 'sqlite') created.push(db)
  if (dialect === 'postgresql') {
    await admin(dialect, `DROP DATABASE IF EXISTS "${db}" WITH (FORCE)`)
    await admin(dialect, `CREATE DATABASE "${db}"`)
    // The extensions schemas use without declaring them (citext columns, gin_trgm_ops indexes,
    // gen_random_uuid() and uuid_generate_v4() defaults), as a provider's database has them.
    await admin(dialect, 'CREATE EXTENSION IF NOT EXISTS citext; CREATE EXTENSION IF NOT EXISTS pg_trgm; CREATE EXTENSION IF NOT EXISTS pgcrypto; CREATE EXTENSION IF NOT EXISTS "uuid-ossp"', db)
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
      notGenerated: Boolean(r.notGenerated),
      checks: r.crud?.counts ?? null,
      findings: r.crud?.findings.length ?? null,
      ddl: r.ddl?.na ? 'na' : r.ddl?.error ? 'error' : r.ddl?.diff === '' ? 'same' : r.ddl ? 'differs' : null,
    })),
  }
}

function sh([cmd, ...rest], options = {}) {
  // No step of a run takes 15 minutes; one that does is stuck.
  const run = spawnSync(cmd, rest, { encoding: 'utf8', maxBuffer: 64 * 1024 * 1024, timeout: 15 * 60_000, killSignal: 'SIGKILL', ...options })
  if (run.signal === 'SIGKILL') return { status: 1, stdout: run.stdout ?? '', stderr: `timed out after 15 minutes\n${run.stderr ?? ''}` }
  if (run.error) return { status: 1, stdout: '', stderr: String(run.error) }
  return run
}

await main()
