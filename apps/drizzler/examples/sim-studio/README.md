# sim-studio

The schema of Sim Studio, an open-source platform for building AI agent workflows: users and Better Auth, workspaces and organizations, workflows and their blocks, executions and logs, knowledge bases with embeddings, credentials, billing, copilot chats and more.

- Source: [simstudioai/sim `packages/db/schema.ts`](https://github.com/simstudioai/sim/blob/ad89acd9726cfeaae2b4ebb0bd74e1e4871ab664/packages/db/schema.ts), [`constants.ts`](https://github.com/simstudioai/sim/blob/ad89acd9726cfeaae2b4ebb0bd74e1e4871ab664/packages/db/constants.ts), [`sealed-text.ts`](https://github.com/simstudioai/sim/blob/ad89acd9726cfeaae2b4ebb0bd74e1e4871ab664/packages/db/sealed-text.ts) and [`packages/utils/src/string.ts`](https://github.com/simstudioai/sim/blob/ad89acd9726cfeaae2b4ebb0bd74e1e4871ab664/packages/utils/src/string.ts) (Apache-2.0)
- Dialect: PostgreSQL
- Covers: 192 tables in one 394 KB file, a `customType` built inside a helper and called on the spot (`sealedText(table, column)`), `vector` columns with HNSW indexes, `tsvector` generated columns, `sql` expressions in indexes and checks, partial indexes, many `pgEnum`s and self-references

## Changes from upstream

- `schema.ts`, `constants.ts` and `sealed-text.ts` are copied as they are. The `@sim/*` imports resolve through tsconfig `paths`: `@sim/utils/string` to a copy of `packages/utils/src/string.ts`, and `@sim/logger`, `@sim/security/encryption` and `@sim/utils/object` to stand-ins in `stubs/sim/` that hold only what `sealed-text.ts` imports (no real sealing).
- `drizzle.config.ts` points at `./src/schema.ts` and drops the `tablesFilter`, the database URL and the migrations folder.

## Generate

```bash
bazel run //apps/drizzler/examples/sim-studio:docs.write
```

The output is written to [`docs/`](docs/):

- [`README.md`](docs/README.md): table docs with the SQL
- [`er.svg`](docs/er.svg), [`er.png`](docs/er.png), [`er-dark.png`](docs/er-dark.png) and `tables/*.svg`: ER diagrams
- [`schema.dbml`](docs/schema.dbml): DBML
- [`er.md`](docs/er.md): Mermaid `erDiagram`

## Known issues

None.
