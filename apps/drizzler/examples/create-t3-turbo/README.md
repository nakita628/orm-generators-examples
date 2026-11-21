# create-t3-turbo

The database package of create-t3-turbo, the T3 Stack as a Turborepo monorepo: a `Post` table and the Better Auth tables.

- Source: [t3-oss/create-t3-turbo `packages/db/src/schema.ts`](https://github.com/t3-oss/create-t3-turbo/blob/8f945b7bb3bfb3ca8358d48b1ff0214079bc11ee/packages/db/src/schema.ts) and [`auth-schema.ts`](https://github.com/t3-oss/create-t3-turbo/blob/8f945b7bb3bfb3ca8358d48b1ff0214079bc11ee/packages/db/src/auth-schema.ts) (MIT)
- Dialect: PostgreSQL
- Covers: the callback form of the table builder (`(t) => ({ id: t.uuid() })`) with columns that have no SQL name, snake_case column names, `` $onUpdateFn(() => sql`now()`) ``, `export *` of the auth tables, and a zod insert schema next to the table

## Changes from upstream

- `pgTable` is `snakeCase.table`. Upstream gets snake_case column names from `casing: "snake_case"` in `drizzle.config.ts`, which drizzle-kit 1.0 no longer reads; `snakeCase.table` is how drizzle-orm 1.0 says the same.
- `createInsertSchema` is imported from `drizzle-orm/zod` instead of `drizzle-zod`, which drizzle-orm 1.0 replaces.
- `drizzle.config.ts` drops `casing`, the database URL and its check.

## Generate

```bash
bazel run //apps/drizzler/examples/create-t3-turbo:docs.write
```

The output is written to [`docs/`](docs/):

- [`README.md`](docs/README.md): table docs with the SQL
- [`er.svg`](docs/er.svg), [`er.png`](docs/er.png), [`er-dark.png`](docs/er-dark.png) and `tables/*.svg`: ER diagrams
- [`schema.dbml`](docs/schema.dbml): DBML
- [`er.md`](docs/er.md): Mermaid `erDiagram`

## Known issues

None.
