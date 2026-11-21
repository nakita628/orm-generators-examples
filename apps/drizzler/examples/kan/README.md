# kan

The schema of Kan, an open-source Trello alternative: workspaces, boards, lists, cards, labels, checklists, comments, permissions, imports, integrations, webhooks and notifications.

- Source: [kanbn/kan `packages/db/src/schema`](https://github.com/kanbn/kan/tree/49a6179591148d36faff5355088664b9ca703fc1/packages/db/src/schema) and [`packages/shared/src/utils/workspacePlans.ts`](https://github.com/kanbn/kan/blob/49a6179591148d36faff5355088664b9ca703fc1/packages/shared/src/utils/workspacePlans.ts) (AGPL-3.0)
- Dialect: PostgreSQL
- Covers: a directory `schema` with an `index.ts` that re-exports every file, `.enableRLS()` on almost every table, `pgEnum` built from an imported `as const` array, `bigserial`/`bigint` keys, `uuid` defaults from `` sql`uuid_generate_v4()` ``, many named v1 relations (`relationName`), and the legacy object-style extra config

## Changes from upstream

- `relations` is imported from `drizzle-orm/_relations` instead of `drizzle-orm`, because drizzle-orm 1.0 moved the v1 relational API there.
- `workspacePlans` and `WorkspacePlan` come from a copy of `@kan/shared`'s `workspacePlans.ts` in `src/shared/`, imported by a relative path.

## Generate

```bash
bazel run //apps/drizzler/examples/kan:docs.write
```

The output is written to [`docs/`](docs/):

- [`README.md`](docs/README.md): table docs with the SQL
- [`er.svg`](docs/er.svg), [`er.png`](docs/er.png), [`er-dark.png`](docs/er-dark.png) and `tables/*.svg`: ER diagrams
- [`schema.dbml`](docs/schema.dbml): DBML
- [`er.md`](docs/er.md): Mermaid `erDiagram`

## Known issues

None.
