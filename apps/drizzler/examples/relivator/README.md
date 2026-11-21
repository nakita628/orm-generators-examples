# relivator

The schema of Relivator, a Next.js e-commerce template with Better Auth and Polar payments: users, sessions, accounts and verifications, two-factor auth, uploads, and Polar customers and subscriptions.

- Source: [reliverse/relivator `src/db/schema`](https://github.com/reliverse/relivator/tree/a1871b006ab09df647b99fc71d3c080acd797e24/src/db/schema) (MIT)
- Dialect: PostgreSQL
- Covers: a directory per feature with separate `tables.ts`, `relations.ts` and `types.ts`, a `pgEnum` whose SQL name is `type`, `schemaFilter` in `drizzle.config.ts`, and v1 `relations` in their own files

## Changes from upstream

- `relations` is imported from `drizzle-orm/_relations` instead of `drizzle-orm`, because drizzle-orm 1.0 moved the v1 relational API there.
- `drizzle.config.ts` drops the env loading and the database URL.

## Generate

```bash
bazel run //apps/drizzler/examples/relivator:docs.write
```

The output is written to [`docs/`](docs/):

- [`README.md`](docs/README.md): table docs with the SQL
- [`er.svg`](docs/er.svg), [`er.png`](docs/er.png), [`er-dark.png`](docs/er-dark.png) and `tables/*.svg`: ER diagrams
- [`schema.dbml`](docs/schema.dbml): DBML
- [`er.md`](docs/er.md): Mermaid `erDiagram`

## Known issues

None.
