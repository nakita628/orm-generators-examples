# infinitunes

The database package of Infinitunes, an open-source music player web app: Better Auth users, accounts, sessions, verifications, rate limits and passkeys, and the app's playlists, favorites and recently played items.

- Source: [rajput-hemant/infinitunes `packages/db/src/schema.ts`](https://github.com/rajput-hemant/infinitunes/blob/89cb78436866307cdc52d81c09f46ae225e60ce5/packages/db/src/schema.ts) and [`table-creator.ts`](https://github.com/rajput-hemant/infinitunes/blob/89cb78436866307cdc52d81c09f46ae225e60ce5/packages/db/src/table-creator.ts) (MIT)
- Dialect: PostgreSQL
- Covers: `pgTableCreator` in its own file with the prefix in a constant (`${projectName}_${name}`), next to plain `pgTable`s, `tablesFilter` in `drizzle.config.ts`, JSDoc table notes, and composite keys

## Changes from upstream

- Both files are copied as they are.
- `drizzle.config.ts` drops the env loading and the database URL.

## Generate

```bash
bazel run //apps/drizzler/examples/infinitunes:docs.write
```

The output is written to [`docs/`](docs/):

- [`README.md`](docs/README.md): table docs with the SQL
- [`er.svg`](docs/er.svg), [`er.png`](docs/er.png), [`er-dark.png`](docs/er-dark.png) and `tables/*.svg`: ER diagrams
- [`schema.dbml`](docs/schema.dbml): DBML
- [`er.md`](docs/er.md): Mermaid `erDiagram`

## Known issues

None.
