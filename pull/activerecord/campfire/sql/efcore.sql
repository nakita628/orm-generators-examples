CREATE TABLE "accounts" (
    "id" INTEGER NOT NULL CONSTRAINT "accounts_pkey" PRIMARY KEY AUTOINCREMENT,
    "custom_styles" TEXT NULL,
    "join_code" TEXT NOT NULL,
    "name" TEXT NOT NULL,
    "singleton_guard" INTEGER NOT NULL DEFAULT 0
);

CREATE TABLE "action_text_rich_texts" (
    "id" INTEGER NOT NULL CONSTRAINT "action_text_rich_texts_pkey" PRIMARY KEY AUTOINCREMENT,
    "body" TEXT NULL,
    "name" TEXT NOT NULL,
    "record_id" INTEGER NOT NULL,
    "record_type" TEXT NOT NULL
);

CREATE TABLE "active_storage_blobs" (
    "id" INTEGER NOT NULL CONSTRAINT "active_storage_blobs_pkey" PRIMARY KEY AUTOINCREMENT,
    "byte_size" INTEGER NOT NULL,
    "checksum" TEXT NULL,
    "content_type" TEXT NULL,
    "filename" TEXT NOT NULL,
    "key" TEXT NOT NULL,
    "metadata" TEXT NULL,
    "service_name" TEXT NOT NULL
);

CREATE TABLE "ar_internal_metadata" (
    "key" TEXT NOT NULL CONSTRAINT "ar_internal_metadata_pkey" PRIMARY KEY,
    "value" TEXT NULL
);

CREATE TABLE "memberships" (
    "id" INTEGER NOT NULL CONSTRAINT "memberships_pkey" PRIMARY KEY AUTOINCREMENT,
    "connections" INTEGER NOT NULL DEFAULT 0,
    "involvement" TEXT NULL DEFAULT 'mentions',
    "room_id" INTEGER NOT NULL,
    "user_id" INTEGER NOT NULL
);

CREATE TABLE "message_search_index_content" (
    "id" INTEGER NOT NULL CONSTRAINT "message_search_index_content_pkey" PRIMARY KEY AUTOINCREMENT
);

CREATE TABLE "message_search_index_data" (
    "id" INTEGER NOT NULL CONSTRAINT "message_search_index_data_pkey" PRIMARY KEY AUTOINCREMENT,
    "block" BLOB NULL
);

CREATE TABLE "message_search_index_docsize" (
    "id" INTEGER NOT NULL CONSTRAINT "message_search_index_docsize_pkey" PRIMARY KEY AUTOINCREMENT,
    "sz" BLOB NULL
);

CREATE TABLE "rooms" (
    "id" INTEGER NOT NULL CONSTRAINT "rooms_pkey" PRIMARY KEY AUTOINCREMENT,
    "creator_id" INTEGER NOT NULL,
    "name" TEXT NULL,
    "type" TEXT NOT NULL,
    "messages_count" INTEGER NOT NULL DEFAULT 0
);

CREATE TABLE "schema_migrations" (
    "version" TEXT NOT NULL CONSTRAINT "schema_migrations_pkey" PRIMARY KEY
);

CREATE TABLE "users" (
    "id" INTEGER NOT NULL CONSTRAINT "users_pkey" PRIMARY KEY AUTOINCREMENT,
    "bio" TEXT NULL,
    "bot_token" TEXT NULL,
    "email_address" TEXT NULL,
    "name" TEXT NOT NULL,
    "password_digest" TEXT NULL,
    "role" INTEGER NOT NULL DEFAULT 0,
    "status" INTEGER NOT NULL DEFAULT 0
);

CREATE TABLE "active_storage_attachments" (
    "id" INTEGER NOT NULL CONSTRAINT "active_storage_attachments_pkey" PRIMARY KEY AUTOINCREMENT,
    "blob_id" INTEGER NOT NULL,
    "name" TEXT NOT NULL,
    "record_id" INTEGER NOT NULL,
    "record_type" TEXT NOT NULL,
    CONSTRAINT "active_storage_attachments_blob_id_fkey" FOREIGN KEY ("blob_id") REFERENCES "active_storage_blobs" ("id")
);

CREATE TABLE "active_storage_variant_records" (
    "id" INTEGER NOT NULL CONSTRAINT "active_storage_variant_records_pkey" PRIMARY KEY AUTOINCREMENT,
    "blob_id" INTEGER NOT NULL,
    "variation_digest" TEXT NOT NULL,
    CONSTRAINT "active_storage_variant_records_blob_id_fkey" FOREIGN KEY ("blob_id") REFERENCES "active_storage_blobs" ("id")
);

CREATE TABLE "bans" (
    "id" INTEGER NOT NULL CONSTRAINT "bans_pkey" PRIMARY KEY AUTOINCREMENT,
    "ip_address" TEXT NOT NULL,
    "user_id" INTEGER NOT NULL,
    CONSTRAINT "bans_user_id_fkey" FOREIGN KEY ("user_id") REFERENCES "users" ("id")
);

CREATE TABLE "messages" (
    "id" INTEGER NOT NULL CONSTRAINT "messages_pkey" PRIMARY KEY AUTOINCREMENT,
    "client_message_id" TEXT NOT NULL,
    "creator_id" INTEGER NOT NULL,
    "room_id" INTEGER NOT NULL,
    CONSTRAINT "messages_creator_id_fkey" FOREIGN KEY ("creator_id") REFERENCES "users" ("id"),
    CONSTRAINT "messages_room_id_fkey" FOREIGN KEY ("room_id") REFERENCES "rooms" ("id")
);

CREATE TABLE "push_subscriptions" (
    "id" INTEGER NOT NULL CONSTRAINT "push_subscriptions_pkey" PRIMARY KEY AUTOINCREMENT,
    "auth_key" TEXT NULL,
    "endpoint" TEXT NULL,
    "p256dh_key" TEXT NULL,
    "user_agent" TEXT NULL,
    "user_id" INTEGER NOT NULL,
    CONSTRAINT "push_subscriptions_user_id_fkey" FOREIGN KEY ("user_id") REFERENCES "users" ("id")
);

CREATE TABLE "searches" (
    "id" INTEGER NOT NULL CONSTRAINT "searches_pkey" PRIMARY KEY AUTOINCREMENT,
    "query" TEXT NOT NULL,
    "user_id" INTEGER NOT NULL,
    CONSTRAINT "searches_user_id_fkey" FOREIGN KEY ("user_id") REFERENCES "users" ("id")
);

CREATE TABLE "sessions" (
    "id" INTEGER NOT NULL CONSTRAINT "sessions_pkey" PRIMARY KEY AUTOINCREMENT,
    "ip_address" TEXT NULL,
    "token" TEXT NOT NULL,
    "user_agent" TEXT NULL,
    "user_id" INTEGER NOT NULL,
    CONSTRAINT "sessions_user_id_fkey" FOREIGN KEY ("user_id") REFERENCES "users" ("id")
);

CREATE TABLE "webhooks" (
    "id" INTEGER NOT NULL CONSTRAINT "webhooks_pkey" PRIMARY KEY AUTOINCREMENT,
    "url" TEXT NULL,
    "user_id" INTEGER NOT NULL,
    CONSTRAINT "webhooks_user_id_fkey" FOREIGN KEY ("user_id") REFERENCES "users" ("id")
);

CREATE TABLE "boosts" (
    "id" INTEGER NOT NULL CONSTRAINT "boosts_pkey" PRIMARY KEY AUTOINCREMENT,
    "booster_id" INTEGER NOT NULL,
    "content" TEXT NOT NULL,
    "message_id" INTEGER NOT NULL,
    CONSTRAINT "boosts_message_id_fkey" FOREIGN KEY ("message_id") REFERENCES "messages" ("id")
);

CREATE UNIQUE INDEX "index_accounts_on_singleton_guard" ON "accounts" ("singleton_guard");

CREATE UNIQUE INDEX "index_action_text_rich_texts_uniqueness" ON "action_text_rich_texts" ("record_type", "record_id", "name");

CREATE INDEX "index_active_storage_attachments_on_blob_id" ON "active_storage_attachments" ("blob_id");

CREATE UNIQUE INDEX "index_active_storage_attachments_uniqueness" ON "active_storage_attachments" ("record_type", "record_id", "name", "blob_id");

CREATE UNIQUE INDEX "index_active_storage_blobs_on_key" ON "active_storage_blobs" ("key");

CREATE UNIQUE INDEX "index_active_storage_variant_records_uniqueness" ON "active_storage_variant_records" ("blob_id", "variation_digest");

CREATE INDEX "index_bans_on_ip_address" ON "bans" ("ip_address");

CREATE INDEX "index_bans_on_user_id" ON "bans" ("user_id");

CREATE INDEX "index_boosts_on_booster_id" ON "boosts" ("booster_id");

CREATE INDEX "index_boosts_on_message_id" ON "boosts" ("message_id");

CREATE INDEX "index_memberships_on_room_id" ON "memberships" ("room_id");

CREATE UNIQUE INDEX "index_memberships_on_room_id_and_user_id" ON "memberships" ("room_id", "user_id");

CREATE INDEX "index_memberships_on_user_id" ON "memberships" ("user_id");

CREATE INDEX "index_messages_on_creator_id" ON "messages" ("creator_id");

CREATE INDEX "index_messages_on_room_id" ON "messages" ("room_id");

CREATE INDEX "idx_on_endpoint_p256dh_key_auth_key_7553014576" ON "push_subscriptions" ("endpoint", "p256dh_key", "auth_key");

CREATE INDEX "index_push_subscriptions_on_user_id" ON "push_subscriptions" ("user_id");

CREATE INDEX "index_searches_on_user_id" ON "searches" ("user_id");

CREATE UNIQUE INDEX "index_sessions_on_token" ON "sessions" ("token");

CREATE INDEX "index_sessions_on_user_id" ON "sessions" ("user_id");

CREATE UNIQUE INDEX "index_users_on_bot_token" ON "users" ("bot_token");

CREATE UNIQUE INDEX "index_users_on_email_address" ON "users" ("email_address");

CREATE INDEX "index_webhooks_on_user_id" ON "webhooks" ("user_id");
