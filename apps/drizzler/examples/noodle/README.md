# noodle

The schema of Noodle, an open-source student productivity platform: modules and early-access sign-ups.

- Source: [noodle-run/noodle `src/db/schema`](https://github.com/noodle-run/noodle/tree/4e26fb7c990ccccce203f1c30f79126ef1bfbcaa/src/db/schema) (AGPL-3.0)
- Dialect: PostgreSQL
- Covers: `uuid().defaultRandom()` keys, zod insert and select schemas extended next to the table, and an `index.ts` of `export *`

## Changes from upstream

- `createInsertSchema` and `createSelectSchema` are imported from `drizzle-orm/zod` instead of `drizzle-zod`, which drizzle-orm 1.0 replaces; `zod` is 4, as `drizzle-orm/zod` requires (upstream uses 3).
- `drizzle.config.ts` drops the env loading and the database URL.

## Generate

```bash
bazel run //apps/drizzler/examples/noodle:docs.write
```

The output is written to [`docs/`](docs/):

- [`README.md`](docs/README.md): table docs with the SQL
- [`er.svg`](docs/er.svg), [`er.png`](docs/er.png), [`er-dark.png`](docs/er-dark.png) and `tables/*.svg`: ER diagrams
- [`schema.dbml`](docs/schema.dbml): DBML
- [`er.md`](docs/er.md): Mermaid `erDiagram`

## Known issues

None.
