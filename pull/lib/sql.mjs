// Holds the SQL of two sets of tables against each other: the tables a project's own ORM created
// (source.sql) and the tables another ORM created from hekireki's code (sql/<orm>.sql), both
// dumped the same way (pg_dump, mysqldump, or SQLite's own schema).
//
// Each side is read into what it says about each table: its columns (type, nullability,
// default, identity), its primary key, uniques and checks, its foreign keys (columns, target,
// options) and its indexes, with the names of constraints and indexes set aside; anything else
// (types, functions, triggers, views) is kept as its statement. Then:
// - identical: the same statements in the same order (once session settings, comments and
//   extensions are left out);
// - the same but for names: the same facts, whatever the constraints and indexes are called and
//   whatever order the columns and statements come in (and with what only spells the same thing
//   two ways taken as one: a unique constraint and a plain unique index on the same columns,
//   PostgreSQL's `timestamp(6)` and `timestamp`, whose default precision is 6);
// - otherwise, table by table, what one side says and the other does not.
import { spawnSync } from 'node:child_process'
import { mkdtempSync, rmSync, writeFileSync } from 'node:fs'
import { createRequire } from 'node:module'
import { tmpdir } from 'node:os'
import { join } from 'node:path'

const require = createRequire(import.meta.url)

/**
 * @param {string} source  SQL of the tables the project's own ORM created
 * @param {string} other   SQL of the tables another ORM created
 * @param {'postgresql' | 'mysql' | 'sqlite'} dialect
 */
export function compareSql(source, other, dialect) {
  const a = read(source, dialect)
  const b = read(other, dialect)
  const identical = a.strict.length === b.strict.length && a.strict.every((s, i) => s === b.strict[i])
  const tables = []
  for (const name of [...new Set([...a.tables.keys(), ...b.tables.keys()])].sort()) {
    const ta = a.tables.get(name)
    const tb = b.tables.get(name)
    if (!ta || !tb) {
      tables.push({ name, only: ta ? 'source' : 'here' })
      continue
    }
    const lines = tableDiff(ta, tb)
    if (lines.length) tables.push({ name, lines })
  }
  const onlySource = [...a.other].filter((s) => !b.other.has(s))
  const onlyOther = [...b.other].filter((s) => !a.other.has(s))
  return {
    identical,
    sameButNames: tables.length === 0 && onlySource.length === 0 && onlyOther.length === 0,
    tables,
    onlySource,
    onlyOther,
    diff: unifiedDiff(a.pretty, b.pretty),
  }
}

// --- reading a dump -----------------------------------------------------------------------------

function emptyTable() {
  return { columns: new Map(), keys: new Set(), fks: new Map(), indexes: new Set(), options: '' }
}

function read(sql, dialect) {
  if (dialect === 'sqlite') return readSqlite(sql)
  const statements = split(sql).filter((s) => !NOT_TABLES.some((re) => re.test(s)))
  const strict = statements.map(collapse)
  const pretty = statements.map((s) => s.split('\n').map((l) => l.trimEnd()).filter((l) => l.trim()).join('\n') + ';').join('\n\n') + '\n'
  const tables = new Map()
  const other = new Set()
  const table = (name) => {
    if (!tables.has(name)) tables.set(name, emptyTable())
    return tables.get(name)
  }
  if (dialect === 'mysql') for (const s of strict) readMysql(s, table, other)
  else {
    const owned = new Set()
    for (const s of strict) {
      const m = /^ALTER SEQUENCE (\S+) OWNED BY \S+$/i.exec(s)
      if (m) owned.add(m[1])
    }
    for (const s of strict) readPg(s, table, other, owned)
    for (const t of tables.values()) for (const [column, def] of t.columns) t.columns.set(column, sameDefault(def))
  }
  return { strict, tables, other, pretty }
}

// Defaults PostgreSQL's tools spell two ways for one value: an empty array (`ARRAY[]::bigint[]`,
// `'{}'::bigint[]`, `ARRAY[]::text[]` on a varchar[] column: the column says the type) and a float's
// whole number with or without its `.0` (`DEFAULT 0.0`, `DEFAULT 0`).
function sameDefault(def) {
  let d = def.replace(/\bARRAY\[\]::[\w ]+(?:\([^)]*\))?\[\]/gi, 'ARRAY[]').replace(/'\{\}'::[\w ]+(?:\([^)]*\))?\[\]/gi, 'ARRAY[]')
  if (/^(double precision|real|float\d*)\b/i.test(d)) d = d.replace(/(\bDEFAULT\s+-?\d+)\.0+\b/i, '$1')
  return d
}

// What a dump holds besides the tables: session settings, extensions, the schema itself.
const NOT_TABLES = [
  /^SET\s/i,
  /^SELECT pg_catalog\.set_config/i,
  /^CREATE EXTENSION/i,
  /^COMMENT ON EXTENSION/i,
  /^CREATE SCHEMA/i,
  /^COMMENT ON SCHEMA/i,
  /^ALTER\s.*\sOWNER TO\s/i,
  /^\/\*!\d+ SET /i,
]

function readPg(s, table, other, owned) {
  let m
  if ((m = /^CREATE (?:UNLOGGED )?TABLE (\S+) \((.*)\)(.*)$/i.exec(s))) {
    const t = table(m[1])
    for (const item of items(m[2])) {
      const c = /^CONSTRAINT \S+ (.*)$/i.exec(item)
      if (c) t.keys.add(c[1])
      else t.columns.set(item.split(' ')[0], item.slice(item.indexOf(' ') + 1).replace(/\b(timestamp|time)\(6\)/gi, '$1'))
    }
    if (m[3].trim()) t.options = m[3].trim()
  } else if ((m = /^ALTER TABLE (?:ONLY )?(\S+) ALTER COLUMN (\S+) SET DEFAULT (.*)$/i.exec(s))) {
    // A serial column's default draws on its own sequence, whatever that sequence is called.
    const nextval = /^nextval\('([^']+)'::regclass\)$/.exec(m[3])
    addToColumn(table(m[1]), m[2], nextval && owned.has(nextval[1]) ? 'DEFAULT nextval(<its sequence>)' : `DEFAULT ${m[3]}`)
  } else if ((m = /^ALTER TABLE (?:ONLY )?(\S+) ALTER COLUMN (\S+) ADD GENERATED (ALWAYS|BY DEFAULT) AS IDENTITY/i.exec(s))) {
    addToColumn(table(m[1]), m[2], `GENERATED ${m[3].toUpperCase()} AS IDENTITY`)
  } else if ((m = /^CREATE SEQUENCE (\S+) /i.exec(s)) && owned.has(m[1])) {
    // A column's own sequence (serial): said by its default.
  } else if (/^ALTER SEQUENCE \S+ OWNED BY /i.test(s)) {
    // ditto
  } else if ((m = /^ALTER TABLE (?:ONLY )?(\S+) ADD CONSTRAINT \S+ FOREIGN KEY \(([^)]*)\) REFERENCES (\S+?)\(([^)]*)\)\s*(.*)$/i.exec(s))) {
    table(m[1]).fks.set(`FOREIGN KEY (${m[2]}) REFERENCES ${m[3]}(${m[4]})`, m[5].trim())
  } else if ((m = /^ALTER TABLE (?:ONLY )?(\S+) ADD CONSTRAINT \S+ (.*)$/i.exec(s))) {
    table(m[1]).keys.add(m[2])
  } else if ((m = /^CREATE (UNIQUE )?INDEX \S+ ON (?:ONLY )?(\S+) (.*)$/i.exec(s))) {
    // A unique btree index on plain columns enforces what a unique constraint on them does.
    const plain = m[1] && /^USING btree \(([\w", ]+)\)$/i.exec(m[3])
    if (plain) table(m[2]).keys.add(`UNIQUE (${plain[1]})`)
    else table(m[2]).indexes.add(`${m[1] ? 'UNIQUE ' : ''}INDEX ${m[3]}`)
  } else other.add(s)
}

function addToColumn(t, column, what) {
  t.columns.set(column, `${t.columns.get(column) ?? '?'} ${what}`)
}

function readMysql(s, table, other) {
  const m = /^CREATE TABLE (`[^`]+`) \((.*)\)(.*)$/i.exec(s)
  if (!m) return other.add(s)
  const t = table(m[1])
  for (const item of items(m[2])) {
    let k
    if ((k = /^`([^`]+)` (.*)$/.exec(item))) t.columns.set(`\`${k[1]}\``, k[2])
    else if ((k = /^CONSTRAINT `[^`]+` FOREIGN KEY \(([^)]*)\) REFERENCES (`[^`]+`) \(([^)]*)\)\s*(.*)$/i.exec(item)))
      t.fks.set(`FOREIGN KEY (${k[1]}) REFERENCES ${k[2]} (${k[3]})`, k[4].trim())
    else if ((k = /^CONSTRAINT `[^`]+` (.*)$/i.exec(item))) t.keys.add(k[1])
    else if ((k = /^(UNIQUE |FULLTEXT |SPATIAL )?KEY `[^`]+` (.*)$/i.exec(item))) {
      if (k[1]?.trim().toUpperCase() === 'UNIQUE') t.keys.add(`UNIQUE ${k[2]}`)
      else t.indexes.add(`${k[1] ?? ''}INDEX ${k[2]}`)
    } else t.keys.add(item)
  }
  t.options = m[3].trim()
}

// SQLite keeps each CREATE as its ORM wrote it, so the tables are read back through SQLite itself
// (PRAGMA table_info, index_list, index_info, foreign_key_list).
function readSqlite(sql) {
  const db = new (require('node:sqlite').DatabaseSync)(':memory:')
  try {
    db.exec(sql.replace(/^--> statement-breakpoint$/gm, ''))
    const lit = (s) => `'${s.replaceAll("'", "''")}'`
    const q = (x) => `"${x.replaceAll('"', '""')}"`
    const shadow = new Set(db.prepare("SELECT name FROM pragma_table_list WHERE type = 'shadow'").all().map((r) => r.name))
    const tables = new Map()
    const other = new Set()
    const strict = []
    const list = db.prepare("SELECT name, sql FROM sqlite_master WHERE type = 'table' AND name NOT LIKE 'sqlite_%' ORDER BY name").all()
    for (const { name, sql: create } of list.filter((t) => !shadow.has(t.name))) {
      if (/^CREATE VIRTUAL TABLE/i.test(create)) {
        other.add(collapse(create))
        strict.push(collapse(create))
        continue
      }
      const t = emptyTable()
      tables.set(q(name), t)
      const columns = db.prepare(`SELECT * FROM pragma_table_info(${lit(name)}) ORDER BY cid`).all()
      const pk = columns.filter((c) => c.pk).sort((x, y) => x.pk - y.pk)
      // An INTEGER PRIMARY KEY is the rowid, never NULL whether or not the SQL says NOT NULL.
      const rowid = pk.length === 1 && String(pk[0].type).toUpperCase() === 'INTEGER' ? pk[0].name : null
      for (const c of columns)
        t.columns.set(q(c.name), `${String(c.type).toUpperCase()}${c.notnull || c.name === rowid ? ' NOT NULL' : ''}${c.dflt_value !== null ? ` DEFAULT ${c.dflt_value}` : ''}`)
      if (pk.length) t.keys.add(`PRIMARY KEY (${pk.map((c) => q(c.name)).join(', ')})${/\bAUTOINCREMENT\b/i.test(create) ? ' AUTOINCREMENT' : ''}`)
      for (const ix of db.prepare(`SELECT * FROM pragma_index_list(${lit(name)}) ORDER BY name`).all()) {
        if (ix.origin === 'pk') continue
        const cols = db.prepare(`SELECT name FROM pragma_index_info(${lit(ix.name)}) ORDER BY seqno`).all().map((r) => (r.name === null ? '<expr>' : q(r.name)))
        const where = ix.partial ? ` WHERE ${/\bWHERE\b(.*)$/is.exec(db.prepare('SELECT sql FROM sqlite_master WHERE name = ?').get(ix.name)?.sql ?? '')?.[1]?.trim() ?? '?'}` : ''
        if (ix.origin === 'u' || (ix.unique && !where && !cols.includes('<expr>'))) t.keys.add(`UNIQUE (${cols.join(', ')})`)
        else t.indexes.add(`${ix.unique ? 'UNIQUE ' : ''}INDEX (${cols.join(', ')})${where}`)
      }
      const fks = db.prepare(`SELECT * FROM pragma_foreign_key_list(${lit(name)}) ORDER BY id, seq`).all()
      for (const id of [...new Set(fks.map((f) => f.id))]) {
        const parts = fks.filter((f) => f.id === id)
        t.fks.set(
          `FOREIGN KEY (${parts.map((p) => q(p.from)).join(', ')}) REFERENCES ${q(parts[0].table)} (${parts.map((p) => (p.to === null ? '<pk>' : q(p.to))).join(', ')})`,
          `ON UPDATE ${parts[0].on_update} ON DELETE ${parts[0].on_delete}`,
        )
      }
      strict.push(collapse(create))
      for (const ix of db.prepare("SELECT sql FROM sqlite_master WHERE type = 'index' AND tbl_name = ? AND sql IS NOT NULL ORDER BY name").all(name)) strict.push(collapse(ix.sql))
    }
    for (const t of db.prepare("SELECT sql FROM sqlite_master WHERE type IN ('view', 'trigger') ORDER BY type, name").all()) {
      other.add(collapse(t.sql))
      strict.push(collapse(t.sql))
    }
    return { strict, tables, other, pretty: strict.join(';\n') + ';\n' }
  } finally {
    db.close()
  }
}

// --- comparing two tables -----------------------------------------------------------------------

// A table's facts against the other side's, as diff lines: `-` what source.sql says, `+` what
// the other says, a column or foreign key written differently as a `-` line and a `+` line.
function tableDiff(a, b) {
  const lines = []
  for (const name of new Set([...a.columns.keys(), ...b.columns.keys()])) {
    const x = a.columns.get(name)
    const y = b.columns.get(name)
    if (x === y) continue
    if (x !== undefined) lines.push(`- ${name} ${x}`)
    if (y !== undefined) lines.push(`+ ${name} ${y}`)
  }
  for (const k of a.keys) if (!b.keys.has(k)) lines.push(`- ${k}`)
  for (const k of b.keys) if (!a.keys.has(k)) lines.push(`+ ${k}`)
  for (const [k, x] of a.fks) {
    const y = b.fks.get(k)
    if (y === x) continue
    lines.push(`- ${k}${x ? ` ${x}` : ''}`)
    if (y !== undefined) lines.push(`+ ${k}${y ? ` ${y}` : ''}`)
  }
  for (const [k, y] of b.fks) if (!a.fks.has(k)) lines.push(`+ ${k}${y ? ` ${y}` : ''}`)
  for (const k of a.indexes) if (!b.indexes.has(k)) lines.push(`- ${k}`)
  for (const k of b.indexes) if (!a.indexes.has(k)) lines.push(`+ ${k}`)
  if (a.options !== b.options) lines.push(...(a.options ? [`- (table) ${a.options}`] : []), ...(b.options ? [`+ (table) ${b.options}`] : []))
  return lines
}

// --- text helpers -------------------------------------------------------------------------------

// Statements, split on `;` outside quotes, dollar-quoted bodies and comments; comments dropped.
function split(sql) {
  const out = []
  let cur = ''
  let quote = null
  let dollar = null
  for (let i = 0; i < sql.length; ) {
    const c = sql[i]
    if (dollar) {
      if (sql.startsWith(dollar, i)) {
        cur += dollar
        i += dollar.length
        dollar = null
      } else {
        cur += c
        i++
      }
      continue
    }
    if (quote) {
      cur += c
      i++
      if (c === quote) {
        if (sql[i] === quote) {
          cur += sql[i++]
          continue
        }
        quote = null
      }
      continue
    }
    if (c === "'" || c === '"' || c === '`') {
      quote = c
      cur += c
      i++
      continue
    }
    if (c === '$') {
      const m = /^\$[A-Za-z_]*\$/.exec(sql.slice(i, i + 64))
      if (m) {
        dollar = m[0]
        cur += dollar
        i += dollar.length
        continue
      }
    }
    if (c === '-' && sql[i + 1] === '-') {
      const nl = sql.indexOf('\n', i)
      i = nl < 0 ? sql.length : nl + 1
      continue
    }
    if (c === ';') {
      if (cur.trim()) out.push(cur.trim())
      cur = ''
      i++
      continue
    }
    cur += c
    i++
  }
  if (cur.trim()) out.push(cur.trim())
  return out
}

function collapse(s) {
  return s.replace(/\s+/g, ' ').replace(/\( /g, '(').replace(/ \)/g, ')').trim()
}

// The comma-separated items of a parenthesized list, at depth one, quotes kept whole.
function items(body) {
  const out = []
  let depth = 0
  let quote = null
  let cur = ''
  for (const c of body) {
    if (quote) {
      if (c === quote) quote = null
    } else if (c === "'" || c === '"' || c === '`') quote = c
    else if (c === '(') depth++
    else if (c === ')') depth--
    else if (c === ',' && depth === 0) {
      out.push(cur.trim())
      cur = ''
      continue
    }
    cur += c
  }
  if (cur.trim()) out.push(cur.trim())
  return out
}

function unifiedDiff(a, b) {
  const dir = mkdtempSync(join(tmpdir(), 'pull-sql-'))
  try {
    writeFileSync(join(dir, 'a.sql'), a)
    writeFileSync(join(dir, 'b.sql'), b)
    const run = spawnSync('diff', ['-u', '--label', 'source.sql', '--label', 'other.sql', join(dir, 'a.sql'), join(dir, 'b.sql')], { encoding: 'utf8', maxBuffer: 64 * 1024 * 1024 })
    return run.stdout
  } finally {
    rmSync(dir, { recursive: true, force: true })
  }
}
