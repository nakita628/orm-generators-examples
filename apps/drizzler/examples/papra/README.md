# papra

The schema of Papra, a self-hosted document management and archiving platform: users and Better Auth, organizations, members and invitations, documents, tags and tagging rules, custom properties, intake emails, share links, webhooks, API keys, plans and subscriptions, AI credits, and a key-value store.

- Source: the `*.table.ts` and `*.tables.ts` files of [papra-hq/papra `apps/papra-server/src/modules`](https://github.com/papra-hq/papra/tree/648c223b807500f0258c822e0321d6ef4964f064/apps/papra-server/src/modules), with the constants, column helpers and id generator they import, and [`packages/webhooks/src/webhooks.constants.ts`](https://github.com/papra-hq/papra/blob/648c223b807500f0258c822e0321d6ef4964f064/packages/webhooks/src/webhooks.constants.ts) (AGPL-3.0)
- Dialect: SQLite (`dialect: "turso"`)
- Covers: a `schema` array of two globs spread over feature folders, `dialect: "turso"`, columns spread from helper functions (`...createPrimaryKeyField({ prefix })`, `...createTimestampColumns()`, which spreads two more helpers), `$default` ids with cuid2, `text({ enum })` from imported constant lists, `integer` with `mode: "timestamp_ms"`, an FTS5 table, and a tsconfig `paths` alias for a workspace package (`@papra/webhooks`)

## Changes from upstream

- Only the table files and what they import at runtime are copied, as they are: the module constants, `shared/db/columns.helpers.ts`, `shared/random/ids.ts`, `shared/random/ids.constants.models.ts`, `app/database/database.constants.ts`, `shared/utils.ts` and `shared/types.ts`.
- `documents.models.ts`, `document-share-links.models.ts` and `organizations.models.ts` are stand-ins that hold only the id generator each table imports, written as upstream writes it. The `*.types.ts` files and `plan-entitlements.registry.ts` the tables import types from are stand-ins with loose types.
- `@papra/webhooks` resolves through tsconfig `paths` to `stubs/papra-webhooks/`, which re-exports a copy of its `webhooks.constants.ts`.
- `drizzle.config.ts` drops the database credentials and the migrations folder.

## Generate

```bash
bazel run //apps/drizzler/examples/papra:docs.write
```

The output is written to [`docs/`](docs/):

- [`README.md`](docs/README.md): table docs with the SQL
- [`er.svg`](docs/er.svg), [`er.png`](docs/er.png), [`er-dark.png`](docs/er-dark.png) and `tables/*.svg`: ER diagrams
- [`schema.dbml`](docs/schema.dbml): DBML
- [`er.md`](docs/er.md): Mermaid `erDiagram`

## Known issues

None.
