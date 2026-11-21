# t4-app

The API database schema of the T4 App, a full-stack template for iOS, Android and the web on Cloudflare (Tamagui, tRPC, Expo, Next.js): users and cars.

- Source: [timothymiller/t4-app `packages/api/src/db/schema.ts`](https://github.com/timothymiller/t4-app/blob/8f7ab4d1f0c2144248a145c46e57d2622175b912/packages/api/src/db/schema.ts) (MIT)
- Dialect: SQLite (Cloudflare D1)
- Covers: `real` columns, comments above the tables, valibot insert and select schemas built next to the tables, and types imported without `type` (`InferInsertModel`)

## Changes from upstream

- `createInsertSchema` and `createSelectSchema` are imported from `drizzle-orm/valibot` instead of `drizzle-valibot`, which drizzle-orm 1.0 replaces; `valibot` is 1.x, as `drizzle-orm/valibot` requires (upstream uses 0.20).
- `tsconfig.json` turns off `verbatimModuleSyntax`, because upstream imports `InferInsertModel` and `InferSelectModel` without `type`.

## Generate

```bash
bazel run //apps/drizzler/examples/t4-app:docs.write
```

The output is written to [`docs/`](docs/):

- [`README.md`](docs/README.md): table docs with the SQL
- [`er.svg`](docs/er.svg), [`er.png`](docs/er.png), [`er-dark.png`](docs/er-dark.png) and `tables/*.svg`: ER diagrams
- [`schema.dbml`](docs/schema.dbml): DBML
- [`er.md`](docs/er.md): Mermaid `erDiagram`

## Known issues

None.
