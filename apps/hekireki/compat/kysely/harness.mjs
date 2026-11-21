// Runs the plan through Kysely typed with the `DB` hekireki wrote, against the database Prisma
// created. hekireki-kysely writes only types, so what is held to the database here is what the
// types promise:
//
// - a column typed `Generated<...>` may be left out of an insert: the database has to fill it.
// - a column not typed `Generated` has to be given: when Prisma Client fills it (cuid(), uuid(),
//   now(), @updatedAt), the harness gives one the way an application would.
// - a value read back has the runtime type the select type declares (`Date`, `bigint`, ...).
//
//   node harness.mjs crud <plan.json> <types.ts> <url>
//
// The dialects are Kysely's own: PostgresDialect over pg, MysqlDialect over mysql2, and
// SqliteDialect over node:sqlite (through the better-sqlite3 shape the dialect expects).
import { readFileSync } from 'node:fs'
import { randomUUID } from 'node:crypto'
import ts from 'typescript'
import { Kysely, MysqlDialect, PostgresDialect, SqliteDialect, sql } from 'kysely'

const [mode, planPath, typesPath, url] = process.argv.slice(2)
if (mode !== 'crud') throw new Error('kysely has no tables of its own: crud only')
const plan = JSON.parse(readFileSync(planPath, 'utf8'))
const types = parseTypes(readFileSync(typesPath, 'utf8'))
const db = new Kysely({ dialect: await dialect(plan.dialect, url) })

const out = { orm: 'kysely', models: [] }
const rows = new Map()
const done = []

for (const m of plan.models) {
  const r = { model: m.name }
  out.models.push(r)
  const iface = types.tables.get(m.table)
  if (!iface) {
    r.missing = true
    continue
  }
  const fields = types.interfaces.get(iface)
  r.unmapped = m.fields.filter((f) => !fields.has(f.column)).map((f) => f.column)
  if (!m.insertable) {
    r.skipped = 'a required key closes a cycle'
    continue
  }
  if (m.pk.length === 0) {
    r.read = { ok: false, error: 'the plan has no key for this model' }
    continue
  }
  r.notes = []
  const values = {}
  for (const f of m.fields) {
    const declared = fields.get(f.column)
    if (!declared) continue
    if (f.write === 'value') values[f.column] = toJs(f, f.value, declared, plan.dialect)
    else if (f.write === 'parent') {
      const parent = f.fk.available ? rows.get(f.fk.model) : null
      values[f.column] = parent ? toJs(f, parent[f.fk.column], declared, plan.dialect) : null
    } else if (f.write === 'orm' && !declared.generated) {
      const given = clientDefault(f)
      if (given === undefined) {
        r.notes.push(`\`${f.column}\` has a default in the database (${f.default}) but is not typed Generated, so the type makes every insert give it`)
        continue
      }
      values[f.column] = toJs(f, given, declared, plan.dialect)
    }
  }
  // The key comes back from the insert: RETURNING on PostgreSQL and SQLite, insertId on MySQL.
  const pk = {}
  let row
  try {
    const insert = db.insertInto(m.table)
    // Every column left to the database: INSERT ... DEFAULT VALUES (MySQL: VALUES ()).
    const q = Object.keys(values).length ? insert.values(values) : plan.dialect === 'mysql' ? insert.values({}) : insert.defaultValues()
    if (plan.dialect === 'mysql') {
      const res = await q.executeTakeFirst()
      for (const c of m.pk) pk[c] = values[c] !== undefined ? values[c] : res.insertId !== undefined ? Number(res.insertId) : undefined
    } else {
      const returned = await q.returningAll().executeTakeFirst()
      for (const c of m.pk) pk[c] = returned[c]
    }
    r.insert = { ok: true }
  } catch (e) {
    r.insert = { ok: false, error: message(e) }
    continue
  }
  try {
    if (m.pk.some((c) => pk[c] === undefined)) throw new Error('the insert did not give the key back')
    row = await m.pk.reduce((acc, c) => acc.where(c, '=', pk[c]), db.selectFrom(m.table).selectAll()).executeTakeFirst()
    if (!row) throw new Error('no row came back for the key the insert returned')
    const read = {}
    const mismatches = []
    for (const [col, declared] of fields) {
      const f = m.fields.find((x) => x.column === col)
      read[col] = normal(row[col], f)
      const got = kind(row[col])
      if (!declared.select.has(got) && !declared.select.has('unknown')) mismatches.push({ column: col, declared: declared.text, got })
    }
    r.read = { ok: true, values: read, typeMismatches: mismatches }
    rows.set(m.name, read)
  } catch (e) {
    r.read = { ok: false, error: message(e) }
    done.push({ m, r, where: (q) => m.pk.reduce((acc, c) => acc.where(c, '=', pk[c]), q) })
    continue
  }
  r.pk = Object.fromEntries(Object.entries(pk).map(([k, v]) => [k, normal(v)]))
  const where = (q) => m.pk.reduce((acc, c) => acc.where(c, '=', pk[c]), q)
  if (m.update) {
    try {
      await where(db.updateTable(m.table).set({ [m.update.column]: toJs(m.fields.find((f) => f.column === m.update.column), m.update.value, fields.get(m.update.column), plan.dialect) })).execute()
      const again = await where(db.selectFrom(m.table).selectAll()).executeTakeFirst()
      r.update = { ok: true, value: normal(again[m.update.column], m.fields.find((f) => f.column === m.update.column)) }
    } catch (e) {
      r.update = { ok: false, error: message(e) }
    }
  }
  r.relations = [] // Kysely has no relations: joins are written by hand.
  done.push({ m, r, where })
}

for (const { m, r, where } of done.reverse()) {
  try {
    await where(db.deleteFrom(m.table)).execute()
    const left = await where(db.selectFrom(m.table).selectAll()).execute()
    r.delete = left.length === 0 ? { ok: true } : { ok: false, error: 'the row is still there' }
  } catch (e) {
    r.delete = { ok: false, error: message(e) }
  }
}

process.stdout.write(JSON.stringify(out))
await db.destroy()

// --- the types --------------------------------------------------------------------------------

// Every interface's fields with the runtime kinds its select type allows, and the `DB` interface's
// table -> interface map. Aliases (`Timestamp`, an enum's union of literals) are followed.
function parseTypes(source) {
  const file = ts.createSourceFile('types.ts', source, ts.ScriptTarget.Latest, true)
  const aliases = new Map()
  const interfaces = new Map()
  for (const node of file.statements) {
    if (ts.isTypeAliasDeclaration(node)) aliases.set(node.name.text, node.type)
    if (ts.isInterfaceDeclaration(node)) {
      const fields = new Map()
      for (const member of node.members) {
        if (!ts.isPropertySignature(member) || !member.type) continue
        const name = member.name.getText(file).replace(/^['"]|['"]$/g, '')
        fields.set(name, member.type)
      }
      interfaces.set(node.name.text, fields)
    }
  }
  const tables = new Map()
  for (const [name, type] of interfaces.get('DB') ?? []) tables.set(name, type.getText(file))

  const kinds = (type, seen = new Set()) => {
    if (ts.isUnionTypeNode(type)) return type.types.flatMap((t) => kinds(t, seen))
    if (ts.isParenthesizedTypeNode(type)) return kinds(type.type, seen)
    if (ts.isArrayTypeNode(type)) return ['array']
    if (ts.isLiteralTypeNode(type)) {
      if (type.literal.kind === ts.SyntaxKind.NullKeyword) return ['null']
      if (ts.isStringLiteral(type.literal)) return ['string']
      if (ts.isNumericLiteral(type.literal)) return ['number']
      return ['boolean']
    }
    switch (type.kind) {
      case ts.SyntaxKind.StringKeyword:
        return ['string']
      case ts.SyntaxKind.NumberKeyword:
        return ['number']
      case ts.SyntaxKind.BooleanKeyword:
        return ['boolean']
      case ts.SyntaxKind.BigIntKeyword:
        return ['bigint']
      case ts.SyntaxKind.UnknownKeyword:
      case ts.SyntaxKind.AnyKeyword:
        return ['unknown']
    }
    if (ts.isTypeReferenceNode(type)) {
      const name = type.typeName.getText(file)
      const args = type.typeArguments ?? []
      if (name === 'Generated' || name === 'ColumnType') return kinds(args[0], seen)
      if (name === 'Date') return ['Date']
      if (name === 'Buffer' || name === 'Uint8Array') return ['Buffer']
      if (name === 'Array') return ['array']
      if (name === 'Record') return ['object']
      if (aliases.has(name) && !seen.has(name)) return kinds(aliases.get(name), new Set([...seen, name]))
      if (interfaces.has(name)) return ['object']
    }
    return ['unknown']
  }
  const resolved = new Map()
  for (const [name, fields] of interfaces) {
    if (name === 'DB') continue
    const out = new Map()
    for (const [col, type] of fields) {
      const text = type.getText(file)
      out.set(col, { text, generated: /^Generated</.test(text), select: new Set(kinds(type)) })
    }
    resolved.set(name, out)
  }
  return { tables, interfaces: resolved }
}

function kind(value) {
  if (value === null || value === undefined) return 'null'
  if (value instanceof Date) return 'Date'
  if (Buffer.isBuffer(value) || value instanceof Uint8Array) return 'Buffer'
  if (Array.isArray(value)) return 'array'
  return typeof value
}

// What an application gives for a column Prisma Client fills on its own.
function clientDefault(f) {
  switch (f.default) {
    case 'uuid':
      return f.native?.[0] === 'Uuid' ? randomUUID() : randomUUID()
    case 'cuid':
      return `c${randomUUID().replace(/-/g, '').slice(0, 24)}`
    case 'nanoid':
    case 'ulid':
      return randomUUID().replace(/-/g, '').slice(0, 21)
    case 'now':
      return new Date().toISOString()
  }
  if (f.updatedAt) return new Date().toISOString()
  return undefined
}

function toJs(f, value, declared, dialect) {
  if (value == null) return value
  if (f.list) return dialect === 'postgresql' ? value.map((v) => toJs({ ...f, list: false }, v, declared, dialect)) : JSON.stringify(value)
  switch (f.type) {
    case 'DateTime':
      // A column typed string (@db.Time) takes the text a time literal is written as.
      if (declared && !declared.select.has('Date') && declared.select.has('string')) return f.native?.[0]?.startsWith('Time') ? value.slice(11, 19) : value
      return new Date(value)
    case 'BigInt':
      return declared?.select.has('bigint') ? BigInt(value) : declared?.select.has('number') ? Number(value) : String(value)
    case 'Decimal':
      return declared?.select.has('number') ? Number(value) : String(value)
    case 'Bytes':
      return Buffer.from(value, 'base64')
    case 'Json':
      // pg serializes an object itself; mysql2 and SQLite take the JSON text, as Kysely's docs say.
      return dialect === 'postgresql' ? JSON.stringify(value) : JSON.stringify(value)
    default:
      return value
  }
}

function normal(value, f) {
  // A time of day read back as text ("12:34:56"), in the plan's form.
  if (typeof value === 'string' && f?.type === 'DateTime' && /^\d{2}:\d{2}:\d{2}/.test(value)) return `1970-01-01T${value.slice(0, 8)}.${(value.split('.')[1] ?? '').padEnd(3, '0').slice(0, 3)}Z`
  if (value instanceof Date) return Number.isNaN(value.getTime()) ? 'Invalid Date' : value.toISOString()
  if (typeof value === 'bigint') return String(value)
  if (Buffer.isBuffer(value) || value instanceof Uint8Array) return Buffer.from(value).toString('base64')
  if (Array.isArray(value)) return value.map((v) => normal(v))
  return value === undefined ? null : value
}

function message(e) {
  return String(e?.message ?? e).slice(0, 600)
}

async function dialect(name, url) {
  if (name === 'postgresql') {
    const { default: pg } = await import('pg')
    return new PostgresDialect({ pool: new pg.Pool({ connectionString: url, max: 2 }) })
  }
  if (name === 'mysql') {
    const { createPool } = await import('mysql2')
    return new MysqlDialect({ pool: createPool({ uri: url, connectionLimit: 2 }) })
  }
  // node:sqlite in the shape SqliteDialect expects of better-sqlite3. Like better-sqlite3, it
  // binds no booleans or dates; an application turns them into 1/0 and ISO text, as here.
  const { DatabaseSync } = await import('node:sqlite')
  const database = new DatabaseSync(url.replace(/^file:/, ''))
  const bind = (params) =>
    params.map((p) => (typeof p === 'boolean' ? (p ? 1 : 0) : p instanceof Date ? p.toISOString() : Buffer.isBuffer(p) ? new Uint8Array(p) : p))
  return new SqliteDialect({
    database: {
      close: () => database.close(),
      prepare(text) {
        const statement = database.prepare(text)
        return {
          reader: statement.columns().length > 0,
          all: (params) => statement.all(...bind(params)),
          run: (params) => statement.run(...bind(params)),
          iterate: (params) => statement.iterate(...bind(params)),
        }
      },
    },
  })
}

void sql
