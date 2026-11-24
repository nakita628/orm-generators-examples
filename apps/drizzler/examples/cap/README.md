# cap

The schema of Cap, an open-source alternative to Loom: users and auth, organizations and spaces, videos and their uploads, comments, folders, shared and public collections, storage integrations, S3 buckets, notifications, developer API videos and more.

- Source: [CapSoftware/Cap `packages/database/schema.ts`](https://github.com/CapSoftware/Cap/blob/c23d12dfdf671c2732eddb101139f0f6f1f8a588/packages/database/schema.ts), [`helpers.ts`](https://github.com/CapSoftware/Cap/blob/c23d12dfdf671c2732eddb101139f0f6f1f8a588/packages/database/helpers.ts) and [`types/`](https://github.com/CapSoftware/Cap/tree/c23d12dfdf671c2732eddb101139f0f6f1f8a588/packages/database/types) (AGPL-3.0)
- Dialect: MySQL
- Covers: 58 MySQL tables in one file, `customType` columns sized from an imported constant (`varchar(${nanoIdLength})`), `int`, `bigint`, `float`, `datetime`, `longtext` and `mediumtext`, `json().$type<...>()` with branded types, `varchar({ enum })` from imported constant lists, composite and `uniqueIndex` keys, `tablesFilter` in `drizzle.config.ts`, and v1 `relations`

## Changes from upstream

- `schema.ts`, `helpers.ts` and the files in `types/` are copied as they are, except that `relations` is imported from `drizzle-orm/_relations` instead of `drizzle-orm/relations`, because drizzle-orm 1.0 moved the v1 relational API there.
- `@cap/web-domain`, which the schema imports only types from, resolves through tsconfig `paths` to a declaration stand-in in `stubs/cap-web-domain/`.
- `drizzle.config.ts` drops `casing: "snake_case"`, which drizzle-kit 1.0 no longer accepts; every column in the schema names its SQL column, so it changes nothing. It also drops the database URL and its checks.

## Generate

```bash
bazel run //apps/drizzler/examples/cap:docs.write
```

The output is written to [`docs/`](docs/):

- [`README.md`](docs/README.md): table docs with the SQL
- [`er.svg`](docs/er.svg), [`er.png`](docs/er.png), [`er-dark.png`](docs/er-dark.png) and `tables/*.svg`: ER diagrams
- [`schema.dbml`](docs/schema.dbml): DBML
- [`er.md`](docs/er.md): Mermaid `erDiagram`
