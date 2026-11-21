# flarestarter

The schema of FlareStarter, an edge-native SaaS starter built with TanStack Start on Cloudflare: Better Auth users, sessions, accounts, verifications and organizations, billing, a waitlist, sponsors and feedback.

- Source: [FlareStarter/flarestarter `src/db`](https://github.com/FlareStarter/flarestarter/tree/446571d2d7312e7f86bf25fced154a9d5e79abb7/src/db) and the [`src/features/*/*.schema.ts`](https://github.com/FlareStarter/flarestarter/tree/446571d2d7312e7f86bf25fced154a9d5e79abb7/src/features) files it re-exports (Apache-2.0)
- Dialect: SQLite (Cloudflare D1)
- Covers: a schema barrel that re-exports table files from feature folders through a tsconfig `paths` alias (`@/features/auth/auth.schema`), and `integer({ mode: "timestamp_ms" })`

## Changes from upstream

- The schema files are copied as they are, in their feature folders, and `tsconfig.json` keeps upstream's `@/*` path alias.

## Generate

```bash
bazel run //apps/drizzler/examples/flarestarter:docs.write
```

The output is written to [`docs/`](docs/):

- [`README.md`](docs/README.md): table docs with the SQL
- [`er.svg`](docs/er.svg), [`er.png`](docs/er.png), [`er-dark.png`](docs/er-dark.png) and `tables/*.svg`: ER diagrams
- [`schema.dbml`](docs/schema.dbml): DBML
- [`er.md`](docs/er.md): Mermaid `erDiagram`

## Known issues

None.
