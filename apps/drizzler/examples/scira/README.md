# scira

The schema of Scira, an open-source AI search engine: Better Auth users, sessions and accounts, chats, messages and streams, subscriptions and payments, usage counters, custom instructions, lookouts (scheduled searches), extreme-search usage and more.

- Source: [zaidmukaddam/scira `lib/db/schema.ts`](https://github.com/zaidmukaddam/scira/blob/e1692f5bdec7ec3f6482a24e0c6cf9b483d810f9/lib/db/schema.ts) (AGPL-3.0)
- Dialect: PostgreSQL
- Covers: `$defaultFn` ids from real third-party packages (`generateId` from `ai`, `v7` from `uuid`), `varchar({ enum })`, `json().$type<...>()`, `real` columns, a single-file schema of 20 tables, and v1 `relations`

## Changes from upstream

- `relations` is imported from `drizzle-orm/_relations` instead of `drizzle-orm`, because drizzle-orm 1.0 moved the v1 relational API there.
- `ai` and `uuid` are installed at the versions upstream uses.
- `tsconfig.json` turns off `verbatimModuleSyntax`, because upstream imports `InferSelectModel` without `type`.

## Generate

```bash
bazel run //apps/drizzler/examples/scira:docs.write
```

The output is written to [`docs/`](docs/):

- [`README.md`](docs/README.md): table docs with the SQL
- [`er.svg`](docs/er.svg), [`er.png`](docs/er.png), [`er-dark.png`](docs/er-dark.png) and `tables/*.svg`: ER diagrams
- [`schema.dbml`](docs/schema.dbml): DBML
- [`er.md`](docs/er.md): Mermaid `erDiagram`

## Known issues

None.
