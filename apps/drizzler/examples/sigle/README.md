# sigle

The server schema of Sigle, an open-source decentralized blogging platform on Stacks: Better Auth users, sessions, accounts and verifications, wallet addresses, drafts, posts and their OpenTimestamps proofs, with drizzle-orm 1.0 relations.

- Source: [sigle/sigle `apps/server-v2/src/db/schema.ts`](https://github.com/sigle/sigle/blob/31767a4d8f16e8448f88c3ff9b3d68efb47541c8/apps/server-v2/src/db/schema.ts) and [`relations.ts`](https://github.com/sigle/sigle/blob/31767a4d8f16e8448f88c3ff9b3d68efb47541c8/apps/server-v2/src/db/relations.ts) (MIT)
- Dialect: PostgreSQL
- Covers: `defineRelations` in its own file, reading the tables through `import * as schema from "@/db/schema"` (a tsconfig `paths` alias), one-to-one and one-to-many relations, and `text[]` columns (`text("tags").array()`)

## Changes from upstream

- `schema.ts` and `relations.ts` are copied as they are.
- `drizzle.config.ts` lists `relations.ts` next to `schema.ts` in `schema`, so drizzler reads the relations too. `tsconfig.json` keeps upstream's `@/*` path alias.

## Generate

```bash
bazel run //apps/drizzler/examples/sigle:docs.write
```

The output is written to [`docs/`](docs/):

- [`README.md`](docs/README.md): table docs with the SQL
- [`er.svg`](docs/er.svg), [`er.png`](docs/er.png), [`er-dark.png`](docs/er-dark.png) and `tables/*.svg`: ER diagrams
- [`schema.dbml`](docs/schema.dbml): DBML
- [`er.md`](docs/er.md): Mermaid `erDiagram`

## Known issues

None.
