# karakeep

The schema of Karakeep (formerly Hoarder), a self-hosted bookmark manager: users, bookmarks and their links, texts and assets, tags, lists, highlights, feeds, rules, webhooks, backups and more.

- Source: [karakeep-app/karakeep `packages/db/schema.ts`](https://github.com/karakeep-app/karakeep/blob/bcc4b90a2ae89f4b96ec23ce395ee9d3d1df6a72/packages/db/schema.ts) (AGPL-3.0)
- Dialect: SQLite
- Covers: 34 tables in one file, columns built by helper functions (`createdAtField()`), `$defaultFn(() => createId())` with cuid2, `text({ enum })`, `integer` with `mode: "timestamp"`, `"timestamp_ms"` and `"boolean"`, `const enum`s, a type imported without `type` (`AnySQLiteColumn`), self-references, composite keys, JSDoc column notes, and v1 `relations`

## Changes from upstream

- `relations` is imported from `drizzle-orm/_relations` instead of `drizzle-orm`, because drizzle-orm 1.0 moved the v1 relational API there.
- The `@karakeep/shared/types/*` imports point at stand-ins in `src/shared/types/` that hold only what the schema imports: `API_KEY_FULL_ACCESS_SCOPE` and `BookmarkTypes` are copied as they are, and the zod-inferred `ZApiKeyScope` and `ZReaderViewReason` are `string`.
- `tsconfig.json` turns off `verbatimModuleSyntax`, because upstream imports `AnySQLiteColumn` and `SQL` without `type`.

## Generate

```bash
bazel run //apps/drizzler/examples/karakeep:docs.write
```

The output is written to [`docs/`](docs/):

- [`README.md`](docs/README.md): table docs with the SQL
- [`er.svg`](docs/er.svg), [`er.png`](docs/er.png), [`er-dark.png`](docs/er-dark.png) and `tables/*.svg`: ER diagrams
- [`schema.dbml`](docs/schema.dbml): DBML
- [`er.md`](docs/er.md): Mermaid `erDiagram`
