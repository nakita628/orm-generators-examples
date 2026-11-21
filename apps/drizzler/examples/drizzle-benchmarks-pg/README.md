# drizzle-benchmarks-pg

The Postgres schema of drizzle-team's benchmarks: the Northwind tables (customers, employees, orders, suppliers, products, order details) written with the drizzle-orm 1.0 style.

- Source: [drizzle-team/drizzle-benchmarks `src/schema.ts`](https://github.com/drizzle-team/drizzle-benchmarks/blob/70eebecff7bafd405d1485f46cc80205b274ad48/src/schema.ts) (no license declared)
- Dialect: PostgreSQL
- Covers: `import * as d from "drizzle-orm/pg-core"`, `d.snakeCase.table` with columns that have no SQL name (`d.varchar()`), a self-referencing `d.foreignKey`, `doublePrecision`, `date({ mode: "date" })`, and indexes

## Changes from upstream

None. `drizzle.config.ts` drops the env loading and the database URL.

## Generate

```bash
bazel run //apps/drizzler/examples/drizzle-benchmarks-pg:docs.write
```

The output is written to [`docs/`](docs/):

- [`README.md`](docs/README.md): table docs with the SQL
- [`er.svg`](docs/er.svg), [`er.png`](docs/er.png), [`er-dark.png`](docs/er-dark.png) and `tables/*.svg`: ER diagrams
- [`schema.dbml`](docs/schema.dbml): DBML
- [`er.md`](docs/er.md): Mermaid `erDiagram`

## Known issues

None.
