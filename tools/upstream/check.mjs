#!/usr/bin/env node
// Reports what the examples are behind on: the upstream projects their schemas come from, and the
// generators they run. The daily upstream sync (tools/upstream/README.md) starts from this report.
//
//   node tools/upstream/check.mjs                       every app, a Markdown report on stdout
//   node tools/upstream/check.mjs --app hekireki        one app (hekireki, drizzler, djangoer, generators)
//   node tools/upstream/check.mjs --name calcom,dub      only these examples
//   node tools/upstream/check.mjs --json report.json    also write the report as JSON
//   node tools/upstream/check.mjs --apply-hekireki      move sources.json to each changed schema's new
//                                                       commit and re-import it (import-schema.mjs)
//
// What each app follows:
//   - hekireki: apps/hekireki/schemas/sources.json pins each schema to a commit. A schema is behind
//     when the file (or the directory, for a multi-file schema) at the default branch's HEAD differs
//     from the one at the pinned commit; a commit that leaves it untouched is not reported.
//   - drizzler: the `- Source:` link of each apps/drizzler/examples/<name>/README.md, the same way.
//   - djangoer: each examples/<name>/uv.lock; an example is behind when `uv lock --upgrade` would
//     move one of the packages its pyproject.toml names (the project, Django, djangoer, ...).
//     Transitive-only updates are listed but not counted.
//   - generators: hekireki and drizzler on npm, djangoer on PyPI, against the versions pinned here.
//
// Pins named in tools/upstream/blocked.json (a commit or version that a generator cannot handle
// yet, with the issue filed for it) are reported as `blocked` instead of `behind`.
import { execFile as execFileCb } from 'node:child_process'
import { existsSync, mkdtempSync, readFileSync, readdirSync, rmSync, writeFileSync } from 'node:fs'
import { tmpdir } from 'node:os'
import { join, resolve } from 'node:path'
import { promisify } from 'node:util'

const execFile = promisify(execFileCb)
const root = resolve(import.meta.dirname, '../..')
const args = process.argv.slice(2)
const flag = (name) => {
  const i = args.indexOf(name)
  return i === -1 ? undefined : (args[i + 1] ?? '')
}
const onlyApp = flag('--app')
const jsonOut = flag('--json')
const onlyNames = flag('--name')?.split(',')
const applyHekireki = args.includes('--apply-hekireki')
const wantName = (name) => !onlyNames || onlyNames.includes(name)
const want = (app) => !onlyApp || onlyApp === app

const blocked = existsSync(join(root, 'tools/upstream/blocked.json'))
  ? JSON.parse(readFileSync(join(root, 'tools/upstream/blocked.json'), 'utf8'))
  : []
const blockedBy = (app, name, target) =>
  blocked.find((b) => b.app === app && b.name === name && (b.target === target || b.target === '*'))

const run = async (cmd, argv, opts = {}) => {
  const { stdout } = await execFile(cmd, argv, { maxBuffer: 64 * 1024 * 1024, ...opts })
  return stdout.trim()
}

// Runs fn over items, `limit` at a time.
async function pool(items, limit, fn) {
  const out = new Array(items.length)
  let next = 0
  await Promise.all(
    Array.from({ length: Math.min(limit, items.length) }, async () => {
      while (next < items.length) {
        const i = next++
        out[i] = await fn(items[i], i)
      }
    }),
  )
  return out
}

// ---- upstream git pins (hekireki, drizzler) ----------------------------------------------------

const work = mkdtempSync(join(tmpdir(), 'upstream-'))
process.on('exit', () => rmSync(work, { recursive: true, force: true }))

// One bare repository per upstream repo, holding the trees (no blobs) of HEAD and of every pinned
// commit, so that a path's object id can be compared between them.
const repos = new Map()
function upstream(repo) {
  if (!repos.has(repo)) {
    repos.set(
      repo,
      (async () => {
        const url = `https://github.com/${repo}`
        const dir = join(work, repo.replace('/', '__'))
        await run('git', ['init', '-q', '--bare', dir])
        const line = await run('git', ['ls-remote', '--symref', url, 'HEAD'])
        const branch = /ref: refs\/heads\/(\S+)\s+HEAD/.exec(line)?.[1]
        const head = /^([0-9a-f]{40})\s+HEAD$/m.exec(line)?.[1]
        if (!head) throw new Error(`no HEAD for ${repo}`)
        return { url, dir, branch, head, fetched: new Set() }
      })(),
    )
  }
  return repos.get(repo)
}

async function fetchCommits(r, commits) {
  const missing = commits.filter((c) => !r.fetched.has(c))
  if (missing.length === 0) return
  await run('git', ['-C', r.dir, 'fetch', '-q', '--depth', '1', '--filter=blob:none', r.url, ...missing])
  for (const c of missing) r.fetched.add(c)
}

async function objectAt(r, commit, path) {
  try {
    return await run('git', ['-C', r.dir, 'rev-parse', `${commit}:${path.replace(/\/$/, '')}`])
  } catch {
    return null
  }
}

// pins: [{ app, name, repo, path, commit, file }] -> the same, with status, head and new_url.
async function checkPins(pins) {
  const byRepo = Object.groupBy(pins, (p) => p.repo)
  const results = []
  await pool(Object.entries(byRepo), 6, async ([repo, list]) => {
    let r
    try {
      r = await upstream(repo)
      await fetchCommits(r, [r.head, ...new Set(list.map((p) => p.commit))])
    } catch (e) {
      for (const p of list) results.push({ ...p, status: 'error', error: String(e.message ?? e).split('\n')[0] })
      return
    }
    for (const p of list) {
      const base = { ...p, branch: r.branch, head: r.head }
      if (p.commit === r.head) {
        results.push({ ...base, status: 'current' })
        continue
      }
      const [was, now] = await Promise.all([objectAt(r, p.commit, p.path), objectAt(r, r.head, p.path)])
      const newUrl = `https://github.com/${repo}/${p.path.endsWith('/') ? 'tree' : 'blob'}/${r.head}/${p.path}`
      const compare = `https://github.com/${repo}/compare/${p.commit}...${r.head}`
      if (now === null) results.push({ ...base, status: 'moved', compare, note: `${p.path} is gone at HEAD` })
      else if (was === now) results.push({ ...base, status: 'current', note: 'HEAD moved, the schema did not' })
      else {
        const b = blockedBy(p.app, p.name, r.head)
        results.push({ ...base, status: b ? 'blocked' : 'behind', new_url: newUrl, compare, ...(b && { issue: b.issue }) })
      }
    }
  })
  return results.sort((a, b) => a.name.localeCompare(b.name))
}

function hekirekiPins() {
  const sources = JSON.parse(readFileSync(join(root, 'apps/hekireki/schemas/sources.json'), 'utf8'))
  return Object.entries(sources).map(([name, s]) => ({
    app: 'hekireki',
    name,
    repo: s.repo,
    path: s.path,
    commit: s.commit,
    file: `apps/hekireki/schemas/${name}/schema.prisma`,
  }))
}

// Every GitHub link with a commit on the `- Source:` line (schema.ts and relations.ts, ...).
const LINK_RE = /\(https:\/\/github\.com\/([^/]+\/[^/]+)\/(blob|tree)\/([0-9a-f]{40})\/([^)\s]+)\)/g
function drizzlerPins() {
  const dir = join(root, 'apps/drizzler/examples')
  const pins = []
  for (const name of readdirSync(dir)) {
    const readme = join(dir, name, 'README.md')
    if (!existsSync(readme)) continue
    const line = /^- Source: .*$/m.exec(readFileSync(readme, 'utf8'))?.[0] ?? ''
    const links = [...line.matchAll(LINK_RE)]
    if (links.length === 0) {
      pins.push({ app: 'drizzler', name, status: 'error', error: 'no `- Source:` link with a commit in README.md' })
      continue
    }
    for (const [, repo, kind, commit, path] of links) {
      pins.push({ app: 'drizzler', name, repo, commit, path: kind === 'tree' ? path.replace(/\/?$/, '/') : path, file: `apps/drizzler/examples/${name}/README.md` })
    }
  }
  return pins
}

// ---- djangoer: uv lock --upgrade --dry-run ------------------------------------------------------

const UV_VERSION = /^UV_VERSION = "([^"]+)"/m.exec(readFileSync(join(root, 'MODULE.bazel'), 'utf8'))[1]
const normalize = (n) => n.toLowerCase().replace(/[-_.]+/g, '-')

async function checkDjangoer() {
  const dir = join(root, 'apps/djangoer/examples')
  const names = readdirSync(dir).filter((n) => wantName(n) && existsSync(join(dir, n, 'uv.lock')))
  return pool(names, 8, async (name) => {
    const project = join(dir, name)
    const pyproject = readFileSync(join(project, 'pyproject.toml'), 'utf8')
    const deps = /^dependencies = \[(.*?)\]/ms.exec(pyproject)?.[1] ?? ''
    const direct = new Set([...deps.matchAll(/"([A-Za-z0-9_.-]+)/g)].map((m) => normalize(m[1])))
    try {
      // uv writes its plan to stderr.
      const { stderr } = await execFile(
        'uvx',
        ['--quiet', '--from', `uv==${UV_VERSION}`, 'uv', 'lock', '--upgrade', '--dry-run', '--project', project],
        { maxBuffer: 16 * 1024 * 1024 },
      )
      const updates = [...stderr.matchAll(/^(Update|Add|Remove) ([A-Za-z0-9_.-]+) (.*)$/gm)].map(([, op, pkg, rest]) => ({
        op,
        pkg,
        change: rest.trim(),
        direct: direct.has(normalize(pkg)),
      }))
      const target = updates.filter((u) => u.direct).map((u) => `${u.pkg} ${u.change}`).join(', ')
      const b = target && blockedBy('djangoer', name, target)
      const status = updates.some((u) => u.direct) ? (b ? 'blocked' : 'behind') : 'current'
      return { app: 'djangoer', name, status, updates, ...(b && { issue: b.issue }) }
    } catch (e) {
      return { app: 'djangoer', name, status: 'error', error: String(e.stderr || e.message).trim().split('\n').slice(-3).join(' ') }
    }
  })
}

// ---- generators -------------------------------------------------------------------------------

async function checkGenerators() {
  const hekirekiPinned = JSON.parse(readFileSync(join(root, 'apps/hekireki/package.json'), 'utf8')).devDependencies?.hekireki
    ?? JSON.parse(readFileSync(join(root, 'apps/hekireki/package.json'), 'utf8')).dependencies?.hekireki
  const drizzlerPinned = /^\s+drizzler: (\S+)$/m.exec(readFileSync(join(root, 'pnpm-workspace.yaml'), 'utf8'))?.[1]
  // djangoer: the oldest version any example's uv.lock holds.
  const dir = join(root, 'apps/djangoer/examples')
  const locked = new Set()
  for (const n of readdirSync(dir)) {
    const lock = join(dir, n, 'uv.lock')
    if (!existsSync(lock)) continue
    const v = /\[\[package\]\]\nname = "djangoer"\nversion = "([^"]+)"/.exec(readFileSync(lock, 'utf8'))?.[1]
    if (v) locked.add(v)
  }
  const npmLatest = (pkg) => run('npm', ['view', pkg, 'version'])
  const pypiLatest = async (pkg) => {
    const out = await run('curl', ['-fsS', `https://pypi.org/pypi/${pkg}/json`])
    return JSON.parse(out).info.version
  }
  const rows = [
    { name: 'hekireki', registry: 'npm', pinned: hekirekiPinned, latest: await npmLatest('hekireki').catch((e) => `error: ${e.message}`) },
    { name: 'drizzler', registry: 'npm', pinned: drizzlerPinned, latest: await npmLatest('drizzler').catch((e) => `error: ${e.message}`) },
    { name: 'djangoer', registry: 'pypi', pinned: [...locked].sort().join(', '), latest: await pypiLatest('djangoer').catch((e) => `error: ${e.message}`) },
  ]
  return rows.map((r) => {
    const behind = !String(r.latest).startsWith('error') && !String(r.pinned).split(', ').every((v) => v === r.latest)
    const b = behind && blockedBy('generators', r.name, r.latest)
    return { app: 'generators', ...r, status: String(r.latest).startsWith('error') ? 'error' : behind ? (b ? 'blocked' : 'behind') : 'current', ...(b && { issue: b.issue }) }
  })
}

// ---- report -----------------------------------------------------------------------------------

const report = { checked_at: new Date().toISOString() }
const pins = [...(want('hekireki') ? hekirekiPins() : []), ...(want('drizzler') ? drizzlerPins() : [])]
const ready = pins.filter((p) => !p.status && wantName(p.name))
const checked = await checkPins(ready)
for (const app of ['hekireki', 'drizzler']) {
  if (want(app)) report[app] = [...pins.filter((p) => p.status && p.app === app && wantName(p.name)), ...checked.filter((p) => p.app === app)]
}
if (want('djangoer')) report.djangoer = await checkDjangoer()
if (want('generators')) report.generators = await checkGenerators()

if (applyHekireki) {
  const path = join(root, 'apps/hekireki/schemas/sources.json')
  const sources = JSON.parse(readFileSync(path, 'utf8'))
  const names = (report.hekireki ?? []).filter((p) => p.status === 'behind').map((p) => p.name)
  for (const name of names) sources[name].commit = report.hekireki.find((p) => p.name === name).head
  if (names.length) {
    writeFileSync(path, `${JSON.stringify(sources, null, 2)}\n`)
    await run('node', [join(root, 'apps/hekireki/tools/import/import-schema.mjs'), ...names], { cwd: join(root, 'apps/hekireki') })
  }
  report.applied_hekireki = names
}

if (jsonOut) writeFileSync(jsonOut, `${JSON.stringify(report, null, 2)}\n`)

const lines = [`# Upstream check (${report.checked_at})`, '']
for (const app of ['hekireki', 'drizzler', 'djangoer', 'generators']) {
  const rows = report[app]
  if (!rows) continue
  const count = (s) => rows.filter((r) => r.status === s).length
  lines.push(`## ${app}: ${count('behind')} behind, ${count('blocked')} blocked, ${count('moved')} moved, ${count('error')} errors, ${count('current')} current`, '')
  for (const r of rows.filter((r) => r.status !== 'current')) {
    if (app === 'generators') lines.push(`- ${r.status}: ${r.name} (${r.registry}) pinned ${r.pinned}, latest ${r.latest}${r.issue ? ` — ${r.issue}` : ''}`)
    else if (app === 'djangoer')
      lines.push(
        `- ${r.status}: ${r.name}${r.error ? ` — ${r.error}` : ''}${r.issue ? ` — ${r.issue}` : ''}`,
        ...(r.updates ?? []).filter((u) => u.direct).map((u) => `  - ${u.op} ${u.pkg} ${u.change}`),
      )
    else lines.push(`- ${r.status}: ${r.name} — ${r.repo ?? ''} ${r.path ?? ''}${r.compare ? ` ${r.compare}` : ''}${r.note ? ` (${r.note})` : ''}${r.error ? ` — ${r.error}` : ''}${r.issue ? ` — ${r.issue}` : ''}`)
  }
  lines.push('')
}
if (report.applied_hekireki) lines.push(`Applied to hekireki: ${report.applied_hekireki.join(', ') || 'nothing'}`, '')
console.log(lines.join('\n'))
