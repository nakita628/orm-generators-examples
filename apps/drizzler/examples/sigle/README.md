# sigle

The server schema of Sigle, an open-source decentralized blogging platform on Stacks: Better Auth users, sessions, accounts and verifications, wallet addresses, drafts, profiles, media images, posts and their OpenTimestamps proofs, with drizzle-orm 1.0 relations.

- Source: [sigle/sigle `apps/server/src/db/schema.ts`](https://github.com/sigle/sigle/blob/51f08cd2158c0b35e57179a0744a5eb574bd3f54/apps/server/src/db/schema.ts) and [`relations.ts`](https://github.com/sigle/sigle/blob/51f08cd2158c0b35e57179a0744a5eb574bd3f54/apps/server/src/db/relations.ts) (MIT)
- Dialect: SQLite
- Covers: `defineRelations` in its own file, reading the tables through `import * as schema from "@/db/schema"` (a tsconfig `paths` alias), one-to-one and one-to-many relations, column helpers that wrap `integer` and `text` (`timestamp`, `boolean`, `textArray`), `check` constraints built by a helper, partial indexes (`.where(sql`...`)`), string arrays stored as JSON text (`text({ mode: "json" }).$type<string[]>()`) and a `blob`

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
