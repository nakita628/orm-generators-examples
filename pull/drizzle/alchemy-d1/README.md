# drizzle/alchemy-d1

The tables drizzle-kit creates from the Drizzle schema of [apps/drizzler/examples/alchemy-d1](../../../apps/drizzler/examples/alchemy-d1), as SQL: [source.sql](source.sql). `prisma db pull` writes them up as [schema.prisma](schema.prisma) (with 3 changes so that Prisma accepts it, listed in its header), and from that write-up Prisma and every ORM hekireki generates create tables of their own; their SQL is in [sql/](sql), dumped the way source.sql is, and held to it here.

`=` identical · `≈` the same but for constraint and index names, and the order of columns and statements · `≠` tables written differently, or statements one side has and the other does not. The last column holds each ORM's SQL to what Prisma itself creates from the same write-up: a difference there is the ORM's or hekireki's, one only against source.sql is something the write-up does not say.

| From the write-up | Creates the tables with | SQL | Against source.sql | Against prisma.sql |
| --- | --- | --- | --- | --- |
| Prisma: schema.prisma | `prisma db push` | [prisma.sql](sql/prisma.sql) | ≈ the same but for names |  |
| hekireki-drizzle | `drizzle-kit export` | [drizzle.sql](sql/drizzle.sql) | ≈ the same but for names | ≈ |
| hekireki-kysely | — | — | writes types only | — |
| hekireki-sqlalchemy | `metadata.create_all()` | [sqlalchemy.sql](sql/sqlalchemy.sql) | ≈ the same but for names | ≈ |
| hekireki-django | `makemigrations`, `migrate` | [django.sql](sql/django.sql) | ≠ 2 tables written differently ([diff](sql/django.diff)) | ≠ 2 |
| hekireki-gorm | `AutoMigrate` | [gorm.sql](sql/gorm.sql) | ≈ the same but for names | ≈ |
| hekireki-sea-orm | `Schema::create_table_from_entity` | [sea-orm.sql](sql/sea-orm.sql) | ≠ 2 tables written differently ([diff](sql/sea-orm.diff)) | ≠ 2 |
| hekireki-ecto | — | — | writes schemas only (no migrations) | — |
| hekireki-activerecord | — | — | writes models only (no migrations) | — |
| hekireki-eloquent | — | — | writes models only (no migrations) | — |
| hekireki-efcore | `EnsureCreated()` | [efcore.sql](sql/efcore.sql) | ≠ 2 tables written differently ([diff](sql/efcore.diff)) | ≠ 2 |
| hekireki-exposed | `SchemaUtils.create` | — | not generated | — |

2 models, sqlite. hekireki 0.52.20, Prisma 8.1.0-dev.9. The same code run against source.sql's tables: [results/alchemy-d1.md](../results/alchemy-d1.md).

<details><summary>hekireki-django: 2 tables written differently</summary>

```diff
@@ "posts"
- "created_at" INTEGER NOT NULL DEFAULT unixepoch()
+ "created_at" INTEGER NOT NULL DEFAULT ((unixepoch()))
- FOREIGN KEY ("user_id") REFERENCES "users" ("id") ON UPDATE NO ACTION ON DELETE CASCADE
+ FOREIGN KEY ("user_id") REFERENCES "users" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION
@@ "users"
- "created_at" INTEGER NOT NULL DEFAULT unixepoch()
+ "created_at" INTEGER NOT NULL DEFAULT ((unixepoch()))
```

</details>

<details><summary>hekireki-sea-orm: 2 tables written differently</summary>

```diff
@@ "posts"
- "title" TEXT NOT NULL
+ "title" VARCHAR NOT NULL
- "body" TEXT NOT NULL
+ "body" VARCHAR NOT NULL
- "created_at" INTEGER NOT NULL DEFAULT unixepoch()
+ "created_at" INTEGER NOT NULL
@@ "users"
- "email" TEXT NOT NULL
+ "email" VARCHAR NOT NULL
- "name" TEXT NOT NULL
+ "name" VARCHAR NOT NULL
- "created_at" INTEGER NOT NULL DEFAULT unixepoch()
+ "created_at" INTEGER NOT NULL
```

</details>

<details><summary>hekireki-efcore: 2 tables written differently</summary>

```diff
@@ "posts"
- "created_at" INTEGER NOT NULL DEFAULT unixepoch()
+ "created_at" INTEGER NOT NULL DEFAULT (unixepoch())
@@ "users"
- "created_at" INTEGER NOT NULL DEFAULT unixepoch()
+ "created_at" INTEGER NOT NULL DEFAULT (unixepoch())
```

</details>

Written by `node pull/pull.mjs`; do not edit.
