# devterms

The schema of DevTerms, a crowdsourced dictionary for developers: users, sessions, definitions, reports and words of the day.

- Source: [aelew/devterms `src/server/db/schema.ts`](https://github.com/aelew/devterms/blob/60a285a308a36baba4b77b582898ca7226ba633a/src/server/db/schema.ts) and [`src/lib/id.ts`](https://github.com/aelew/devterms/blob/60a285a308a36baba4b77b582898ca7226ba633a/src/lib/id.ts) (MIT)
- Dialect: SQLite (`dialect: "turso"`)
- Covers: columns with no SQL name and snake_case names, `text({ length, enum })`, relations declared only with v1 `relations` (no foreign keys), `$defaultFn` ids from `@oslojs/crypto`, and the legacy object-style extra config

## Changes from upstream

- `sqliteTable` is `snakeCase.table`. Upstream gets snake_case column names from `casing: "snake_case"` in `drizzle.config.ts`, which drizzle-kit 1.0 no longer reads; `snakeCase.table` is how drizzle-orm 1.0 says the same.
- `relations` is imported from `drizzle-orm/_relations` instead of `drizzle-orm`, because drizzle-orm 1.0 moved the v1 relational API there.
- `src/lib/id.ts` is copied as it is, with `@oslojs/crypto`, and `tsconfig.json` keeps upstream's `@/*` path alias.
- `drizzle.config.ts` drops `casing` and the database credentials.

## Generate

```bash
bazel run //apps/drizzler/examples/devterms:docs.write
```

The output is written to [`docs/`](docs/):

- [`README.md`](docs/README.md): table docs with the SQL
- [`er.svg`](docs/er.svg), [`er.png`](docs/er.png), [`er-dark.png`](docs/er-dark.png) and `tables/*.svg`: ER diagrams
- [`schema.dbml`](docs/schema.dbml): DBML
- [`er.md`](docs/er.md): Mermaid `erDiagram`

## Known issues

None.
