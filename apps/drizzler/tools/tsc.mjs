#!/usr/bin/env node
// `tsc -p tsconfig.json` in one example (js_test's chdir), with the TypeScript its package.json
// resolves to.
import { spawnSync } from 'node:child_process'
import { createRequire } from 'node:module'
import { join } from 'node:path'

const require = createRequire(join(process.cwd(), 'package.json'))
const tsc = require.resolve('typescript/bin/tsc')
const result = spawnSync(process.execPath, [tsc, '-p', 'tsconfig.json'], { stdio: 'inherit' })
process.exit(result.status ?? 1)
