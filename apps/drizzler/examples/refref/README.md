# refref

The core database schema of RefRef, an open-source referral and affiliate marketing platform: Better Auth users, sessions, organizations and members, products, programs, participants, referral codes and links, events, rewards and their rules, and API keys.

- Source: [amicalhq/refref `packages/coredb/src/schema.ts`](https://github.com/amicalhq/refref/blob/81af934fec3b20990a4d9af7ed472d0d14d73a82/packages/coredb/src/schema.ts) and [`packages/id/src/index.ts`](https://github.com/amicalhq/refref/blob/81af934fec3b20990a4d9af7ed472d0d14d73a82/packages/id/src/index.ts) (AGPL-3.0)
- Dialect: PostgreSQL
- Covers: 21 tables that each start with `...baseFields("user")`, an arrow function returning the key and timestamps, prefixed ids from a workspace package (`@refref/id`), `jsonb().$type<...>()` config columns, ``default(sql`now()`)``, many foreign keys, and v1 `relations`

## Changes from upstream

- `relations` is imported from `drizzle-orm/_relations` instead of `drizzle-orm`, because drizzle-orm 1.0 moved the v1 relational API there.
- `@refref/id` resolves through tsconfig `paths` to a copy of `packages/id/src/index.ts`, with `@paralleldrive/cuid2` at its version. `@refref/types`, which the schema imports only types from, resolves to a stand-in with loose types.
- `drizzle.config.ts` drops the database URL and its check.

## Generate

```bash
bazel run //apps/drizzler/examples/refref:docs.write
```

The output is written to [`docs/`](docs/):

- [`README.md`](docs/README.md): table docs with the SQL
- [`er.svg`](docs/er.svg), [`er.png`](docs/er.png), [`er-dark.png`](docs/er-dark.png) and `tables/*.svg`: ER diagrams
- [`schema.dbml`](docs/schema.dbml): DBML
- [`er.md`](docs/er.md): Mermaid `erDiagram`

## Known issues

None.
