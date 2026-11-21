# midday

The schema of Midday, a business assistant for freelancers: teams, bank accounts, transactions, invoices, the inbox, documents, time tracking, OAuth apps and more. It was pulled from Supabase with `drizzle-kit pull` and has grown by hand since.

- Source: [midday-ai/midday `packages/db/src/schema.ts`](https://github.com/midday-ai/midday/blob/51587319f26a0ffaa9dfccab1920373cb65689b7/packages/db/src/schema.ts) and [`app-config.ts`](https://github.com/midday-ai/midday/blob/51587319f26a0ffaa9dfccab1920373cb65689b7/packages/db/src/app-config.ts) (AGPL-3.0)
- Dialect: PostgreSQL
- Covers: 50 tables in one file, about 40 `pgEnum`s, `customType()` (`numericCasted`, `tsvector`), `vector({ dimensions })`, `` generatedAlwaysAs(sql`...`) ``, `pgPolicy` row-level security, `index().using("gin", ...)` with operator classes, composite primary and foreign keys, a self-referencing foreign key, a dotted table name (`auth.users`), and v1 `relations`

## Changes from upstream

- `relations` is imported from `drizzle-orm/_relations` instead of `drizzle-orm`, because drizzle-orm 1.0 moved the v1 relational API there.
- The check constraint on `auth.users` is written as `` check("users_email_change_confirm_status_check", sql`...`) ``. Upstream returns a plain `{ kind: "check", name, expression }` object from the extra config, which drizzle-orm 0.x ignored and drizzle-orm 1.0 rejects (`Invalid extra config value for table "auth.users"`).

## Generate

```bash
bazel run //apps/drizzler/examples/midday:docs.write
```

The output is written to [`docs/`](docs/):

- [`README.md`](docs/README.md): table docs with the SQL
- [`er.svg`](docs/er.svg), [`er.png`](docs/er.png), [`er-dark.png`](docs/er-dark.png) and `tables/*.svg`: ER diagrams
- [`schema.dbml`](docs/schema.dbml): DBML
- [`er.md`](docs/er.md): Mermaid `erDiagram`

## Known issues

None.
