# create-t3-app-better-auth

The Drizzle + Better Auth schema that `create-t3-app` scaffolds for PostgreSQL.

- Source: [t3-oss/create-t3-app `with-better-auth-postgres.ts`](https://github.com/t3-oss/create-t3-app/blob/4709861f7e67a15564c0460c13e7b4b6cfcae40d/cli/template/extras/src/server/db/schema-drizzle/with-better-auth-postgres.ts) (MIT)
- Dialect: PostgreSQL
- Covers: `pgTableCreator` and plain `pgTable` side by side, Better Auth's core tables (`user`, `session`, `account`, `verification`)

## Changes from upstream

- `relations` is imported from `drizzle-orm/_relations` instead of `drizzle-orm`, because drizzle-orm 1.0 moved the v1 relational API there.

## Generate

```bash
bazel run //apps/drizzler/examples/create-t3-app-better-auth:docs.write
```

The output is written to [`docs/`](docs/):

- [`README.md`](docs/README.md): table docs with the SQL
- [`er.svg`](docs/er.svg), [`er.png`](docs/er.png), [`er-dark.png`](docs/er-dark.png) and `tables/*.svg`: ER diagrams
- [`schema.dbml`](docs/schema.dbml): DBML
- [`er.md`](docs/er.md): Mermaid `erDiagram`

## Known issues

None.
