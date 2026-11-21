# create-t3-app-mysql

The Drizzle + NextAuth.js (Auth.js) schema that `create-t3-app` scaffolds for MySQL.

- Source: [t3-oss/create-t3-app `with-auth-mysql.ts`](https://github.com/t3-oss/create-t3-app/blob/4709861f7e67a15564c0460c13e7b4b6cfcae40d/cli/template/extras/src/server/db/schema-drizzle/with-auth-mysql.ts) (MIT)
- Dialect: MySQL
- Covers: `mysqlTableCreator`, MySQL column types, Auth.js adapter tables

## Changes from upstream

- `relations` is imported from `drizzle-orm/_relations` instead of `drizzle-orm`, because drizzle-orm 1.0 moved the v1 relational API there.

## Generate

```bash
bazel run //apps/drizzler/examples/create-t3-app-mysql:docs.write
```

The output is written to [`docs/`](docs/):

- [`README.md`](docs/README.md): table docs with the SQL
- [`er.svg`](docs/er.svg), [`er.png`](docs/er.png), [`er-dark.png`](docs/er-dark.png) and `tables/*.svg`: ER diagrams
- [`schema.dbml`](docs/schema.dbml): DBML
- [`er.md`](docs/er.md): Mermaid `erDiagram`

## Known issues

None.
