# cove

The database schema of Cove, a minimal TanStack Start template with Better Auth: users, sessions, accounts and verifications, with drizzle-orm 1.0 relations.

- Source: [mugnavo/cove `src/lib/db/schema`](https://github.com/mugnavo/cove/tree/7d3cbfb03ea098558faf1797a1dc915b2069f436/src/lib/db/schema) (Unlicense)
- Dialect: PostgreSQL
- Covers: `defineRelationsPart` next to the tables it relates, an `index.ts` of `export *`, a `relations.ts` that reads the schema through `import * as schema from "./"`, indexes, and `$onUpdate`

## Changes from upstream

- The three files are copied as they are.
- `drizzle.config.ts` drops `breakpoints`, `verbose` and the database URL read through `varlock`.

## Generate

```bash
bazel run //apps/drizzler/examples/cove:docs.write
```

The output is written to [`docs/`](docs/):

- [`README.md`](docs/README.md): table docs with the SQL
- [`er.svg`](docs/er.svg), [`er.png`](docs/er.png), [`er-dark.png`](docs/er-dark.png) and `tables/*.svg`: ER diagrams
- [`schema.dbml`](docs/schema.dbml): DBML
- [`er.md`](docs/er.md): Mermaid `erDiagram`

## Known issues

None.
