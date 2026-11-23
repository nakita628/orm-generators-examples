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
     queries, SQLAlchemy's `relationship`, Django's related managers, GORM's `Preload`). A relation
     has to load a row exactly when the row on the other side holds the matching key.
   - **Delete:** children first, through the ORM.
2. **Tables (DDL).** The ORM creates the tables from the generated code into an empty database
   (`drizzle-kit export`, `metadata.create_all()`, `makemigrations` + `migrate`, `AutoMigrate`,
   `Schema::create_table_from_entity`, `EnsureCreated()`, `SchemaUtils.create`). `prisma migrate diff` then prints
   the SQL that would turn those tables into the ones the Prisma schema describes. An empty diff
   means the ORM can own the migrations. A diff means the tables should come from Prisma (`prisma
   migrate`), with the ORM used only at runtime.

hekireki-kysely writes only types, so for Kysely the runtime check holds the types to the
database instead: a column typed `Generated<...>` is left out of the insert (the database has to
fill it), every other column is given (for `cuid()`, `uuid()`, `now()` and `@updatedAt` the
harness gives a value, as an application would), and every value read back has to have the runtime
type its select type declares (`Date`, `bigint`, `boolean`, ...). Kysely has no relations and no
tables of its own.

The harnesses only report what happened. [lib/compare.mjs](lib/compare.mjs) judges every ORM the
same way, and [lib/report.mjs](lib/report.mjs) writes `results/<schema>.md` and the matrix in
[results/README.md](results/README.md).

| ORM | Harness | Driver |
| --- | --- | --- |
| drizzle | [drizzle/harness.mjs](drizzle/harness.mjs) | drizzle-orm 1.0.0-rc.5 with postgres.js, mysql2, node:sqlite |
| kysely | [kysely/harness.mjs](kysely/harness.mjs) | Kysely 0.29 with its own dialects over pg, mysql2, node:sqlite |
| sqlalchemy | [sqlalchemy/harness.py](sqlalchemy/harness.py) | SQLAlchemy 2.0 with psycopg 3, PyMySQL, sqlite3 |
| django | [django/harness.py](django/harness.py) | Django 5.2 with psycopg 3, mysqlclient, sqlite3 |
| gorm | [gorm/main.go](gorm/main.go) | GORM 1.31 with its postgres, mysql and sqlite drivers |
| sea-orm | [sea-orm/main.rs](sea-orm/main.rs) | SeaORM 1.1 with its sqlx drivers (Docker: Rust) |
| ecto | [ecto/harness.exs](ecto/harness.exs) | Ecto 3.14 with postgrex, myxql, ecto_sqlite3 (Docker: Elixir 1.18) |
| activerecord | [activerecord/harness.rb](activerecord/harness.rb) | Active Record 8.1 with pg, mysql2, sqlite3 (Docker: Ruby 3.4) |
| eloquent | [eloquent/harness.php](eloquent/harness.php) | illuminate/database 12 with PDO (Docker: PHP 8.3) |
| efcore | [efcore/Program.cs](efcore/Program.cs) | EF Core 10 with Npgsql, Pomelo 9, Microsoft.Data.Sqlite (Docker: .NET 10 SDK) |
| exposed | [exposed/Main.kt](exposed/Main.kt) | Exposed 1.5 with the PostgreSQL JDBC driver (Docker: Gradle, JDK 21) |

SeaORM, Ecto, Active Record, Eloquent, EF Core and Exposed run in an image of their own (each directory's
Dockerfile, built on the first run), because their versions need newer runtimes than a machine may
have. SeaORM, EF Core and Exposed build the generated code with the harness in the container; the
build is kept in a Docker volume (`hekireki-compat-sea-orm`, which also keeps the crates' build
across runs, `hekireki-compat-efcore`, `hekireki-compat-exposed`). SeaORM's harness is generic over
the entity; run.mjs writes `registry.rs`, which calls it for each generated entity, and a new row
starts from `ActiveModelBehavior::new()`, where hekireki puts the client-side defaults.
The plan, the generated code and the harness go into the container on stdin, and the container
joins the databases' Docker network. Ecto, Active Record and Eloquent write no tables of their
own (their migrations are written by hand), so they have the runtime check only, as Kysely does.
Active Record and Eloquent read the columns from the database, so their mapping check is the
table only.

## Usage

Docker, Node, [uv](https://docs.astral.sh/uv/) and Go are needed (and the MySQL client library for
mysqlclient). run.mjs starts PostgreSQL 17 (the pgvector image) and MySQL 8.4 in Docker (`hekireki-compat-pg` on port 55432, `hekireki-compat-mysql` on 53306) when
they are not already running. SQLite uses files.

```sh
pnpm install
cd apps/hekireki/compat

# Every ORM over some schemas: writes results/<schema>.md and results/README.md
node run.mjs opensaas taxonomy nextauth

# Every schema
node run.mjs $(ls ../schemas | grep -v -e '\.' -e BUILD)

# Some ORMs only
node run.mjs --orm drizzle,gorm calcom
```

What a run leaves behind (the plan, the schema without generators, SQLite files, the copies of the
generated code) is in `~/.cache/hekireki-compat` (or `$COMPAT_WORK`). Stop the databases with
`docker rm -f hekireki-compat-pg hekireki-compat-mysql`.

This is not a Bazel target, for the same reason the `lang` checks are not: it needs Docker and the
toolchains from PATH.

## Scope

Every ORM hekireki generates (drizzle, kysely, sqlalchemy, django, gorm, sea-orm, ecto,
activerecord, eloquent, efcore and exposed) over every schema. exposed covers the PostgreSQL ones
(hekireki-exposed writes PostgreSQL only), and eloquent leaves out the schemas hekireki leaves it
out of.

SeaORM's harness does not build for the six largest schemas (calcom, inboxzero, langfuse,
langwatch, litellm, postiz: 85 to 213 models) on a machine with 8 GB of memory: rustc runs out of it
compiling the generic harness over every entity. Those show "did not run" in the results. Its
tables check creates the PostgreSQL enum types from Prisma's DDL first, and still fails where
SeaORM writes an enum type's name unquoted (PostgreSQL folds it to lower case).

A model whose required key closes a cycle with another model's is not inserted (it would need the
other row first, and that row needs it). The report lists those models. A view is left out.
Exposed's relations are loaded through the DAO by a single id, so a model with a composite key
has its relations skipped.

## Pulled schemas: other ORMs' tables, read back into Prisma

**[Results](pull/results/README.md)**

The schemas above are Prisma schemas as their projects wrote them. [pull.mjs](pull.mjs) turns the
direction around: a well-known project's **own** ORM creates its tables from the project's own
definitions, `prisma db pull` reads them into a Prisma schema, hekireki generates every ORM's code
from that, and run.mjs runs the code against the tables the project's ORM created (not tables
from `prisma db push`). The generated Django models then have to work on the tables Django itself
made from Wagtail's models, the generated Active Record models on Mastodon's `db/schema.rb`, and
so on for every ORM, on schemas no one wrote for Prisma.

| Source | Projects | How the tables are made |
| --- | --: | --- |
| `django/<example>` | every [djangoer example](../../djangoer/examples) | Django's schema editor creates every model's table ([pull/django/create_tables.py](pull/django/create_tables.py)), on PostgreSQL |
| `drizzle/<example>` | every [drizzler example](../../drizzler/examples) | `drizzle-kit export`, in the example's dialect |
| `rails/<app>` | Mastodon, Forem, Chatwoot, Lobsters, Maybe, Loomio, Campfire ([pull/sources.json](pull/sources.json)) | Active Record loads the app's `db/schema.rb` at a pinned commit ([pull/activerecord/load_schema.rb](pull/activerecord/load_schema.rb)) |
| `sqlalchemy/<project>` | MLflow | the project's own command (`mlflow db upgrade`), or `metadata.create_all()` ([pull/sqlalchemy/create_all.py](pull/sqlalchemy/create_all.py)) |

For each source, `pull/schemas/<kind>-<name>/` keeps:

- `source.sql`: the tables as the project's ORM created them (pg_dump, mysqldump, sqlite_master).
  run.mjs makes the crud database from it.
- `schema.prisma`: what `prisma db pull` wrote, under hekireki's generator blocks. Where Prisma
  writes a schema it refuses itself (an integer foreign key on a bigint id, a default on an
  `Unsupported` field), pull.mjs makes the smallest change that validates and lists it in the
  header and in `source.json`'s `fixes`.
- `prisma-diff.sql`: what `prisma migrate diff` would run to turn the tables into the Prisma
  schema's, that is, what a Prisma schema cannot say about them (left out when there is nothing).
  A difference listed there is Prisma's, not hekireki's.
- `source.json`: where the tables came from, the versions, and when they were last checked.

`generated/` beside them is hekireki's output, written again by every run (not committed); a
generator that fails writes `generated/<orm>.error` instead, and the results show it.

```sh
node pull.mjs --list                          # every source, and when it was last checked
node pull.mjs django/wagtail rails/mastodon   # pull, generate and check these
node pull.mjs --random 3                      # three: never checked first, then the oldest
node pull.mjs --orm django,activerecord ...   # only these ORMs
node pull.mjs --pull-only drizzle/northwind   # only the pull: schema.prisma and source.sql
```

It needs what run.mjs needs, plus uv for the Django and SQLAlchemy sources, the pnpm install of
apps/drizzler for drizzle-kit, and the Active Record image for the Rails ones. A source that cannot
be pulled at all goes in `skip` of [pull/sources.json](pull/sources.json), with the reason.

A Claude Code scheduled task runs `node pull.mjs --random 3` every day, files what is hekireki's
on [hekireki](https://github.com/nakita628/hekireki/issues), and pushes the results; see
[.claude/skills/hekireki-pull-check](../../../.claude/skills/hekireki-pull-check/SKILL.md).

## In a Claude Code cloud session

On your own machine, Docker is used as it is. In a cloud session, run
[cloud/prepare.sh](cloud/prepare.sh) first and export what it prints:

```sh
eval "$(bash apps/hekireki/compat/cloud/prepare.sh | tail -1)"
```

It starts dockerd with mirror.gcr.io in front of Docker Hub, re-tags each Dockerfile's base image
with the session proxy's CA in it ([cloud/base.Dockerfile](cloud/base.Dockerfile)), builds every
image with the host's network, and prints `COMPAT_DOCKER_BUILD_ARGS`, which run.mjs and pull.mjs
add to their own `docker build` so that it finds every layer in the cache. The Dockerfiles
themselves stay as they are. The environment has to allow `deb.debian.org`, `repo.hex.pm` and
`builds.hex.pm`; Maven Central comes from Google's mirror of it
([cloud/maven-mirror.gradle](cloud/maven-mirror.gradle)), and Composer installs from git.
