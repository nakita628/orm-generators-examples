# onestopshop

The schema of OneStopShop, an online marketplace built with the Next.js 13 App Router and PlanetScale: stores, products, carts, payments, orders and addresses.

- Source: [jackblatch/OneStopShop `db/schema.ts`](https://github.com/jackblatch/OneStopShop/blob/31208da07043b369e02f55d3a77ecebb64056577/db/schema.ts) (MIT)
- Dialect: MySQL (PlanetScale)
- Covers: MySQL tables with no foreign keys (PlanetScale), `serial` keys, `decimal`, `json` and `boolean`, and `uniqueIndex` in the legacy object-style extra config

## Changes from upstream

- The `InferModel<typeof table>` types, which drizzle-orm 1.0 removed, are written as `typeof table.$inferSelect`.
- `drizzle.config.ts` is added, pointing at `db/schema.ts`; upstream has none.

## Generate

```bash
bazel run //apps/drizzler/examples/onestopshop:docs.write
```

The output is written to [`docs/`](docs/):

- [`README.md`](docs/README.md): table docs with the SQL
- [`er.svg`](docs/er.svg), [`er.png`](docs/er.png), [`er-dark.png`](docs/er-dark.png) and `tables/*.svg`: ER diagrams
- [`schema.dbml`](docs/schema.dbml): DBML
- [`er.md`](docs/er.md): Mermaid `erDiagram`

## Known issues

None.
