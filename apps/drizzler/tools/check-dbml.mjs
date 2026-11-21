#!/usr/bin/env node
// Parses a schema.dbml drizzler wrote with @dbml/core, the parser dbdiagram.io uses.
//
//   check-dbml.mjs <schema.dbml> [--expect-failure]
//
// With --expect-failure the parse has to fail: the example has a known drizzler issue (its
// README's "Known issues"), and a pass means the issue is fixed and the flag can go.
import { readFileSync } from 'node:fs'
import { Parser } from '@dbml/core'

const [file, flag] = process.argv.slice(2)
const expectFailure = flag === '--expect-failure'
let messages = []
try {
  new Parser().parse(readFileSync(file, 'utf8'), 'dbmlv2')
} catch (error) {
  messages = [...new Set((error.diags ?? [error]).map((d) => d.message))]
}
for (const message of messages) console.log(message)
if (expectFailure && messages.length === 0) {
  console.error(`${file} parses: the known issue is fixed, drop dbml_known_issue from BUILD.bazel`)
  process.exit(1)
}
process.exit(!expectFailure && messages.length > 0 ? 1 : 0)
