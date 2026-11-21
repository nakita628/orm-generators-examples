# zws

The API database schema of zws, which shortens URLs using invisible zero-width spaces: blocked hostnames, URLs and visits.

- Source: [zws-im/zws `apps/api/src/db/schema.ts`](https://github.com/zws-im/zws/blob/9f2fcbd39464e2d3affbf961e35daf9a1eb5e1a5/apps/api/src/db/schema.ts) (Apache-2.0)
- Dialect: PostgreSQL
- Covers: a text primary key referenced with `onUpdate: "cascade"`, an index with no name (`index().on(...)`) in the legacy object-style extra config, `timestamp({ withTimezone: true })`, and a column named `timestamp`

## Changes from upstream

- `schema.ts` is copied as it is, under its upstream path.
- `drizzle.config.ts` drops the database URL read through `envalid`.

## Generate

```bash
bazel run //apps/drizzler/examples/zws:docs.write
```

The output is written to [`docs/`](docs/):

- [`README.md`](docs/README.md): table docs with the SQL
- [`er.svg`](docs/er.svg), [`er.png`](docs/er.png), [`er-dark.png`](docs/er-dark.png) and `tables/*.svg`: ER diagrams
- [`schema.dbml`](docs/schema.dbml): DBML
- [`er.md`](docs/er.md): Mermaid `erDiagram`

## Known issues

None.
