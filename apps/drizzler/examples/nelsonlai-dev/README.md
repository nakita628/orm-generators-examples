# nelsonlai-dev

The schema of nelsonlai.dev, Nelson Lai's popular open-source personal website: Better Auth users (with the admin plugin's role and ban columns), sessions, accounts and verifications, posts and their likes, comments and their votes, guestbook messages, comment notification preferences and comment reply unsubscribes, with drizzle-orm 1.0 relations.

- Source: [nelsonlaidev/nelsonlai.dev `packages/db/src/schema`](https://github.com/nelsonlaidev/nelsonlai.dev/tree/3cf0c572d3d6b20ebf0092435efb70901d18af8f/packages/db/src/schema) and [`relations.ts`](https://github.com/nelsonlaidev/nelsonlai.dev/blob/3cf0c572d3d6b20ebf0092435efb70901d18af8f/packages/db/src/relations.ts) (MIT)
- Dialect: PostgreSQL
- Covers: `defineRelations` in its own file (one-to-one, one-to-many, and an `alias` for the self-referencing comments), indexes on `.desc().nullsFirst()` columns, a `check` constraint, composite primary keys, a self-referencing `foreignKey` with `onDelete('restrict')`, `$defaultFn(() => createId())` with cuid2, and column helpers that wrap `timestamp` (`timestamptz`, `createdAt`, `updatedAt` with `$onUpdateFn`)

## Changes from upstream

- `schema/` and `relations.ts` are copied as they are, into `src/`.
- `drizzle.config.ts` lists `relations.ts` next to `schema/index.ts` in `schema`, so drizzler reads the relations too, and drops the env loading and the database URL.

## Generate

```bash
bazel run //apps/drizzler/examples/nelsonlai-dev:docs.write
```

The output is written to [`docs/`](docs/):

- [`README.md`](docs/README.md): table docs with the SQL
- [`er.svg`](docs/er.svg), [`er.png`](docs/er.png), [`er-dark.png`](docs/er-dark.png) and `tables/*.svg`: ER diagrams
- [`schema.dbml`](docs/schema.dbml): DBML
- [`er.md`](docs/er.md): Mermaid `erDiagram`

## Known issues

None.
