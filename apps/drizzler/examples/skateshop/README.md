# skateshop

The schema of Skateshop, sadmann7's open-source e-commerce skateshop built with Next.js: stores, products, categories and subcategories, carts, orders, payments, addresses, customers, notifications and newsletter subscriptions.

- Source: [sadmann7/skateshop `src/db/schema`](https://github.com/sadmann7/skateshop/tree/e954d54379a7a6f798d3f49ffc6664449220adfc/src/db/schema), with [`src/lib/id.ts`](https://github.com/sadmann7/skateshop/blob/e954d54379a7a6f798d3f49ffc6664449220adfc/src/lib/id.ts) and [`src/lib/validations/cart.ts`](https://github.com/sadmann7/skateshop/blob/e954d54379a7a6f798d3f49ffc6664449220adfc/src/lib/validations/cart.ts) (MIT)
- Dialect: PostgreSQL
- Covers: a directory of one file per table behind an `index.ts`, a shared `lifecycleDates` object spread into tables, `$defaultFn(() => generateId("store"))` ids from `nanoid`, `pgEnum`, `decimal`, `json().$type<...>()` with zod-inferred types, and v1 `relations`

## Changes from upstream

- `relations` is imported from `drizzle-orm/_relations` instead of `drizzle-orm`, because drizzle-orm 1.0 moved the v1 relational API there.
- `src/lib/id.ts` and `src/lib/validations/cart.ts` are copied as they are, with `nanoid` and `zod` at upstream's versions. `src/types/index.ts` is a stand-in that holds only `StoredFile`. `tsconfig.json` keeps upstream's `@/*` path alias.
- `drizzle.config.ts` drops the database URL.

## Generate

```bash
bazel run //apps/drizzler/examples/skateshop:docs.write
```

The output is written to [`docs/`](docs/):

- [`README.md`](docs/README.md): table docs with the SQL
- [`er.svg`](docs/er.svg), [`er.png`](docs/er.png), [`er-dark.png`](docs/er-dark.png) and `tables/*.svg`: ER diagrams
- [`schema.dbml`](docs/schema.dbml): DBML
- [`er.md`](docs/er.md): Mermaid `erDiagram`

## Known issues

None.
