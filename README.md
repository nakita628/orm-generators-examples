# orm-generators-examples

Generation examples for three schema tools, run over the schemas of well-known open source
projects, in one [Bazel](https://bazel.build/) monorepo. Every example commits what its generator
wrote, and Bazel holds that copy to a fresh generation.

| App | Tool | Input | Examples | Output |
| --- | --- | --- | --: | --- |
| [apps/hekireki](apps/hekireki) | [hekireki](https://github.com/nakita628/hekireki) | Prisma schema | 31 | Drizzle, Kysely, SQLAlchemy, Django, GORM, SeaORM, Ecto, Active Record, EF Core, Eloquent, Exposed, ER diagram, docs |
| [apps/drizzler](apps/drizzler) | [drizzler](https://github.com/nakita628/drizzler) | Drizzle schema | 62 | ER diagrams (SVG, PNG, Mermaid), DBML, table docs with SQL |
| [apps/djangoer](apps/djangoer) | [Djangoer](https://github.com/nakita628/djangoer) | Django models | 224 | ER diagrams (SVG, Mermaid), DBML, model docs |

Each app was its own repository before: [hekireki-examples](https://github.com/nakita628/hekireki-examples),
[drizzler-examples](https://github.com/nakita628/drizzler-examples) and
[djangoer-examples](https://github.com/nakita628/djangoer-examples).

## Layout

```text
MODULE.bazel           Bazel modules: rules_js and Node for hekireki and drizzler, uv for djangoer
.bazelrc               shared flags; `--config=djangoer`
package.json           the root of the one pnpm workspace
pnpm-workspace.yaml    its packages, the drizzler catalog, allowBuilds
pnpm-lock.yaml         every npm package of every app, read by rules_js
tools/golden/          `golden()`: the `.write` / `.test` pair every example's output goes through
tools/uv/              the uv binary of MODULE.bazel for the host platform
tools/upstream/        the daily upstream sync: check.mjs, blocked.json (see .claude/skills/upstream-sync)
tools/cloud/           Docker in a Claude Code cloud session, for hekireki's compat and pull/
apps/hekireki/         schemas/<name>/ (schema.prisma, generated/), checks/<lang>/, compat/, tools/
apps/drizzler/         examples/<name>/ (src/, drizzle.config.ts, docs/), packages/tsconfig, tools/
apps/djangoer/         examples/<name>/ (pyproject.toml, uv.lock, settings.py, outputs), scripts/, tools/
pull/                  <orm>/<project>/: the SQL other ORMs generate (source.sql), written up in Prisma (schema.prisma)
```

## Usage

[bazelisk](https://github.com/bazelbuild/bazelisk) picks the Bazel of `.bazelversion`.

```sh
# hekireki and drizzler: generate every example and compare with the committed output, tsc and
# the DBML parser on every drizzler example
bazel test //...

# One app, or one example
bazel test //apps/drizzler/...
bazel test //apps/drizzler/examples/northwind:all
bazel test //apps/hekireki/schemas/calcom:all

# Rewrite one example's committed output
bazel run //apps/hekireki/schemas/calcom:generate.write
bazel run //apps/drizzler/examples/northwind:docs.write
bazel run //apps/djangoer/examples/wagtail:generate.write

# djangoer: uv installs each example's locked environment from PyPI (left out of //...)
bazel test --config=djangoer //apps/djangoer/...

# hekireki's language checks: the toolchain comes from PATH (left out of //...)
bazel test --test_tag_filters=lang-gorm //apps/hekireki/checks/...
```

Every example has the same pair of targets, from [tools/golden](tools/golden/defs.bzl):

| Target | |
| --- | --- |
| `<name>.test` | fails when the committed output differs from a fresh generation |
| `<name>.write` | `bazel run` it to write a fresh generation over the committed output |

| App | Generation | Committed output | More tests |
| --- | --- | --- | --- |
| hekireki | `//apps/hekireki/schemas/<name>:generate` | `generated/` | `//apps/hekireki/checks/<lang>:<name>` (tag `lang`) |
| drizzler | `//apps/drizzler/examples/<name>:docs` | `docs/` | `:typecheck`, `:dbml` |
| djangoer | `//apps/djangoer/examples/<name>:generate.test` | `er.md`, `er.svg`, `schema.dbml`, `docs/`, `models/` | `//apps/djangoer:lint` (tag `djangoer`) |

Machine-specific flags go in `user.bazelrc` (ignored by git), for example
`build --disk_cache=~/.cache/bazel-disk` or `build --jobs=4`.

## Versions

- npm: hekireki, prisma and the ORMs hekireki's output is checked against are pinned in
  [apps/hekireki/package.json](apps/hekireki/package.json); drizzler, drizzle-orm and drizzle-kit
  in the `catalog` of [pnpm-workspace.yaml](pnpm-workspace.yaml). After changing a version, run
  `pnpm install --lockfile-only` and `bazel run` the `.write` targets.
- Python: each Django example has its own `pyproject.toml` and `uv.lock`; Python 3.14 is set in
  [apps/djangoer/.python-version](apps/djangoer/.python-version).
- Node, uv and the Bazel rules: [MODULE.bazel](MODULE.bazel).

## Following upstream

A scheduled task runs once a day: [`tools/upstream/check.mjs`](tools/upstream/check.mjs) finds the
schemas whose upstream project changed and the generator releases not yet pinned, and each one is
followed, regenerated and tested, then pushed to `main`. A failure that is the generator's becomes
an issue on hekireki, drizzler or Djangoer, and the example waits at its last good pin
([`tools/upstream/blocked.json`](tools/upstream/blocked.json)). See
[tools/upstream](tools/upstream/README.md).

## pull: other ORMs' SQL, written up in Prisma

[pull/](pull) goes the other way: well-known projects written for other ORMs (every djangoer and
drizzler example, Rails apps such as Mastodon and Forem, MLflow) have their own ORM create their
tables, the tables are kept as SQL, and `prisma db pull` writes them up in Prisma, one layer per
ORM (`pull/django/`, `pull/drizzle/`, `pull/activerecord/`, `pull/sqlalchemy/`). Every ORM
hekireki generates from a write-up then runs against the project's own tables. A scheduled task
does three projects a day and files what is hekireki's on hekireki. It runs on your own Docker:

```sh
pnpm install
node pull/pull.mjs --random 3      # or: node pull/pull.mjs django/wagtail activerecord/mastodon
```

## License

Each schema belongs to the project it comes from, under that project's license; see each app's
README.
