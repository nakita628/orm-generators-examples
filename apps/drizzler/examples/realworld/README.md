# realworld

The [RealWorld](https://github.com/gothinkster/realworld) (Conduit) schema from bedstack, a RealWorld backend built with ElysiaJS and Drizzle.

- Source: [bedstack/elysia-drizzle-realworld-example `src/*/*.schema.ts`](https://github.com/bedstack/elysia-drizzle-realworld-example/tree/f642e24f8703cf61c66890dda4979fc55f442a4e/src) (MIT)
- Dialect: PostgreSQL
- Covers: a schema split across feature folders, a glob `schema` in `drizzle.config.ts`, imports between schema files, composite primary keys, `pgTable` imported under an alias (`pgTable as table`), named v1 relations (`relationName`)

## Changes from upstream

- `relations` is imported from `drizzle-orm/_relations` instead of `drizzle-orm`, because drizzle-orm 1.0 moved the v1 relational API there.
- The `@/…` path aliases are rewritten as relative imports (`../users/users.schema.ts`).

## Generate

```bash
bazel run //apps/drizzler/examples/realworld:docs.write
```

The output is written to [`docs/`](docs/):

- [`README.md`](docs/README.md): table docs with the SQL
- [`er.svg`](docs/er.svg), [`er.png`](docs/er.png), [`er-dark.png`](docs/er-dark.png) and `tables/*.svg`: ER diagrams
- [`schema.dbml`](docs/schema.dbml): DBML
- [`er.md`](docs/er.md): Mermaid `erDiagram`

## Known issues

None.
