# drizzle-benchmarks-sqlite

The SQLite schema of drizzle-team's benchmarks: the Northwind tables (customers, employees, orders, suppliers, products, order details).

- Source: [drizzle-team/drizzle-benchmarks `src/sqlite/schema.ts`](https://github.com/drizzle-team/drizzle-benchmarks/blob/70eebecff7bafd405d1485f46cc80205b274ad48/src/sqlite/schema.ts) (no license declared)
- Dialect: SQLite
- Covers: a self-referencing `foreignKey` in the legacy object-style extra config, `numeric` columns, indexes, and v1 `relations` with a self-relation

## Changes from upstream

- `relations` is imported from `drizzle-orm/_relations` instead of `drizzle-orm`, because drizzle-orm 1.0 moved the v1 relational API there.
- The `InferModel<typeof table>` and `InferModel<typeof table, 'insert'>` types, which drizzle-orm 1.0 removed, are written as `typeof table.$inferSelect` and `typeof table.$inferInsert`.

## Generate

```bash
bazel run //apps/drizzler/examples/drizzle-benchmarks-sqlite:docs.write
```

The output is written to [`docs/`](docs/):

- [`README.md`](docs/README.md): table docs with the SQL
- [`er.svg`](docs/er.svg), [`er.png`](docs/er.png), [`er-dark.png`](docs/er-dark.png) and `tables/*.svg`: ER diagrams
- [`schema.dbml`](docs/schema.dbml): DBML
- [`er.md`](docs/er.md): Mermaid `erDiagram`

## Known issues

None.
