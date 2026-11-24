# dyad

The local SQLite schema of Dyad, a local, open-source AI app builder: apps, chats, messages, prompts, versions, MCP servers, custom models and providers, and more.

- Source: [dyad-sh/dyad `src/db/schema.ts`](https://github.com/dyad-sh/dyad/blob/bc4f2af5c8739d3179e135b2dc874f87cf002213/src/db/schema.ts), [`src/shared/inference_source.ts`](https://github.com/dyad-sh/dyad/blob/bc4f2af5c8739d3179e135b2dc874f87cf002213/src/shared/inference_source.ts) and [`src/shared/chat_turn_outcome.ts`](https://github.com/dyad-sh/dyad/blob/bc4f2af5c8739d3179e135b2dc874f87cf002213/src/shared/chat_turn_outcome.ts) (Apache-2.0)
- Dialect: SQLite
- Covers: imports through a tsconfig `paths` alias (`@/shared/...`), `text({ enum })` built from imported `as const` arrays, `integer` with `mode: "timestamp"` and `` .default(sql`(unixepoch())`) ``, `primaryKey({ autoIncrement: true })`, `unique` and `uniqueIndex`, long comments between columns, and v1 `relations`

## Changes from upstream

- `relations` is imported from `drizzle-orm/_relations` instead of `drizzle-orm`, because drizzle-orm 1.0 moved the v1 relational API there.
- `src/shared/inference_source.ts` and `src/shared/chat_turn_outcome.ts` are copied as they are. The types the schema imports from `@/lib/schemas`, `@/chat_stream/transport` and `ai` come from stand-ins (`src/lib/schemas.ts`, `src/chat_stream/transport.ts`, `src/types/ai.d.ts`).
- `tsconfig.json` keeps only the `@/*` path alias.

## Generate

```bash
bazel run //apps/drizzler/examples/dyad:docs.write
```

The output is written to [`docs/`](docs/):

- [`README.md`](docs/README.md): table docs with the SQL
- [`er.svg`](docs/er.svg), [`er.png`](docs/er.png), [`er-dark.png`](docs/er-dark.png) and `tables/*.svg`: ER diagrams
- [`schema.dbml`](docs/schema.dbml): DBML
- [`er.md`](docs/er.md): Mermaid `erDiagram`

## Known issues

None.
