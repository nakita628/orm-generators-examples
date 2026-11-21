# mylists

The schema of MyLists, a self-hosted app to organize series, anime, movies, games, books and manga: Better Auth users and OAuth, per-media tables (the media, user lists, genres, people, tags), achievements, notifications, collections, imports, admin tasks, feature votes and more.

- Source: [vincedelmas/MyLists `src/lib/server/database`](https://github.com/vincedelmas/MyLists/tree/324deef6197f205fd41e205df89d8e08a4ab894f/src/lib/server/database) (`schema/` and `custom-types.ts`), with [`src/lib/utils/enums.ts`](https://github.com/vincedelmas/MyLists/blob/324deef6197f205fd41e205df89d8e08a4ab894f/src/lib/utils/enums.ts) and [`constants.ts`](https://github.com/vincedelmas/MyLists/blob/324deef6197f205fd41e205df89d8e08a4ab894f/src/lib/utils/constants.ts) (GPL-3.0)
- Dialect: SQLite
- Covers: 82 tables across 22 files, media tables built from shared helpers (`...commonMediaListCols(books.id, MediaType.BOOKS)`) whose foreign keys point at a parameter, indexes and `check` constraints returned by a function in the third argument, `customType` columns built by functions that take arguments (`imageUrl(name, coverType)`), `as const` enum objects, a tsconfig `paths` alias (`@/*`), and v1 `relations`

## Changes from upstream

- `sqliteTable` is `snakeCase.table`. Upstream gets snake_case names for its unnamed columns from `casing: "snake_case"` in `drizzle.config.ts`, which drizzle-kit 1.0 no longer reads.
- `relations` is imported from `drizzle-orm/_relations` instead of `drizzle-orm/relations`, because drizzle-orm 1.0 moved the v1 relational API there.
- `enums.ts` and `constants.ts` are copied as they are. The type files the schema imports from (`src/lib/types/*.types.ts`, `src/lib/schemas/dynamic-lists.schema.ts`) are stand-ins with loose types, and `src/lib/server/core/images/image-url.ts` is a stand-in for the two helpers the custom types call when reading and writing.
- `tsconfig.json` turns off `verbatimModuleSyntax`, because upstream imports types such as `SQLiteColumn` without `type`.
- `drizzle.config.ts` drops `casing`, `breakpoints`, `verbose`, `strict` and the database URL.

## Generate

```bash
bazel run //apps/drizzler/examples/mylists:docs.write
```

The output is written to [`docs/`](docs/):

- [`README.md`](docs/README.md): table docs with the SQL
- [`er.svg`](docs/er.svg), [`er.png`](docs/er.png), [`er-dark.png`](docs/er-dark.png) and `tables/*.svg`: ER diagrams
- [`schema.dbml`](docs/schema.dbml): DBML
- [`er.md`](docs/er.md): Mermaid `erDiagram`

## Known issues

- [nakita628/drizzler#61](https://github.com/nakita628/drizzler/issues/61): the indexes of the `*_list` tables, named from `MediaType.BOOKS` and the like, have no names in the DBML.
