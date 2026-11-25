# openstatus

The schema of openstatus, an open-source synthetic monitoring and status page platform: workspaces and users, monitors and their runs, statuses, tags and groups, status pages and their components, status reports, incidents and maintenances, notifications, subscribers, private locations, API keys, OAuth, and more.

- Source: [openstatusHQ/openstatus `packages/db/src/schema`](https://github.com/openstatusHQ/openstatus/tree/f2f3cd92512443f5d8a415fc32d5850c8a487efd/packages/db/src/schema), with the workspace packages it imports: [`assertions`](https://github.com/openstatusHQ/openstatus/tree/f2f3cd92512443f5d8a415fc32d5850c8a487efd/packages/assertions), [`locales`](https://github.com/openstatusHQ/openstatus/tree/f2f3cd92512443f5d8a415fc32d5850c8a487efd/packages/locales), [`regions`](https://github.com/openstatusHQ/openstatus/tree/f2f3cd92512443f5d8a415fc32d5850c8a487efd/packages/regions), [`theme-store`](https://github.com/openstatusHQ/openstatus/tree/f2f3cd92512443f5d8a415fc32d5850c8a487efd/packages/theme-store) and [`utils`](https://github.com/openstatusHQ/openstatus/tree/f2f3cd92512443f5d8a415fc32d5850c8a487efd/packages/utils) (AGPL-3.0)
- Dialect: SQLite (`dialect: "turso"`)
- Covers: 62 tables in a directory of feature folders, each with an `index.ts` that re-exports the table, its constants and its zod validation built with `createInsertSchema` / `createSelectSchema`, workspace packages resolved through tsconfig `paths`, `text({ enum })` from imported constant lists, `check` constraints, composite keys, and v1 `relations`

## Changes from upstream

- `relations` is imported from `drizzle-orm/_relations` instead of `drizzle-orm`, and `createInsertSchema` / `createSelectSchema` from `drizzle-orm/zod` instead of `drizzle-zod`, because drizzle-orm 1.0 moved and replaced them.
- The workspace packages are copied as they are, without their tests, and resolve through tsconfig `paths`. `zod`, `date-fns`, `jsonpath-plus` and `next-auth` are installed at upstream's catalog versions (`next-auth` at this repo's).
- The tests in `schema/` are left out.
- `drizzle.config.ts` drops the database credentials and `strict`.

## Generate

```bash
bazel run //apps/drizzler/examples/openstatus:docs.write
```

The output is written to [`docs/`](docs/):

- [`README.md`](docs/README.md): table docs with the SQL
- [`er.svg`](docs/er.svg), [`er.png`](docs/er.png), [`er-dark.png`](docs/er-dark.png) and `tables/*.svg`: ER diagrams
- [`schema.dbml`](docs/schema.dbml): DBML
- [`er.md`](docs/er.md): Mermaid `erDiagram`

## Known issues

None.
