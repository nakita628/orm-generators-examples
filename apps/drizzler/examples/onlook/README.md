# onlook

The schema of Onlook, an open-source visual editor for React apps ("Cursor for designers"): users on top of Supabase's `auth.users`, projects, branches, canvases and frames, chats and messages, invitations, custom and preview domains, deployments, and Stripe products, prices, subscriptions and usage.

- Source: [onlook-dev/onlook `packages/db/src/schema`](https://github.com/onlook-dev/onlook/tree/993068a859decf9e93152e8a9f43d379d0d708c0/packages/db/src/schema) (Apache-2.0), with the enums it imports from [`packages/models`](https://github.com/onlook-dev/onlook/tree/993068a859decf9e93152e8a9f43d379d0d708c0/packages/models/src) and [`packages/stripe`](https://github.com/onlook-dev/onlook/tree/993068a859decf9e93152e8a9f43d379d0d708c0/packages/stripe/src)
- Dialect: PostgreSQL (Supabase)
- Covers: a directory `schema` of nested `index.ts` files, a table in another Postgres schema (`pgSchema("auth").table("users", ...)`) referenced from `public.users`, `pgEnum` built from TypeScript enums (`pgEnum("project_role", ProjectRole)`), `entities.roles` and `schemaFilter` in `drizzle.config.ts`, zod insert and update schemas built next to the tables, and v1 `relations`

## Changes from upstream

- `relations` is imported from `drizzle-orm/_relations` instead of `drizzle-orm`, because drizzle-orm 1.0 moved the v1 relational API there.
- `createInsertSchema` and `createUpdateSchema` are imported from `drizzle-orm/zod` instead of `drizzle-zod`, which drizzle-orm 1.0 replaces.
- `@onlook/models` and `@onlook/stripe` resolve through tsconfig `paths` to stand-ins in `stubs/`: the TypeScript enums the schema imports are copied as they are, and the types it imports are loose. The `LanguageModelUsage` type from `ai` comes from `stubs/types/ai.d.ts`.
- `drizzle.config.ts` drops the database URL, the migrations folder and `verbose`.

## Generate

```bash
bazel run //apps/drizzler/examples/onlook:docs.write
```

The output is written to [`docs/`](docs/):

- [`README.md`](docs/README.md): table docs with the SQL
- [`er.svg`](docs/er.svg), [`er.png`](docs/er.png), [`er-dark.png`](docs/er-dark.png) and `tables/*.svg`: ER diagrams
- [`schema.dbml`](docs/schema.dbml): DBML
- [`er.md`](docs/er.md): Mermaid `erDiagram`
