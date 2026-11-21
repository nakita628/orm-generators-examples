# better-auth-cloudflare

The database schema of the Hono example of better-auth-cloudflare, which integrates Better Auth with Cloudflare Workers, D1 and KV: users, sessions, accounts, verifications and user files.

- Source: [zpg6/better-auth-cloudflare `examples/hono/src/db`](https://github.com/zpg6/better-auth-cloudflare/tree/c8af01e9e4926c970c0d3fc3acbe2f1776fe2409/examples/hono/src/db) (MIT)
- Dialect: SQLite (Cloudflare D1)
- Covers: a schema entry (`index.ts`) that also re-exports all of `drizzle-orm` (`export * from "drizzle-orm"`), a `schema` object spread from `import * as authSchema`, the Better Auth tables, and v1 `relations`

## Changes from upstream

- The three files are copied as they are, except that `relations` is imported from `drizzle-orm/_relations` instead of `drizzle-orm`, because drizzle-orm 1.0 moved the v1 relational API there.
- `drizzle.config.ts` drops the D1 credentials.

## Generate

```bash
bazel run //apps/drizzler/examples/better-auth-cloudflare:docs.write
```

The output is written to [`docs/`](docs/):

- [`README.md`](docs/README.md): table docs with the SQL
- [`er.svg`](docs/er.svg), [`er.png`](docs/er.png), [`er-dark.png`](docs/er-dark.png) and `tables/*.svg`: ER diagrams
- [`schema.dbml`](docs/schema.dbml): DBML
- [`er.md`](docs/er.md): Mermaid `erDiagram`

## Known issues

None.
