# canva-clone

The schema of a popular open-source Canva clone built with Next.js and Fabric.js: Auth.js users, accounts, sessions, verification tokens and authenticators, projects and subscriptions.

- Source: [Davronov-Alimardon/canva-clone `src/db/schema.ts`](https://github.com/Davronov-Alimardon/canva-clone/blob/53f2704646714e70339ffcad88a4e96d9b2138c9/src/db/schema.ts) (Apache-2.0)
- Dialect: PostgreSQL
- Covers: the Auth.js tables with composite primary keys, `$defaultFn(() => crypto.randomUUID())`, a zod insert schema next to the table, and v1 `relations`

## Changes from upstream

- `relations` is imported from `drizzle-orm/_relations` instead of `drizzle-orm`, and `createInsertSchema` from `drizzle-orm/zod` instead of `drizzle-zod`, because drizzle-orm 1.0 moved and replaced them.
- `drizzle.config.ts` drops the env loading and the database URL.

## Generate

```bash
bazel run //apps/drizzler/examples/canva-clone:docs.write
```

The output is written to [`docs/`](docs/):

- [`README.md`](docs/README.md): table docs with the SQL
- [`er.svg`](docs/er.svg), [`er.png`](docs/er.png), [`er-dark.png`](docs/er-dark.png) and `tables/*.svg`: ER diagrams
- [`schema.dbml`](docs/schema.dbml): DBML
- [`er.md`](docs/er.md): Mermaid `erDiagram`

## Known issues

None.
