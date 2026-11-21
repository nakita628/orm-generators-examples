# openbio

The schema of OpenBio, an open-source link-in-bio page builder: Auth.js users, accounts, sessions and verification tokens, bento links, link views and clicks, and email subscribers.

- Source: [vanxh/openbio `src/server/db/schema`](https://github.com/vanxh/openbio/tree/a0ae99a58bc75a985be566d01b05f003840c6bba/src/server/db/schema) and [`src/types.ts`](https://github.com/vanxh/openbio/blob/a0ae99a58bc75a985be566d01b05f003840c6bba/src/types.ts) (AGPL-3.0)
- Dialect: PostgreSQL
- Covers: a directory of one file per table behind an `index.ts` of `export *`, zod schemas re-exported from a schema file (`export { BentoSchema } from '@/types'`) through a tsconfig `paths` alias, `json().$type<z.infer<...>[]>()`, and v1 `relations`

## Changes from upstream

- `relations` is imported from `drizzle-orm/_relations` instead of `drizzle-orm`, because drizzle-orm 1.0 moved the v1 relational API there.
- `src/types.ts` is copied as it is, with `zod` 3. `tsconfig.json` keeps upstream's `@/*` path alias.
- `drizzle.config.ts` drops the database URL.

## Generate

```bash
bazel run //apps/drizzler/examples/openbio:docs.write
```

The output is written to [`docs/`](docs/):

- [`README.md`](docs/README.md): table docs with the SQL
- [`er.svg`](docs/er.svg), [`er.png`](docs/er.png), [`er-dark.png`](docs/er-dark.png) and `tables/*.svg`: ER diagrams
- [`schema.dbml`](docs/schema.dbml): DBML
- [`er.md`](docs/er.md): Mermaid `erDiagram`

## Known issues

None.
