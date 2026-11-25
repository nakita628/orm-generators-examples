// Turns a Prisma schema into the plan every ORM harness runs: the models in an order where each
// one's foreign keys point at rows inserted before it, and for every column what the harness
// writes, what it leaves to the ORM, and how the value reads back.
//
// The plan speaks in table and column names, never in an ORM's attribute names: each harness
// finds the attribute that maps to a column through its ORM's own metadata, so a column the ORM
// does not map at all shows up as a finding instead of a crash.
import { getDMMF } from '@prisma/get-dmmf'

/** @param {string} source the schema.prisma text */
export async function makePlan(source) {
  // Line comments out first: a schema may keep another datasource commented out above its own.
  const code = source.replace(/\/\/.*$/gm, '')
  const provider = /datasource\s+\w+\s*\{[^}]*provider\s*=\s*"(\w+)"/s.exec(code)?.[1]
  if (!provider) throw new Error('no datasource provider in the schema')
  const dialect = { postgresql: 'postgresql', postgres: 'postgresql', mysql: 'mysql', sqlite: 'sqlite' }[provider]
  if (!dialect) throw new Error(`unsupported provider ${provider}`)

  const { datamodel } = await getDMMF({ datamodel: source })
  // A view is in the DMMF's models with nothing to tell it apart; it is read-only, so it is left out.
  const views = new Set([...code.matchAll(/^\s*view\s+(\w+)\s*\{/gm)].map((m) => m[1]))
  datamodel.models = datamodel.models.filter((m) => !views.has(m.name))
  const enums = Object.fromEntries(datamodel.enums.map((e) => [e.name, e.values.map((v) => v.dbName ?? v.name)]))
  const byName = new Map(datamodel.models.map((m) => [m.name, m]))
  const table = (m) => m.dbName ?? m.name
  const column = (m, fieldName) => {
    const f = m.fields.find((x) => x.name === fieldName)
    return f.dbName ?? f.name
  }

  const models = datamodel.models.map((m) => {
    const fkOf = new Map()
    const relations = []
    for (const f of m.fields.filter((x) => x.kind === 'object')) {
      const target = byName.get(f.type)
      if (f.relationFromFields?.length) {
        f.relationFromFields.forEach((from, i) =>
          fkOf.set(from, { model: f.type, column: column(target, f.relationToFields[i]) }),
        )
        relations.push({
          name: f.name,
          kind: 'toOne',
          model: f.type,
          columns: f.relationFromFields.map((x) => column(m, x)),
          references: f.relationToFields.map((x) => column(target, x)),
          required: f.isRequired,
        })
      } else {
        // The other side holds the key: find it to know which columns join the two.
        const back = target.fields.find(
          (x) => x.kind === 'object' && x.relationName === f.relationName && x.relationFromFields?.length && x !== f,
        )
        if (!back) continue // an implicit many-to-many: Prisma's own join table, which no ORM maps
        relations.push({
          name: f.name,
          kind: f.isList ? 'toMany' : 'toOneBack',
          model: f.type,
          columns: back.relationToFields.map((x) => column(m, x)),
          references: back.relationFromFields.map((x) => column(target, x)),
          required: false,
        })
      }
    }
    // The row's key: the primary key, or for a model without one (Prisma allows it with a unique
    // instead) the first unique, which is how Prisma Client addresses the row too.
    let pk = m.primaryKey?.fields ?? m.fields.filter((f) => f.isId).map((f) => f.name)
    if (pk.length === 0) pk = m.uniqueFields?.[0] ?? m.fields.filter((f) => f.isUnique).slice(0, 1).map((f) => f.name)
    const fields = m.fields
      .filter((f) => f.kind !== 'object')
      .map((f) => {
        const fk = fkOf.get(f.name) ?? null
        const defaultKind = !f.hasDefaultValue
          ? null
          : typeof f.default === 'object' && f.default !== null && !Array.isArray(f.default)
            ? f.default.name // autoincrement, now, cuid, uuid, dbgenerated, ...
            : 'value'
        return {
          name: f.name,
          column: f.dbName ?? f.name,
          type: f.kind === 'enum' ? 'Enum' : f.type,
          enum: f.kind === 'enum' ? f.type : undefined,
          list: f.isList,
          required: f.isRequired,
          id: pk.includes(f.name),
          primary: Boolean(m.primaryKey?.fields?.includes(f.name) || f.isId),
          unique: f.isUnique,
          default: defaultKind,
          updatedAt: f.isUpdatedAt,
          native: f.nativeType ?? null,
          fk,
        }
      })
    return { name: m.name, table: table(m), pk: pk.map((x) => column(m, x)), fields, relations }
  })

  const ordered = order(models)
  // A column of a type Prisma does not support (`Unsupported("tsvector")`, ...) is not in the DMMF,
  // so no harness knows to write it. A model with a required one and no default cannot be
  // inserted (Prisma Client cannot create it either), and neither can a model whose required key
  // points at it: both are left out, as a model whose key closes a cycle is.
  const unsupported = new Map()
  for (const [, name, body] of code.matchAll(/^\s*model\s+(\w+)\s*\{([^]*?)^\s*\}/gm)) {
    for (const [, field, type, list, optional, rest] of body.matchAll(/^[ \t]*(\w+)[ \t]+(Unsupported\("[^"]*"\))(\[\])?(\?)?([^\n]*)$/gm)) {
      if (!optional && !/@default\(/.test(rest)) unsupported.set(name, [...(unsupported.get(name) ?? []), `\`${field}\` ${type}${list ?? ''}`])
    }
  }
  const byPlanName = new Map(ordered.map((m) => [m.name, m]))
  for (const m of ordered) {
    if (unsupported.has(m.name)) {
      m.insertable = false
      m.skipReason = `a required column of a type Prisma does not support, without a default (${unsupported.get(m.name).join(', ')})`
      continue
    }
    const parent = m.fields.find((f) => f.fk && f.required && f.fk.model !== m.name && byPlanName.get(f.fk.model)?.skipReason)
    if (parent) {
      m.insertable = false
      m.skipReason = `its required key points at \`${parent.fk.model}\`, which is not inserted`
    }
  }
  // Columns another model's key points at: changing one after the insert would leave the rows
  // inserted later pointing at a value that is gone.
  const referenced = new Set(models.flatMap((m) => m.fields.filter((f) => f.fk).map((f) => `${f.fk.model}.${f.fk.column}`)))
  let n = 0
  for (const m of ordered) {
    for (const f of m.fields) {
      // What the harness writes. Left out: a key filled from a parent row, and any column with a
      // default or @updatedAt, which is the ORM's to fill (the client-side ones, cuid() and now()
      // and @updatedAt, are not in the database: Prisma Client fills them, so the ORM has to).
      if (f.fk) f.write = 'parent'
      else if (f.default || f.updatedAt) f.write = 'orm'
      else if (f.type === 'Unsupported') f.write = 'skip'
      else {
        f.write = 'value'
        f.value = sample(f, enums, ++n, dialect)
      }
    }
    // One column the harness changes after the insert, to see an update go through.
    const target = m.fields.find(
      (f) =>
        f.write === 'value' &&
        !f.id &&
        !f.unique &&
        !f.list &&
        !referenced.has(`${m.name}.${f.column}`) &&
        f.native?.[0] !== 'Char' &&
        ['String', 'Int', 'Boolean'].includes(f.type),
    )
    m.update = target ? { column: target.column, value: updated(target, enums, dialect) } : null
  }
  return { dialect, enums, models: ordered }
}

// Parents first. A required key to another model is a hard edge; an optional one is kept when it
// can be, and left null when it closes a cycle (or points at the model itself).
function order(models) {
  const byName = new Map(models.map((m) => [m.name, m]))
  const out = []
  const state = new Map()
  const visit = (m) => {
    if (state.get(m.name) === 'done') return
    state.set(m.name, 'visiting')
    for (const f of m.fields.filter((x) => x.fk)) {
      if (f.fk.model === m.name) continue
      const s = state.get(f.fk.model)
      if (s === 'visiting') continue
      visit(byName.get(f.fk.model))
    }
    state.set(m.name, 'done')
    out.push(m)
  }
  for (const m of models) visit(m)
  const position = new Map(out.map((m, i) => [m.name, i]))
  for (const m of out) {
    for (const f of m.fields.filter((x) => x.fk)) {
      // A key whose parent row is not there yet: written as null, which only an optional one can be.
      f.fk.available = position.get(f.fk.model) < position.get(m.name)
    }
    m.insertable = m.fields.every((f) => !f.fk || f.fk.available || !f.required)
  }
  return out
}

// A value of the column's type. Typed so that each harness can turn it into its language's own
// value: DateTime is an ISO string, BigInt and Decimal are strings, Bytes is base64.
function sample(f, enums, n, dialect) {
  const one = () => {
    const [nativeName, nativeArgs] = f.native ?? []
    switch (f.type) {
      case 'String': {
        if (nativeName === 'Uuid') return `00000000-0000-4000-8000-${String(n).padStart(12, '0')}`
        if (nativeName === 'Inet') return `10.0.0.${n % 250}`
        const size = nativeArgs?.[0] ? Number(nativeArgs[0]) : null
        const text = `v${n}`
        if (nativeName === 'Char' && size) return text.padEnd(size, 'x').slice(0, size)
        return size ? text.slice(0, size) : text
      }
      case 'Int':
        return n
      case 'BigInt':
        return String(9007199254740000 + n)
      case 'Float':
        return n + 0.5
      case 'Decimal': {
        // As many fraction digits as the column keeps: @db.Decimal(p, s) rounds to s.
        const scale = nativeArgs?.[1] !== undefined ? Number(nativeArgs[1]) : 2
        return scale === 0 ? String(n) : `${n}.${'25'.padEnd(scale, '0').slice(0, scale)}`
      }
      case 'Boolean':
        return true
      case 'DateTime':
        if (nativeName === 'Date') return `2024-01-${String((n % 27) + 1).padStart(2, '0')}T00:00:00.000Z`
        if (nativeName === 'Time' || nativeName === 'Timetz') return '1970-01-01T12:34:56.000Z'
        // A precision below 3 (@db.Timestamp(0), @db.Timestamptz(1)) keeps fewer digits.
        {
          const precision = nativeArgs?.[0] !== undefined ? Math.min(3, Number(nativeArgs[0])) : 3
          const ms = String(n % 1000).padStart(3, '0').slice(0, precision).padEnd(3, '0')
          return `2024-01-02T03:04:05.${ms}Z`
        }
      case 'Json':
        return { n, tags: ['a', 'b'], nested: { ok: true } }
      case 'Bytes':
        return Buffer.from(`bytes-${n}`).toString('base64')
      case 'Enum':
        return enums[f.enum][0]
      default:
        throw new Error(`no sample for ${f.type}`)
    }
  }
  if (f.list) return dialect === 'postgresql' ? [one()] : null
  return one()
}

function updated(f, enums, dialect) {
  switch (f.type) {
    case 'String':
      // A column that only takes text of its own kind: a UUID, an address.
      if (f.native?.[0] === 'Uuid') return '00000000-0000-4000-8000-ffffffffffff'
      if (f.native?.[0] === 'Inet') return '10.9.9.9'
      return 'u1'
    case 'Int':
      return 4242
    case 'Boolean':
      return false
  }
}
