# nuxt-ui-chat

The database schema of the Nuxt UI Chat template, Nuxt's official AI chatbot template with the Vercel AI SDK on NuxtHub: users, chats, messages and votes.

- Source: [nuxt-ui-templates/chat `server/db/schema.ts`](https://github.com/nuxt-ui-templates/chat/blob/4530cb6b16a5e920b9e7c043a612e735f0c46e65/server/db/schema.ts) (MIT)
- Dialect: SQLite (NuxtHub / Cloudflare D1)
- Covers: a shared `timestamps` object spread into tables, `text({ enum })`, `text({ mode: "json" })`, composite primary keys, `uniqueIndex`, and v1 `relations`; NuxtHub reads the schema itself, so upstream has no `drizzle.config.ts`

## Changes from upstream

- `relations` is imported from `drizzle-orm/_relations` instead of `drizzle-orm`, because drizzle-orm 1.0 moved the v1 relational API there.
- `drizzle.config.ts` is added, pointing at `server/db/schema.ts`.

## Generate

```bash
bazel run //apps/drizzler/examples/nuxt-ui-chat:docs.write
```

The output is written to [`docs/`](docs/):

- [`README.md`](docs/README.md): table docs with the SQL
- [`er.svg`](docs/er.svg), [`er.png`](docs/er.png), [`er-dark.png`](docs/er-dark.png) and `tables/*.svg`: ER diagrams
- [`schema.dbml`](docs/schema.dbml): DBML
- [`er.md`](docs/er.md): Mermaid `erDiagram`

## Known issues

None.
