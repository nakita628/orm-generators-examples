# next-saas-stripe-starter

The schema of mickasmt's Next SaaS Stripe Starter, an open-source SaaS starter with user roles, organizations and an admin panel: Better Auth users, sessions, accounts and verifications, organizations, members and invitations, and Stripe subscriptions.

- Source: [mickasmt/next-saas-stripe-starter `lib/db/schema`](https://github.com/mickasmt/next-saas-stripe-starter/tree/c6df229ae73c207e87356d5e592779c511a9b56d/lib/db/schema) (MIT)
- Dialect: PostgreSQL
- Covers: the Better Auth organization plugin tables with indexes, an `index.ts` of `export *`, and v1 `relations`

## Changes from upstream

- `relations` is imported from `drizzle-orm/_relations` instead of `drizzle-orm`, because drizzle-orm 1.0 moved the v1 relational API there.
- `drizzle.config.ts` drops the env loading and the database URL.

## Generate

```bash
bazel run //apps/drizzler/examples/next-saas-stripe-starter:docs.write
```

The output is written to [`docs/`](docs/):

- [`README.md`](docs/README.md): table docs with the SQL
- [`er.svg`](docs/er.svg), [`er.png`](docs/er.png), [`er-dark.png`](docs/er-dark.png) and `tables/*.svg`: ER diagrams
- [`schema.dbml`](docs/schema.dbml): DBML
- [`er.md`](docs/er.md): Mermaid `erDiagram`

## Known issues

None.
