# drizzler-examples

Well-known [Drizzle](https://orm.drizzle.team/) schemas from the wild, run through [drizzler](https://github.com/nakita628/drizzler) to check what it generates: ER diagrams, table docs and SQL.

Each app holds one upstream schema with as few changes as possible, plus the docs drizzler generated from it. Every app's README lists its source, what was changed and any known drizzler issues.

## Examples

| App | Source | Dialect | Known issues |
| --- | --- | --- | --- |
| [northwind](examples/northwind) | drizzle-team/drizzle-northwind-benchmarks-pg | PostgreSQL | — |
| [drizzle-benchmarks-pg](examples/drizzle-benchmarks-pg) | drizzle-team/drizzle-benchmarks | PostgreSQL | — |
| [drizzle-benchmarks-sqlite](examples/drizzle-benchmarks-sqlite) | drizzle-team/drizzle-benchmarks | SQLite | — |
| [realworld](examples/realworld) | bedstack/elysia-drizzle-realworld-example | PostgreSQL | — |
| [create-t3-app](examples/create-t3-app) | t3-oss/create-t3-app (NextAuth) | PostgreSQL | — |
| [create-t3-app-sqlite](examples/create-t3-app-sqlite) | t3-oss/create-t3-app (NextAuth) | SQLite | — |
| [create-t3-app-mysql](examples/create-t3-app-mysql) | t3-oss/create-t3-app (NextAuth) | MySQL | — |
| [create-t3-app-better-auth](examples/create-t3-app-better-auth) | t3-oss/create-t3-app (Better Auth) | PostgreSQL | — |
| [react-starter-kit](examples/react-starter-kit) | kriasoft/react-starter-kit | PostgreSQL | — |
| [authjs-drizzle-sqlite](examples/authjs-drizzle-sqlite) | nextauthjs/next-auth (Drizzle adapter) | SQLite | — |
| [authjs-drizzle-pg](examples/authjs-drizzle-pg) | nextauthjs/next-auth (Drizzle adapter) | PostgreSQL | — |
| [authjs-drizzle-mysql](examples/authjs-drizzle-mysql) | nextauthjs/next-auth (Drizzle adapter) | MySQL | — |
| [vercel-ai-chatbot](examples/vercel-ai-chatbot) | vercel/ai-chatbot | PostgreSQL | — |
| [saas-starter](examples/saas-starter) | nextjs/saas-starter | PostgreSQL | — |
| [midday](examples/midday) | midday-ai/midday | PostgreSQL | — |
| [kan](examples/kan) | kanbn/kan | PostgreSQL | — |
| [karakeep](examples/karakeep) | karakeep-app/karakeep | SQLite | [#60](https://github.com/nakita628/drizzler/issues/60) |
| [dyad](examples/dyad) | dyad-sh/dyad | SQLite | — |
| [zero](examples/zero) | Mail-0/Zero | PostgreSQL | — |
| [dokploy](examples/dokploy) | Dokploy/dokploy | PostgreSQL | [#60](https://github.com/nakita628/drizzler/issues/60), [#62](https://github.com/nakita628/drizzler/issues/62) |
| [sim-studio](examples/sim-studio) | simstudioai/sim | PostgreSQL | — |
| [morphic](examples/morphic) | miurla/morphic | PostgreSQL | — |
| [scira](examples/scira) | zaidmukaddam/scira | PostgreSQL | — |
| [papra](examples/papra) | papra-hq/papra | SQLite | — |
| [onlook](examples/onlook) | onlook-dev/onlook | PostgreSQL | [#60](https://github.com/nakita628/drizzler/issues/60) |
| [cap](examples/cap) | CapSoftware/Cap | MySQL | [#60](https://github.com/nakita628/drizzler/issues/60) |
| [pages-cms](examples/pages-cms) | pages-cms/pages-cms | PostgreSQL | — |
| [turborepo-elysia](examples/turborepo-elysia) | vercel/turborepo (with-nextjs-elysia) | PostgreSQL | — |
| [sigle](examples/sigle) | sigle/sigle | PostgreSQL | — |
| [create-t3-turbo](examples/create-t3-turbo) | t3-oss/create-t3-turbo | PostgreSQL | — |
| [alchemy-d1](examples/alchemy-d1) | alchemy-run/alchemy (cloudflare-d1-drizzle) | SQLite | — |
| [alchemy-planetscale-mysql](examples/alchemy-planetscale-mysql) | alchemy-run/alchemy (cloudflare-planetscale-mysql-drizzle) | MySQL | — |
| [alchemy-dsql](examples/alchemy-dsql) | alchemy-run/alchemy (aws-dsql-drizzle) | PostgreSQL | — |
| [justship](examples/justship) | ocluf/justship | SQLite | — |
| [cove](examples/cove) | mugnavo/cove | PostgreSQL | — |
| [t4-app](examples/t4-app) | timothymiller/t4-app | SQLite | — |
| [duolingo-clone](examples/duolingo-clone) | sanidhyy/duolingo-clone | PostgreSQL | — |
| [canva-clone](examples/canva-clone) | Davronov-Alimardon/canva-clone | PostgreSQL | — |
| [refref](examples/refref) | amicalhq/refref | PostgreSQL | — |
| [openbio](examples/openbio) | vanxh/openbio | PostgreSQL | — |
| [infinitunes](examples/infinitunes) | rajput-hemant/infinitunes | PostgreSQL | — |
| [lipi](examples/lipi) | rajput-hemant/lipi | PostgreSQL | — |
| [rdt-li](examples/rdt-li) | nrjdalal/rdt-li | PostgreSQL | — |
| [devterms](examples/devterms) | aelew/devterms | SQLite | — |
| [mylists](examples/mylists) | vincedelmas/MyLists | SQLite | [#61](https://github.com/nakita628/drizzler/issues/61) |
| [survev](examples/survev) | survev/survev | PostgreSQL | — |
| [openstatus](examples/openstatus) | openstatusHQ/openstatus | SQLite | — |
| [hyprnote](examples/hyprnote) | fastrepl/hyprnote | SQLite | — |
| [skateshop](examples/skateshop) | sadmann7/skateshop | PostgreSQL | — |
| [relivator](examples/relivator) | reliverse/relivator | PostgreSQL | — |
| [next-saas-stripe-starter](examples/next-saas-stripe-starter) | mickasmt/next-saas-stripe-starter | PostgreSQL | — |
| [nelsonlai-dev](examples/nelsonlai-dev) | nelsonlaidev/nelsonlai.dev | PostgreSQL | — |
| [nuxt-ui-chat](examples/nuxt-ui-chat) | nuxt-ui-templates/chat | SQLite | — |
| [onestopshop](examples/onestopshop) | jackblatch/OneStopShop | MySQL | — |
| [noodle](examples/noodle) | noodle-run/noodle | PostgreSQL | — |
| [orbitkit](examples/orbitkit) | ixahmedxi/orbitkit | PostgreSQL | — |
| [zws](examples/zws) | zws-im/zws | PostgreSQL | — |
| [ai-fullstack-saas](examples/ai-fullstack-saas) | alan345/AI-Fullstack-SaaS-Boilerplate | PostgreSQL | — |
| [deft](examples/deft) | Maneek21/Deft | PostgreSQL | — |
| [cut-it](examples/cut-it) | mehrabmp/cut-it | SQLite | [#60](https://github.com/nakita628/drizzler/issues/60) |
| [flarestarter](examples/flarestarter) | FlareStarter/flarestarter | SQLite | — |
| [better-auth-cloudflare](examples/better-auth-cloudflare) | zpg6/better-auth-cloudflare (Hono example) | SQLite | — |

## Output

Every app writes the same set of files to `docs/`:

| File | Format |
| --- | --- |
| `README.md` | Table docs, with the DDL and query SQL drizzle-kit and drizzle-orm build |
| `er.svg`, `er.png`, `er-dark.png`, `tables/*.svg` | ER diagrams |
| `schema.dbml` | [DBML](https://dbml.dbdiagram.io/docs/) |
| `er.md` | [Mermaid](https://mermaid.js.org/syntax/entityRelationshipDiagram.html) `erDiagram` (rendered by GitHub) |

`bazel test` parses every `schema.dbml` with `@dbml/core`, the parser behind dbdiagram.io.

Two things keep a fresh generation comparable with the committed `docs/`:

- drizzler fills the example queries of `README.md` by calling each column's `$defaultFn`, and
  defaults such as cuid2's `createId()` or a random slug draw on `Math.random` and `node:crypto`.
  [`tools/deterministic.mjs`](tools/deterministic.mjs) is preloaded into every drizzler process
  and replaces both with a fixed-seed generator, so the `-- params:` lines are the same on every
  run.
- `er.png` and `er-dark.png` are rasterized by resvg with the platform's own build and fonts, so
  their pixels differ from one machine to another. They only have to keep their format and size;
  `:docs.write` keeps a committed PNG that does, and every other file has to match byte for byte.

## Layout

```text
apps/drizzler/
  examples/<example>/
    BUILD.bazel         # drizzler_example() from tools/defs.bzl
    package.json        # a pnpm workspace package; versions from the root catalog
    drizzle.config.ts   # dialect and schema path; drizzler reads it
    src/…               # the upstream schema, kept in its original layout
    docs/               # drizzler output (committed)
  packages/tsconfig/    # shared tsconfig (base.json)
  tools/                # the Bazel macro and the scripts its targets run
```

The versions of drizzler, drizzle-orm and drizzle-kit are pinned once in the `catalog` of the root
[`pnpm-workspace.yaml`](../../pnpm-workspace.yaml).

## Usage

[Bazel](https://bazel.build/) through bazelisk, from the root of the repository. The npm packages
come from the root `pnpm-lock.yaml` through rules_js. Each example has:

| Target | What it does |
| --- | --- |
| `:docs` | runs drizzler (the same five commands as `pnpm generate`) into bazel-bin |
| `:docs.write` | `bazel run` it to write the result into `docs/` |
| `:docs.test` | fails when the committed `docs/` differs from a fresh generation |
| `:typecheck` | `tsc` on the schema, against the drizzle-orm of the lockfile |
| `:dbml` | `@dbml/core` parses the generated `schema.dbml`; for an example with a known drizzler issue that breaks it (`dbml_known_issue` in its `BUILD.bazel`), it has to fail, so the fix is noticed |

```bash
bazel test //apps/drizzler/...                                  # every example
bazel test //apps/drizzler/examples/northwind:all               # one
bazel run //apps/drizzler/examples/northwind:docs.write         # regenerate one into docs/
```

Without Bazel, `pnpm install` at the root and `pnpm -F @drizzler-examples/northwind generate` still
work.

## Adding an example

1. Copy an existing example under `examples/` and rename it in `package.json`.
2. Put the upstream schema in its original layout and point `schema` in `drizzle.config.ts` at it.
3. Make only the changes it needs to load with drizzle-orm 1.0 (for example `relations` from `drizzle-orm/_relations`), and note them in the example's README.
4. Run `pnpm install` at the root (to add it to `pnpm-lock.yaml` and `.bazelignore`'s list of `node_modules`), then `bazel run //apps/drizzler/examples/<name>:docs.write`, and record anything drizzler gets wrong under "Known issues".

## License

Each schema keeps the license of its source repository, linked from the app's README. drizzle-team/drizzle-northwind-benchmarks-pg, drizzle-team/drizzle-benchmarks and nextjs/saas-starter do not declare a license.
