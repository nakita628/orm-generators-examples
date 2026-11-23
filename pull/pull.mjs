#!/usr/bin/env node
// Writes up, in Prisma, the SQL that other ORMs generate. A well-known project's own ORM creates
// its tables from the project's definitions (Django models, a Drizzle schema, a Rails
// db/schema.rb, ...), the tables are kept as SQL, and `prisma db pull` writes them up as a Prisma
// schema. Then hekireki generates every ORM's code from that schema, and apps/hekireki/compat
// runs the code against the tables the project's own ORM created: whether what hekireki writes
// from a schema it did not see written for Prisma is valid for each ORM, and works.
//
//   node pull/pull.mjs <orm>/<project> [...]     django/wagtail, drizzle/northwind, activerecord/mastodon
//   node pull/pull.mjs --random 3                three: never pulled first, then the oldest
//   node pull/pull.mjs --list                    every source, with when it was last checked
//   node pull/pull.mjs --pull-only <...>         only the write-up: source.sql and schema.prisma
//   node pull/pull.mjs --orm django,gorm <...>   check only these of hekireki's ORMs
//   node pull/pull.mjs --hekireki latest <...>   check with another hekireki release (installed in
//                                                the cache); nothing under pull/ changes, the
//                                                results go to the cache
//
// The layers, one per ORM that wrote the SQL (pull/<orm>/README.md for each):
//   django/<example>        every apps/djangoer/examples/<example>: Django's schema editor creates
//                           every model's table (django/create_tables.py), on PostgreSQL
//   drizzle/<example>       every apps/drizzler/examples/<example>: `drizzle-kit export`, in the
//                           example's dialect
//   activerecord/<app>      activerecord/sources.json: a Rails app's db/schema.rb at a pinned
//                           commit, loaded by Active Record (activerecord/load_schema.rb)
//   sqlalchemy/<project>    sqlalchemy/sources.json: a project from PyPI, by its own command or
//                           metadata.create_all() (sqlalchemy/create_all.py)
// A layer's sources.json may also `skip` a source that cannot be pulled, with the reason.
//
// What it writes, in pull/<orm>/<project>/ (committed):
//   source.sql       the tables as the project's ORM created them (pg_dump, mysqldump, sqlite_master)
//   schema.prisma    `prisma db pull` of those tables: the write-up
//   prisma-diff.sql  what `prisma migrate diff` would run to turn the tables into the schema's,
//                    that is what the Prisma schema cannot say about them (left out when empty)
//   source.json      where it came from, the versions, and when it was checked
// and generated/ (hekireki's output, not committed; a generator that fails leaves
// generated/<orm>.error). The checks go to pull/<orm>/results/.
import { spawnSync } from 'node:child_process'
import { cpSync, existsSync, mkdirSync, readdirSync, readFileSync, rmSync, writeFileSync } from 'node:fs'
import { createRequire } from 'node:module'
import { homedir } from 'node:os'
import { dirname, join, relative, resolve } from 'node:path'
import { fileURLToPath } from 'node:url'

const here = dirname(fileURLToPath(import.meta.url))
const repo = resolve(here, '..')
// Prisma and hekireki are apps/hekireki's, so that a pull is checked with the versions the rest
// of the repository uses; the compat harnesses are apps/hekireki/compat's.
const hekirekiRequire = createRequire(join(repo, 'apps/hekireki/package.json'))
// node:sqlite only when a SQLite source needs it: loading it prints an ExperimentalWarning.
const sqlite = (file) => new (createRequire(import.meta.url)('node:sqlite').DatabaseSync)(file)
const prismaCli = hekirekiRequire.resolve('prisma/build/index.js')
const compat = join(repo, 'apps/hekireki/compat')
const { renderIndex } = await import(join(compat, 'lib/report.mjs'))
let generateMjs = join(repo, 'apps/hekireki/tools/hekireki/generate.mjs')
let layersDir = here
const workRoot = process.env.PULL_WORK ?? join(homedir(), '.cache', 'orm-generators-examples', 'pull')
const djangoerExamples = join(repo, 'apps/djangoer/examples')
const drizzlerExamples = join(repo, 'apps/drizzler/examples')

const LAYERS = ['django', 'drizzle', 'activerecord', 'sqlalchemy']
const ORM_NAMES = { django: 'Django', drizzle: 'drizzle-kit', activerecord: 'Active Record', sqlalchemy: 'SQLAlchemy' }
const ORMS = ['drizzle', 'kysely', 'sqlalchemy', 'django', 'gorm', 'sea-orm', 'ecto', 'activerecord', 'eloquent', 'efcore', 'exposed']
// The same containers, ports and network as apps/hekireki/compat/run.mjs, which then finds them up.
const PG = { port: 55432, image: 'pgvector/pgvector:pg17', name: 'hekireki-compat-pg' }
const MYSQL = { port: 53306, image: 'mysql:8.4', name: 'hekireki-compat-mysql' }
const NETWORK = 'hekireki-compat'
// The extensions a project's tables may use without creating them (Django's HStoreField and
// GinIndex on scalars, citext, trigram and vector indexes, uuid and crypto defaults).
const PG_EXTENSIONS = ['citext', 'pg_trgm', 'pgcrypto', 'uuid-ossp', 'hstore', 'btree_gin', 'btree_gist', 'unaccent', 'vector', 'ltree']

// --- arguments ----------------------------------------------------------------------------------

const args = process.argv.slice(2)
let orms = ORMS
let random = 0
let list = false
let pullOnly = false
let hekirekiVersion = null
const wanted = []
for (let i = 0; i < args.length; i++) {
  if (args[i] === '--orm') orms = args[++i].split(',')
  else if (args[i] === '--random') random = Number(args[++i])
  else if (args[i] === '--list') list = true
  else if (args[i] === '--pull-only') pullOnly = true
  else if (args[i] === '--hekireki') hekirekiVersion = args[++i]
  else wanted.push(args[i])
}

// --- sources ------------------------------------------------------------------------------------

function layerConfig(layer) {
  const file = join(here, layer, 'sources.json')
  return existsSync(file) ? JSON.parse(readFileSync(file, 'utf8')) : { sources: {}, skip: {} }
}

function allSources() {
  const sources = []
  for (const name of readdirSync(djangoerExamples).sort()) {
    if (!existsSync(join(djangoerExamples, name, 'pyproject.toml'))) continue
    sources.push({ layer: 'django', name, dialect: 'postgresql', about: `the models of [apps/djangoer/examples/${name}](../../../apps/djangoer/examples/${name})` })
  }
  for (const name of readdirSync(drizzlerExamples).sort()) {
    const config = join(drizzlerExamples, name, 'drizzle.config.ts')
    if (!existsSync(config)) continue
    const dialect = /dialect:\s*["'](\w+)["']/.exec(readFileSync(config, 'utf8'))?.[1]
    sources.push({
      layer: 'drizzle',
      name,
      dialect: { postgresql: 'postgresql', mysql: 'mysql', sqlite: 'sqlite', turso: 'sqlite' }[dialect] ?? dialect,
      about: `the Drizzle schema of [apps/drizzler/examples/${name}](../../../apps/drizzler/examples/${name})`,
    })
  }
  for (const layer of ['activerecord', 'sqlalchemy']) {
    for (const [name, source] of Object.entries(layerConfig(layer).sources)) sources.push({ layer, name, ...source })
  }
  const skips = Object.fromEntries(LAYERS.map((l) => [l, layerConfig(l).skip ?? {}]))
  for (const s of sources) {
    s.key = `${s.layer}/${s.name}`
    s.skip = skips[s.layer][s.name]
    const info = join(here, s.layer, s.name, 'source.json')
    s.checkedAt = existsSync(info) ? JSON.parse(readFileSync(info, 'utf8')).checked_at : undefined
  }
  return sources
}

// Never pulled first, in an order fixed by the day (the same day picks the same ones), then the
// ones checked longest ago.
function pick(candidates, n) {
  let seed = [...new Date().toISOString().slice(0, 10)].reduce((h, c) => (h * 31 + c.charCodeAt(0)) >>> 0, 7)
  const rand = () => ((seed = (seed * 1103515245 + 12345) >>> 0) / 2 ** 32)
  const never = candidates.filter((s) => !s.checkedAt)
  for (let i = never.length - 1; i > 0; i--) {
    const j = Math.floor(rand() * (i + 1))
    ;[never[i], never[j]] = [never[j], never[i]]
  }
  const checked = candidates.filter((s) => s.checkedAt).sort((a, b) => a.checkedAt.localeCompare(b.checkedAt))
  return [...never, ...checked].slice(0, n)
}

// --- main ---------------------------------------------------------------------------------------

async function main() {
  const sources = allSources()
  if (list) {
    for (const s of sources) console.log(`${s.key.padEnd(48)} ${s.dialect.padEnd(10)} ${s.skip ? `skipped: ${s.skip}` : (s.checkedAt ?? 'never checked')}`)
    return
  }
  const chosen = wanted.map((w) => {
    const s = sources.find((x) => x.key === w)
    if (!s) throw new Error(`no source ${w}; see node pull/pull.mjs --list`)
    return s
  })
  if (random > 0) chosen.push(...pick(sources.filter((s) => !s.skip && !chosen.includes(s)), random))
  if (chosen.length === 0) throw new Error('usage: node pull/pull.mjs [--orm a,b] [--pull-only] [--hekireki <version>] (<orm>/<project> ... | --random N | --list)')

  mkdirSync(workRoot, { recursive: true })
  if (hekirekiVersion) useHekireki(hekirekiVersion, chosen)
  const failed = []
  for (const source of chosen) {
    console.log(`== ${source.key} (${source.dialect})`)
    try {
      if (!hekirekiVersion) await pullSource(source)
      if (!pullOnly) {
        generate(source)
        check(source)
      }
    } catch (e) {
      failed.push(source.key)
      console.error(`${source.key}: ${String(e?.message ?? e).slice(0, 4000)}`)
      recordFailure(source, e)
    }
  }
  for (const layer of new Set(chosen.map((s) => s.layer))) writeIndex(layer)
  if (hekirekiVersion) console.log(`results with hekireki ${hekirekiVersion}: ${layersDir}/<orm>/results/`)
  if (failed.length) {
    console.error(`could not pull: ${failed.join(', ')}`)
    process.exitCode = 1
  }
}

// --- 1. the project's ORM creates the tables; 2. they are kept as SQL and written up ------------

async function pullSource(source) {
  const dir = join(here, source.layer, source.name)
  const work = join(workRoot, source.layer, source.name)
  rmSync(work, { recursive: true, force: true })
  mkdirSync(work, { recursive: true })
  mkdirSync(dir, { recursive: true })

  const url = await freshDatabase(source.dialect, work, `pull_${source.layer}_${source.name}`)
  if (source.before_sql) execSql(source.dialect, url, source.before_sql)
  const created = CREATE[source.layer](source, url, work)
  if (source.after_sql) execSql(source.dialect, url, source.after_sql)

  writeFileSync(join(dir, 'source.sql'), dump(source.dialect, url))

  // `prisma db pull` into a schema of the datasource alone.
  const pulled = join(work, 'schema.prisma')
  writeFileSync(pulled, `datasource db {\n  provider = "${source.dialect}"\n}\n`)
  const pull = prisma(work, url, ['db', 'pull'])
  if (pull.status !== 0) throw new Error(`prisma db pull: ${tail(pull)}`)
  const pulledModels = readFileSync(pulled, 'utf8').replace(/^datasource\s+\w+\s*\{[^}]*\}\s*/m, '').trim()
  const { models, features, fixes } = validate(source, pulledModels, work, url)
  writeFileSync(join(dir, 'schema.prisma'), writeUp(source, models, features, fixes))

  // What the Prisma schema cannot say about the tables.
  const diff = prisma(work, url, ['migrate', 'diff', '--from-config-datasource', '--to-schema', join(dir, 'schema.prisma'), '--script'])
  const diffSql = diff.status === 0 ? diff.stdout.trim() : `-- prisma migrate diff failed:\n-- ${tail(diff).replaceAll('\n', '\n-- ')}`
  const diffFile = join(dir, 'prisma-diff.sql')
  if (diffSql && !/^-- This is an empty migration\.?$/.test(diffSql)) writeFileSync(diffFile, diffSql + '\n')
  else rmSync(diffFile, { force: true })

  const info = {
    source: source.key,
    about: source.about,
    ...(source.repo && { repo: source.repo, commit: source.commit, path: source.path, license: source.license }),
    ...(source.package && { package: source.package }),
    dialect: source.dialect,
    tables: created,
    models: [...models.matchAll(/^model \w+/gm)].length,
    ...(fixes.length && { fixes }),
    prisma: hekirekiRequire('prisma/package.json').version,
    pulled_at: new Date().toISOString(),
  }
  writeFileSync(join(dir, 'source.json'), JSON.stringify(info, null, 2) + '\n')
  console.log(`${source.key}: ${created} tables, ${info.models} models${existsSync(diffFile) ? ', prisma-diff.sql' : ''}`)
}

const CREATE = {
  django(source, url, work) {
    const example = join(djangoerExamples, source.name)
    const project = join(work, 'project')
    const notInputs = new Set(['er.md', 'schema.dbml', 'er.svg', 'docs', 'models', 'BUILD.bazel', '.venv', '__pycache__'])
    cpSync(example, project, { recursive: true, filter: (src) => !(dirname(src) === example && notInputs.has(src.slice(example.length + 1))) })
    const pyproject = readFileSync(join(example, 'pyproject.toml'), 'utf8')
    const settings = /^settings\s*=\s*"([^"]+)"/m.exec(pyproject)?.[1] ?? 'settings'
    const envLine = /^env\s*=\s*\{(.*)\}\s*$/m.exec(pyproject)?.[1] ?? ''
    const extraEnv = Object.fromEntries([...envLine.matchAll(/(\w+)\s*=\s*"([^"]*)"/g)].map((m) => [m[1], m[2]]))
    const packagePaths = [...(/^package_paths\s*=\s*\[([^\]]*)\]/m.exec(pyproject)?.[1] ?? '').matchAll(/"([^"]+)"/g)].map((m) => m[1])
    writeFileSync(
      join(project, 'pull_settings.py'),
      `# Written by pull.mjs: the example's settings, with the database to create the tables in.
import os
from urllib.parse import urlparse

from ${settings} import *  # noqa: F401,F403

_url = urlparse(os.environ["PULL_DATABASE_URL"])
DATABASES = {
    "default": {
        "ENGINE": "django.db.backends.postgresql",
        "NAME": _url.path.lstrip("/"),
        "USER": _url.username,
        "PASSWORD": _url.password,
        "HOST": _url.hostname,
        "PORT": _url.port,
    }
}
`,
    )
    // The example's own locked environment, in the cache Djangoer's Bazel targets use.
    const cache = process.env.DJANGOER_CACHE ?? join(homedir(), '.cache', 'orm-generators-examples', 'djangoer')
    const env = {
      ...process.env,
      UV_LINK_MODE: 'copy',
      UV_PROJECT_ENVIRONMENT: join(cache, 'venvs', source.name),
      UV_PYTHON: readFileSync(join(repo, 'apps/djangoer/.python-version'), 'utf8').trim(),
      UV_PYTHON_PREFERENCE: 'only-managed',
      UV_NO_PROGRESS: '1',
      ...extraEnv,
    }
    const uv = (...a) => sh(['uv', 'run', '--locked', '--quiet', '--with', 'psycopg[binary]', '--project', project, ...a], { cwd: project, env })
    const pythonpath = [project]
    for (const pkg of packagePaths) {
      const found = uv('python', '-c', `import importlib.util, pathlib; print(pathlib.Path(importlib.util.find_spec(${JSON.stringify(pkg)}).origin).parent)`)
      if (found.status !== 0) throw new Error(`find ${pkg}: ${tail(found)}`)
      pythonpath.push(found.stdout.trim())
    }
    env.PYTHONPATH = pythonpath.join(':')
    env.DJANGO_SETTINGS_MODULE = 'pull_settings'
    env.PULL_DATABASE_URL = url
    const run = uv('python', join(here, 'django', 'create_tables.py'))
    if (run.status !== 0) throw new Error(`Django could not create the tables: ${tail(run)}`)
    return run.stdout.trim().split('\n').filter(Boolean).length
  },

  drizzle(source, url, work) {
    const example = join(drizzlerExamples, source.name)
    const kit = join(example, 'node_modules', '.bin', 'drizzle-kit')
    if (!existsSync(kit)) throw new Error(`no drizzle-kit in ${relative(repo, example)}/node_modules: run pnpm install at the root`)
    const exported = sh([kit, 'export', '--config', 'drizzle.config.ts'], { cwd: example })
    if (exported.status !== 0) throw new Error(`drizzle-kit export: ${tail(exported)}`)
    writeFileSync(join(work, 'drizzle-kit-export.sql'), exported.stdout)
    execSql(source.dialect, url, exported.stdout)
    return (exported.stdout.match(/^CREATE TABLE/gim) ?? []).length
  },

  activerecord(source, url, work) {
    const file = join(work, source.path.split('/').pop())
    download(`https://raw.githubusercontent.com/${source.repo}/${source.commit}/${source.path}`, file)
    if (file.endsWith('.sql')) {
      execSql(source.dialect, url, readFileSync(file, 'utf8'))
      return countTables(source.dialect, url)
    }
    // Active Record runs in apps/hekireki/compat's image of it (Ruby, Active Record, pg, mysql2,
    // sqlite3).
    const stage = join(work, 'stage')
    mkdirSync(stage, { recursive: true })
    cpSync(file, join(stage, 'schema.rb'))
    cpSync(join(here, 'activerecord', 'load_schema.rb'), join(stage, 'load_schema.rb'))
    const image = dockerBuild('activerecord')
    const run = dockerRun(image, stage, ['bundle', 'exec', '--gemfile=/harness/Gemfile', 'ruby', 'load_schema.rb', '{url}', 'schema.rb'], url)
    if (run.status !== 0) throw new Error(`Active Record could not load ${source.path}: ${tail(run)}`)
    return run.stdout.trim().split('\n').filter(Boolean).length
  },

  sqlalchemy(source, url) {
    const withs = (source.with ?? []).flatMap((p) => ['--with', p])
    const base = ['uv', 'run', '--quiet', '--no-project', '--python', source.python ?? '3.13', '--with', source.package, ...withs, '--with', 'psycopg[binary]', '--with', 'pymysql']
    const sqlalchemyUrl = url.replace(/^postgresql:/, 'postgresql+psycopg:').replace(/^mysql:/, 'mysql+pymysql:')
    const command = source.command
      ? [...base, ...source.command.map((a) => a.replace('{url}', sqlalchemyUrl))]
      : [...base, 'python', join(here, 'sqlalchemy', 'create_all.py'), url, source.base, ...source.modules]
    const run = sh(command, { cwd: workRoot })
    if (run.status !== 0) throw new Error(`${source.package} could not create the tables: ${tail(run)}`)
    return countTables(source.dialect, url)
  },
}

// The write-up as it is committed: the datasource, a generator block only for the preview
// features the schema needs, and the models as `prisma db pull` wrote them.
function writeUp(source, models, features, fixes = []) {
  const header = `// ${source.key}: the tables that ${ORM_NAMES[source.layer]} creates from ${source.about.replace(/\]\([^)]*\)/g, ']').replace(/\[|\]/g, '')},
// written up by \`prisma db pull\`. Written by pull/pull.mjs; do not edit.${
    fixes.length
      ? `
//
// Changed from what \`prisma db pull\` wrote, which \`prisma validate\` refuses:
${fixes.map((f) => `// - ${f}`).join('\n')}`
      : ''
  }`
  const generator = features.length
    ? `\n\n// The preview features the schema needs.\ngenerator client {\n  provider        = "prisma-client"\n  output          = "./generated/prisma"\n  previewFeatures = [${features.map((f) => `"${f}"`).join(', ')}]\n}`
    : ''
  return `${header}\n\ndatasource db {\n  provider = "${source.dialect}"\n}${generator}\n\n${models}\n`
}

// `prisma validate` on the pulled schema, and the smallest change that makes it valid where
// `prisma db pull` writes what Prisma itself refuses:
// - a preview feature the schema needs (partialIndexes, ...) goes on a generator block;
// - a foreign key whose type is not its referenced key's (an integer column pointing at a bigint
//   id, common in Rails apps) takes the referenced key's type;
// - a default on an Unsupported(...) field, which only dbgenerated() may be, is dropped.
// Each change is listed in the schema's header and in source.json. Anything else is a schema
// Prisma cannot read, and the pull stops there.
function validate(source, pulledModels, work, url) {
  const lines = pulledModels.split('\n')
  const features = []
  const fixes = []
  for (let round = 0; round < 10; round++) {
    const header = writeUp(source, '', features).trimEnd()
    const offset = header.split('\n').length + 1
    const file = join(work, 'validate.prisma')
    writeFileSync(file, `${header}\n${lines.join('\n')}\n`)
    const run = prisma(work, url, ['validate', '--schema', file])
    if (run.status === 0) return { models: lines.join('\n'), features, fixes }
    const out = run.stdout + run.stderr
    const missing = /Add "(\w+)" to previewFeatures/.exec(out)?.[1]
    if (missing && !features.includes(missing)) {
      features.push(missing)
      continue
    }
    let changed = false
    for (const [, message, at] of out.matchAll(/error: (.*?)\n\s*-->\s*validate\.prisma:(\d+)/g)) {
      const index = Number(at) - offset
      const mismatch = /The type of the field `(\w+)` in the model `(\w+)` is not matching the type of the referenced field `(\w+)` in model `(\w+)`/.exec(message)
      if (mismatch) {
        const [, field, model, refField, refModel] = mismatch
        const ref = fieldLine(lines, refModel, refField)
        const own = fieldLine(lines, model, field)
        if (ref < 0 || own < 0) continue
        const refType = lines[ref].trim().split(/\s+/)[1].replace(/[?[\]]/g, '')
        const ownType = lines[own].trim().split(/\s+/)[1]
        const newType = refType + ownType.replace(/^\w+/, '')
        lines[own] = lines[own].replace(/^(\s*\w+\s+)\S+/, `$1${newType}`).replace(/\s@db\.\w+(\([^)]*\))?/, '')
        fixes.push(`${model}.${field}: ${ownType} -> ${newType}, the type of ${refModel}.${refField} it references`)
        changed = true
      } else if (/Only @default\(dbgenerated\("\.\.\."\)\) can be used for Unsupported types/.test(message) && index >= 0 && index < lines.length) {
        const before = lines[index]
        lines[index] = dropDefault(lines[index])
        if (lines[index] !== before) {
          fixes.push(`${modelOf(lines, index)}.${before.trim().split(/\s+/)[0]}: its default dropped (a default on an Unsupported field has to be dbgenerated())`)
          changed = true
        }
      }
    }
    if (!changed) throw new Error(`the pulled schema does not validate: ${tail(run)}`)
  }
  throw new Error('the pulled schema still does not validate after 10 rounds of changes')
}

function fieldLine(lines, model, field) {
  const start = lines.findIndex((l) => new RegExp(`^model ${model} \\{`).test(l))
  if (start < 0) return -1
  for (let i = start + 1; i < lines.length && !/^\}/.test(lines[i]); i++) if (new RegExp(`^\\s+${field}\\s`).test(lines[i])) return i
  return -1
}

function modelOf(lines, index) {
  for (let i = index; i >= 0; i--) {
    const m = /^model (\w+) \{/.exec(lines[i])
    if (m) return m[1]
  }
  return '?'
}

// The line without its @default(...), the parentheses balanced.
function dropDefault(line) {
  const at = line.indexOf('@default(')
  if (at < 0) return line
  let depth = 0
  for (let i = at + '@default'.length; i < line.length; i++) {
    if (line[i] === '(') depth++
    else if (line[i] === ')' && --depth === 0) return (line.slice(0, at).trimEnd() + line.slice(i + 1)).trimEnd()
  }
  return line
}

// --- 3. every hekireki generator on the write-up, one at a time ---------------------------------

const GENERATORS = {
  drizzle: () => `generator Hekireki-Drizzle {\n  provider = "hekireki-drizzle"\n  output   = "generated/drizzle/schema.ts"\n}`,
  kysely: () => `generator Hekireki-Kysely {\n  provider = "hekireki-kysely"\n  output   = "generated/kysely/types.ts"\n}`,
  sqlalchemy: () => `generator Hekireki-SQLAlchemy {\n  provider = "hekireki-sqlalchemy"\n  output   = "generated/sqlalchemy"\n}`,
  django: () => `generator Hekireki-Django {\n  provider = "hekireki-django"\n  output   = "generated/django"\n}`,
  gorm: () => `generator Hekireki-GORM {\n  provider = "hekireki-gorm"\n  output   = "generated/gorm"\n  package  = "model"\n}`,
  'sea-orm': () => `generator Hekireki-SeaORM {\n  provider = "hekireki-sea-orm"\n  output   = "generated/sea-orm"\n}`,
  ecto: (s) => `generator Hekireki-Ecto {\n  provider = "hekireki-ecto"\n  output   = "generated/ecto"\n  app      = "${pascal(s.name)}"\n}`,
  activerecord: () => `generator Hekireki-ActiveRecord {\n  provider = "hekireki-activerecord"\n  output   = "generated/activerecord"\n}`,
  efcore: (s) =>
    `generator Hekireki-EFCore {\n  provider  = "hekireki-efcore"\n  output    = "generated/efcore"\n  namespace = "${pascal(s.name)}.Models"\n  context   = "${pascal(s.name)}Context"\n}`,
  eloquent: () => `generator Hekireki-Eloquent {\n  provider  = "hekireki-eloquent"\n  output    = "generated/eloquent"\n  namespace = "App\\\\Models"\n}`,
  exposed: (s) => `generator Hekireki-Exposed {\n  provider = "hekireki-exposed"\n  output   = "generated/exposed"\n  package  = "${s.name.toLowerCase().replace(/[^a-z0-9]/g, '')}.models"\n}`,
}

// One generator per run, so that one generator failing does not hide what the others write.
// hekireki-exposed writes PostgreSQL only and is left out of the others, as in apps/hekireki.
function generate(source) {
  const dir = join(layersDir, source.layer, source.name)
  const generated = join(dir, 'generated')
  rmSync(generated, { recursive: true, force: true })
  mkdirSync(generated, { recursive: true })
  const schema = readFileSync(join(dir, 'schema.prisma'), 'utf8')
  const features = /previewFeatures = \[([^\]]*)\]/.exec(schema)?.[1]
  const models = schema.slice(schema.search(/^(model|enum|view|type) /m))
  for (const orm of orms) {
    if (orm === 'exposed' && source.dialect !== 'postgresql') continue
    let block = GENERATORS[orm](source)
    if (features) block = block.replace(/\n\}$/, `\n  previewFeatures = [${features}]\n}`)
    const work = join(workRoot, source.layer, source.name, 'generate', orm)
    rmSync(work, { recursive: true, force: true })
    mkdirSync(work, { recursive: true })
    const file = join(work, 'schema.prisma')
    writeFileSync(file, `datasource db {\n  provider = "${source.dialect}"\n}\n\n${block}\n\n${models}`)
    const run = sh(['node', generateMjs, '--schema', file, '--out', `generated:${join(work, 'out')}`], { cwd: work })
    if (run.status === 0 && existsSync(join(work, 'out', orm))) cpSync(join(work, 'out', orm), join(generated, orm), { recursive: true })
    else writeFileSync(join(generated, `${orm}.error`), tail(run, 3000) + '\n')
  }
}

// --- 4. the generated code, run against tables made from source.sql -----------------------------

// apps/hekireki/compat/run.mjs over this layer: every ORM's harness loads the generated code and
// runs it against a database made from source.sql (not from `prisma db push`).
function check(source) {
  const layer = join(layersDir, source.layer)
  const run = spawnSync('node', [join(compat, 'run.mjs'), '--schemas', layer, '--results', join(layer, 'results'), '--orm', orms.join(','), source.name], { stdio: 'inherit' })
  if (run.status !== 0) throw new Error(`run.mjs exited with ${run.status}`)
  const infoFile = join(layer, source.name, 'source.json')
  const info = JSON.parse(readFileSync(infoFile, 'utf8'))
  info.hekireki = hekirekiVersion ?? hekirekiRequire('hekireki/package.json').version
  info.checked_at = new Date().toISOString()
  writeFileSync(infoFile, JSON.stringify(info, null, 2) + '\n')
}

function recordFailure(source, e) {
  const dir = join(layersDir, source.layer, source.name)
  mkdirSync(dir, { recursive: true })
  const infoFile = join(dir, 'source.json')
  const info = existsSync(infoFile) ? JSON.parse(readFileSync(infoFile, 'utf8')) : { source: source.key, about: source.about, dialect: source.dialect }
  info.error = String(e?.message ?? e).slice(0, 3000)
  info.checked_at = new Date().toISOString()
  writeFileSync(infoFile, JSON.stringify(info, null, 2) + '\n')
}

// pull/<orm>/results/README.md: the matrix of the layer's sources, and the ones that could not be
// pulled.
function writeIndex(layer) {
  const layerDir = join(layersDir, layer)
  const resultsDir = join(layerDir, 'results')
  mkdirSync(resultsDir, { recursive: true })
  const all = readdirSync(resultsDir)
    .filter((f) => f.endsWith('.json'))
    .map((f) => ({ name: f.replace(/\.json$/, ''), ...JSON.parse(readFileSync(join(resultsDir, f), 'utf8')) }))
  const info = (name) => {
    const file = join(layerDir, name, 'source.json')
    return existsSync(file) ? JSON.parse(readFileSync(file, 'utf8')) : {}
  }
  const notPulled = readdirSync(layerDir)
    .filter((name) => existsSync(join(layerDir, name, 'source.json')) && info(name).error)
    .map((name) => ({ name, ...info(name) }))
  const index = renderIndex(all, {
    title: `${ORM_NAMES[layer]}: what hekireki generates from the write-ups`,
    intro: `Each project's tables as ${ORM_NAMES[layer]} created them (\`<project>/source.sql\`), written up by \`prisma db pull\` (\`<project>/schema.prisma\`); every ORM hekireki generates from that write-up runs against those tables. Each cell: the runtime checks that passed (insert, read back, update, relations, delete, per model) and whether the tables the ORM creates on its own are the written-up ones. Generated by \`node pull/pull.mjs\`; do not edit.`,
    columns: [
      { title: 'Models', value: (s) => String(info(s.name).models ?? '') },
      { title: 'Prisma diff', value: (s) => (existsSync(join(layerDir, s.name, 'prisma-diff.sql')) ? `[yes](../${s.name}/prisma-diff.sql)` : 'none') },
      { title: 'hekireki', value: (s) => info(s.name).hekireki ?? '' },
    ],
  })
  const failures = notPulled.length ? `\n## Not pulled\n\n${notPulled.map((f) => `- **${layer}/${f.name}**: ${f.error.split('\n')[0].slice(0, 300)}`).join('\n')}\n` : ''
  writeFileSync(join(resultsDir, 'README.md'), index + failures)
}

// --hekireki: another release, installed beside the work directory with the Prisma this
// repository uses, and its own copy of generate.mjs, which resolves hekireki from there. The
// write-ups are copied, so that generation, results and checked_at all stay out of pull/.
function useHekireki(version, chosen) {
  if (version === 'latest') version = sh(['npm', 'view', 'hekireki', 'version']).stdout.trim()
  hekirekiVersion = version
  const dir = join(workRoot, `hekireki-${version}`)
  if (!existsSync(join(dir, 'node_modules', 'hekireki', 'package.json'))) {
    mkdirSync(dir, { recursive: true })
    writeFileSync(join(dir, 'package.json'), '{ "private": true, "type": "module" }\n')
    const prismaVersion = hekirekiRequire('prisma/package.json').version
    const run = sh(
      ['npm', 'install', '--no-audit', '--no-fund', '--legacy-peer-deps', '--loglevel=error', `hekireki@${version}`, `prisma@${prismaVersion}`, `@prisma/client@${prismaVersion}`],
      { cwd: dir },
    )
    if (run.status !== 0) throw new Error(`npm install hekireki@${version}: ${tail(run)}`)
  }
  cpSync(generateMjs, join(dir, 'generate.mjs'))
  generateMjs = join(dir, 'generate.mjs')
  layersDir = join(dir, 'layers')
  for (const source of chosen) {
    const from = join(here, source.layer, source.name)
    if (!existsSync(join(from, 'schema.prisma'))) throw new Error(`${source.key} has not been pulled yet: run node pull/pull.mjs ${source.key} first`)
    const to = join(layersDir, source.layer, source.name)
    rmSync(to, { recursive: true, force: true })
    mkdirSync(to, { recursive: true })
    for (const f of ['schema.prisma', 'source.sql', 'source.json', 'prisma-diff.sql']) if (existsSync(join(from, f))) cpSync(join(from, f), join(to, f))
  }
}

// --- databases ----------------------------------------------------------------------------------

async function freshDatabase(dialect, work, name) {
  const db = name.toLowerCase().replace(/[^a-z0-9_]/g, '_').slice(0, 60)
  if (dialect === 'postgresql') {
    up(PG)
    psql('postgres', `DROP DATABASE IF EXISTS "${db}" WITH (FORCE)`)
    psql('postgres', `CREATE DATABASE "${db}"`)
    psql(db, PG_EXTENSIONS.map((e) => `CREATE EXTENSION IF NOT EXISTS "${e}";`).join(' '))
    return `postgresql://postgres:compat@127.0.0.1:${PG.port}/${db}`
  }
  if (dialect === 'mysql') {
    up(MYSQL)
    mysql(`DROP DATABASE IF EXISTS \`${db}\`; CREATE DATABASE \`${db}\``)
    return `mysql://root:compat@127.0.0.1:${MYSQL.port}/${db}`
  }
  const file = join(work, `${db}.db`)
  rmSync(file, { force: true })
  return `file:${file}`
}

function up({ name, image, port }) {
  sh(['docker', 'network', 'create', NETWORK])
  const running = sh(['docker', 'ps', '--format', '{{.Image}}', '-f', `name=^${name}$`]).stdout.trim()
  if (running !== image) {
    sh(['docker', 'rm', '-f', '-v', name])
    const env = image.startsWith('mysql') ? ['-e', 'MYSQL_ROOT_PASSWORD=compat'] : ['-e', 'POSTGRES_PASSWORD=compat']
    const started = sh(['docker', 'run', '-d', '--rm', '--name', name, '--network', NETWORK, ...env, '-p', `${port}:${image.startsWith('mysql') ? 3306 : 5432}`, image])
    if (started.status !== 0) throw new Error(`docker run ${image}: ${started.stderr}`)
  }
  sh(['docker', 'network', 'connect', NETWORK, name])
  for (let i = 0; i < 90; i++) {
    const ok = image.startsWith('mysql') ? sh(['docker', 'exec', name, 'mysql', '-uroot', '-pcompat', '-e', 'SELECT 1']) : sh(['docker', 'exec', name, 'pg_isready', '-U', 'postgres'])
    if (ok.status === 0) return
    spawnSync('sleep', ['1'])
  }
  throw new Error(`${image} did not come up`)
}

function psql(db, sql) {
  const run = sh(['docker', 'exec', '-i', PG.name, 'psql', '-q', '-v', 'ON_ERROR_STOP=1', '-U', 'postgres', '-d', db], { input: sql })
  if (run.status !== 0) throw new Error(`psql: ${tail(run)}`)
  return run
}

function mysql(sql, db) {
  const run = sh(['docker', 'exec', '-i', MYSQL.name, 'mysql', '-uroot', '-pcompat', ...(db ? [db] : [])], { input: sql })
  if (run.status !== 0) throw new Error(`mysql: ${tail(run)}`)
  return run
}

function dbName(url) {
  return new URL(url).pathname.slice(1)
}

// A script as a whole, the way its own tool would run it (psql, the mysql client, SQLite).
function execSql(dialect, url, sql) {
  if (dialect === 'postgresql') psql(dbName(url), sql.replace(/^\\(un)?restrict .*$/gm, ''))
  else if (dialect === 'mysql') mysql(sql, dbName(url))
  else {
    const db = sqlite(url.slice('file:'.length))
    try {
      db.exec(sql.replace(/^--> statement-breakpoint$/gm, ''))
    } finally {
      db.close()
    }
  }
}

function countTables(dialect, url) {
  if (dialect === 'postgresql')
    return Number(sh(['docker', 'exec', PG.name, 'psql', '-tAq', '-U', 'postgres', '-d', dbName(url), '-c', "SELECT count(*) FROM information_schema.tables WHERE table_schema = 'public' AND table_type = 'BASE TABLE'"]).stdout.trim())
  if (dialect === 'mysql') return Number(mysql(`SELECT count(*) FROM information_schema.tables WHERE table_schema = '${dbName(url)}' AND table_type = 'BASE TABLE'`).stdout.trim().split('\n').pop())
  const db = sqlite(url.slice('file:'.length))
  try {
    return db.prepare("SELECT count(*) AS n FROM sqlite_master WHERE type = 'table' AND name NOT LIKE 'sqlite_%'").get().n
  } finally {
    db.close()
  }
}

// The tables as SQL, without what changes from one run to the next (versions, dates).
function dump(dialect, url) {
  if (dialect === 'postgresql') {
    const run = sh(['docker', 'exec', PG.name, 'pg_dump', '-U', 'postgres', '--schema-only', '--no-owner', '--no-privileges', dbName(url)])
    if (run.status !== 0) throw new Error(`pg_dump: ${tail(run)}`)
    return run.stdout
      .replace(/^\\(un)?restrict .*\n/gm, '')
      .replace(/^-- Dumped (from database|by pg_dump) version .*\n/gm, '')
      .replace(/\n{3,}/g, '\n\n')
  }
  if (dialect === 'mysql') {
    const run = sh(['docker', 'exec', MYSQL.name, 'mysqldump', '-uroot', '-pcompat', '--no-data', '--skip-comments', '--compact', '--skip-dump-date', '--triggers', '--routines', dbName(url)])
    if (run.status !== 0) throw new Error(`mysqldump: ${tail(run)}`)
    return run.stdout.replace(/ AUTO_INCREMENT=\d+/g, '')
  }
  const db = sqlite(url.slice('file:'.length))
  try {
    // A virtual table (FTS5, ...) creates its shadow tables itself; they are left out.
    const shadow = new Set(db.prepare("SELECT name FROM pragma_table_list WHERE type = 'shadow'").all().map((r) => r.name))
    const rows = db.prepare("SELECT name, sql FROM sqlite_master WHERE sql IS NOT NULL AND name NOT LIKE 'sqlite_%' ORDER BY rowid").all()
    return rows.filter((r) => !shadow.has(r.name)).map((r) => `${r.sql};`).join('\n\n') + '\n'
  } finally {
    db.close()
  }
}

function prisma(cwd, url, command) {
  const config = join(cwd, 'prisma.config.mjs')
  if (!existsSync(config)) writeFileSync(config, `export default { schema: './schema.prisma', datasource: { url: process.env.DATABASE_URL } }\n`)
  return sh(['node', prismaCli, ...command, '--config', config], {
    cwd,
    env: { ...process.env, DATABASE_URL: url, CHECKPOINT_DISABLE: '1', PRISMA_HIDE_UPDATE_MESSAGE: '1', NO_COLOR: '1', FORCE_COLOR: '0' },
  })
}

// --- docker -------------------------------------------------------------------------------------

// apps/hekireki/compat's image of an ORM; COMPAT_DOCKER_BUILD_ARGS as in run.mjs (tools/cloud).
function dockerBuild(orm) {
  const image = `hekireki-compat-${orm}`
  const extra = (process.env.COMPAT_DOCKER_BUILD_ARGS ?? '').split(/\s+/).filter(Boolean)
  const built = sh(['docker', 'build', '-q', ...extra, '-t', image, join(compat, orm)])
  if (built.status !== 0) throw new Error(`docker build ${orm}: ${tail(built)}`)
  return image
}

// Runs a command in an image on the databases' network, with the stage directory as /work (sent
// on stdin, since the daemon may not see this machine's paths); `{url}` in the command is the
// database's URL as seen from the container. A SQLite database comes back out.
function dockerRun(image, stage, command, url) {
  let inner = url.replace(`127.0.0.1:${PG.port}`, `${PG.name}:5432`).replace(`127.0.0.1:${MYSQL.port}`, `${MYSQL.name}:3306`)
  const sqlite = url.startsWith('file:') ? url.slice('file:'.length) : null
  if (sqlite) {
    if (existsSync(sqlite)) cpSync(sqlite, join(stage, 'db.sqlite'))
    inner = 'sqlite3:/work/db.sqlite'
  }
  const name = `pull-${process.pid}`
  sh(['docker', 'rm', '-f', '-v', name])
  const script = `mkdir -p /work && tar -C /work -xf - && cd /work && exec ${command.map((a) => `'${a.replace('{url}', inner)}'`).join(' ')}`
  const run = sh(['bash', '-c', `tar -C '${stage}' -cf - . | docker run -i --name ${name} --network ${NETWORK} ${image} sh -c "$0"`, script])
  if (sqlite) sh(['docker', 'cp', `${name}:/work/db.sqlite`, sqlite])
  sh(['docker', 'rm', '-f', '-v', name])
  return run
}

// --- helpers ------------------------------------------------------------------------------------

function download(url, file) {
  const run = sh(['curl', '-fsSL', '-o', file, url])
  if (run.status !== 0) throw new Error(`download ${url}: ${tail(run)}`)
}

function pascal(name) {
  const p = name.replace(/(^|[^a-zA-Z0-9]+)([a-zA-Z0-9])/g, (_, __, c) => c.toUpperCase()).replace(/[^a-zA-Z0-9]/g, '')
  return /^[0-9]/.test(p) ? `App${p}` : p
}

// The start and the end of what a command printed: a Ruby or Python error says what went wrong
// first and where last.
function tail(run, n = 2000) {
  const text = ((run.stderr ?? '') + (run.stdout ?? '')).trim()
  return text.length <= n ? text : `${text.slice(0, n / 2)}\n...\n${text.slice(-n / 2)}`
}

function sh([cmd, ...rest], options = {}) {
  const run = spawnSync(cmd, rest, { encoding: 'utf8', maxBuffer: 256 * 1024 * 1024, timeout: 20 * 60_000, killSignal: 'SIGKILL', ...options })
  if (run.signal === 'SIGKILL') return { status: 1, stdout: run.stdout ?? '', stderr: `timed out after 20 minutes\n${run.stderr ?? ''}` }
  if (run.error) return { status: 1, stdout: '', stderr: String(run.error) }
  return run
}

await main()
