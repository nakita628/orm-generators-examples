# nelsonlai-dev

The schema of nelsonlai.dev, Nelson Lai's popular open-source personal website: Better Auth users, sessions, accounts and verifications, posts, likes, comments and their votes, guestbook messages, settings and unsubscribes.

- Source: [nelsonlaidev/nelsonlai.dev `src/db/schemas`](https://github.com/nelsonlaidev/nelsonlai.dev/tree/c0d625e37123f57703ef94977d92b08c0f8ed4ac/src/db/schemas) and [`src/db/constants.ts`](https://github.com/nelsonlaidev/nelsonlai.dev/blob/c0d625e37123f57703ef94977d92b08c0f8ed4ac/src/db/constants.ts) (MIT)
- Dialect: PostgreSQL
- Covers: indexes on `.desc()` columns and partial indexes (`.where()`), a gin index on `to_tsvector(...)`, `pgEnum`, `check` constraints, `$defaultFn(() => createId())` with cuid2, self-referencing comments, and v1 `relations`

## Changes from upstream

- `relations` is imported from `drizzle-orm/_relations` instead of `drizzle-orm`, because drizzle-orm 1.0 moved the v1 relational API there.
- `src/db/constants.ts` is copied as it is, with `@paralleldrive/cuid2` at upstream's version.
- `drizzle.config.ts` drops the env loading and the database URL.

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
