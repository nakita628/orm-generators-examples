# justship

The schema of JustShip, a Svelte 5 and SvelteKit SaaS boilerplate: users, sessions, email verification codes and password reset tokens for Lucia-style auth.

- Source: [ocluf/justship `src/lib/server/database/schema.ts`](https://github.com/ocluf/justship/blob/891bd37540aa2fe2866e495a187b0f5b112526f7/src/lib/server/database/schema.ts) (MIT)
- Dialect: SQLite (`dialect: "turso"`)
- Covers: `integer({ mode: "timestamp" })` and `integer({ mode: "boolean" })`, a foreign key, and `dialect: "turso"`

## Changes from upstream

- `schema.ts` is copied as it is.
- `drizzle.config.ts` drops the database credentials.

## Generate

```bash
bazel run //apps/drizzler/examples/justship:docs.write
```

The output is written to [`docs/`](docs/):

- [`README.md`](docs/README.md): table docs with the SQL
- [`er.svg`](docs/er.svg), [`er.png`](docs/er.png), [`er-dark.png`](docs/er-dark.png) and `tables/*.svg`: ER diagrams
- [`schema.dbml`](docs/schema.dbml): DBML
- [`er.md`](docs/er.md): Mermaid `erDiagram`

## Known issues

None.
