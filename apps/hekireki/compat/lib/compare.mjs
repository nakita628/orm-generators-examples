// Holds what a harness read back to what the plan wrote. A harness only reports; every judgement
// is made here, the same way for every ORM.
import { isDeepStrictEqual } from 'node:util'

const CHECKS = ['model', 'insert', 'read', 'update', 'relations', 'delete']

export function compare(plan, out) {
  const findings = []
  const counts = Object.fromEntries(CHECKS.map((c) => [c, { ok: 0, total: 0 }]))
  const tally = (check, ok) => {
    counts[check].total++
    if (ok) counts[check].ok++
  }
  const byModel = new Map(out.models.map((r) => [r.model, r]))
  const rows = new Map()
  const pending = []

  for (const m of plan.models) {
    const r = byModel.get(m.name) ?? { missing: true }
    const add = (check, text, detail) => findings.push({ model: m.name, check, text, ...(detail ? { detail } : {}) })
    const mapped = !r.missing && !r.unmapped?.length
    tally('model', mapped)
    if (r.missing) {
      add('model', `no model maps the table \`${m.table}\``)
      continue
    }
    for (const col of r.unmapped ?? []) add('model', `no attribute maps the column \`${col}\``)
    if (r.skipped) {
      add('insert', `not inserted: ${r.skipped}`)
      continue
    }

    const orphan = !r.insert?.ok && m.fields.find((f) => f.fk?.available && f.required && !rows.has(f.fk.model))
    if (orphan) {
      tally('insert', false)
      add('insert', `not inserted: the parent row (\`${orphan.fk.model}\`) did not go in`)
      continue
    }
    tally('insert', r.insert?.ok)
    if (!r.insert?.ok) {
      add('insert', 'insert failed', r.insert?.error)
      continue
    }
    tally('read', r.read?.ok)
    if (!r.read?.ok) {
      add('read', 'read back failed', r.read?.error)
      continue
    }
    const values = r.read.values
    rows.set(m.name, values)
    let readOk = true
    for (const f of m.fields) {
      if (!(f.column in values)) continue
      const got = values[f.column]
      if (f.write === 'value') {
        const problem = differs(f, f.value, got)
        if (problem) {
          readOk = false
          add('read', `\`${f.column}\` (${typeName(f)}) ${problem}`)
        }
      } else if (f.write === 'parent') {
        // A parent that did not go in leaves nothing to hold the key to (a database without
        // foreign keys takes whatever the ORM wrote).
        if (f.fk.available && !rows.has(f.fk.model)) continue
        const parent = f.fk.available ? rows.get(f.fk.model) : null
        const expected = parent ? parent[f.fk.column] : null
        const problem = differs(f, expected, got)
        if (problem) {
          readOk = false
          add('read', `key \`${f.column}\` ${problem}`)
        }
      } else if (f.write === 'orm' && f.required && got == null) {
        readOk = false
        add('read', `\`${f.column}\` was left null: the ORM did not fill its default (${f.updatedAt ? '@updatedAt' : `${f.default}()`})`)
      } else if (f.write === 'orm' && got != null) {
        const problem = badDefault(f, got)
        if (problem) {
          readOk = false
          add('read', `\`${f.column}\` default ${problem}`)
        }
      }
    }
    for (const t of r.read.typeMismatches ?? []) {
      readOk = false
      add('read', `\`${t.column}\` is typed \`${t.declared}\` but came back as ${t.got}`)
    }
    for (const note of r.notes ?? []) add('types', note)
    if (!readOk) counts.read.ok--

    if (m.update) {
      const f = m.fields.find((x) => x.column === m.update.column)
      const ok = r.update?.ok && !differs(f, m.update.value, r.update.value)
      tally('update', ok)
      if (!r.update?.ok) add('update', `update of \`${m.update.column}\` failed`, r.update?.error)
      else if (!ok) add('update', `\`${m.update.column}\` ${differs(f, m.update.value, r.update.value)} after the update`)
    }

    pending.push({ m, r, add })

    if (r.delete) {
      tally('delete', r.delete.ok)
      if (!r.delete.ok) add('delete', 'delete failed', r.delete.error)
    }
  }
  // Relations, once every row is known: a relation has to load a row exactly when the row on the
  // other side holds the matching key.
  for (const { m, r, add } of pending) {
    const mine = rows.get(m.name)
    for (const rel of r.relations ?? []) {
      if (rel.skipped) continue
      if (!rel.found) {
        tally('relations', false)
        add('relations', `no relation for \`${rel.name}\` (${rel.kind})`)
        continue
      }
      const def = m.relations.find((x) => x.name === rel.name)
      const other = rows.get(def.model)
      const expected =
        Boolean(mine && other) &&
        def.columns.length > 0 &&
        def.columns.every((c, i) => mine[c] != null && String(mine[c]) === String(other[def.references[i]]))
      const loaded = rel.loaded ?? 0
      const ok = !rel.error && (expected ? loaded > 0 : loaded === 0)
      tally('relations', ok)
      if (rel.error) add('relations', `relation \`${rel.name}\` failed to load`, rel.error)
      else if (!ok) add('relations', expected ? `relation \`${rel.name}\` loaded nothing, though the related row is there` : `relation \`${rel.name}\` loaded ${loaded} row(s), though none matches`)
    }
  }
  return { counts, findings }
}

function typeName(f) {
  return `${f.type === 'Enum' ? f.enum : f.type}${f.list ? '[]' : ''}${f.native ? ` @db.${f.native[0]}` : ''}`
}

// Null when the value read back is the value written; otherwise what is wrong with it.
export function differs(f, expected, got) {
  if (expected == null && got == null) return null
  if (expected == null) return `came back as ${show(got)}, expected null`
  if (got == null) return `came back null, expected ${show(expected)}`
  if (f.list) {
    if (!Array.isArray(got)) return `came back as ${show(got)}, not a list`
    if (got.length !== expected.length) return `came back as ${show(got)}, expected ${show(expected)}`
    for (let i = 0; i < got.length; i++) {
      const p = differs({ ...f, list: false }, expected[i], got[i])
      if (p) return p
    }
    return null
  }
  const mismatch = () => `came back as ${show(got)}, expected ${show(expected)}`
  switch (f.type) {
    case 'DateTime': {
      if (typeof got !== 'string') return `came back as ${typeof got} ${show(got)}, not a date`
      const a = Date.parse(expected)
      const b = Date.parse(got)
      if (Number.isNaN(b)) return mismatch()
      if (a === b) return null
      if (Math.abs(a - b) % 3_600_000 === 0) return `${mismatch()}: shifted by ${(b - a) / 3_600_000}h (a time zone)`
      return mismatch()
    }
    case 'Int':
    case 'Float':
      return typeof got === 'number' && got === expected ? null : mismatch()
    case 'BigInt':
      return String(got) === String(expected) ? (typeof got === 'string' || typeof got === 'number' ? null : mismatch()) : mismatch()
    case 'Decimal':
      return Number(got) === Number(expected) ? null : mismatch()
    case 'Boolean':
      return got === expected ? null : mismatch()
    case 'Json': {
      if (typeof got === 'string') {
        try {
          if (isDeepStrictEqual(JSON.parse(got), expected)) return 'came back as a JSON string, not the parsed value'
        } catch {}
        return mismatch()
      }
      return isDeepStrictEqual(got, expected) ? null : mismatch()
    }
    case 'String':
      if (f.native?.[0] === 'Char' && typeof got === 'string' && got.trimEnd() === String(expected).trimEnd()) return null
      return got === expected ? null : mismatch()
    default:
      return got === expected ? null : mismatch()
  }
}

// A default the ORM filled, held to what Prisma would have written.
function badDefault(f, got) {
  if (f.type === 'DateTime' && (f.default === 'now' || f.updatedAt)) {
    const t = Date.parse(got)
    if (Number.isNaN(t)) return `is ${show(got)}, not a date`
    const drift = Math.abs(Date.now() - t)
    const allowed = f.native?.[0] === 'Date' ? 26 * 3_600_000 : 10 * 60_000 // a date drops the time of day
    if (drift > allowed) return `is ${got}, ${Math.round(drift / 3_600_000)}h away from now (a time zone)`
  }
  if (f.default === 'uuid' && typeof got === 'string' && !/^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$/i.test(got)) return `is ${show(got)}, not a UUID`
  if (f.default === 'cuid' && typeof got === 'string' && !/^[a-z0-9]{20,32}$/.test(got)) return `is ${show(got)}, not a cuid`
  return null
}

function show(v) {
  const s = JSON.stringify(v)
  return s.length > 80 ? s.slice(0, 77) + '...' : s
}
