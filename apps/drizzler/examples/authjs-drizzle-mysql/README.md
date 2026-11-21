# authjs-drizzle-mysql

The MySQL schema from Auth.js's official Drizzle adapter: users, accounts, sessions, verification tokens, and WebAuthn authenticators. The sibling of [authjs-drizzle-sqlite](../authjs-drizzle-sqlite).

- Source: [nextauthjs/next-auth `packages/adapter-drizzle/src/lib/mysql.ts`](https://github.com/nextauthjs/next-auth/blob/a1a16a5a7780488c7449feece410033f445d0b31/packages/adapter-drizzle/src/lib/mysql.ts) (ISC)
- Dialect: MySQL
- Covers: tables built inside `defineTables()`, the legacy object-style extra config, `varchar({ length })` keys, `timestamp` with `mode: "date", fsp: 3`, `int` and `boolean`, composite primary keys, and `onDelete: "cascade"`

## Changes from upstream

- The adapter implementation and the `satisfies` table types are left out. The table calls inside `defineTables` are unchanged.
- `AdapterAccountType` is declared locally instead of imported from `@auth/core/adapters`. It is only used as a `$type<>` argument.
- `defineTables()` is called at the bottom of the module and its tables are exported, so the tables exist when the file loads. Upstream builds them only when the adapter runs.

## Generate

```bash
bazel run //apps/drizzler/examples/authjs-drizzle-mysql:docs.write
```

The output is written to [`docs/`](docs/):

- [`README.md`](docs/README.md): table docs with the SQL
- [`er.svg`](docs/er.svg), [`er.png`](docs/er.png), [`er-dark.png`](docs/er-dark.png) and `tables/*.svg`: ER diagrams
- [`schema.dbml`](docs/schema.dbml): DBML
- [`er.md`](docs/er.md): Mermaid `erDiagram`

## Known issues

None.
