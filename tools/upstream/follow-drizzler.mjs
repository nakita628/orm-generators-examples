#!/usr/bin/env node
// Follows one drizzler example to a newer upstream commit with a 3-way merge per file, so that the
// example's "Changes from upstream" are kept without being re-applied by hand.
//
//   node tools/upstream/follow-drizzler.mjs <example> <newCommit> [--exact] [--apply]
//
// For every file under the README's `- Source:` links: base = the upstream file at the pinned
// commit, theirs = the same file at <newCommit>, ours = the example's copy. Without --apply it
// only reports what would happen: `merged` (clean), `conflicts` (the file keeps git's markers once
// applied: resolve them by hand), `added` and `deleted` upstream files, `unmapped` ones the example
// does not copy and `unmappedChanged` among those that changed (the tests of a package, usually).
// With --apply it writes the merged files, copies added files as upstream has them, and moves
// every pinned commit of the README to <newCommit>; then re-apply the README's conventions to the
// added files (the `relations` import, ...), and run docs.write and the tests as the skill says.
//
// A copied upstream file is found in the example by the longest common path suffix. Examples that
// keep upstream's layout (openstatus, survev: no src/) need --exact, which maps a path to itself.
// <newCommit> is `head` of the example in `node tools/upstream/check.mjs --json report.json`.
import { execFileSync, spawnSync } from 'node:child_process'
import { mkdtempSync, readFileSync, writeFileSync, existsSync, mkdirSync, rmSync, readdirSync, statSync } from 'node:fs'
import { join, dirname, basename } from 'node:path'
import { tmpdir } from 'node:os'
import { fileURLToPath } from 'node:url'

const [name, newCommit, ...flags] = process.argv.slice(2)
const apply = flags.includes('--apply')
const exact = flags.includes('--exact')
const root = join(dirname(fileURLToPath(import.meta.url)), '..', '..')
const exDir = join(root, 'apps/drizzler/examples', name)
const readmePath = join(exDir, 'README.md')
const readme = readFileSync(readmePath, 'utf8')
const LINK_RE = /\(https:\/\/github\.com\/([^/]+\/[^/]+)\/(blob|tree)\/([0-9a-f]{40})\/([^)\s]+)\)/g
const line = /^- Source: .*$/m.exec(readme)?.[0] ?? ''
const links = [...line.matchAll(LINK_RE)].map((m) => ({ repo: m[1], kind: m[2], commit: m[3], path: m[4] }))
if (!links.length) throw new Error('no source links')
const repo = links[0].repo
const oldCommits = [...new Set(links.map((l) => l.commit))]
if (oldCommits.length !== 1) console.error('WARNING: several old commits', oldCommits)
const oldCommit = oldCommits[0]

const work = mkdtempSync(join(tmpdir(), 'follow-'))
process.on('exit', () => rmSync(work, { recursive: true, force: true }))
const g = (...a) => execFileSync('git', ['-C', work, ...a], { encoding: 'utf8', maxBuffer: 1 << 28 })
execFileSync('git', ['init', '-q', '--bare', work])
g('remote', 'add', 'origin', `https://github.com/${repo}`)
g('config', 'remote.origin.promisor', 'true')
g('config', 'remote.origin.partialclonefilter', 'blob:none')
g('fetch', '-q', '--depth', '1', '--filter=blob:none', 'origin', oldCommit, newCommit)

function listFiles(commit, path, kind) {
  const p = path.replace(/\/$/, '')
  const out = g('ls-tree', '-r', '--name-only', commit, '--', p)
  return out.split('\n').filter(Boolean)
}
function blobId(commit, path) {
  try { return g('rev-parse', `${commit}:${path}`).trim() } catch { return null }
}
function blob(commit, path) {
  return execFileSync('git', ['-C', work, 'cat-file', 'blob', `${commit}:${path}`], { maxBuffer: 1 << 28 })
}

// local files
const srcRoot = exact ? exDir : join(exDir, 'src')
function walk(d, base = '') {
  const out = []
  for (const e of readdirSync(d)) {
    const rel = base ? `${base}/${e}` : e
    const full = join(d, e)
    if (statSync(full).isDirectory()) out.push(...walk(full, rel))
    else out.push(rel)
  }
  return out
}
const localFiles = walk(srcRoot).filter((f) => !exact || !(f.startsWith('node_modules/') || f.startsWith('docs/')))

function suffixLen(a, b) {
  const x = a.split('/').reverse(), y = b.split('/').reverse()
  let n = 0
  while (n < x.length && n < y.length && x[n] === y[n]) n++
  return n
}
function mapToLocal(upPath) {
  if (exact) return localFiles.includes(upPath) ? { local: upPath } : null
  let best = null, bestN = 0, tie = false
  for (const lf of localFiles) {
    const n = suffixLen(upPath, lf)
    if (n > bestN) { best = lf; bestN = n; tie = false }
    else if (n === bestN && n > 0) tie = true
  }
  if (!best) return null
  if (tie) return { local: best, ambiguous: true }
  return { local: best }
}

const oldFiles = new Map() // up path -> local path
const newFilesAll = new Set()
const report = { merged: [], same: [], added: [], deleted: [], unmapped: [], conflicts: [], localOnlyChange: [] }
const dirMap = new Map() // upstream dir -> local dir

for (const l of links) {
  const o = listFiles(oldCommit, l.path, l.kind)
  const n = listFiles(newCommit, l.path, l.kind)
  for (const f of o) {
    const m = mapToLocal(f)
    if (!m) { report.unmapped.push(f); continue }
    if (m.ambiguous) console.error('ambiguous mapping', f, m.local)
    oldFiles.set(f, m.local)
    dirMap.set(dirname(f), dirname(m.local))
  }
  for (const f of n) newFilesAll.add(f)
}

const tmp = join(work, '_m')
mkdirSync(tmp, { recursive: true })

for (const [up, local] of oldFiles) {
  const ob = blobId(oldCommit, up)
  const nb = blobId(newCommit, up)
  const localPath = join(srcRoot, local)
  if (nb === null) { report.deleted.push({ up, local }); continue }
  if (ob === nb) { report.same.push(up); continue }
  const base = join(tmp, 'base'), theirs = join(tmp, 'theirs'), ours = join(tmp, 'ours')
  writeFileSync(base, blob(oldCommit, up))
  writeFileSync(theirs, blob(newCommit, up))
  const oursBuf = readFileSync(localPath)
  writeFileSync(ours, oursBuf)
  const same = Buffer.compare(oursBuf, readFileSync(base)) === 0
  if (same) {
    if (apply) writeFileSync(localPath, readFileSync(theirs))
    report.merged.push({ up, local, how: 'copy' })
    continue
  }
  const r = spawnSync('git', ['merge-file', '-p', '-L', 'local', '-L', 'upstream-old', '-L', 'upstream-new', ours, base, theirs], { maxBuffer: 1 << 28 })
  if (r.status < 0 || r.status === null) throw new Error('merge-file failed for ' + up)
  if (apply) writeFileSync(localPath, r.stdout)
  if (r.status > 0) report.conflicts.push({ up, local, n: r.status })
  else report.merged.push({ up, local, how: '3-way' })
}

for (const f of newFilesAll) {
  if (oldFiles.has(f) || report.unmapped.includes(f)) continue
  // new upstream file (or one that was unmapped before)
  const d = dirname(f)
  const ld = exact ? d : dirMap.get(d)
  const skip = exact && /(\.test\.ts|\/__tests__\/|package\.json|tsconfig\.json|vitest\.config\.ts|README\.md)$/.test(f)
  report.added.push({ up: f, localDir: ld ?? null, local: !skip && ld != null ? join(ld, basename(f)) : null })
}

// Which of the "unmapped" files still exist upstream and changed?
const unmappedChanged = report.unmapped.filter((f) => blobId(oldCommit, f) !== blobId(newCommit, f))

console.log(JSON.stringify({ name, repo, oldCommit, newCommit, apply, ...Object.fromEntries(Object.entries(report).map(([k, v]) => [k, k === 'same' ? v.length : v])), unmappedChanged }, null, 2))

if (apply) {
  // Fetch the added ones raw as-is into src/ (the caller re-applies local conventions by hand).
  for (const a of report.added) {
    if (!a.local) continue
    const dest = join(srcRoot, a.local)
    mkdirSync(dirname(dest), { recursive: true })
    writeFileSync(dest, blob(newCommit, a.up))
  }
  // README: every old commit -> new commit
  writeFileSync(readmePath, readme.split(oldCommit).join(newCommit))
}
