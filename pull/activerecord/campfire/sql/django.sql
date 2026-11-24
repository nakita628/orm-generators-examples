CREATE TABLE "accounts" ("id" integer NOT NULL PRIMARY KEY AUTOINCREMENT, "custom_styles" text NULL, "join_code" text NOT NULL, "name" text NOT NULL, "singleton_guard" integer NOT NULL UNIQUE);

CREATE TABLE "active_storage_blobs" ("id" integer NOT NULL PRIMARY KEY AUTOINCREMENT, "byte_size" bigint NOT NULL, "checksum" text NULL, "content_type" text NULL, "filename" text NOT NULL, "key" text NOT NULL UNIQUE, "metadata" text NULL, "service_name" text NOT NULL);

CREATE TABLE "ar_internal_metadata" ("key" text NOT NULL PRIMARY KEY, "value" text NULL);

CREATE TABLE "message_search_index_content" ("id" integer NOT NULL PRIMARY KEY AUTOINCREMENT);

CREATE TABLE "message_search_index_data" ("id" integer NOT NULL PRIMARY KEY AUTOINCREMENT, "block" BLOB NULL);

CREATE TABLE "message_search_index_docsize" ("id" integer NOT NULL PRIMARY KEY AUTOINCREMENT, "sz" BLOB NULL);

CREATE TABLE "rooms" ("id" integer NOT NULL PRIMARY KEY AUTOINCREMENT, "creator_id" bigint NOT NULL, "name" text NULL, "type" text NOT NULL, "messages_count" integer NOT NULL);

CREATE TABLE "schema_migrations" ("version" text NOT NULL PRIMARY KEY);

CREATE TABLE "users" ("id" integer NOT NULL PRIMARY KEY AUTOINCREMENT, "bio" text NULL, "bot_token" text NULL UNIQUE, "email_address" text NULL UNIQUE, "name" text NOT NULL, "password_digest" text NULL, "role" integer NOT NULL, "status" integer NOT NULL);

CREATE TABLE "action_text_rich_texts" ("id" integer NOT NULL PRIMARY KEY AUTOINCREMENT, "body" text NULL, "name" text NOT NULL, "record_id" bigint NOT NULL, "record_type" text NOT NULL, CONSTRAINT "index_action_text_rich_texts_uniqueness" UNIQUE ("record_type", "record_id", "name"));

CREATE TABLE "memberships" ("id" integer NOT NULL PRIMARY KEY AUTOINCREMENT, "connections" integer NOT NULL, "involvement" text NULL, "room_id" integer NOT NULL, "user_id" integer NOT NULL, CONSTRAINT "index_memberships_on_room_id_and_user_id" UNIQUE ("room_id", "user_id"));

CREATE TABLE "messages" ("id" integer NOT NULL PRIMARY KEY AUTOINCREMENT, "client_message_id" text NOT NULL, "room_id" integer NOT NULL REFERENCES "rooms" ("id") DEFERRABLE INITIALLY DEFERRED, "creator_id" integer NOT NULL REFERENCES "users" ("id") DEFERRABLE INITIALLY DEFERRED);

CREATE TABLE "bans" ("id" integer NOT NULL PRIMARY KEY AUTOINCREMENT, "ip_address" text NOT NULL, "user_id" integer NOT NULL REFERENCES "users" ("id") DEFERRABLE INITIALLY DEFERRED);

CREATE TABLE "webhooks" ("id" integer NOT NULL PRIMARY KEY AUTOINCREMENT, "url" text NULL, "user_id" integer NOT NULL REFERENCES "users" ("id") DEFERRABLE INITIALLY DEFERRED);

CREATE TABLE "active_storage_attachments" ("id" integer NOT NULL PRIMARY KEY AUTOINCREMENT, "name" text NOT NULL, "record_id" bigint NOT NULL, "record_type" text NOT NULL, "blob_id" integer NOT NULL REFERENCES "active_storage_blobs" ("id") DEFERRABLE INITIALLY DEFERRED, CONSTRAINT "index_active_storage_attachments_uniqueness" UNIQUE ("record_type", "record_id", "name", "blob_id"));

CREATE TABLE "active_storage_variant_records" ("id" integer NOT NULL PRIMARY KEY AUTOINCREMENT, "variation_digest" text NOT NULL, "blob_id" integer NOT NULL REFERENCES "active_storage_blobs" ("id") DEFERRABLE INITIALLY DEFERRED, CONSTRAINT "index_active_storage_variant_records_uniqueness" UNIQUE ("blob_id", "variation_digest"));

CREATE TABLE "boosts" ("id" integer NOT NULL PRIMARY KEY AUTOINCREMENT, "booster_id" integer NOT NULL, "content" text NOT NULL, "message_id" integer NOT NULL REFERENCES "messages" ("id") DEFERRABLE INITIALLY DEFERRED);

CREATE TABLE "sessions" ("id" integer NOT NULL PRIMARY KEY AUTOINCREMENT, "ip_address" text NULL, "token" text NOT NULL UNIQUE, "user_agent" text NULL, "user_id" integer NOT NULL REFERENCES "users" ("id") DEFERRABLE INITIALLY DEFERRED);

CREATE TABLE "searches" ("id" integer NOT NULL PRIMARY KEY AUTOINCREMENT, "query" text NOT NULL, "user_id" integer NOT NULL REFERENCES "users" ("id") DEFERRABLE INITIALLY DEFERRED);

CREATE TABLE "push_subscriptions" ("id" integer NOT NULL PRIMARY KEY AUTOINCREMENT, "auth_key" text NULL, "endpoint" text NULL, "p256dh_key" text NULL, "user_agent" text NULL, "user_id" integer NOT NULL REFERENCES "users" ("id") DEFERRABLE INITIALLY DEFERRED);

CREATE INDEX "index_messages_on_room_id" ON "messages" ("room_id");

CREATE INDEX "index_messages_on_creator_id" ON "messages" ("creator_id");

CREATE INDEX "index_bans_on_user_id" ON "bans" ("user_id");

CREATE INDEX "index_bans_on_ip_address" ON "bans" ("ip_address");

CREATE INDEX "index_webhooks_on_user_id" ON "webhooks" ("user_id");

CREATE INDEX "index_memberships_on_user_id" ON "memberships" ("user_id");

CREATE INDEX "index_memberships_on_room_id" ON "memberships" ("room_id");

CREATE INDEX "active_stor_blob_id_ca5c43_idx" ON "active_storage_attachments" ("blob_id");

CREATE INDEX "index_boosts_on_message_id" ON "boosts" ("message_id");

CREATE INDEX "index_boosts_on_booster_id" ON "boosts" ("booster_id");

CREATE INDEX "index_sessions_on_user_id" ON "sessions" ("user_id");

CREATE INDEX "index_searches_on_user_id" ON "searches" ("user_id");

CREATE INDEX "push_subscr_user_id_41d090_idx" ON "push_subscriptions" ("user_id");

CREATE INDEX "push_subscr_endpoin_acc3e3_idx" ON "push_subscriptions" ("endpoint", "p256dh_key", "auth_key");
