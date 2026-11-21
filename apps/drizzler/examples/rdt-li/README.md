# rdt-li

The schema of rdt.li, a self-hostable URL shortener: Auth.js users, accounts, sessions and verification tokens, and short URLs.

- Source: [nrjdalal/rdt-li `src/lib/db/schema.ts`](https://github.com/nrjdalal/rdt-li/blob/f04e5023523260bb6ae047b43e6f83bcefdd2f67/src/lib/db/schema.ts) (MIT)
- Dialect: PostgreSQL
- Covers: the Auth.js tables with composite primary keys, `$type<AdapterAccount['type']>()` from `@auth/core`, and a nullable foreign key

## Changes from upstream

- `schema.ts` is copied as it is. `@auth/core` is installed for its types.
- `drizzle.config.ts` is rewritten for drizzle-kit 1.0: upstream's uses the old `driver: 'pg'` and `connectionString` options.

## Generate

```bash
bazel run //apps/drizzler/examples/rdt-li:docs.write
```

The output is written to [`docs/`](docs/):

- [`README.md`](docs/README.md): table docs with the SQL
- [`er.svg`](docs/er.svg), [`er.png`](docs/er.png), [`er-dark.png`](docs/er-dark.png) and `tables/*.svg`: ER diagrams
- [`schema.dbml`](docs/schema.dbml): DBML
- [`er.md`](docs/er.md): Mermaid `erDiagram`

## Known issues

None.
