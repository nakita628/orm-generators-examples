# lipi

The schema of Lipi, a work-in-progress SaaS Notion replica with real-time collaboration: Auth.js users, accounts and verification tokens, and the app's workspaces, folders, files, collaborators, and Stripe customers, products, prices and subscriptions.

- Source: [rajput-hemant/lipi `lib/db/schema`](https://github.com/rajput-hemant/lipi/tree/eac9a79adf601491372d8bc376b36def3c50d67b/lib/db/schema), [`lib/db/table-creator.ts`](https://github.com/rajput-hemant/lipi/blob/eac9a79adf601491372d8bc376b36def3c50d67b/lib/db/table-creator.ts) and [`config/site.ts`](https://github.com/rajput-hemant/lipi/blob/eac9a79adf601491372d8bc376b36def3c50d67b/config/site.ts) (MIT)
- Dialect: PostgreSQL
- Covers: a directory `schema`, `pgTableCreator` whose prefix is computed from an imported config (`siteConfig.name.toLowerCase()`), a separate `enums.ts` of `pgEnum`s, `$type<AdapterAccount["type"]>()` from `@auth/core`, and v1 `relations`

## Changes from upstream

- `relations` is imported from `drizzle-orm/_relations` instead of `drizzle-orm`, because drizzle-orm 1.0 moved the v1 relational API there.
- `config/site.ts` is copied as it is, and `tsconfig.json` keeps upstream's `@/*` path alias. `@auth/core` is installed at upstream's version for its types.
- `drizzle.config.ts` writes the `tablesFilter` upstream computes from `siteConfig` as `lipi_*`, and drops the database URL.

## Generate

```bash
bazel run //apps/drizzler/examples/lipi:docs.write
```

The output is written to [`docs/`](docs/):

- [`README.md`](docs/README.md): table docs with the SQL
- [`er.svg`](docs/er.svg), [`er.png`](docs/er.png), [`er-dark.png`](docs/er-dark.png) and `tables/*.svg`: ER diagrams
- [`schema.dbml`](docs/schema.dbml): DBML
- [`er.md`](docs/er.md): Mermaid `erDiagram`

## Known issues

None.
