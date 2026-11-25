#!/usr/bin/env node
// Writes up, in Prisma, the SQL that other ORMs generate, and holds the SQL of every ORM hekireki
// generates from the write-up to it. A well-known project's own ORM creates its tables from the
// project's definitions (Django models, a Drizzle schema, a Rails db/schema.rb, ...), the tables
// are kept as SQL, and `prisma db pull` writes them up as a Prisma schema. Then hekireki generates
// every ORM's code from that schema, and apps/hekireki/compat runs the code against the tables the
// project's own ORM created, and has each ORM create tables of its own: whether the code is valid
// for each ORM and works, and whether the SQL each ORM creates from it is the project's own SQL.
//
//   node pull/pull.mjs <orm>/<project> [...]     django/wagtail, drizzle/northwind, activerecord/mastodon
//   node pull/pull.mjs --random 3                three: never pulled first, then the oldest
//   node pull/pull.mjs --list                    every source, with when it was last checked
//   node pull/pull.mjs --pull-only <...>         only the write-up: source.sql and schema.prisma
//   node pull/pull.mjs --orm django,gorm <...>   check only these of hekireki's ORMs
//   node pull/pull.mjs --compare-only <...>      hold the SQL already in sql/ to source.sql again,
//                                                and write the READMEs (no database)
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
//   source.sql       the SQL of the tables the project's own ORM created (pg_dump, mysqldump, or
//                    SQLite's own schema)
//   schema.prisma    `prisma db pull` of those tables: the write-up
//   sql/<orm>.sql    the SQL of the tables each ORM creates on its own from the write-up, dumped the
//                    way source.sql is: Prisma's (`prisma db push`), and every ORM hekireki
//                    generates that creates tables; sql/<orm>.diff, the whole diff against
//                    source.sql, beside each one that differs (lib/sql.mjs)
//   README.md        each sql/<orm>.sql against source.sql
//   source.json      where it came from, the versions, when it was checked, the comparison
// and generated/ (hekireki's output, not committed; a generator that fails leaves
// generated/<orm>.error). The runtime checks go to pull/<orm>/results/, with the matrix of the
// layer's comparisons.
import { spawnSync } from 'node:child_process'
import { cpSync, existsSync, mkdirSync, readdirSync, readFileSync, rmSync, writeFileSync } from 'node:fs'
import { createRequire } from 'node:module'
import { homedir } from 'node:os'
import { dirname, join, relative, resolve } from 'node:path'
import { fileURLToPath } from 'node:url'
import { compareSql } from './lib/sql.mjs'

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
let compareOnly = false
const wanted = []
for (let i = 0; i < args.length; i++) {
  if (args[i] === '--orm') orms = args[++i].split(',')
  else if (args[i] === '--random') random = Number(args[++i])
  else if (args[i] === '--list') list = true
  else if (args[i] === '--pull-only') pullOnly = true
  else if (args[i] === '--hekireki') hekirekiVersion = args[++i]
  else if (args[i] === '--compare-only') compareOnly = true
  else wanted.push(args[i])
}

// --- sources ------------------------------------------------------------------------------------

function layerConfig(layer) {
  const file = join(here, layer, 'sources.json')
  return existsSync(file) ? JSON.parse(readFileSync(file, 'utf8')) : { sources: {}, skip: {} }
}

function allSources() {
  const sources = []
  // `sources.json` of the django and drizzle layers may hold, per example, what the project's ORM
  // needs and the example's own files do not say: `before_sql` / `after_sql` (a sequence or a
  // function the app creates outside its models).
  const extra = (layer, name) => layerConfig(layer).sources?.[name] ?? {}
  for (const name of readdirSync(djangoerExamples).sort()) {
    if (!existsSync(join(djangoerExamples, name, 'pyproject.toml'))) continue
    sources.push({ layer: 'django', name, dialect: 'postgresql', about: `the models of [apps/djangoer/examples/${name}](../../../apps/djangoer/examples/${name})`, ...extra('django', name) })
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
      ...extra('drizzle', name),
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
  if (compareOnly) {
    for (const source of chosen) {
      const infoFile = join(layersDir, source.layer, source.name, 'source.json')
      if (!existsSync(infoFile)) continue
      const info = JSON.parse(readFileSync(infoFile, 'utf8'))
      compareAll(source, info, Object.keys(CREATES).filter((o) => existsSync(join(layersDir, source.layer, source.name, 'sql', `${o}.sql`)) || info.sql?.[o]))
      writeFileSync(infoFile, JSON.stringify(info, null, 2) + '\n')
      writeReadme(source, info)
    }
    for (const layer of new Set(chosen.map((s) => s.layer))) writeIndex(layer)
    return
  }
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

  // The tables Prisma itself creates from the write-up, dumped the same way: sql/prisma.sql.
  rmSync(join(dir, 'prisma-diff.sql'), { force: true })
  const sqlDir = join(dir, 'sql')
  rmSync(sqlDir, { recursive: true, force: true })
  mkdirSync(sqlDir, { recursive: true })
  const pushUrl = await freshDatabase(source.dialect, work, `pull_${source.layer}_${source.name}_prisma`)
  // What the app creates outside its schema and a default calls (Mastodon's timestamp_id()).
  if (source.before_sql) execSql(source.dialect, pushUrl, source.before_sql)
  const push = prisma(work, pushUrl, ['db', 'push', '--schema', join(dir, 'schema.prisma')])
  writeFileSync(join(sqlDir, 'prisma.sql'), push.status === 0 ? dump(source.dialect, pushUrl) : `-- prisma db push failed:\n-- ${tail(push).replaceAll('\n', '\n-- ')}\n`)
  dropDatabase(source.dialect, pushUrl)
  dropDatabase(source.dialect, url)

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
  compareAll(source, info, ['prisma'])
  writeFileSync(join(dir, 'source.json'), JSON.stringify(info, null, 2) + '\n')
  writeReadme(source, info)
  console.log(`${source.key}: ${created} tables, ${info.models} models`)
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
// Changed from what \`prisma db pull\` wrote, which \`prisma validate\` or \`prisma db push\` refuses:
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
// - a default on an Unsupported(...) field, which only dbgenerated() may be, is dropped;
// - a Gin index over column types Prisma does not take together is dropped;
// - on SQLite, the parentheses around an expression default (dbgenerated("(unixepoch())")) are put back,
//   and the sqlite_autoindex_* name of a unique index, which SQLite reserves, is dropped.
// Each change is listed in the schema's header and in source.json. Anything else is a schema
// Prisma cannot read, and the pull stops there.
function validate(source, pulledModels, work, url) {
  const lines = pulledModels.split('\n')
  const features = []
  const fixes = []
  // SQLite takes an expression as a default only in parentheses (`DEFAULT (unixepoch())`).
  // `prisma db pull` writes the expression without them, dbgenerated("unixepoch()"), and
  // `prisma db push` then writes `DEFAULT unixepoch()`, a syntax error: the parentheses go back.
  if (source.dialect === 'sqlite') {
    for (let i = 0; i < lines.length; i++) {
      const m = /@default\(dbgenerated\("((?:[^"\\]|\\.)*)"\)\)/.exec(lines[i])
      if (!m || m[1].startsWith('(') || /^(-?\d+(\.\d+)?|'[^']*'|NULL|TRUE|FALSE|CURRENT_(TIME|DATE|TIMESTAMP))$/i.test(m[1])) continue
      lines[i] = lines[i].replace(m[0], `@default(dbgenerated("(${m[1]})"))`)
      fixes.push(`${modelOf(lines, i)}.${lines[i].trim().split(/\s+/)[0]}: dbgenerated("${m[1]}") -> dbgenerated("(${m[1]})") (an SQLite expression default has to be in parentheses, which \`prisma db push\` leaves out)`)
    }
    // The index SQLite makes for a UNIQUE column is named sqlite_autoindex_<table>_<n>, and
    // `prisma db pull` writes that name as the map: of the @unique. SQLite reserves the prefix, so
    // no ORM can create an index of that name: the name goes, and the unique stays.
    for (let i = 0; i < lines.length; i++) {
      const name = /map:\s*"(sqlite_autoindex_[^"]*)"/.exec(lines[i])?.[1]
      if (!name) continue
      const before = lines[i]
      lines[i] = lines[i].replace(/\(map:\s*"sqlite_autoindex_[^"]*"\)/, '').replace(/,\s*map:\s*"sqlite_autoindex_[^"]*"/, '')
      if (lines[i] !== before) fixes.push(`${modelOf(lines, i)}: the name "${name}" of ${/@@/.test(before) ? 'an @@index / @@unique' : `the unique of ${before.trim().split(/\s+/)[0]}`} dropped (SQLite reserves the sqlite_autoindex_ prefix, so no ORM can create an index called that)`)
    }
  }
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
    const gin = new Set()
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
      } else if (/The Gin index type does not support the type of the field/.test(message) && /^\s*@@index\(/.test(lines[index] ?? '')) {
        gin.add(index)
      }
    }
    // A Gin index over text and integer columns together (Django's GinIndex with btree_gin):
    // Prisma writes it up and then refuses it, so the index goes.
    for (const index of [...gin].sort((a, b) => a - b)) {
      const name = /map:\s*"([^"]+)"/.exec(lines[index])?.[1] ?? lines[index].trim()
      fixes.push(`${modelOf(lines, index)}: @@index "${name}" dropped (a Gin index over these column types, which Prisma writes up and then refuses)`)
      changed = true
    }
    for (const index of [...gin].sort((a, b) => b - a)) lines.splice(index, 1)
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
// runs it against a database made from source.sql (not from `prisma db push`), and each ORM that
// creates tables creates them on its own, which --ddl-sql keeps as sql/<orm>.sql.
function check(source) {
  const layer = join(layersDir, source.layer)
  const dir = join(layer, source.name)
  const ddlSql = join(workRoot, source.layer, source.name, 'ddl-sql')
  rmSync(ddlSql, { recursive: true, force: true })
  const run = spawnSync(
    'node',
    [join(compat, 'run.mjs'), '--schemas', layer, '--results', join(layer, 'results'), '--ddl-sql', ddlSql, '--orm', orms.join(','), source.name],
    { stdio: 'inherit' },
  )
  if (run.status !== 0) throw new Error(`run.mjs exited with ${run.status}`)
  const sqlDir = join(dir, 'sql')
  mkdirSync(sqlDir, { recursive: true })
  // This run's ORMs only: a run with --orm keeps what an earlier one left for the others.
  for (const orm of orms) for (const ext of ['sql', 'diff']) rmSync(join(sqlDir, `${orm}.${ext}`), { force: true })
  const dumped = join(ddlSql, source.name)
  for (const f of existsSync(dumped) ? readdirSync(dumped) : []) cpSync(join(dumped, f), join(sqlDir, f))
  const infoFile = join(dir, 'source.json')
  const info = JSON.parse(readFileSync(infoFile, 'utf8'))
  info.hekireki = hekirekiVersion ?? hekirekiRequire('hekireki/package.json').version
  info.checked_at = new Date().toISOString()
  compareAll(source, info, ['prisma', ...orms])
  writeFileSync(infoFile, JSON.stringify(info, null, 2) + '\n')
  writeReadme(source, info)
}

// The ORMs that create tables of their own, how, and what the others do instead.
const CREATES = {
  prisma: { from: 'Prisma: schema.prisma', how: '`prisma db push`' },
  drizzle: { from: 'hekireki-drizzle', how: '`drizzle-kit export`' },
  kysely: { from: 'hekireki-kysely', none: 'writes types only' },
  sqlalchemy: { from: 'hekireki-sqlalchemy', how: '`metadata.create_all()`' },
  django: { from: 'hekireki-django', how: '`makemigrations`, `migrate`' },
  gorm: { from: 'hekireki-gorm', how: '`AutoMigrate`' },
  'sea-orm': { from: 'hekireki-sea-orm', how: '`Schema::create_table_from_entity`' },
  ecto: { from: 'hekireki-ecto', none: 'writes schemas only (no migrations)' },
  activerecord: { from: 'hekireki-activerecord', none: 'writes models only (no migrations)' },
  eloquent: { from: 'hekireki-eloquent', none: 'writes models only (no migrations)' },
  efcore: { from: 'hekireki-efcore', how: '`EnsureCreated()`' },
  exposed: { from: 'hekireki-exposed', how: '`SchemaUtils.create`' },
}

// Every sql/<orm>.sql against source.sql: info.sql[orm] is `identical`, `names` (the same but for
// constraint and index names, and order), `differs` (with how many statements each side alone
// has), or why there is no SQL; sql/<orm>.diff beside every one that is not identical.
function compareAll(source, info, which = Object.keys(CREATES)) {
  const dir = join(layersDir, source.layer, source.name)
  const sqlDir = join(dir, 'sql')
  const original = readFileSync(join(dir, 'source.sql'), 'utf8')
  const resultsFile = join(layersDir, source.layer, 'results', `${source.name}.json`)
  const results = existsSync(resultsFile) ? (JSON.parse(readFileSync(resultsFile, 'utf8')).results ?? []) : []
  info.sql = Object.fromEntries(Object.entries(info.sql ?? {}).filter(([orm]) => !which.includes(orm)))
  for (const orm of Object.keys(CREATES).filter((o) => which.includes(o))) {
    const file = join(sqlDir, `${orm}.sql`)
    const diffFile = join(sqlDir, `${orm}.diff`)
    rmSync(diffFile, { force: true })
    const r = results.find((x) => x.orm === orm)
    if (!existsSync(file)) {
      if (CREATES[orm].none) info.sql[orm] = { status: 'none' }
      else if (orm === 'exposed' && source.dialect !== 'postgresql') info.sql[orm] = { status: 'not-generated' }
      else if (r?.generatorError || r?.notGenerated) info.sql[orm] = { status: 'not-generated' }
      else if (r?.ddl?.error) info.sql[orm] = { status: 'error', error: errorLine(r.ddl.error) }
      else if (r?.fatal) info.sql[orm] = { status: 'error', error: 'the harness did not run' }
      else if (r || orm === 'prisma') info.sql[orm] = { status: 'missing' }
      continue
    }
    const sql = readFileSync(file, 'utf8')
    if (!sql.trim()) {
      info.sql[orm] = { status: 'error', error: 'no tables' }
      continue
    }
    if (/^-- (prisma db push failed|could not dump)/.test(sql)) {
      info.sql[orm] = { status: 'error', error: errorLine(sql) }
      continue
    }
    const c = compareSql(original, sql, source.dialect)
    info.sql[orm] = c.identical
      ? { status: 'identical' }
      : c.sameButNames
        ? { status: 'names' }
        : { status: 'differs', tables: c.tables, onlySource: c.onlySource, onlyHere: c.onlyOther }
    // The whole diff where the SQL says something else; not for names alone.
    if (!c.identical && !c.sameButNames) writeFileSync(diffFile, c.diff.replace(/^\+\+\+ other\.sql$/m, `+++ sql/${orm}.sql`))
    // Against what Prisma itself creates from the same write-up: a difference there is the ORM's
    // or hekireki's, one only against source.sql is what the write-up does not say.
    const prismaFile = join(sqlDir, 'prisma.sql')
    if (orm !== 'prisma' && existsSync(prismaFile) && !/^-- prisma db push failed/.test(readFileSync(prismaFile, 'utf8'))) {
      const p = compareSql(readFileSync(prismaFile, 'utf8'), sql, source.dialect)
      const n = p.tables.length + p.onlySource.length + p.onlyOther.length
      info.sql[orm].prisma = p.identical ? 'identical' : p.sameButNames ? 'names' : `differs:${n}`
    }
  }
}

// The line of an error that says what went wrong: a Python traceback's last exception line, else
// the first line that says ERROR or Error, else the first line; tools' own chatter left out.
function errorLine(text) {
  const lines = text
    .split('\n')
    .map((l) => l.replace(/^--\s?/, '').trim())
    .filter((l) => l && !/^(Loaded Prisma config|Prisma schema loaded|Picked up JAVA_TOOL_OPTIONS|prisma db push failed:?)/i.test(l))
  const traceback = text.includes('Traceback') ? [...lines].reverse().find((l) => /^[A-Za-z_][\w.]*: \S/.test(l) && !l.startsWith('File ')) : null
  const said = lines.find((l) => /\b(ERROR|Error|Exception)\b/.test(l))
  return (traceback ?? said ?? lines[0] ?? '').slice(0, 300)
}

const MARK = { identical: '= identical', names: '≈ the same but for names', differs: '≠', error: '❌ could not create its tables', none: '—', 'not-generated': 'not generated', missing: '?' }

function sqlCell(c, short = false) {
  if (!c) return ''
  if (c.status === 'differs') {
    const tables = c.tables ?? []
    const changed = tables.filter((t) => t.lines).length
    const onlySource = tables.filter((t) => t.only === 'source').length
    const onlyHere = tables.filter((t) => t.only === 'here').length
    const n = tables.length + c.onlySource.length + c.onlyHere.length
    if (short) return `≠ ${n}`
    const plural = (k, one, many = `${one}s`) => `${k} ${k === 1 ? one : many}`
    return `≠ ${[
      changed ? `${plural(changed, 'table')} written differently` : null,
      onlySource ? `${plural(onlySource, 'table')} only in source.sql` : null,
      onlyHere ? `${plural(onlyHere, 'table')} only here` : null,
      c.onlySource.length ? `${plural(c.onlySource.length, 'other statement')} only in source.sql` : null,
      c.onlyHere.length ? `${plural(c.onlyHere.length, 'other statement')} only here` : null,
    ]
      .filter(Boolean)
      .join(', ')}`
  }
  if (short) return { identical: '=', names: '≈', error: '❌', none: '—', 'not-generated': 'not generated', missing: '?' }[c.status]
  return MARK[c.status]
}

// pull/<orm>/<project>/README.md: where the tables came from and every ORM's SQL against them.
function writeReadme(source, info) {
  const dir = join(layersDir, source.layer, source.name)
  const about = info.about ?? source.about
  const rows = Object.entries(CREATES).map(([orm, c]) => {
    const r = info.sql?.[orm]
    const file = existsSync(join(dir, 'sql', `${orm}.sql`)) ? `[${orm}.sql](sql/${orm}.sql)` : '—'
    const against = !r ? 'not checked yet' : r.status === 'none' ? c.none : `${sqlCell(r)}${existsSync(join(dir, 'sql', `${orm}.diff`)) ? ` ([diff](sql/${orm}.diff))` : ''}${r.error ? `: ${r.error.replaceAll('|', '\\|')}` : ''}`
    const vsPrisma = orm === 'prisma' ? '' : !r?.prisma ? '—' : r.prisma === 'identical' ? '=' : r.prisma === 'names' ? '≈' : `≠ ${r.prisma.split(':')[1]}`
    return `| ${c.from} | ${c.how ?? '—'} | ${file} | ${against} | ${vsPrisma} |`
  })
  // What differs, as a diff: the tables both sides have, column by column, then the statements
  // only one side has.
  const cut = (x) => (x.length > 300 ? `${x.slice(0, 300)} …` : x)
  const details = Object.entries(info.sql ?? {})
    .filter(([, r]) => r.status === 'differs')
    .flatMap(([orm, r]) => {
      const body = [
        ...(r.tables ?? []).flatMap((t) =>
          t.only === 'source' ? [`- ${t.name} (the table is only in source.sql)`] : t.only === 'here' ? [`+ ${t.name} (the table is only in sql/${orm}.sql)`] : [`@@ ${t.name}`, ...t.lines.map(cut)],
        ),
        ...(r.onlySource.length ? ['@@ only in source.sql', ...r.onlySource.map((x) => `- ${cut(x)}`)] : []),
        ...(r.onlyHere.length ? [`@@ only in sql/${orm}.sql`, ...r.onlyHere.map((x) => `+ ${cut(x)}`)] : []),
      ]
      const shown = body.slice(0, 120)
      return [
        '',
        `<details><summary>${CREATES[orm].from}: ${sqlCell(r).replace(/^≠ /, '')}</summary>`,
        '',
        '```diff',
        ...shown,
        ...(body.length > shown.length ? [`@@ … ${body.length - shown.length} more lines in sql/${orm}.diff`] : []),
        '```',
        '',
        '</details>',
      ]
    })
  const lines = [
    `# ${source.key}`,
    '',
    `The tables ${ORM_NAMES[source.layer]} creates from ${about}, as SQL: [source.sql](source.sql). \`prisma db pull\` writes them up as [schema.prisma](schema.prisma)${info.fixes?.length ? ` (with ${info.fixes.length} change${info.fixes.length === 1 ? '' : 's'} so that Prisma accepts it, listed in its header)` : ''}, and from that write-up Prisma and every ORM hekireki generates create tables of their own; their SQL is in [sql/](sql), dumped the way source.sql is, and held to it here.`,
    '',
    `\`=\` identical · \`≈\` the same but for constraint and index names, and the order of columns and statements · \`≠\` tables written differently, or statements one side has and the other does not. The last column holds each ORM's SQL to what Prisma itself creates from the same write-up: a difference there is the ORM's or hekireki's, one only against source.sql is something the write-up does not say.`,
    '',
    '| From the write-up | Creates the tables with | SQL | Against source.sql | Against prisma.sql |',
    '| --- | --- | --- | --- | --- |',
    ...rows,
    '',
    `${info.models ?? '?'} models, ${info.dialect}. ${info.hekireki ? `hekireki ${info.hekireki}, ` : ''}Prisma ${info.prisma ?? '?'}. The same code run against source.sql's tables: [results/${source.name}.md](../results/${source.name}.md).`,
    ...details,
    '',
    'Written by `node pull/pull.mjs`; do not edit.',
    '',
  ]
  writeFileSync(join(dir, 'README.md'), lines.join('\n'))
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
      { title: 'hekireki', value: (s) => info(s.name).hekireki ?? '' },
    ],
  })
  const failures = notPulled.length ? `\n## Not pulled\n\n${notPulled.map((f) => `- **${layer}/${f.name}**: ${f.error.split('\n')[0].slice(0, 300)}`).join('\n')}\n` : ''
  // The SQL each ORM creates from the write-up, against source.sql (pull/<orm>/<project>/README.md).
  const creating = Object.keys(CREATES).filter((o) => !CREATES[o].none)
  const compared = readdirSync(layerDir)
    .filter((name) => existsSync(join(layerDir, name, 'source.json')) && info(name).sql)
    .sort()
  const vsPrisma = (c) => (!c?.prisma ? (c ? sqlCell(c, true) : '') : c.prisma === 'identical' ? '=' : c.prisma === 'names' ? '≈' : `≠ ${c.prisma.split(':')[1]}`)
  const ormColumns = creating.filter((o) => o !== 'prisma')
  const sqlMatrix = compared.length
    ? [
        '',
        '## The SQL each ORM creates from the write-up',
        '',
        '`=` identical · `≈` the same but for constraint and index names, and order · `≠ n` n tables (or other statements) that differ · `❌` could not create its tables. Each project\'s README shows what differs.',
        '',
        '### Against source.sql, the project\'s own SQL',
        '',
        `| Project | ${creating.join(' | ')} |`,
        `| --- | ${creating.map(() => '---').join(' | ')} |`,
        ...compared.map((name) => `| [${name}](../${name}/README.md) | ${creating.map((o) => sqlCell(info(name).sql[o], true) ?? '').join(' | ')} |`),
        '',
        '### Against prisma.sql, what Prisma itself creates from the same write-up',
        '',
        'What the write-up cannot say (Django\'s identity columns become serial ones, check constraints go) is the same for every ORM, and is left out here: a difference here is the ORM\'s or hekireki\'s.',
        '',
        `| Project | ${ormColumns.join(' | ')} |`,
        `| --- | ${ormColumns.map(() => '---').join(' | ')} |`,
        ...compared.map((name) => `| [${name}](../${name}/README.md) | ${ormColumns.map((o) => vsPrisma(info(name).sql[o]) ?? '').join(' | ')} |`),
        '',
      ].join('\n')
    : ''
  writeFileSync(join(resultsDir, 'README.md'), index + sqlMatrix + failures)
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
    for (const f of ['schema.prisma', 'source.sql', 'source.json', 'sql']) if (existsSync(join(from, f))) cpSync(join(from, f), join(to, f), { recursive: true })
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

function dropDatabase(dialect, url) {
  if (dialect === 'postgresql') psql('postgres', `DROP DATABASE IF EXISTS "${dbName(url)}" WITH (FORCE)`)
  else if (dialect === 'mysql') mysql(`DROP DATABASE IF EXISTS \`${dbName(url)}\``)
  else rmSync(url.slice('file:'.length), { force: true })
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
  // Over TCP: while the image initializes, a temporary server answers on the socket only.
  for (let i = 0; i < 90; i++) {
    const ok = image.startsWith('mysql')
      ? sh(['docker', 'exec', name, 'mysql', '-h127.0.0.1', '-uroot', '-pcompat', '-e', 'SELECT 1'])
      : sh(['docker', 'exec', name, 'pg_isready', '-h', '127.0.0.1', '-U', 'postgres'])
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
    // pg_dump's own comment lines (`-- Name: ...; Type: ...`) go: the statements stay.
    return run.stdout
      .replace(/^\\(un)?restrict .*\n/gm, '')
      .replace(/^--.*\n/gm, '')
      .replace(/\n{3,}/g, '\n\n')
      .replace(/^\n+/, '')
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
