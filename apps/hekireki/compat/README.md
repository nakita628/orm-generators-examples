# hekireki compat

Whether the ORM code [hekireki](https://github.com/nakita628/hekireki) generates from a Prisma schema
works against the database that Prisma itself creates from the same schema. [checks/](../checks)
compiles the generated code. This directory runs it against a real database.

**[Results](results/README.md)**

## What runs

For every schema and ORM, [run.mjs](run.mjs) runs two checks:

1. **Runtime.** `prisma db push` creates the database, the same way the project's own Prisma would.
   The ORM's harness then runs a plan made from the schema's DMMF ([lib/plan.mjs](lib/plan.mjs))
   through the generated code, one row per model, parents first:
   - **Mapping:** the ORM has a model for every table and an attribute for every column. Matching
     goes through the ORM's own metadata (column names, not attribute names).
   - **Insert:** a value of the column's type goes into every column that has no default. Every
     column with a default (`cuid()`, `uuid()`, `now()`, `autoincrement()`, `@updatedAt`, a
     literal) is left to the ORM. Prisma Client fills `cuid()`, `uuid()`, `now()` and
     `@updatedAt` itself, and they are not in the database, so the ORM has to fill them.
   - **Read back:** the row read back through the ORM holds what was written. A default the ORM
     filled is checked as well: not null, a date near now, a UUID or cuid of the right form.
   - **Update:** one column is changed through the ORM and read back.
   - **Relations:** every relation is loaded through the ORM from both sides (drizzle's relational
     queries, SQLAlchemy's `relationship`, GORM's `Preload`).
   - **Delete:** children first, through the ORM.
2. **Tables (DDL).** The ORM creates the tables from the generated code into an empty database
   (`drizzle-kit export`, `metadata.create_all()`, `AutoMigrate`). `prisma migrate diff` then prints
   the SQL that would turn those tables into the ones the Prisma schema describes. An empty diff
   means the ORM can own the migrations. A diff means the tables should come from Prisma (`prisma
   migrate`), with the ORM used only at runtime.

The harnesses only report what happened. [lib/compare.mjs](lib/compare.mjs) judges every ORM the
same way, and [lib/report.mjs](lib/report.mjs) writes `results/<schema>.md` and the matrix in
[results/README.md](results/README.md).

| ORM | Harness | Driver |
| --- | --- | --- |
| drizzle | [drizzle/harness.mjs](drizzle/harness.mjs) | drizzle-orm 1.0.0-rc.5 with postgres.js, mysql2, node:sqlite |
| sqlalchemy | [sqlalchemy/harness.py](sqlalchemy/harness.py) | SQLAlchemy 2.0 with psycopg 3, PyMySQL, sqlite3 |
| gorm | [gorm/main.go](gorm/main.go) | GORM 1.31 with its postgres, mysql and sqlite drivers |

## Usage

Docker, Node, [uv](https://docs.astral.sh/uv/) and Go are needed. run.mjs starts PostgreSQL 17 and
MySQL 8.4 in Docker (`hekireki-compat-pg` on port 55432, `hekireki-compat-mysql` on 53306) when
they are not already running. SQLite uses files.

```sh
pnpm install
cd apps/hekireki/compat

# Every ORM over some schemas: writes results/<schema>.md and results/README.md
node run.mjs opensaas taxonomy nextauth

# Some ORMs only
node run.mjs --orm drizzle,gorm calcom
```

What a run leaves behind (the plan, the schema without generators, SQLite files, the copies of the
generated code) is in `.work/`, which git ignores. Stop the databases with
`docker rm -f hekireki-compat-pg hekireki-compat-mysql`.

This is not a Bazel target, for the same reason the `lang` checks are not: it needs Docker and the
toolchains from PATH.

## Scope

This is a first step: three ORMs (drizzle, sqlalchemy, gorm) over three schemas, one per database
(opensaas on PostgreSQL, taxonomy on MySQL, nextauth on SQLite). More ORMs (django, sea-orm, efcore
and exposed can create tables; kysely, ecto, activerecord and eloquent run at runtime only) and more
schemas are the next step.
