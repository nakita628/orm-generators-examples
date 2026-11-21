# alchemy-d1

The Drizzle schema of Alchemy's `cloudflare-d1-drizzle` example. Alchemy is infrastructure-as-code in TypeScript.

- Source: [alchemy-run/alchemy `examples/cloudflare-d1-drizzle/src/schema.ts`](https://github.com/alchemy-run/alchemy/blob/876a0e2276eaafe19adc3a9b9975ee1f45204a48/examples/cloudflare-d1-drizzle/src/schema.ts) (Apache-2.0)
- Dialect: SQLite (Cloudflare D1)
- Covers: a users and posts schema with `integer` keys using `primaryKey({ autoIncrement: true })`, `integer({ mode: "timestamp" })` with `` .default(sql`(unixepoch())`) ``, a foreign key, and `defineRelations` over an object of tables (`{ Users, Posts }`) in the same file

## Changes from upstream

None.

## Generate

```bash
bazel run //apps/drizzler/examples/alchemy-d1:docs.write
```

The output is written to [`docs/`](docs/):

- [`README.md`](docs/README.md): table docs with the SQL
- [`er.svg`](docs/er.svg), [`er.png`](docs/er.png), [`er-dark.png`](docs/er-dark.png) and `tables/*.svg`: ER diagrams
- [`schema.dbml`](docs/schema.dbml): DBML
- [`er.md`](docs/er.md): Mermaid `erDiagram`

## Known issues

None.
