# alchemy-dsql

The Drizzle schema of Alchemy's `aws-dsql-drizzle` example. Alchemy is infrastructure-as-code in TypeScript.

- Source: [alchemy-run/alchemy `examples/aws-dsql-drizzle/src/schema.ts`](https://github.com/alchemy-run/alchemy/blob/876a0e2276eaafe19adc3a9b9975ee1f45204a48/examples/aws-dsql-drizzle/src/schema.ts) (Apache-2.0)
- Dialect: PostgreSQL (Aurora DSQL)
- Covers: a table in another Postgres schema, built by calling `.table()` on `pgSchema("app")` in place

## Changes from upstream

None.

## Generate

```bash
bazel run //apps/drizzler/examples/alchemy-dsql:docs.write
```

The output is written to [`docs/`](docs/):

- [`README.md`](docs/README.md): table docs with the SQL
- [`er.svg`](docs/er.svg), [`er.png`](docs/er.png), [`er-dark.png`](docs/er-dark.png) and `tables/*.svg`: ER diagrams
- [`schema.dbml`](docs/schema.dbml): DBML
- [`er.md`](docs/er.md): Mermaid `erDiagram`

## Known issues

None.
