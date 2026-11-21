# orbitkit

The database package of Orbitkit, a monorepo starter kit with Lucia-style auth: users, sessions and OAuth accounts.

- Source: [ixahmedxi/orbitkit `packages/db/src/schema`](https://github.com/ixahmedxi/orbitkit/tree/0207fdfddf22d0f31fc88c51d275e5b4a56e0c98/packages/db/src/schema) (MIT)
- Dialect: PostgreSQL
- Covers: tables that import each other through their own `index.ts` (`import { userTable } from '.'`), composite primary keys, and `timestamp({ withTimezone: true })`

## Changes from upstream

- The schema files are copied as they are.
- `drizzle.config.ts` drops the database URL.

## Generate

```bash
bazel run //apps/drizzler/examples/orbitkit:docs.write
```

The output is written to [`docs/`](docs/):

- [`README.md`](docs/README.md): table docs with the SQL
- [`er.svg`](docs/er.svg), [`er.png`](docs/er.png), [`er-dark.png`](docs/er-dark.png) and `tables/*.svg`: ER diagrams
- [`schema.dbml`](docs/schema.dbml): DBML
- [`er.md`](docs/er.md): Mermaid `erDiagram`

## Known issues

None.
