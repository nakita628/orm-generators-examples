# activerecord/maybe

The tables Active Record creates from [Maybe](https://github.com/maybe-finance/maybe)'s db/schema.rb, as SQL: [source.sql](source.sql). `prisma db pull` writes them up as [schema.prisma](schema.prisma), and from that write-up Prisma and every ORM hekireki generates create tables of their own; their SQL is in [sql/](sql), dumped the way source.sql is, and held to it here.

`=` identical · `≈` the same but for constraint and index names, and the order of columns and statements · `≠` tables written differently, or statements one side has and the other does not. The last column holds each ORM's SQL to what Prisma itself creates from the same write-up: a difference there is the ORM's or hekireki's, one only against source.sql is something the write-up does not say.

| From the write-up | Creates the tables with | SQL | Against source.sql | Against prisma.sql |
| --- | --- | --- | --- | --- |
| Prisma: schema.prisma | `prisma db push` | [prisma.sql](sql/prisma.sql) | ❌ could not create its tables: Error: ERROR: cannot use column reference in DEFAULT expression |  |
| hekireki-drizzle | `drizzle-kit export` | — | ❌ could not create its tables: Error: ERROR: syntax error at or near "," | — |
| hekireki-kysely | — | — | writes types only | — |
| hekireki-sqlalchemy | `metadata.create_all()` | — | ❌ could not create its tables: psycopg.errors.FeatureNotSupported: cannot use column reference in DEFAULT expression | — |
| hekireki-django | `makemigrations`, `migrate` | — | ❌ could not create its tables: psycopg.errors.FeatureNotSupported: cannot use column reference in DEFAULT expression | — |
| hekireki-gorm | `AutoMigrate` | — | ❌ could not create its tables: ERROR: relation "users" does not exist (SQLSTATE 42P01) | — |
| hekireki-sea-orm | `Schema::create_table_from_entity` | — | ❌ could not create its tables: "chats": Execution Error: error returned from database: relation "users" does not exist | — |
| hekireki-ecto | — | — | writes schemas only (no migrations) | — |
| hekireki-activerecord | — | — | writes models only (no migrations) | — |
| hekireki-eloquent | — | — | writes models only (no migrations) | — |
| hekireki-efcore | `EnsureCreated()` | — | ❌ could not create its tables: 0A000: cannot use column reference in DEFAULT expression | — |
| hekireki-exposed | `SchemaUtils.create` | — | ❌ could not create its tables: Exception in thread "main" org.jetbrains.exposed.v1.exceptions.ExposedSQLException: org.postgresql.util.PSQLException: ERROR: cannot use column reference in DEFAULT expression | — |

61 models, postgresql. hekireki 0.52.2, Prisma 8.1.0-dev.9. The same code run against source.sql's tables: [results/maybe.md](../results/maybe.md).

Written by `node pull/pull.mjs`; do not edit.
