# saas-starter

The schema of the Next.js SaaS Starter: users, teams, team members, activity logs and invitations.

- Source: [nextjs/saas-starter `lib/db/schema.ts`](https://github.com/nextjs/saas-starter/blob/6e33e58b1e553a41fe22e6b941a7229a002de361/lib/db/schema.ts) (no license declared)
- Dialect: PostgreSQL
- Covers: `serial` primary keys, snake_case SQL names, v1 `relations` with a renamed `one()` (`invitedBy`), a nullable foreign key, and a TypeScript `enum` exported next to the tables

## Changes from upstream

- `relations` is imported from `drizzle-orm/_relations` instead of `drizzle-orm`, because drizzle-orm 1.0 moved the v1 relational API there.

## Generate

```bash
bazel run //apps/drizzler/examples/saas-starter:docs.write
```

The output is written to [`docs/`](docs/):

- [`README.md`](docs/README.md): table docs with the SQL
- [`er.svg`](docs/er.svg), [`er.png`](docs/er.png), [`er-dark.png`](docs/er-dark.png) and `tables/*.svg`: ER diagrams
- [`schema.dbml`](docs/schema.dbml): DBML
- [`er.md`](docs/er.md): Mermaid `erDiagram`

## Known issues

None.
