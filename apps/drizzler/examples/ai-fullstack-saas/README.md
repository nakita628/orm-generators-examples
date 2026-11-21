# ai-fullstack-saas

The database package of alan345's AI Fullstack SaaS Boilerplate (tRPC, Fastify, React, Better Auth): users, sessions, accounts, verifications, posts and devices.

- Source: [alan345/AI-Fullstack-SaaS-Boilerplate `packages/drizzle/src/db/schema.ts`](https://github.com/alan345/AI-Fullstack-SaaS-Boilerplate/blob/b8a8194a916a5c1c4df4125a3910f807e6c326c7/packages/drizzle/src/db/schema.ts) (MIT)
- Dialect: PostgreSQL
- Covers: the Better Auth tables, `uuid` keys, foreign keys with `onDelete`, and v1 `relations`

## Changes from upstream

- `relations` is imported from `drizzle-orm/_relations` instead of `drizzle-orm`, because drizzle-orm 1.0 moved the v1 relational API there.
- `drizzle.config.ts` drops the env loading and the database URL.

## Generate

```bash
bazel run //apps/drizzler/examples/ai-fullstack-saas:docs.write
```

The output is written to [`docs/`](docs/):

- [`README.md`](docs/README.md): table docs with the SQL
- [`er.svg`](docs/er.svg), [`er.png`](docs/er.png), [`er-dark.png`](docs/er-dark.png) and `tables/*.svg`: ER diagrams
- [`schema.dbml`](docs/schema.dbml): DBML
- [`er.md`](docs/er.md): Mermaid `erDiagram`

## Known issues

None.
