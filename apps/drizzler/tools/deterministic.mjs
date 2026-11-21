// Preloaded into every drizzler process (node --import) by generate.mjs.
//
// drizzler calls each column's $defaultFn to fill the example queries of the table docs, and a
// default such as cuid2's createId() or a faker slug draws on Math.random and node:crypto: the
// docs would change on every run, and the committed copy could never be held to a fresh one.
// Both are replaced here with one fixed-seed generator (mulberry32), so a run is reproducible.
import crypto from 'node:crypto'
import { syncBuiltinESMExports } from 'node:module'

let state = 0x5eed1234
function next() {
  state = (state + 0x6d2b79f5) | 0
  let t = state
  t = Math.imul(t ^ (t >>> 15), t | 1)
  t ^= t + Math.imul(t ^ (t >>> 7), t | 61)
  return ((t ^ (t >>> 14)) >>> 0) / 4294967296
}

function fill(view) {
  const bytes = new Uint8Array(view.buffer, view.byteOffset, view.byteLength)
  for (let i = 0; i < bytes.length; i++) bytes[i] = Math.floor(next() * 256)
  return view
}

function uuid() {
  const b = fill(new Uint8Array(16))
  b[6] = (b[6] & 0x0f) | 0x40
  b[8] = (b[8] & 0x3f) | 0x80
  const h = Buffer.from(b).toString('hex')
  return `${h.slice(0, 8)}-${h.slice(8, 12)}-${h.slice(12, 16)}-${h.slice(16, 20)}-${h.slice(20)}`
}

// Some of these are getters or prototype methods (webcrypto's), so each is defined on the object
// itself rather than assigned. node:crypto's getRandomValues cannot be redefined; it calls
// webcrypto's, which is replaced below.
function replace(object, key, value) {
  const own = Object.getOwnPropertyDescriptor(object, key)
  if (own && !own.configurable) return
  Object.defineProperty(object, key, { value, configurable: true, writable: true, enumerable: true })
}

replace(Math, 'random', next)
replace(crypto, 'getRandomValues', fill)
replace(crypto, 'randomFillSync', (buffer, offset = 0, size = buffer.byteLength - offset) => {
  fill(new Uint8Array(buffer.buffer ?? buffer, (buffer.byteOffset ?? 0) + offset, size))
  return buffer
})
replace(crypto, 'randomBytes', (size, callback) => {
  const buffer = fill(Buffer.alloc(size))
  if (callback) {
    callback(null, buffer)
    return undefined
  }
  return buffer
})
replace(crypto, 'randomUUID', uuid)
for (const webcrypto of new Set([crypto.webcrypto, globalThis.crypto])) {
  replace(webcrypto, 'getRandomValues', fill)
  replace(webcrypto, 'randomUUID', uuid)
}

// So that `import { randomBytes } from 'node:crypto'` sees the replacements too.
syncBuiltinESMExports()
