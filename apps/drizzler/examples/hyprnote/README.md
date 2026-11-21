# hyprnote

The local SQLite schema of Hyprnote, an open-source AI notepad for meetings: sessions and their transcripts, participants, templates, chats, calendars and events, tags, the local library and its attachments, and app settings.

- Source: [fastrepl/hyprnote `packages/db/src/schema.ts`](https://github.com/fastrepl/hyprnote/blob/9c1cbe287019e7ca199ee77bd5f6d095cf076ad8/packages/db/src/schema.ts) and [`drizzle.config.ts`](https://github.com/fastrepl/hyprnote/blob/9c1cbe287019e7ca199ee77bd5f6d095cf076ad8/packages/db/drizzle.config.ts) (MIT)
- Dialect: SQLite
- Covers: 35 tables with no foreign keys, a `sql` default held in a constant (`.default(currentTimestamp)`), partial unique indexes (`` uniqueIndex().on().where(sql`...`) ``), composite primary keys, `integer({ mode: "boolean" })`, and a `drizzle.config.ts` with no `schema`

## Changes from upstream

- `schema.ts` is copied as it is.
- `drizzle.config.ts` drops the database path; like upstream's, it has no `schema`, so `generate` passes `src/schema.ts` to drizzler (without it, drizzler stops with `No schema to read`).

## Generate

```bash
bazel run //apps/drizzler/examples/hyprnote:docs.write
```

The output is written to [`docs/`](docs/):

- [`README.md`](docs/README.md): table docs with the SQL
- [`er.svg`](docs/er.svg), [`er.png`](docs/er.png), [`er-dark.png`](docs/er-dark.png) and `tables/*.svg`: ER diagrams
- [`schema.dbml`](docs/schema.dbml): DBML
- [`er.md`](docs/er.md): Mermaid `erDiagram`

## Known issues

None.
