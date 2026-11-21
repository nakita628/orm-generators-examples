# alchemy-planetscale-mysql

The Drizzle schema of Alchemy's `cloudflare-planetscale-mysql-drizzle` example. Alchemy is infrastructure-as-code in TypeScript.

- Source: [alchemy-run/alchemy `examples/cloudflare-planetscale-mysql-drizzle/src/schema.ts`](https://github.com/alchemy-run/alchemy/blob/876a0e2276eaafe19adc3a9b9975ee1f45204a48/examples/cloudflare-planetscale-mysql-drizzle/src/schema.ts) (Apache-2.0)
- Dialect: MySQL (PlanetScale)
- Covers: a users and posts schema with no foreign key: the relation between them is only declared with `defineRelations`, as on PlanetScale, plus `int().autoincrement()` keys and `varchar({ length })`

## Changes from upstream

None.

## Generate

```bash
bazel run //apps/drizzler/examples/alchemy-planetscale-mysql:docs.write
```

The output is written to [`docs/`](docs/):

- [`README.md`](docs/README.md): table docs with the SQL
- [`er.svg`](docs/er.svg), [`er.png`](docs/er.png), [`er-dark.png`](docs/er-dark.png) and `tables/*.svg`: ER diagrams
- [`schema.dbml`](docs/schema.dbml): DBML
- [`er.md`](docs/er.md): Mermaid `erDiagram`

## Known issues

None.
