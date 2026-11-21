# morphic

The schema of Morphic, an open-source AI answer engine with a generative UI: chats, messages and their parts (text, reasoning, files, sources, tool calls), spaces, feedback, and the Better Auth tables.

- Source: [miurla/morphic `lib/db/schema.ts`](https://github.com/miurla/morphic/blob/244ea609ab0d604f6c7020fee214ac087f0129c9/lib/db/schema.ts) and [`lib/auth/better-auth/schema.ts`](https://github.com/miurla/morphic/blob/244ea609ab0d604f6c7020fee214ac087f0129c9/lib/auth/better-auth/schema.ts) (Apache-2.0)
- Dialect: PostgreSQL
- Covers: `export *` of a second schema file, `.enableRLS()` with `pgPolicy`, a `tsvector` `customType`, `$defaultFn(() => createId())` with cuid2, `varchar({ length })` keys, `check` constraints, many nullable tool-specific `json` columns, and types imported without `type` (`InferSelectModel`)

## Changes from upstream

- Both files are copied as they are.
- `tsconfig.json` turns off `verbatimModuleSyntax`, because upstream imports `InferSelectModel` without `type`.

## Generate

```bash
bazel run //apps/drizzler/examples/morphic:docs.write
```

The output is written to [`docs/`](docs/):

- [`README.md`](docs/README.md): table docs with the SQL
- [`er.svg`](docs/er.svg), [`er.png`](docs/er.png), [`er-dark.png`](docs/er-dark.png) and `tables/*.svg`: ER diagrams
- [`schema.dbml`](docs/schema.dbml): DBML
- [`er.md`](docs/er.md): Mermaid `erDiagram`

## Known issues

None.
