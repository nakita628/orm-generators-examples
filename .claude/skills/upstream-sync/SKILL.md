---
name: upstream-sync
description: The daily upstream sync of orm-generators-examples. Follow every example's upstream project and the generators (hekireki, drizzler, djangoer) to their latest, regenerate, test, push what passes to main, and file an issue on the generator's repository for what fails because of the generator. Use when the scheduled daily task fires, or when asked to "sync upstream", "follow upstream" or "追従".
---

# Daily upstream sync

Goal: every example in `apps/` runs against the newest upstream schema and the newest generator
release. What passes goes straight to `main`. What a generator gets wrong becomes an issue on that
generator's repository, and the example stays at its last good pin until the issue is closed.

Work from the root of the repository. Read `tools/upstream/README.md` once for the files involved.
The daily scheduled task runs this sync first and the pull (`.claude/skills/pull-check`) after it,
in the same session, whether the sync followed anything or not.

## 0. Set up

In a cloud session `gh` has no valid token (only `gh api -X GET search/issues -f q='repo:… is:issue …'`
works); read, file and comment on issues with the GitHub MCP tools (`issue_read`, `list_issues`,
`issue_write`, `add_issue_comment`) wherever this file says `gh api`.

1. `git checkout main && git pull --ff-only`.
2. `bazel version` must work. If it is missing or Bazel cannot download, the environment's setup
   script did not run: see "Machine setup" in `tools/upstream/README.md`, run those steps, and go on.
3. Attach the generator repositories so issues can be filed: `nakita628/hekireki`,
   `nakita628/drizzler`, `nakita628/djangoer` (add_repo, access `push`; no need to clone them). When
   the session has no `add_repo` tool, they are already attached (cloned under `/home/user`) and the
   GitHub tools reach them.

## 1. Unblock

For each entry of `tools/upstream/blocked.json` with an `issue`, read its state:
`gh api repos/nakita628/<generator>/issues/<number> --jq .state`. When it is `closed`, remove the
entry (the example is retried below with the newest generator). Entries with `"issue": null` (an
upstream problem, not the generator's) are retried every run: remove them, and add them back if
they still fail.

## 2. Check

```bash
node tools/upstream/check.mjs --json /tmp/upstream.json
```

`behind` = to follow today. `moved` = the upstream file is gone at HEAD. `blocked` = waiting on an
issue. `error` = the check itself failed (network…): retry once, then report it.

Nothing behind and nothing moved: the sync ends here; report "nothing to follow" (the daily task
then goes on with the pull).

## 3. Generators first

A new generator release goes in before the schemas, so that schema updates are tested against it.
One commit per generator, in the existing style (`Bump hekireki to 0.52.17 and regenerate`).

- **hekireki**: set the version in `apps/hekireki/package.json` and add `hekireki@<version>` to
  `minimumReleaseAgeExclude` in `pnpm-workspace.yaml` (replace the old hekireki entry), then
  `pnpm install --lockfile-only`. Regenerate every schema:
  `bazel query 'kind(sh_binary, //apps/hekireki/schemas/...)' | grep '\.write$' | xargs -n1 bazel run`,
  then `bazel test //apps/hekireki/...`.
- **drizzler**: set `drizzler` in the `catalog` of `pnpm-workspace.yaml`, add the version to the
  drizzler entry of `minimumReleaseAgeExclude`, `pnpm install --lockfile-only`, run every
  `//apps/drizzler/examples/*:docs.write`, then `bazel test //apps/drizzler/...`.
- **djangoer**: for each example, `uvx --from uv==<UV_VERSION of MODULE.bazel> uv lock
  --upgrade-package djangoer --project apps/djangoer/examples/<name>`; run every
  `//apps/djangoer/examples/<name>:generate.write` (plain `bazel run`: `--config=djangoer` is a `test`
  config), then `bazel test --config=djangoer //apps/djangoer/...`. Every example's environment together takes
  about 24 GB (`~/.cache/orm-generators-examples/djangoer/venvs`); the session's disk is limited,
  so remove that directory once the djangoer tests are done.

Then read the diff of the generated output (`git diff --stat`, and the files that changed most).
A release is expected to change output; a change that is plainly a regression (a column, relation,
index or enum that disappears; a type that becomes wrong; output that no longer parses) is a
generator problem even though the tests pass: see step 6.

If the bump breaks examples: keep the bump for the examples it does not break only when that is
simple; otherwise revert it entirely (`git checkout -- .`), file the issue(s), and add a
`generators` entry with `target` = the new version to `blocked.json`.

## 4. Follow each upstream schema

Do the examples one at a time. After each, run its tests; commit only what passes.

### hekireki

```bash
node tools/upstream/check.mjs --app hekireki --apply-hekireki   # sources.json -> HEAD, re-import
bazel run //apps/hekireki/schemas/<name>:generate.write
bazel test //apps/hekireki/schemas/<name>:all
```

Then in `apps/hekireki/README.md` update the row of the schema: the commit in the source link and
the model count (`grep -c '^model ' apps/hekireki/schemas/<name>/schema.prisma`; keep the table
sorted by it). `--apply-hekireki` moves every behind schema at once; for one that you then cannot
finish, restore its `sources.json` commit and `schema.prisma` with `git checkout`.

### drizzler

0. `node tools/upstream/follow-drizzler.mjs <name> <head of the check's JSON> [--exact]` reports, and
   with `--apply` does, steps 2 to 4 below as a 3-way merge per file (base = the upstream file at the
   pinned commit, theirs = at the new one, ours = the example's copy), and lists conflicts, added and
   deleted files. Then do what is left of them by hand: resolve the conflicts (when upstream now does
   what a "Changes from upstream" line did, take upstream's file and drop that line), apply the list
   to the added files, and add a stand-in or a copy for a new import of an app module.
1. Read the example's `README.md`: the `- Source:` links and `## Changes from upstream`.
2. Download every linked file (or directory) at the new HEAD commit from
   `https://raw.githubusercontent.com/<repo>/<commit>/<path>` (for a directory, list it with a
   tree-only `git fetch --depth 1 --filter=blob:none` and `git ls-tree -r`) into `src/`, in the
   layout already there.
3. Re-apply every change listed under "Changes from upstream". If upstream changed in a way the
   list does not cover (a new import of an app module, a new dependency), make the smallest change
   that keeps the schema as written and add a line for it to the list.
4. Update the commit in every `- Source:` link (and in any other link of that README).
5. `bazel run //apps/drizzler/examples/<name>:docs.write` and
   `bazel test //apps/drizzler/examples/<name>:all`. A new npm dependency of the schema goes in the
   example's `package.json`, then `pnpm install --lockfile-only`.
6. If the example's description or "Covers" line in its README, or its row in
   `apps/drizzler/README.md`, no longer holds, update it.

### djangoer

```bash
uvx --from uv==<UV_VERSION> uv lock --upgrade --project apps/djangoer/examples/<name>
bazel run //apps/djangoer/examples/<name>:generate.write      # no --config: it is defined for `test` only
bazel test --config=djangoer //apps/djangoer/examples/<name>:all
```

Then `uv run --project apps/djangoer apps/djangoer/scripts/readme.py --check` and
`bazel test --config=djangoer //apps/djangoer:lint`.

### moved

The file is gone at upstream's HEAD. Look for where it went (`git ls-tree -r --name-only HEAD` in
the tree-only fetch, `git log --diff-filter=D` is not available in a shallow fetch: search by file
name and by a table name from the schema). If it moved, follow it like `behind` and update the
path in `sources.json` / the README link. If the project dropped Drizzle/Prisma or the file for
good, leave the example pinned, add a `blocked.json` entry with `"issue": null`, `target: "*"`
and the reason, and mention it in the report. Never file a generator issue for this.

## 5. Push

After each app (or each generator bump), when its tests pass:

```bash
git add -A && git commit -m "<what>"        # e.g. "Follow upstream: calcom, dub (hekireki)"
git pull --rebase && git push origin main
```

Never push a red test, a half-applied example, or `user.bazelrc`. Retry a failed push up to 4
times with backoff.

## 6. When it fails: whose problem is it?

Read the error and the generated output before deciding.

**The generator's** — file an issue on `nakita628/<generator>`:
- the generator crashes, hangs, or refuses a schema that its ORM (Prisma, drizzle-kit, Django)
  accepts;
- its output does not compile / typecheck / parse (`:typecheck` of hekireki's output, `:dbml`,
  the `lang` checks), or differs between two runs;
- its output is wrong for the schema (a lost field, relation, index, default, enum value; a wrong
  type or nullability), including regressions seen in step 3's diff.

**Upstream's** — no generator issue:
- `prisma validate` or `tsc` fails on the upstream schema itself (missing generator-only syntax,
  an import of an app module the copy does not have, a broken upstream commit);
- the file moved or was deleted.
Adapt when the existing "Changes from upstream" conventions cover it; otherwise keep the old pin
and add a `blocked.json` entry with `"issue": null` and the reason.

**The machine's** — a download refused, a timeout: retry once, then report it. Not an issue.

### Filing an issue

1. Search first, open and closed:
   `gh api 'repos/nakita628/<generator>/issues?state=all&per_page=100' --jq '.[] | "\(.number) \(.state) \(.title)"'`
   (page with `&page=2`… as needed) and read any that look like the same problem. Same problem
   and open: add a comment with the new case instead of a new issue. Same problem and closed in a
   release older than the one used: it regressed; open a new issue that links the old one.
2. Reduce it: the smallest schema snippet (a model or two) that shows it. Run it through the
   generator to confirm the snippet alone reproduces it.
3. Write it in English, in the style of the existing issues: the title names the generator /
   output and the symptom (`hekireki-ecto: a Json field on MySQL loads as a string`); the body has
   the generator version, the example and the upstream commit link, the snippet, the command,
   what it wrote, what it should write, and why (the ORM's docs if relevant).
   `gh api repos/nakita628/<generator>/issues -f title='…' -f body='…'`
4. Record it in `tools/upstream/blocked.json` and keep that example at its old pin:
   ```json
   { "app": "hekireki", "name": "calcom", "target": "<new commit | version | '*'>", "issue": "https://github.com/nakita628/hekireki/issues/123", "reason": "one line" }
   ```
   For djangoer, `target` is the direct updates as the check prints them (`wagtail v7.2 -> v7.3`).
   For a generator bump, `app` is `generators` and `name` is the generator.
   For a drizzler example whose `dbml` check fails, the repo's own convention is
   `dbml_known_issue` in the example's `BUILD.bazel` plus the "Known issues" column of
   `apps/drizzler/README.md`; use it (with the issue link) instead of holding the example back.

Never file more than one issue per root cause per run, and never file one without the reduced
reproduction.

## 7. Report

End with a short report (Japanese, for the owner): what was followed and pushed (commits), what
is blocked and why (issue links, new and old), what was skipped or errored. If nothing changed,
say so in one line.
