# survev

The API database schema of survev, an open-source recreation of the 2D battle royale surviv.io: users and sessions, unlocked items, match data, IP logs, bans and reports.

- Source: [survev/survev `server/src/api/db/schema.ts`](https://github.com/survev/survev/blob/3edee7e8bd07fda8a63905aab4c45b595b684a80/server/src/api/db/schema.ts) and the [`shared/`](https://github.com/survev/survev/tree/3edee7e8bd07fda8a63905aab4c45b595b684a80/shared) directory it imports from (GPL-3.0)
- Dialect: PostgreSQL
- Covers: defaults computed at load time from game data in another package (`.default(loadout.validate({} as Loadout))`), defaults and `$type<>` from TypeScript enums declared in other files (`ItemStatus.New`, `TeamMode.Solo`), `.ts` extensions in relative imports, `json`, `bigint`, `uuid`, composite primary keys and `uniqueIndex`

## Changes from upstream

- `schema.ts` and the whole `shared/` directory are copied as they are, in their original relative layout, with `zod` at upstream's version.
- `tsconfig.json` adds `allowJs` and the DOM lib, as the shared code imports `.js` files and uses browser types.
- `drizzle.config.ts` drops the database settings read from `../config.ts`.

## Generate

```bash
bazel run //apps/drizzler/examples/survev:docs.write
```

The output is written to [`docs/`](docs/):

- [`README.md`](docs/README.md): table docs with the SQL
- [`er.svg`](docs/er.svg), [`er.png`](docs/er.png), [`er-dark.png`](docs/er-dark.png) and `tables/*.svg`: ER diagrams
- [`schema.dbml`](docs/schema.dbml): DBML
- [`er.md`](docs/er.md): Mermaid `erDiagram`

## Known issues

None.
