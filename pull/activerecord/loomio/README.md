# activerecord/loomio

The tables Active Record creates from [Loomio](https://github.com/loomio/loomio)'s db/schema.rb, as SQL: [source.sql](source.sql). `prisma db pull` writes them up as [schema.prisma](schema.prisma) (with 23 changes so that Prisma accepts it, listed in its header), and from that write-up Prisma and every ORM hekireki generates create tables of their own; their SQL is in [sql/](sql), dumped the way source.sql is, and held to it here.

`=` identical · `≈` the same but for constraint and index names, and the order of columns and statements · `≠` tables written differently, or statements one side has and the other does not. The last column holds each ORM's SQL to what Prisma itself creates from the same write-up: a difference there is the ORM's or hekireki's, one only against source.sql is something the write-up does not say.

| From the write-up | Creates the tables with | SQL | Against source.sql | Against prisma.sql |
| --- | --- | --- | --- | --- |
| Prisma: schema.prisma | `prisma db push` | [prisma.sql](sql/prisma.sql) | ❌ could not create its tables: Error: ERROR: cannot use column reference in DEFAULT expression |  |
| hekireki-drizzle | `drizzle-kit export` | — | ❌ could not create its tables: Error: ERROR: syntax error at or near "," | — |
| hekireki-kysely | — | — | writes types only | — |
| hekireki-sqlalchemy | `metadata.create_all()` | — | ❌ could not create its tables: NameError: name 'parent_id' is not defined | — |
| hekireki-django | `makemigrations`, `migrate` | — | ❌ could not create its tables: psycopg.errors.FeatureNotSupported: cannot use column reference in DEFAULT expression | — |
| hekireki-gorm | `AutoMigrate` | — | ❌ could not create its tables: ERROR: cannot use column reference in DEFAULT expression (SQLSTATE 0A000) | — |
| hekireki-sea-orm | `Schema::create_table_from_entity` | — | ❌ could not create its tables: "groups": Execution Error: error returned from database: cannot use column reference in DEFAULT expression | — |
| hekireki-ecto | — | — | writes schemas only (no migrations) | — |
| hekireki-activerecord | — | — | writes models only (no migrations) | — |
| hekireki-eloquent | — | — | writes models only (no migrations) | — |
| hekireki-efcore | `EnsureCreated()` | — | ❌ could not create its tables: Property 'TrialInvitationRecipients.IpAddress' cannot be used as a key because it has type 'NpgsqlInet' which does not implement 'IComparable<T>', 'IComparable' or 'IStructuralComparable'. Use 'HasConversion' in 'OnModelCreating' to wrap 'NpgsqlInet' with a type that can be compared. | — |
| hekireki-exposed | `SchemaUtils.create` | — | ❌ could not create its tables: Exception in thread "main" org.jetbrains.exposed.v1.exceptions.ExposedSQLException: org.postgresql.util.PSQLException: ERROR: cannot use column reference in DEFAULT expression | — |

81 models, postgresql. hekireki 0.55.0, Prisma 8.1.0-dev.9. The same code run against source.sql's tables: [results/loomio.md](../results/loomio.md).

Written by `node pull/pull.mjs`; do not edit.
