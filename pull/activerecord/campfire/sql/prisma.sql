CREATE TABLE "accounts" (
    "id" INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
    "created_at" datetime(6) NOT NULL,
    "custom_styles" TEXT,
    "join_code" TEXT NOT NULL,
    "name" TEXT NOT NULL,
    "settings" json,
    "singleton_guard" INTEGER NOT NULL DEFAULT 0,
    "updated_at" datetime(6) NOT NULL
);

CREATE TABLE "action_text_rich_texts" (
    "id" INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
    "body" TEXT,
    "created_at" datetime(6) NOT NULL,
    "name" TEXT NOT NULL,
    "record_id" BIGINT NOT NULL,
    "record_type" TEXT NOT NULL,
    "updated_at" datetime(6) NOT NULL
);

CREATE TABLE "active_storage_attachments" (
    "id" INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
    "blob_id" INTEGER NOT NULL,
    "created_at" datetime(6) NOT NULL,
    "name" TEXT NOT NULL,
    "record_id" BIGINT NOT NULL,
    "record_type" TEXT NOT NULL,
    CONSTRAINT "active_storage_attachments_blob_id_fkey" FOREIGN KEY ("blob_id") REFERENCES "active_storage_blobs" ("id") ON DELETE NO ACTION ON UPDATE NO ACTION
);

CREATE TABLE "active_storage_blobs" (
    "id" INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
    "byte_size" BIGINT NOT NULL,
    "checksum" TEXT,
    "content_type" TEXT,
    "created_at" datetime(6) NOT NULL,
    "filename" TEXT NOT NULL,
    "key" TEXT NOT NULL,
    "metadata" TEXT,
    "service_name" TEXT NOT NULL
);

CREATE TABLE "active_storage_variant_records" (
    "id" INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
    "blob_id" INTEGER NOT NULL,
    "variation_digest" TEXT NOT NULL,
    CONSTRAINT "active_storage_variant_records_blob_id_fkey" FOREIGN KEY ("blob_id") REFERENCES "active_storage_blobs" ("id") ON DELETE NO ACTION ON UPDATE NO ACTION
);

CREATE TABLE "ar_internal_metadata" (
    "key" TEXT NOT NULL PRIMARY KEY,
    "value" TEXT,
    "created_at" datetime(6) NOT NULL,
    "updated_at" datetime(6) NOT NULL
);

CREATE TABLE "bans" (
    "id" INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
    "created_at" datetime(6) NOT NULL,
    "ip_address" TEXT NOT NULL,
    "updated_at" datetime(6) NOT NULL,
    "user_id" INTEGER NOT NULL,
    CONSTRAINT "bans_user_id_fkey" FOREIGN KEY ("user_id") REFERENCES "users" ("id") ON DELETE NO ACTION ON UPDATE NO ACTION
);

CREATE TABLE "boosts" (
    "id" INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
    "booster_id" INTEGER NOT NULL,
    "content" TEXT NOT NULL,
    "created_at" datetime(6) NOT NULL,
    "message_id" INTEGER NOT NULL,
    "updated_at" datetime(6) NOT NULL,
    CONSTRAINT "boosts_message_id_fkey" FOREIGN KEY ("message_id") REFERENCES "messages" ("id") ON DELETE NO ACTION ON UPDATE NO ACTION
);

CREATE TABLE "memberships" (
    "id" INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
    "connected_at" datetime(6),
    "connections" INTEGER NOT NULL DEFAULT 0,
    "created_at" datetime(6) NOT NULL,
    "involvement" TEXT DEFAULT 'mentions',
    "room_id" INTEGER NOT NULL,
    "unread_at" datetime(6),
    "updated_at" datetime(6) NOT NULL,
    "user_id" INTEGER NOT NULL
);

CREATE TABLE "message_search_index" (
    "body" 
);

CREATE TABLE "message_search_index_config" (
    "k"  NOT NULL PRIMARY KEY,
    "v" 
);

CREATE TABLE "message_search_index_content" (
    "id" INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
    "c0" 
);

CREATE TABLE "message_search_index_data" (
    "id" INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
    "block" BLOB
);

CREATE TABLE "message_search_index_docsize" (
    "id" INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
    "sz" BLOB
);

CREATE TABLE "message_search_index_idx" (
    "segid"  NOT NULL,
    "term"  NOT NULL,
    "pgno" ,

    PRIMARY KEY ("segid", "term")
);

CREATE TABLE "messages" (
    "id" INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
    "client_message_id" TEXT NOT NULL,
    "created_at" datetime(6) NOT NULL,
    "creator_id" INTEGER NOT NULL,
    "room_id" INTEGER NOT NULL,
    "updated_at" datetime(6) NOT NULL,
    CONSTRAINT "messages_creator_id_fkey" FOREIGN KEY ("creator_id") REFERENCES "users" ("id") ON DELETE NO ACTION ON UPDATE NO ACTION,
    CONSTRAINT "messages_room_id_fkey" FOREIGN KEY ("room_id") REFERENCES "rooms" ("id") ON DELETE NO ACTION ON UPDATE NO ACTION
);

CREATE TABLE "push_subscriptions" (
    "id" INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
    "auth_key" TEXT,
    "created_at" datetime(6) NOT NULL,
    "endpoint" TEXT,
    "p256dh_key" TEXT,
    "updated_at" datetime(6) NOT NULL,
    "user_agent" TEXT,
    "user_id" INTEGER NOT NULL,
    CONSTRAINT "push_subscriptions_user_id_fkey" FOREIGN KEY ("user_id") REFERENCES "users" ("id") ON DELETE NO ACTION ON UPDATE NO ACTION
);

CREATE TABLE "rooms" (
    "id" INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
    "created_at" datetime(6) NOT NULL,
    "creator_id" BIGINT NOT NULL,
    "name" TEXT,
    "type" TEXT NOT NULL,
    "updated_at" datetime(6) NOT NULL,
    "messages_count" INTEGER NOT NULL DEFAULT 0
);

CREATE TABLE "schema_migrations" (
    "version" TEXT NOT NULL PRIMARY KEY
);

CREATE TABLE "searches" (
    "id" INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
    "created_at" datetime(6) NOT NULL,
    "query" TEXT NOT NULL,
    "updated_at" datetime(6) NOT NULL,
    "user_id" INTEGER NOT NULL,
    CONSTRAINT "searches_user_id_fkey" FOREIGN KEY ("user_id") REFERENCES "users" ("id") ON DELETE NO ACTION ON UPDATE NO ACTION
);

CREATE TABLE "sessions" (
    "id" INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
    "created_at" datetime(6) NOT NULL,
    "ip_address" TEXT,
    "last_active_at" datetime(6) NOT NULL,
    "token" TEXT NOT NULL,
    "updated_at" datetime(6) NOT NULL,
    "user_agent" TEXT,
    "user_id" INTEGER NOT NULL,
    CONSTRAINT "sessions_user_id_fkey" FOREIGN KEY ("user_id") REFERENCES "users" ("id") ON DELETE NO ACTION ON UPDATE NO ACTION
);

CREATE TABLE "users" (
    "id" INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
    "bio" TEXT,
    "bot_token" TEXT,
    "created_at" datetime(6) NOT NULL,
    "email_address" TEXT,
    "name" TEXT NOT NULL,
    "password_digest" TEXT,
    "role" INTEGER NOT NULL DEFAULT 0,
    "status" INTEGER NOT NULL DEFAULT 0,
    "updated_at" datetime(6) NOT NULL
);

CREATE TABLE "webhooks" (
    "id" INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
    "created_at" datetime(6) NOT NULL,
    "updated_at" datetime(6) NOT NULL,
    "url" TEXT,
    "user_id" INTEGER NOT NULL,
    CONSTRAINT "webhooks_user_id_fkey" FOREIGN KEY ("user_id") REFERENCES "users" ("id") ON DELETE NO ACTION ON UPDATE NO ACTION
);

CREATE UNIQUE INDEX "index_accounts_on_singleton_guard" ON "accounts"("singleton_guard");

CREATE UNIQUE INDEX "index_action_text_rich_texts_uniqueness" ON "action_text_rich_texts"("record_type", "record_id", "name");

CREATE INDEX "index_active_storage_attachments_on_blob_id" ON "active_storage_attachments"("blob_id");

CREATE UNIQUE INDEX "index_active_storage_attachments_uniqueness" ON "active_storage_attachments"("record_type", "record_id", "name", "blob_id");

CREATE UNIQUE INDEX "index_active_storage_blobs_on_key" ON "active_storage_blobs"("key");

CREATE UNIQUE INDEX "index_active_storage_variant_records_uniqueness" ON "active_storage_variant_records"("blob_id", "variation_digest");

CREATE INDEX "index_bans_on_user_id" ON "bans"("user_id");

CREATE INDEX "index_bans_on_ip_address" ON "bans"("ip_address");

CREATE INDEX "index_boosts_on_message_id" ON "boosts"("message_id");

CREATE INDEX "index_boosts_on_booster_id" ON "boosts"("booster_id");

CREATE INDEX "index_memberships_on_user_id" ON "memberships"("user_id");

CREATE INDEX "index_memberships_on_room_id" ON "memberships"("room_id");

CREATE INDEX "index_memberships_on_room_id_and_created_at" ON "memberships"("room_id", "created_at");

CREATE UNIQUE INDEX "index_memberships_on_room_id_and_user_id" ON "memberships"("room_id", "user_id");

CREATE INDEX "index_messages_on_room_id" ON "messages"("room_id");

CREATE INDEX "index_messages_on_room_id_and_updated_at" ON "messages"("room_id", "updated_at");

CREATE INDEX "index_messages_on_room_id_and_created_at" ON "messages"("room_id", "created_at");

CREATE INDEX "index_messages_on_creator_id" ON "messages"("creator_id");

CREATE INDEX "index_push_subscriptions_on_user_id" ON "push_subscriptions"("user_id");

CREATE INDEX "idx_on_endpoint_p256dh_key_auth_key_7553014576" ON "push_subscriptions"("endpoint", "p256dh_key", "auth_key");

CREATE INDEX "index_searches_on_user_id" ON "searches"("user_id");

CREATE UNIQUE INDEX "index_sessions_on_token" ON "sessions"("token");

CREATE INDEX "index_sessions_on_user_id" ON "sessions"("user_id");

CREATE UNIQUE INDEX "index_users_on_bot_token" ON "users"("bot_token");

CREATE UNIQUE INDEX "index_users_on_email_address" ON "users"("email_address");

CREATE INDEX "index_webhooks_on_user_id" ON "webhooks"("user_id");
