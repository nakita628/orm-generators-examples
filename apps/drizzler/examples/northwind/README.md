# northwind

The Northwind Traders schema (PostgreSQL) from Drizzle's own ORM benchmarks.

- Source: [drizzle-team/drizzle-northwind-benchmarks-pg `src/drizzle/schema.ts`](https://github.com/drizzle-team/drizzle-northwind-benchmarks-pg/blob/c225260a8c75c88e905592a8ff7995940abb4e0e/src/drizzle/schema.ts)
- Dialect: PostgreSQL
- Covers: a self-referencing foreign key declared with `foreignKey()` in the legacy object-style extra config, `date` / `doublePrecision` columns, v1 `relations()`

## Changes from upstream

- `relations` is imported from `drizzle-orm/_relations` instead of `drizzle-orm`, because drizzle-orm 1.0 moved the v1 relational API there.

## Generate

```bash
bazel run //apps/drizzler/examples/northwind:docs.write
```

The output is written to [`docs/`](docs/):

- [`README.md`](docs/README.md): table docs with the SQL
- [`er.svg`](docs/er.svg), [`er.png`](docs/er.png), [`er-dark.png`](docs/er-dark.png) and `tables/*.svg`: ER diagrams
- [`schema.dbml`](docs/schema.dbml): DBML
- [`er.md`](docs/er.md): Mermaid `erDiagram`

## Known issues

None.
