# turborepo-elysia

The Drizzle schema of Vercel's Turborepo example `with-nextjs-elysia`: users, roles and the join table between them, with their v1 `relations`.

- Source: [vercel/turborepo `examples/with-nextjs-elysia/packages/contract/src/drizzle`](https://github.com/vercel/turborepo/tree/51d6ddfcbb35d7cc264ee9b6e0ef8f4f82fc23d1/examples/with-nextjs-elysia/packages/contract/src/drizzle) (MIT)
- Dialect: PostgreSQL
- Covers: `import * as p from "drizzle-orm/pg-core"` (`p.pgTable`, `p.uuid`), column builders held in constants and reused (`id: idUuid`) or spread through an object (`...Audit`), v1 `relations` over a join table (`many` from users and roles, `one` back from the join table), and a relations file next to the tables

## Changes from upstream

- `index.ts`, `table.schema.ts` and `table.relation.ts` are copied as they are, except that `relations` is imported from `drizzle-orm/_relations` instead of `drizzle-orm`, because drizzle-orm 1.0 moved the v1 relational API there.
- `drizzle.config.ts` lists `table.relation.ts` next to `table.schema.ts` in `schema`, so drizzler reads the relations too. It drops `casing: "snake_case"`, which drizzle-kit 1.0 no longer accepts; every column names its SQL column. It also drops the database URL.

## Generate

```bash
bazel run //apps/drizzler/examples/turborepo-elysia:docs.write
```

The output is written to [`docs/`](docs/):

- [`README.md`](docs/README.md): table docs with the SQL
- [`er.svg`](docs/er.svg), [`er.png`](docs/er.png), [`er-dark.png`](docs/er-dark.png) and `tables/*.svg`: ER diagrams
- [`schema.dbml`](docs/schema.dbml): DBML
- [`er.md`](docs/er.md): Mermaid `erDiagram`

## Known issues

None.
