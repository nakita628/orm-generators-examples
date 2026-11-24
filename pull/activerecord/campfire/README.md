# activerecord/campfire

The tables Active Record creates from [Campfire](https://github.com/basecamp/once-campfire)'s db/schema.rb, as SQL: [source.sql](source.sql). `prisma db pull` writes them up as [schema.prisma](schema.prisma) (with 2 changes so that Prisma accepts it, listed in its header), and from that write-up Prisma and every ORM hekireki generates create tables of their own; their SQL is in [sql/](sql), dumped the way source.sql is, and held to it here.

`=` identical · `≈` the same but for constraint and index names, and the order of columns and statements · `≠` tables written differently, or statements one side has and the other does not. The last column holds each ORM's SQL to what Prisma itself creates from the same write-up: a difference there is the ORM's or hekireki's, one only against source.sql is something the write-up does not say.

| From the write-up | Creates the tables with | SQL | Against source.sql | Against prisma.sql |
| --- | --- | --- | --- | --- |
| Prisma: schema.prisma | `prisma db push` | [prisma.sql](sql/prisma.sql) | ≠ 17 tables written differently, 6 tables only here, 1 other statement only in source.sql ([diff](sql/prisma.diff)) |  |
| hekireki-drizzle | `drizzle-kit export` | [drizzle.sql](sql/drizzle.sql) | ≠ 17 tables written differently, 3 tables only here, 1 other statement only in source.sql ([diff](sql/drizzle.diff)) | ≠ 20 |
| hekireki-kysely | — | — | writes types only | — |
| hekireki-sqlalchemy | `metadata.create_all()` | [sqlalchemy.sql](sql/sqlalchemy.sql) | ≠ 17 tables written differently, 3 tables only here, 1 other statement only in source.sql ([diff](sql/sqlalchemy.diff)) | ≠ 19 |
| hekireki-django | `makemigrations`, `migrate` | [django.sql](sql/django.sql) | ≠ 17 tables written differently, 3 tables only here, 1 other statement only in source.sql ([diff](sql/django.diff)) | ≠ 19 |
| hekireki-gorm | `AutoMigrate` | [gorm.sql](sql/gorm.sql) | ≠ 17 tables written differently, 3 tables only here, 1 other statement only in source.sql ([diff](sql/gorm.diff)) | ≠ 20 |
| hekireki-sea-orm | `Schema::create_table_from_entity` | [sea-orm.sql](sql/sea-orm.sql) | ≠ 16 tables written differently, 3 tables only here, 1 other statement only in source.sql ([diff](sql/sea-orm.diff)) | ≠ 23 |
| hekireki-ecto | — | — | writes schemas only (no migrations) | — |
| hekireki-activerecord | — | — | writes models only (no migrations) | — |
| hekireki-eloquent | — | — | writes models only (no migrations) | — |
| hekireki-efcore | `EnsureCreated()` | [efcore.sql](sql/efcore.sql) | ≠ 17 tables written differently, 3 tables only here, 1 other statement only in source.sql ([diff](sql/efcore.diff)) | ≠ 19 |
| hekireki-exposed | `SchemaUtils.create` | — | not generated | — |

23 models, sqlite. hekireki 0.52.2, Prisma 8.1.0-dev.9. The same code run against source.sql's tables: [results/campfire.md](../results/campfire.md).

<details><summary>Prisma: schema.prisma: 17 tables written differently, 6 tables only here, 1 other statement only in source.sql</summary>

```diff
@@ "accounts"
- "join_code" VARCHAR NOT NULL
+ "join_code" TEXT NOT NULL
- "name" VARCHAR NOT NULL
+ "name" TEXT NOT NULL
@@ "action_text_rich_texts"
- "name" VARCHAR NOT NULL
+ "name" TEXT NOT NULL
- "record_type" VARCHAR NOT NULL
+ "record_type" TEXT NOT NULL
@@ "active_storage_attachments"
- "blob_id" BIGINT NOT NULL
+ "blob_id" INTEGER NOT NULL
- "name" VARCHAR NOT NULL
+ "name" TEXT NOT NULL
- "record_type" VARCHAR NOT NULL
+ "record_type" TEXT NOT NULL
@@ "active_storage_blobs"
- "checksum" VARCHAR
+ "checksum" TEXT
- "content_type" VARCHAR
+ "content_type" TEXT
- "filename" VARCHAR NOT NULL
+ "filename" TEXT NOT NULL
- "key" VARCHAR NOT NULL
+ "key" TEXT NOT NULL
- "service_name" VARCHAR NOT NULL
+ "service_name" TEXT NOT NULL
@@ "active_storage_variant_records"
- "blob_id" BIGINT NOT NULL
+ "blob_id" INTEGER NOT NULL
- "variation_digest" VARCHAR NOT NULL
+ "variation_digest" TEXT NOT NULL
@@ "ar_internal_metadata"
- "key" VARCHAR NOT NULL
+ "key" TEXT NOT NULL
- "value" VARCHAR
+ "value" TEXT
@@ "bans"
- "ip_address" VARCHAR NOT NULL
+ "ip_address" TEXT NOT NULL
@@ "boosts"
- "content" VARCHAR(16) NOT NULL
+ "content" TEXT NOT NULL
@@ "memberships"
- "involvement" VARCHAR DEFAULT 'mentions'
+ "involvement" TEXT DEFAULT 'mentions'
+ "message_search_index" (the table is only in sql/prisma.sql)
+ "message_search_index_config" (the table is only in sql/prisma.sql)
+ "message_search_index_content" (the table is only in sql/prisma.sql)
+ "message_search_index_data" (the table is only in sql/prisma.sql)
+ "message_search_index_docsize" (the table is only in sql/prisma.sql)
+ "message_search_index_idx" (the table is only in sql/prisma.sql)
@@ "messages"
- "client_message_id" VARCHAR NOT NULL
+ "client_message_id" TEXT NOT NULL
@@ "push_subscriptions"
- "auth_key" VARCHAR
+ "auth_key" TEXT
- "endpoint" VARCHAR
+ "endpoint" TEXT
- "p256dh_key" VARCHAR
+ "p256dh_key" TEXT
- "user_agent" VARCHAR
+ "user_agent" TEXT
@@ "rooms"
- "name" VARCHAR
+ "name" TEXT
- "type" VARCHAR NOT NULL
+ "type" TEXT NOT NULL
@@ "schema_migrations"
- "version" VARCHAR NOT NULL
+ "version" TEXT NOT NULL
@@ "searches"
- "query" VARCHAR NOT NULL
+ "query" TEXT NOT NULL
@@ "sessions"
- "ip_address" VARCHAR
+ "ip_address" TEXT
- "token" VARCHAR NOT NULL
+ "token" TEXT NOT NULL
- "user_agent" VARCHAR
+ "user_agent" TEXT
@@ "users"
- "bot_token" VARCHAR
+ "bot_token" TEXT
- "email_address" VARCHAR
+ "email_address" TEXT
- "name" VARCHAR NOT NULL
+ "name" TEXT NOT NULL
- "password_digest" VARCHAR
+ "password_digest" TEXT
@@ "webhooks"
- "url" VARCHAR
+ "url" TEXT
@@ only in source.sql
- CREATE VIRTUAL TABLE message_search_index USING fts5 (body, tokenize=porter)
```

</details>

<details><summary>hekireki-drizzle: 17 tables written differently, 3 tables only here, 1 other statement only in source.sql</summary>

```diff
@@ "accounts"
- "created_at" DATETIME(6) NOT NULL
- "join_code" VARCHAR NOT NULL
+ "join_code" TEXT NOT NULL
- "name" VARCHAR NOT NULL
+ "name" TEXT NOT NULL
- "settings" JSON
- "updated_at" DATETIME(6) NOT NULL
@@ "action_text_rich_texts"
- "created_at" DATETIME(6) NOT NULL
- "name" VARCHAR NOT NULL
+ "name" TEXT NOT NULL
- "record_type" VARCHAR NOT NULL
+ "record_type" TEXT NOT NULL
- "updated_at" DATETIME(6) NOT NULL
@@ "active_storage_attachments"
- "blob_id" BIGINT NOT NULL
+ "blob_id" INTEGER NOT NULL
- "created_at" DATETIME(6) NOT NULL
- "name" VARCHAR NOT NULL
+ "name" TEXT NOT NULL
- "record_type" VARCHAR NOT NULL
+ "record_type" TEXT NOT NULL
@@ "active_storage_blobs"
- "checksum" VARCHAR
+ "checksum" TEXT
- "content_type" VARCHAR
+ "content_type" TEXT
- "created_at" DATETIME(6) NOT NULL
- "filename" VARCHAR NOT NULL
+ "filename" TEXT NOT NULL
- "key" VARCHAR NOT NULL
+ "key" TEXT NOT NULL
- "service_name" VARCHAR NOT NULL
+ "service_name" TEXT NOT NULL
@@ "active_storage_variant_records"
- "blob_id" BIGINT NOT NULL
+ "blob_id" INTEGER NOT NULL
- "variation_digest" VARCHAR NOT NULL
+ "variation_digest" TEXT NOT NULL
@@ "ar_internal_metadata"
- "key" VARCHAR NOT NULL
+ "key" TEXT
- "value" VARCHAR
+ "value" TEXT
- "created_at" DATETIME(6) NOT NULL
- "updated_at" DATETIME(6) NOT NULL
@@ "bans"
- "created_at" DATETIME(6) NOT NULL
- "ip_address" VARCHAR NOT NULL
+ "ip_address" TEXT NOT NULL
- "updated_at" DATETIME(6) NOT NULL
@@ "boosts"
- "content" VARCHAR(16) NOT NULL
+ "content" TEXT NOT NULL
- "created_at" DATETIME(6) NOT NULL
- "updated_at" DATETIME(6) NOT NULL
@@ "memberships"
- "connected_at" DATETIME(6)
- "created_at" DATETIME(6) NOT NULL
- "involvement" VARCHAR DEFAULT 'mentions'
+ "involvement" TEXT DEFAULT 'mentions'
- "unread_at" DATETIME(6)
- "updated_at" DATETIME(6) NOT NULL
- INDEX ("room_id", "created_at")
+ "message_search_index_content" (the table is only in sql/drizzle.sql)
+ "message_search_index_data" (the table is only in sql/drizzle.sql)
+ "message_search_index_docsize" (the table is only in sql/drizzle.sql)
@@ "messages"
- "client_message_id" VARCHAR NOT NULL
+ "client_message_id" TEXT NOT NULL
- "created_at" DATETIME(6) NOT NULL
- "updated_at" DATETIME(6) NOT NULL
- INDEX ("room_id", "created_at")
- INDEX ("room_id", "updated_at")
@@ "push_subscriptions"
- "auth_key" VARCHAR
+ "auth_key" TEXT
- "created_at" DATETIME(6) NOT NULL
- "endpoint" VARCHAR
+ "endpoint" TEXT
- "p256dh_key" VARCHAR
+ "p256dh_key" TEXT
- "updated_at" DATETIME(6) NOT NULL
- "user_agent" VARCHAR
+ "user_agent" TEXT
@@ "rooms"
- "created_at" DATETIME(6) NOT NULL
- "name" VARCHAR
+ "name" TEXT
- "type" VARCHAR NOT NULL
+ "type" TEXT NOT NULL
- "updated_at" DATETIME(6) NOT NULL
@@ "schema_migrations"
- "version" VARCHAR NOT NULL
+ "version" TEXT
@@ "searches"
- "created_at" DATETIME(6) NOT NULL
- "query" VARCHAR NOT NULL
+ "query" TEXT NOT NULL
- "updated_at" DATETIME(6) NOT NULL
@@ "sessions"
- "created_at" DATETIME(6) NOT NULL
- "ip_address" VARCHAR
+ "ip_address" TEXT
- "last_active_at" DATETIME(6) NOT NULL
- "token" VARCHAR NOT NULL
+ "token" TEXT NOT NULL
- "updated_at" DATETIME(6) NOT NULL
- "user_agent" VARCHAR
+ "user_agent" TEXT
@@ "users"
- "bot_token" VARCHAR
+ "bot_token" TEXT
- "created_at" DATETIME(6) NOT NULL
- "email_address" VARCHAR
+ "email_address" TEXT
- "name" VARCHAR NOT NULL
+ "name" TEXT NOT NULL
- "password_digest" VARCHAR
@@ … 9 more lines in sql/drizzle.diff
```

</details>

<details><summary>hekireki-sqlalchemy: 17 tables written differently, 3 tables only here, 1 other statement only in source.sql</summary>

```diff
@@ "accounts"
- "created_at" DATETIME(6) NOT NULL
- "join_code" VARCHAR NOT NULL
+ "join_code" TEXT NOT NULL
- "name" VARCHAR NOT NULL
+ "name" TEXT NOT NULL
- "settings" JSON
- "updated_at" DATETIME(6) NOT NULL
@@ "action_text_rich_texts"
- "created_at" DATETIME(6) NOT NULL
- "name" VARCHAR NOT NULL
+ "name" TEXT NOT NULL
- "record_type" VARCHAR NOT NULL
+ "record_type" TEXT NOT NULL
- "updated_at" DATETIME(6) NOT NULL
@@ "active_storage_attachments"
- "blob_id" BIGINT NOT NULL
+ "blob_id" INTEGER NOT NULL
- "created_at" DATETIME(6) NOT NULL
- "name" VARCHAR NOT NULL
+ "name" TEXT NOT NULL
- "record_type" VARCHAR NOT NULL
+ "record_type" TEXT NOT NULL
@@ "active_storage_blobs"
- "checksum" VARCHAR
+ "checksum" TEXT
- "content_type" VARCHAR
+ "content_type" TEXT
- "created_at" DATETIME(6) NOT NULL
- "filename" VARCHAR NOT NULL
+ "filename" TEXT NOT NULL
- "key" VARCHAR NOT NULL
+ "key" TEXT NOT NULL
- "service_name" VARCHAR NOT NULL
+ "service_name" TEXT NOT NULL
@@ "active_storage_variant_records"
- "blob_id" BIGINT NOT NULL
+ "blob_id" INTEGER NOT NULL
- "variation_digest" VARCHAR NOT NULL
+ "variation_digest" TEXT NOT NULL
@@ "ar_internal_metadata"
- "key" VARCHAR NOT NULL
+ "key" TEXT NOT NULL
- "value" VARCHAR
+ "value" TEXT
- "created_at" DATETIME(6) NOT NULL
- "updated_at" DATETIME(6) NOT NULL
@@ "bans"
- "created_at" DATETIME(6) NOT NULL
- "ip_address" VARCHAR NOT NULL
+ "ip_address" TEXT NOT NULL
- "updated_at" DATETIME(6) NOT NULL
@@ "boosts"
- "content" VARCHAR(16) NOT NULL
+ "content" TEXT NOT NULL
- "created_at" DATETIME(6) NOT NULL
- "updated_at" DATETIME(6) NOT NULL
@@ "memberships"
- "connected_at" DATETIME(6)
- "created_at" DATETIME(6) NOT NULL
- "involvement" VARCHAR DEFAULT 'mentions'
+ "involvement" TEXT DEFAULT 'mentions'
- "unread_at" DATETIME(6)
- "updated_at" DATETIME(6) NOT NULL
- INDEX ("room_id", "created_at")
+ "message_search_index_content" (the table is only in sql/sqlalchemy.sql)
+ "message_search_index_data" (the table is only in sql/sqlalchemy.sql)
+ "message_search_index_docsize" (the table is only in sql/sqlalchemy.sql)
@@ "messages"
- "client_message_id" VARCHAR NOT NULL
+ "client_message_id" TEXT NOT NULL
- "created_at" DATETIME(6) NOT NULL
- "updated_at" DATETIME(6) NOT NULL
- INDEX ("room_id", "created_at")
- INDEX ("room_id", "updated_at")
@@ "push_subscriptions"
- "auth_key" VARCHAR
+ "auth_key" TEXT
- "created_at" DATETIME(6) NOT NULL
- "endpoint" VARCHAR
+ "endpoint" TEXT
- "p256dh_key" VARCHAR
+ "p256dh_key" TEXT
- "updated_at" DATETIME(6) NOT NULL
- "user_agent" VARCHAR
+ "user_agent" TEXT
@@ "rooms"
- "created_at" DATETIME(6) NOT NULL
- "name" VARCHAR
+ "name" TEXT
- "type" VARCHAR NOT NULL
+ "type" TEXT NOT NULL
- "updated_at" DATETIME(6) NOT NULL
@@ "schema_migrations"
- "version" VARCHAR NOT NULL
+ "version" TEXT NOT NULL
@@ "searches"
- "created_at" DATETIME(6) NOT NULL
- "query" VARCHAR NOT NULL
+ "query" TEXT NOT NULL
- "updated_at" DATETIME(6) NOT NULL
@@ "sessions"
- "created_at" DATETIME(6) NOT NULL
- "ip_address" VARCHAR
+ "ip_address" TEXT
- "last_active_at" DATETIME(6) NOT NULL
- "token" VARCHAR NOT NULL
+ "token" TEXT NOT NULL
- "updated_at" DATETIME(6) NOT NULL
- "user_agent" VARCHAR
+ "user_agent" TEXT
@@ "users"
- "bot_token" VARCHAR
+ "bot_token" TEXT
- "created_at" DATETIME(6) NOT NULL
- "email_address" VARCHAR
+ "email_address" TEXT
- "name" VARCHAR NOT NULL
+ "name" TEXT NOT NULL
- "password_digest" VARCHAR
@@ … 9 more lines in sql/sqlalchemy.diff
```

</details>

<details><summary>hekireki-django: 17 tables written differently, 3 tables only here, 1 other statement only in source.sql</summary>

```diff
@@ "accounts"
- "created_at" DATETIME(6) NOT NULL
- "join_code" VARCHAR NOT NULL
+ "join_code" TEXT NOT NULL
- "name" VARCHAR NOT NULL
+ "name" TEXT NOT NULL
- "settings" JSON
- "singleton_guard" INTEGER NOT NULL DEFAULT 0
+ "singleton_guard" INTEGER NOT NULL
- "updated_at" DATETIME(6) NOT NULL
@@ "action_text_rich_texts"
- "created_at" DATETIME(6) NOT NULL
- "name" VARCHAR NOT NULL
+ "name" TEXT NOT NULL
- "record_type" VARCHAR NOT NULL
+ "record_type" TEXT NOT NULL
- "updated_at" DATETIME(6) NOT NULL
@@ "active_storage_attachments"
- "blob_id" BIGINT NOT NULL
+ "blob_id" INTEGER NOT NULL
- "created_at" DATETIME(6) NOT NULL
- "name" VARCHAR NOT NULL
+ "name" TEXT NOT NULL
- "record_type" VARCHAR NOT NULL
+ "record_type" TEXT NOT NULL
@@ "active_storage_blobs"
- "checksum" VARCHAR
+ "checksum" TEXT
- "content_type" VARCHAR
+ "content_type" TEXT
- "created_at" DATETIME(6) NOT NULL
- "filename" VARCHAR NOT NULL
+ "filename" TEXT NOT NULL
- "key" VARCHAR NOT NULL
+ "key" TEXT NOT NULL
- "service_name" VARCHAR NOT NULL
+ "service_name" TEXT NOT NULL
@@ "active_storage_variant_records"
- "blob_id" BIGINT NOT NULL
+ "blob_id" INTEGER NOT NULL
- "variation_digest" VARCHAR NOT NULL
+ "variation_digest" TEXT NOT NULL
@@ "ar_internal_metadata"
- "key" VARCHAR NOT NULL
+ "key" TEXT NOT NULL
- "value" VARCHAR
+ "value" TEXT
- "created_at" DATETIME(6) NOT NULL
- "updated_at" DATETIME(6) NOT NULL
@@ "bans"
- "created_at" DATETIME(6) NOT NULL
- "ip_address" VARCHAR NOT NULL
+ "ip_address" TEXT NOT NULL
- "updated_at" DATETIME(6) NOT NULL
@@ "boosts"
- "content" VARCHAR(16) NOT NULL
+ "content" TEXT NOT NULL
- "created_at" DATETIME(6) NOT NULL
- "updated_at" DATETIME(6) NOT NULL
@@ "memberships"
- "connected_at" DATETIME(6)
- "connections" INTEGER NOT NULL DEFAULT 0
+ "connections" INTEGER NOT NULL
- "created_at" DATETIME(6) NOT NULL
- "involvement" VARCHAR DEFAULT 'mentions'
+ "involvement" TEXT
- "unread_at" DATETIME(6)
- "updated_at" DATETIME(6) NOT NULL
- INDEX ("room_id", "created_at")
+ "message_search_index_content" (the table is only in sql/django.sql)
+ "message_search_index_data" (the table is only in sql/django.sql)
+ "message_search_index_docsize" (the table is only in sql/django.sql)
@@ "messages"
- "client_message_id" VARCHAR NOT NULL
+ "client_message_id" TEXT NOT NULL
- "created_at" DATETIME(6) NOT NULL
- "updated_at" DATETIME(6) NOT NULL
- INDEX ("room_id", "created_at")
- INDEX ("room_id", "updated_at")
@@ "push_subscriptions"
- "auth_key" VARCHAR
+ "auth_key" TEXT
- "created_at" DATETIME(6) NOT NULL
- "endpoint" VARCHAR
+ "endpoint" TEXT
- "p256dh_key" VARCHAR
+ "p256dh_key" TEXT
- "updated_at" DATETIME(6) NOT NULL
- "user_agent" VARCHAR
+ "user_agent" TEXT
@@ "rooms"
- "created_at" DATETIME(6) NOT NULL
- "name" VARCHAR
+ "name" TEXT
- "type" VARCHAR NOT NULL
+ "type" TEXT NOT NULL
- "updated_at" DATETIME(6) NOT NULL
- "messages_count" INTEGER NOT NULL DEFAULT 0
+ "messages_count" INTEGER NOT NULL
@@ "schema_migrations"
- "version" VARCHAR NOT NULL
+ "version" TEXT NOT NULL
@@ "searches"
- "created_at" DATETIME(6) NOT NULL
- "query" VARCHAR NOT NULL
+ "query" TEXT NOT NULL
- "updated_at" DATETIME(6) NOT NULL
@@ "sessions"
- "created_at" DATETIME(6) NOT NULL
- "ip_address" VARCHAR
+ "ip_address" TEXT
- "last_active_at" DATETIME(6) NOT NULL
- "token" VARCHAR NOT NULL
+ "token" TEXT NOT NULL
- "updated_at" DATETIME(6) NOT NULL
- "user_agent" VARCHAR
+ "user_agent" TEXT
@@ "users"
- "bot_token" VARCHAR
+ "bot_token" TEXT
@@ … 19 more lines in sql/django.diff
```

</details>

<details><summary>hekireki-gorm: 17 tables written differently, 3 tables only here, 1 other statement only in source.sql</summary>

```diff
@@ "accounts"
- "created_at" DATETIME(6) NOT NULL
- "join_code" VARCHAR NOT NULL
+ "join_code" TEXT NOT NULL
- "name" VARCHAR NOT NULL
+ "name" TEXT NOT NULL
- "settings" JSON
- "updated_at" DATETIME(6) NOT NULL
@@ "action_text_rich_texts"
- "created_at" DATETIME(6) NOT NULL
- "name" VARCHAR NOT NULL
+ "name" TEXT NOT NULL
- "record_id" BIGINT NOT NULL
+ "record_id" INTEGER NOT NULL
- "record_type" VARCHAR NOT NULL
+ "record_type" TEXT NOT NULL
- "updated_at" DATETIME(6) NOT NULL
- UNIQUE ("record_type", "record_id", "name")
+ UNIQUE ("name", "record_id", "record_type")
@@ "active_storage_attachments"
- "blob_id" BIGINT NOT NULL
+ "blob_id" INTEGER NOT NULL
- "created_at" DATETIME(6) NOT NULL
- "name" VARCHAR NOT NULL
+ "name" TEXT NOT NULL
- "record_id" BIGINT NOT NULL
+ "record_id" INTEGER NOT NULL
- "record_type" VARCHAR NOT NULL
+ "record_type" TEXT NOT NULL
- UNIQUE ("record_type", "record_id", "name", "blob_id")
+ UNIQUE ("blob_id", "name", "record_id", "record_type")
@@ "active_storage_blobs"
- "byte_size" BIGINT NOT NULL
+ "byte_size" INTEGER NOT NULL
- "checksum" VARCHAR
+ "checksum" TEXT
- "content_type" VARCHAR
+ "content_type" TEXT
- "created_at" DATETIME(6) NOT NULL
- "filename" VARCHAR NOT NULL
+ "filename" TEXT NOT NULL
- "key" VARCHAR NOT NULL
+ "key" TEXT NOT NULL
- "service_name" VARCHAR NOT NULL
+ "service_name" TEXT NOT NULL
@@ "active_storage_variant_records"
- "blob_id" BIGINT NOT NULL
+ "blob_id" INTEGER NOT NULL
- "variation_digest" VARCHAR NOT NULL
+ "variation_digest" TEXT NOT NULL
@@ "ar_internal_metadata"
- "key" VARCHAR NOT NULL
+ "key" TEXT
- "value" VARCHAR
+ "value" TEXT
- "created_at" DATETIME(6) NOT NULL
- "updated_at" DATETIME(6) NOT NULL
@@ "bans"
- "created_at" DATETIME(6) NOT NULL
- "ip_address" VARCHAR NOT NULL
+ "ip_address" TEXT NOT NULL
- "updated_at" DATETIME(6) NOT NULL
@@ "boosts"
- "content" VARCHAR(16) NOT NULL
+ "content" TEXT NOT NULL
- "created_at" DATETIME(6) NOT NULL
- "updated_at" DATETIME(6) NOT NULL
@@ "memberships"
- "connected_at" DATETIME(6)
- "created_at" DATETIME(6) NOT NULL
- "involvement" VARCHAR DEFAULT 'mentions'
+ "involvement" TEXT DEFAULT "mentions"
- "unread_at" DATETIME(6)
- "updated_at" DATETIME(6) NOT NULL
- INDEX ("room_id", "created_at")
+ "message_search_index_content" (the table is only in sql/gorm.sql)
+ "message_search_index_data" (the table is only in sql/gorm.sql)
+ "message_search_index_docsize" (the table is only in sql/gorm.sql)
@@ "messages"
- "client_message_id" VARCHAR NOT NULL
+ "client_message_id" TEXT NOT NULL
- "created_at" DATETIME(6) NOT NULL
- "updated_at" DATETIME(6) NOT NULL
- INDEX ("room_id", "created_at")
- INDEX ("room_id", "updated_at")
@@ "push_subscriptions"
- "auth_key" VARCHAR
+ "auth_key" TEXT
- "created_at" DATETIME(6) NOT NULL
- "endpoint" VARCHAR
+ "endpoint" TEXT
- "p256dh_key" VARCHAR
+ "p256dh_key" TEXT
- "updated_at" DATETIME(6) NOT NULL
- "user_agent" VARCHAR
+ "user_agent" TEXT
- INDEX ("endpoint", "p256dh_key", "auth_key")
+ INDEX ("auth_key", "endpoint", "p256dh_key")
@@ "rooms"
- "created_at" DATETIME(6) NOT NULL
- "creator_id" BIGINT NOT NULL
+ "creator_id" INTEGER NOT NULL
- "name" VARCHAR
+ "name" TEXT
- "type" VARCHAR NOT NULL
+ "type" TEXT NOT NULL
- "updated_at" DATETIME(6) NOT NULL
@@ "schema_migrations"
- "version" VARCHAR NOT NULL
+ "version" TEXT
@@ "searches"
- "created_at" DATETIME(6) NOT NULL
- "query" VARCHAR NOT NULL
+ "query" TEXT NOT NULL
- "updated_at" DATETIME(6) NOT NULL
@@ "sessions"
- "created_at" DATETIME(6) NOT NULL
- "ip_address" VARCHAR
+ "ip_address" TEXT
- "last_active_at" DATETIME(6) NOT NULL
@@ … 23 more lines in sql/gorm.diff
```

</details>

<details><summary>hekireki-sea-orm: 16 tables written differently, 3 tables only here, 1 other statement only in source.sql</summary>

```diff
@@ "accounts"
- "created_at" DATETIME(6) NOT NULL
- "custom_styles" TEXT
+ "custom_styles" VARCHAR
- "settings" JSON
- "updated_at" DATETIME(6) NOT NULL
@@ "action_text_rich_texts"
- "body" TEXT
+ "body" VARCHAR
- "created_at" DATETIME(6) NOT NULL
- "updated_at" DATETIME(6) NOT NULL
- UNIQUE ("record_type", "record_id", "name")
@@ "active_storage_attachments"
- "blob_id" BIGINT NOT NULL
+ "blob_id" INTEGER NOT NULL
- "created_at" DATETIME(6) NOT NULL
- UNIQUE ("record_type", "record_id", "name", "blob_id")
- INDEX ("blob_id")
@@ "active_storage_blobs"
- "created_at" DATETIME(6) NOT NULL
- "metadata" TEXT
+ "metadata" VARCHAR
@@ "active_storage_variant_records"
- "blob_id" BIGINT NOT NULL
+ "blob_id" INTEGER NOT NULL
- UNIQUE ("blob_id", "variation_digest")
@@ "ar_internal_metadata"
- "created_at" DATETIME(6) NOT NULL
- "updated_at" DATETIME(6) NOT NULL
@@ "bans"
- "created_at" DATETIME(6) NOT NULL
- "updated_at" DATETIME(6) NOT NULL
- INDEX ("ip_address")
- INDEX ("user_id")
@@ "boosts"
- "content" VARCHAR(16) NOT NULL
+ "content" VARCHAR NOT NULL
- "created_at" DATETIME(6) NOT NULL
- "updated_at" DATETIME(6) NOT NULL
- INDEX ("booster_id")
- INDEX ("message_id")
@@ "memberships"
- "connected_at" DATETIME(6)
- "created_at" DATETIME(6) NOT NULL
- "unread_at" DATETIME(6)
- "updated_at" DATETIME(6) NOT NULL
- UNIQUE ("room_id", "user_id")
- INDEX ("room_id")
- INDEX ("room_id", "created_at")
- INDEX ("user_id")
+ "message_search_index_content" (the table is only in sql/sea-orm.sql)
+ "message_search_index_data" (the table is only in sql/sea-orm.sql)
+ "message_search_index_docsize" (the table is only in sql/sea-orm.sql)
@@ "messages"
- "created_at" DATETIME(6) NOT NULL
- "updated_at" DATETIME(6) NOT NULL
- INDEX ("creator_id")
- INDEX ("room_id")
- INDEX ("room_id", "created_at")
- INDEX ("room_id", "updated_at")
@@ "push_subscriptions"
- "created_at" DATETIME(6) NOT NULL
- "updated_at" DATETIME(6) NOT NULL
- INDEX ("endpoint", "p256dh_key", "auth_key")
- INDEX ("user_id")
@@ "rooms"
- "created_at" DATETIME(6) NOT NULL
- "updated_at" DATETIME(6) NOT NULL
@@ "searches"
- "created_at" DATETIME(6) NOT NULL
- "updated_at" DATETIME(6) NOT NULL
- INDEX ("user_id")
@@ "sessions"
- "created_at" DATETIME(6) NOT NULL
- "last_active_at" DATETIME(6) NOT NULL
- "updated_at" DATETIME(6) NOT NULL
- INDEX ("user_id")
@@ "users"
- "bio" TEXT
+ "bio" VARCHAR
- "created_at" DATETIME(6) NOT NULL
- "updated_at" DATETIME(6) NOT NULL
@@ "webhooks"
- "created_at" DATETIME(6) NOT NULL
- "updated_at" DATETIME(6) NOT NULL
- INDEX ("user_id")
@@ only in source.sql
- CREATE VIRTUAL TABLE message_search_index USING fts5 (body, tokenize=porter)
```

</details>

<details><summary>hekireki-efcore: 17 tables written differently, 3 tables only here, 1 other statement only in source.sql</summary>

```diff
@@ "accounts"
- "created_at" DATETIME(6) NOT NULL
- "join_code" VARCHAR NOT NULL
+ "join_code" TEXT NOT NULL
- "name" VARCHAR NOT NULL
+ "name" TEXT NOT NULL
- "settings" JSON
- "updated_at" DATETIME(6) NOT NULL
@@ "action_text_rich_texts"
- "created_at" DATETIME(6) NOT NULL
- "name" VARCHAR NOT NULL
+ "name" TEXT NOT NULL
- "record_id" BIGINT NOT NULL
+ "record_id" INTEGER NOT NULL
- "record_type" VARCHAR NOT NULL
+ "record_type" TEXT NOT NULL
- "updated_at" DATETIME(6) NOT NULL
@@ "active_storage_attachments"
- "blob_id" BIGINT NOT NULL
+ "blob_id" INTEGER NOT NULL
- "created_at" DATETIME(6) NOT NULL
- "name" VARCHAR NOT NULL
+ "name" TEXT NOT NULL
- "record_id" BIGINT NOT NULL
+ "record_id" INTEGER NOT NULL
- "record_type" VARCHAR NOT NULL
+ "record_type" TEXT NOT NULL
@@ "active_storage_blobs"
- "byte_size" BIGINT NOT NULL
+ "byte_size" INTEGER NOT NULL
- "checksum" VARCHAR
+ "checksum" TEXT
- "content_type" VARCHAR
+ "content_type" TEXT
- "created_at" DATETIME(6) NOT NULL
- "filename" VARCHAR NOT NULL
+ "filename" TEXT NOT NULL
- "key" VARCHAR NOT NULL
+ "key" TEXT NOT NULL
- "service_name" VARCHAR NOT NULL
+ "service_name" TEXT NOT NULL
@@ "active_storage_variant_records"
- "blob_id" BIGINT NOT NULL
+ "blob_id" INTEGER NOT NULL
- "variation_digest" VARCHAR NOT NULL
+ "variation_digest" TEXT NOT NULL
@@ "ar_internal_metadata"
- "key" VARCHAR NOT NULL
+ "key" TEXT NOT NULL
- "value" VARCHAR
+ "value" TEXT
- "created_at" DATETIME(6) NOT NULL
- "updated_at" DATETIME(6) NOT NULL
@@ "bans"
- "created_at" DATETIME(6) NOT NULL
- "ip_address" VARCHAR NOT NULL
+ "ip_address" TEXT NOT NULL
- "updated_at" DATETIME(6) NOT NULL
@@ "boosts"
- "content" VARCHAR(16) NOT NULL
+ "content" TEXT NOT NULL
- "created_at" DATETIME(6) NOT NULL
- "updated_at" DATETIME(6) NOT NULL
@@ "memberships"
- "connected_at" DATETIME(6)
- "created_at" DATETIME(6) NOT NULL
- "involvement" VARCHAR DEFAULT 'mentions'
+ "involvement" TEXT DEFAULT 'mentions'
- "unread_at" DATETIME(6)
- "updated_at" DATETIME(6) NOT NULL
- INDEX ("room_id", "created_at")
+ "message_search_index_content" (the table is only in sql/efcore.sql)
+ "message_search_index_data" (the table is only in sql/efcore.sql)
+ "message_search_index_docsize" (the table is only in sql/efcore.sql)
@@ "messages"
- "client_message_id" VARCHAR NOT NULL
+ "client_message_id" TEXT NOT NULL
- "created_at" DATETIME(6) NOT NULL
- "updated_at" DATETIME(6) NOT NULL
- INDEX ("room_id", "created_at")
- INDEX ("room_id", "updated_at")
@@ "push_subscriptions"
- "auth_key" VARCHAR
+ "auth_key" TEXT
- "created_at" DATETIME(6) NOT NULL
- "endpoint" VARCHAR
+ "endpoint" TEXT
- "p256dh_key" VARCHAR
+ "p256dh_key" TEXT
- "updated_at" DATETIME(6) NOT NULL
- "user_agent" VARCHAR
+ "user_agent" TEXT
@@ "rooms"
- "created_at" DATETIME(6) NOT NULL
- "creator_id" BIGINT NOT NULL
+ "creator_id" INTEGER NOT NULL
- "name" VARCHAR
+ "name" TEXT
- "type" VARCHAR NOT NULL
+ "type" TEXT NOT NULL
- "updated_at" DATETIME(6) NOT NULL
@@ "schema_migrations"
- "version" VARCHAR NOT NULL
+ "version" TEXT NOT NULL
@@ "searches"
- "created_at" DATETIME(6) NOT NULL
- "query" VARCHAR NOT NULL
+ "query" TEXT NOT NULL
- "updated_at" DATETIME(6) NOT NULL
@@ "sessions"
- "created_at" DATETIME(6) NOT NULL
- "ip_address" VARCHAR
+ "ip_address" TEXT
- "last_active_at" DATETIME(6) NOT NULL
- "token" VARCHAR NOT NULL
+ "token" TEXT NOT NULL
- "updated_at" DATETIME(6) NOT NULL
- "user_agent" VARCHAR
+ "user_agent" TEXT
@@ "users"
@@ … 17 more lines in sql/efcore.diff
```

</details>

Written by `node pull/pull.mjs`; do not edit.
