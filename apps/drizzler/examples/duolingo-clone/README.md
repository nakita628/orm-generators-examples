# duolingo-clone

The schema of Lingo, a popular open-source Duolingo clone built with Next.js: courses, units, lessons, challenges and their options, user progress and subscriptions.

- Source: [sanidhyy/duolingo-clone `db/schema.ts`](https://github.com/sanidhyy/duolingo-clone/blob/e2093c71be67288d661cfccee5c57f9855108ab4/db/schema.ts) and [`constants.ts`](https://github.com/sanidhyy/duolingo-clone/blob/e2093c71be67288d661cfccee5c57f9855108ab4/constants.ts) (MIT)
- Dialect: PostgreSQL
- Covers: `serial` keys, a `pgEnum` whose SQL name is `type`, a default from an imported constant (`.default(MAX_HEARTS)`) through a tsconfig `paths` alias (`@/constants`), `onDelete: "cascade"`, and v1 `relations`

## Changes from upstream

- `relations` is imported from `drizzle-orm/_relations` instead of `drizzle-orm`, because drizzle-orm 1.0 moved the v1 relational API there.
- `constants.ts` is copied as it is, and `tsconfig.json` keeps upstream's `@/*` path alias.
- `drizzle.config.ts` drops the env loading and the database URL.

## Generate

```bash
bazel run //apps/drizzler/examples/duolingo-clone:docs.write
```

The output is written to [`docs/`](docs/):

- [`README.md`](docs/README.md): table docs with the SQL
- [`er.svg`](docs/er.svg), [`er.png`](docs/er.png), [`er-dark.png`](docs/er-dark.png) and `tables/*.svg`: ER diagrams
- [`schema.dbml`](docs/schema.dbml): DBML
- [`er.md`](docs/er.md): Mermaid `erDiagram`

## Known issues

None.
