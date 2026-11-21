# pages-cms

The schema of Pages CMS, an open-source CMS for static sites on GitHub: Better Auth users, sessions, accounts and verifications, GitHub installation tokens, collaborators and invites, repository configs, the file and permission caches, and action runs.

- Source: [pages-cms/pages-cms `db/schema.ts`](https://github.com/pages-cms/pages-cms/blob/6f4e860a35d934406580287e7042e5e111e207a1/db/schema.ts) (MIT)
- Dialect: PostgreSQL
- Covers: tables declared as unexported `const`s and exported together at the end (`export { userTable, ... }`), `serial` and `bigint` keys, many composite indexes and `uniqueIndex`, and the legacy object-style extra config

## Changes from upstream

- `schema.ts` is copied as it is.
- `drizzle.config.ts` drops `strict` and `verbose`, which drizzle-kit 1.0 no longer accepts, and the env loading and database URL.

## Generate

```bash
bazel run //apps/drizzler/examples/pages-cms:docs.write
```

The output is written to [`docs/`](docs/):

- [`README.md`](docs/README.md): table docs with the SQL
- [`er.svg`](docs/er.svg), [`er.png`](docs/er.png), [`er-dark.png`](docs/er-dark.png) and `tables/*.svg`: ER diagrams
- [`schema.dbml`](docs/schema.dbml): DBML
- [`er.md`](docs/er.md): Mermaid `erDiagram`

## Known issues

None.
