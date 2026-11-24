# dokploy

The schema of Dokploy, a self-hosted PaaS (an open-source alternative to Vercel, Netlify and Heroku): organizations, projects and environments, applications, compose stacks, the database services (Postgres, MySQL, MariaDB, Mongo, Redis, libSQL), domains, deployments, backups, servers, git providers, notifications and Better Auth.

- Source: [Dokploy/dokploy `packages/server/src/db/schema`](https://github.com/Dokploy/dokploy/tree/a23108fb52b7c68d15c8f9be96f272715820b124/packages/server/src/db/schema) and [`validations`](https://github.com/Dokploy/dokploy/tree/a23108fb52b7c68d15c8f9be96f272715820b124/packages/server/src/db/validations), plus [`utils/git-branch-validation.ts`](https://github.com/Dokploy/dokploy/blob/a23108fb52b7c68d15c8f9be96f272715820b124/packages/server/src/utils/git-branch-validation.ts), [`utils/hostname-validation.ts`](https://github.com/Dokploy/dokploy/blob/a23108fb52b7c68d15c8f9be96f272715820b124/packages/server/src/utils/hostname-validation.ts) and [`utils/password-validation.ts`](https://github.com/Dokploy/dokploy/blob/a23108fb52b7c68d15c8f9be96f272715820b124/packages/server/src/utils/password-validation.ts) (Apache-2.0)
- Dialect: PostgreSQL
- Covers: 69 tables across 53 files behind an `index.ts` of `export *`, imports through a tsconfig `paths` alias (`@dokploy/server/*`), about 40 `pgEnum`s with PascalCase SQL names, a `customType` that encrypts at rest (`encryptedText`), `$defaultFn(() => nanoid())` keys, zod insert schemas built with `createInsertSchema` next to every table, and v1 `relations`

## Changes from upstream

- `relations` is imported from `drizzle-orm/_relations` instead of `drizzle-orm`, because drizzle-orm 1.0 moved the v1 relational API there.
- `createInsertSchema` is imported from `drizzle-orm/zod` instead of `drizzle-zod`, which drizzle-orm 1.0 replaces.
- `dbml.ts` is left out. It is not exported from `index.ts` and only runs `drizzle-dbml-generator`.
- The modules the schema imports from the rest of `@dokploy/server` are stand-ins: `lib/encryption.ts` (identity `encryptValue` / `decryptValue`), `templates/index.ts` (`generatePassword`), `constants/index.ts` (`paths`) and the types in `services/*.ts`. The three `utils` files are copied as they are.

## Generate

```bash
bazel run //apps/drizzler/examples/dokploy:docs.write
```

The output is written to [`docs/`](docs/):

- [`README.md`](docs/README.md): table docs with the SQL
- [`er.svg`](docs/er.svg), [`er.png`](docs/er.png), [`er-dark.png`](docs/er-dark.png) and `tables/*.svg`: ER diagrams
- [`schema.dbml`](docs/schema.dbml): DBML
- [`er.md`](docs/er.md): Mermaid `erDiagram`

## Known issues

- [nakita628/drizzler#62](https://github.com/nakita628/drizzler/issues/62): the `appName` defaults, built with Faker, change the insert examples in `docs/README.md` on every run.
