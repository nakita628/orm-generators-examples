# cut-it

The schema of Cut It, a link shortener built with the Next.js App Router and server actions: Auth.js users, accounts, sessions and verification tokens, links and link visits.

- Source: [mehrabmp/cut-it `src/server/db/schema.ts`](https://github.com/mehrabmp/cut-it/blob/9733e9808a0ea6f1bf5f5f80a466de7afb1c6d95/src/server/db/schema.ts) (MIT)
- Dialect: SQLite (Turso)
- Covers: the Auth.js tables with composite primary keys, `$defaultFn(() => createId())` with cuid2, `$type<AdapterAccount["type"]>()` from `next-auth/adapters`, indexes, and v1 `relations`

## Changes from upstream

- `relations` is imported from `drizzle-orm/_relations` instead of `drizzle-orm`, because drizzle-orm 1.0 moved the v1 relational API there.
- `next-auth` is this repo's v5 beta (upstream uses v4) for the `AdapterAccount` type, and `@paralleldrive/cuid2` is at upstream's version.
- `drizzle.config.ts` is rewritten for drizzle-kit 1.0: upstream's uses the old `driver: "turso"` and `dbCredentials` options.

## Generate

```bash
bazel run //apps/drizzler/examples/cut-it:docs.write
```

The output is written to [`docs/`](docs/):

- [`README.md`](docs/README.md): table docs with the SQL
- [`er.svg`](docs/er.svg), [`er.png`](docs/er.png), [`er-dark.png`](docs/er-dark.png) and `tables/*.svg`: ER diagrams
- [`schema.dbml`](docs/schema.dbml): DBML
- [`er.md`](docs/er.md): Mermaid `erDiagram`
