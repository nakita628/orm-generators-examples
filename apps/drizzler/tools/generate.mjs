#!/usr/bin/env node
// Runs drizzler in one example and writes every format it has into one directory:
//
//   generate.mjs <out dir> [<schema>]
//
// The same five commands as each example's `pnpm generate`: the table docs (README.md, er.svg,
// tables/*.svg), er.png, er-dark.png, schema.dbml and er.md. The working directory is the
// example (js_run_binary's chdir), where drizzler finds drizzle.config.ts, and drizzler is the
// version the example's own package.json resolves to.
//
// Run it twice and the output is the same: see deterministic.mjs.
import { spawnSync } from 'node:child_process'
import { createRequire } from 'node:module'
import { join, resolve } from 'node:path'
import { pathToFileURL } from 'node:url'

const [out, schema] = process.argv.slice(2)
if (!out) throw new Error('usage: generate.mjs <out dir> [<schema>]')

const require = createRequire(join(process.cwd(), 'package.json'))
const pkg = require('drizzler/package.json')
const bin = typeof pkg.bin === 'string' ? pkg.bin : Object.values(pkg.bin)[0]
const cli = resolve(require.resolve('drizzler/package.json'), '..', bin)

// Fixed-seed randomness for the $defaultFn values in the example queries; see deterministic.mjs.
const preload = pathToFileURL(join(import.meta.dirname, 'deterministic.mjs')).href

const runs = [
  [out],
  [join(out, 'er.png')],
  [join(out, 'er-dark.png'), '--theme', 'dark'],
  [join(out, 'schema.dbml')],
  [join(out, 'er.md')],
]
for (const [target, ...options] of runs) {
  const args = ['--import', preload, cli, ...(schema ? [schema] : []), '-o', target, ...options]
  const result = spawnSync(process.execPath, args, { stdio: 'inherit' })
  if (result.status !== 0) {
    console.error(`drizzler ${args.slice(3).join(' ')} exited with ${result.status ?? result.signal}`)
    process.exit(result.status ?? 1)
  }
}
