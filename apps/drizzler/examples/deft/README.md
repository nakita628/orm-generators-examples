# deft

The schema of Deft, a self-hostable workspace where humans and AI agents share chat, tasks, projects, email and apps: auth, orgs and users, spaces and messages, tasks and projects, agents and their runs, a full email client, app integrations, events, and more.

- Source: [Maneek21/Deft `packages/db/src`](https://github.com/Maneek21/Deft/tree/6b8965cb81f3349b8b46e1d4944458dd1c7c9629/packages/db/src) (AGPL-3.0)
- Dialect: PostgreSQL
- Covers: 237 tables across 31 files re-exported from `schema.ts` with `.js` extensions, column helpers spread into tables (`...id()`, `...orgId()`, `...scope()`), many composite primary and foreign keys, `check` constraints, `vector` and `bytea` (`customType`) columns, and `pgEnum`

## Changes from upstream

- `relations` is imported from `drizzle-orm/_relations` instead of `drizzle-orm`, because drizzle-orm 1.0 moved the v1 relational API there.
- `src/index.ts`, which creates the `node-postgres` client, is left out.
- `drizzle.config.ts` drops the database URL.

## Generate

```bash
bazel run //apps/drizzler/examples/deft:docs.write
```

The output is written to [`docs/`](docs/):

- [`README.md`](docs/README.md): table docs with the SQL
- [`er.svg`](docs/er.svg), [`er.png`](docs/er.png), [`er-dark.png`](docs/er-dark.png) and `tables/*.svg`: ER diagrams
- [`schema.dbml`](docs/schema.dbml): DBML
- [`er.md`](docs/er.md): Mermaid `erDiagram`

## Known issues

None.
