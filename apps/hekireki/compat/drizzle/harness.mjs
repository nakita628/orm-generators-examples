// Runs the plan through the Drizzle schema hekireki wrote, against the database Prisma created.
//
//   node harness.mjs crud <plan.json> <schema.ts> <url>   insert, read, update, relations, delete
//   node harness.mjs ddl <plan.json> <schema.ts>          print the CREATE TABLEs drizzle-kit export writes
//
// The schema is a copy under compat/.work, so that it resolves drizzle-orm from this package:
// one drizzle-orm instance for the schema and the harness.
import { readFileSync } from 'node:fs'
import { and, eq, getTableName, is, Table } from 'drizzle-orm'

const [mode, planPath, schemaPath, url] = process.argv.slice(2)
const plan = JSON.parse(readFileSync(planPath, 'utf8'))
const schema = await import(schemaPath)

const tables = new Map()
for (const [key, value] of Object.entries(schema)) {
  if (is(value, Table)) tables.set(getTableName(value), { key, table: value })
}

if (mode === 'ddl') {
  // drizzle-kit's own CREATE TABLEs for the schema (`drizzle-kit export`), as `generate` would
  // write the first migration.
  const { spawnSync } = await import('node:child_process')
  const kit = new URL('../node_modules/.bin/drizzle-kit', import.meta.url).pathname
  const run = spawnSync(kit, ['export', '--dialect', plan.dialect, '--schema', schemaPath], { encoding: 'utf8' })
  if (run.status !== 0 || !/CREATE TABLE/i.test(run.stdout)) {
    process.stderr.write(run.stderr || run.stdout)
    process.exit(1)
  }
  process.stdout.write(run.stdout.replace(/--> statement-breakpoint/g, ''))
  process.exit(0)
}

const db = await connect(plan.dialect, url)
const out = { orm: 'drizzle', models: [] }
const rows = new Map() // model name -> the row as read back, by column

for (const m of plan.models) {
  const r = { model: m.name }
  out.models.push(r)
  const found = tables.get(m.table)
  if (!found) {
    r.missing = true
    continue
  }
  const { key, table } = found
  const columns = Object.entries(table).length ? columnsOf(table) : new Map()
  r.unmapped = m.fields.filter((f) => !columns.has(f.column)).map((f) => f.column)
  if (!m.insertable) {
    r.skipped = 'a required key closes a cycle'
    continue
  }

  const values = {}
  for (const f of m.fields) {
    const c = columns.get(f.column)
    if (!c) continue
    if (f.write === 'value') values[c.key] = toJs(f, f.value, c.column)
    else if (f.write === 'parent') {
      const parent = f.fk.available ? rows.get(f.fk.model) : null
      values[c.key] = parent ? toJs(f, parent[f.fk.column], c.column) : null
    }
  }

  let pk
  try {
    pk = await insert(db, plan.dialect, table, values, m, columns)
    r.insert = { ok: true }
  } catch (e) {
    r.insert = { ok: false, error: message(e) }
    continue
  }
  r.pk = pk
  if (m.pk.length === 0) {
    r.read = { ok: false, error: 'the plan has no key for this model' }
    continue
  }
  const where = () => and(...m.pk.map((col) => eq(columns.get(col).column, toJs(m.fields.find((f) => f.column === col), pk[col], columns.get(col).column))))
  try {
    const [row] = await db.select().from(table).where(where())
    if (!row) throw new Error('no row came back for the key the insert returned')
    const read = {}
    for (const [col, c] of columns) read[col] = normal(row[c.key])
    r.read = { ok: true, values: read }
    rows.set(m.name, read)
  } catch (e) {
    r.read = { ok: false, error: message(e) }
    continue
  }
  if (m.update) {
    const c = columns.get(m.update.column)
    try {
      if (!c) throw new Error(`no attribute maps ${m.update.column}`)
      await db.update(table).set({ [c.key]: m.update.value }).where(where())
      const [row] = await db.select().from(table).where(where())
      r.update = { ok: true, value: normal(row[c.key]) }
    } catch (e) {
      r.update = { ok: false, error: message(e) }
    }
  }
  r.where = where
  r.key = key
}

// Relations: through the relational query API (`db.query.<table>.findFirst({ with })`), every
// relation the plan has, from both sides.
for (const m of plan.models) {
  const r = out.models.find((x) => x.model === m.name)
  if (!r.read?.ok) continue
  r.relations = []
  for (const rel of m.relations) {
    const result = { name: rel.name, kind: rel.kind }
    r.relations.push(result)
    const target = out.models.find((x) => x.model === rel.model)
    if (!target?.read?.ok) {
      result.skipped = 'no row on the other side'
      continue
    }
    const name = findRelation(schema.relations, r.key, rel.name)
    if (!name) {
      result.found = false
      continue
    }
    result.found = true
    try {
      const pkWhere = Object.fromEntries(m.pk.map((col) => [columnsOf(tables.get(m.table).table).get(col).key, toJs(m.fields.find((f) => f.column === col), r.pk[col], columnsOf(tables.get(m.table).table).get(col).column)]))
      const row = await db.query[r.key].findFirst({ where: pkWhere, with: { [name]: true } })
      const loaded = row?.[name]
      const expected = rel.kind === 'toOne' ? rel.columns.every((c) => r.read.values[c] != null) : true
      const count = loaded == null ? 0 : Array.isArray(loaded) ? loaded.length : 1
      result.ok = expected ? count > 0 : count === 0
      result.loaded = count
    } catch (e) {
      result.ok = false
      result.error = message(e)
    }
  }
}

for (const m of [...plan.models].reverse()) {
  const r = out.models.find((x) => x.model === m.name)
  if (!r.where) continue
  const { table } = tables.get(m.table)
  try {
    await db.delete(table).where(r.where())
    const left = await db.select().from(table).where(r.where())
    r.delete = { ok: left.length === 0, ...(left.length ? { error: 'the row is still there' } : {}) }
  } catch (e) {
    r.delete = { ok: false, error: message(e) }
  }
  delete r.where
  delete r.key
}

process.stdout.write(JSON.stringify(out))
await db.$client?.end?.()
process.exit(0)

function columnsOf(table) {
  const out = new Map()
  for (const [key, column] of Object.entries(table)) {
    if (column && typeof column === 'object' && 'name' in column && 'columnType' in column) out.set(column.name, { key, column })
  }
  return out
}

async function insert(db, dialect, table, values, m, columns) {
  const pk = {}
  if (dialect === 'mysql') {
    const [ids] = await db.insert(table).values(values).$returningId()
    for (const col of m.pk) {
      const c = columns.get(col)
      pk[col] = normal(values[c.key] ?? ids?.[c.key])
    }
    // A key no one wrote and $returningId does not know (a composite one with a default): read
    // it back by the columns that were written.
    return pk
  }
  const [row] = await db.insert(table).values(values).returning()
  for (const col of m.pk) pk[col] = normal(row[columns.get(col).key])
  return pk
}

function findRelation(relations, tableKey, prismaName) {
  const config = relations?.[tableKey]?.relations
  if (!config) return null
  if (prismaName in config) return prismaName
  const squash = (s) => s.toLowerCase().replace(/_/g, '')
  return Object.keys(config).find((k) => squash(k) === squash(prismaName)) ?? null
}

function toJs(f, value, column) {
  if (value == null) return value
  if (f.list) return value.map((v) => toJs({ ...f, list: false }, v, column))
  const dataType = String(column.dataType ?? '')
  switch (f.type) {
    case 'DateTime':
      return dataType.startsWith('string') ? value : new Date(value)
    case 'BigInt':
      return dataType.includes('bigint') ? BigInt(value) : dataType.startsWith('string') ? String(value) : Number(value)
    case 'Decimal':
      return dataType.startsWith('number') ? Number(value) : String(value)
    case 'Bytes':
      return Buffer.from(value, 'base64')
    default:
      return value
  }
}

function normal(value) {
  if (value instanceof Date) return Number.isNaN(value.getTime()) ? 'Invalid Date' : value.toISOString()
  if (typeof value === 'bigint') return String(value)
  if (value instanceof Uint8Array) return Buffer.from(value).toString('base64')
  if (Array.isArray(value)) return value.map(normal)
  return value === undefined ? null : value
}

function message(e) {
  // drizzle wraps the driver's error: the driver's reason first, then the query it ran.
  const cause = e?.cause?.message ?? (e?.cause ? String(e.cause) : '')
  return (cause ? `${cause} — ${e?.message ?? e}` : String(e?.message ?? e)).slice(0, 600)
}

async function connect(dialect, url) {
  if (dialect === 'postgresql') {
    const { default: postgres } = await import('postgres')
    const { drizzle } = await import('drizzle-orm/postgres-js')
    return drizzle({ client: postgres(url, { onnotice: () => {} }), relations: schema.relations })
  }
  if (dialect === 'mysql') {
    const mysql = await import('mysql2/promise')
    const { drizzle } = await import('drizzle-orm/mysql2')
    return drizzle({ client: mysql.createPool(url), relations: schema.relations, mode: 'default' })
  }
  const { DatabaseSync } = await import('node:sqlite')
  const { drizzle } = await import('drizzle-orm/node-sqlite')
  return drizzle({ client: new DatabaseSync(url.replace(/^file:/, '')), relations: schema.relations })
}
