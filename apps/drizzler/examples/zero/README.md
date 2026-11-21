# zero

The server schema of Zero (Mail-0), an open-source AI email client: Better Auth users, sessions and accounts, mail connections (Gmail, Outlook, IMAP), user settings, notes, writing styles, hotkeys and the assistant.

- Source: [Mail-0/Zero `apps/server/src/db/schema.ts`](https://github.com/Mail-0/Zero/blob/4b642dfd9890a1fdcb1b98a177fe74ca3142c445/apps/server/src/db/schema.ts) and [`apps/server/src/lib/schemas.ts`](https://github.com/Mail-0/Zero/blob/4b642dfd9890a1fdcb1b98a177fe74ca3142c445/apps/server/src/lib/schemas.ts) (MIT)
- Dialect: PostgreSQL
- Covers: `pgTableCreator` with a `mail0_` prefix, `tablesFilter` in `drizzle.config.ts`, a `jsonb` default built from a zod-typed object in another file, JSDoc table notes, composite primary keys, `unique().on(...)` and indexes on `jsonb`

## Changes from upstream

- `src/lib/schemas.ts` is copied as it is, with `zod` 3 as a dependency.
- `drizzle.config.ts` drops the database URL and the migrations folder.

## Generate

```bash
bazel run //apps/drizzler/examples/zero:docs.write
```

The output is written to [`docs/`](docs/):

- [`README.md`](docs/README.md): table docs with the SQL
- [`er.svg`](docs/er.svg), [`er.png`](docs/er.png), [`er-dark.png`](docs/er-dark.png) and `tables/*.svg`: ER diagrams
- [`schema.dbml`](docs/schema.dbml): DBML
- [`er.md`](docs/er.md): Mermaid `erDiagram`

## Known issues

None.
