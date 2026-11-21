# vercel-ai-chatbot

The schema of Vercel's AI Chatbot template: users, chats, messages, votes, documents, suggestions and resumable streams.

- Source: [vercel/ai-chatbot `lib/db/schema.ts`](https://github.com/vercel/ai-chatbot/blob/a8989b522fa3672a2ae52f053105e40993b36ca1/lib/db/schema.ts) (Apache-2.0)
- Dialect: PostgreSQL
- Covers: PascalCase and versioned table names (`Message_v2`), a column whose SQL name differs from its key (`kind: varchar("text", ...)`), `varchar` with `enum`, `uuid().defaultRandom()`, composite primary keys in the legacy object-style extra config, a composite `foreignKey()` (`Suggestion` → `Document`), and a single-column `foreignKey()`

## Changes from upstream

None. `drizzle.config.ts` drops the `dotenv` loading and the database URL, which drizzler does not need.

## Generate

```bash
bazel run //apps/drizzler/examples/vercel-ai-chatbot:docs.write
```

The output is written to [`docs/`](docs/):

- [`README.md`](docs/README.md): table docs with the SQL
- [`er.svg`](docs/er.svg), [`er.png`](docs/er.png), [`er-dark.png`](docs/er-dark.png) and `tables/*.svg`: ER diagrams
- [`schema.dbml`](docs/schema.dbml): DBML
- [`er.md`](docs/er.md): Mermaid `erDiagram`

## Known issues

None.
