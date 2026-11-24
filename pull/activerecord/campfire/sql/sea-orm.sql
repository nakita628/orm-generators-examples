CREATE TABLE "accounts" ( "id" integer NOT NULL PRIMARY KEY AUTOINCREMENT, "custom_styles" varchar, "join_code" varchar NOT NULL, "name" varchar NOT NULL, "singleton_guard" integer NOT NULL UNIQUE DEFAULT 0 );

CREATE TABLE "action_text_rich_texts" ( "id" integer NOT NULL PRIMARY KEY AUTOINCREMENT, "body" varchar, "name" varchar NOT NULL, "record_id" bigint NOT NULL, "record_type" varchar NOT NULL );

CREATE TABLE "active_storage_blobs" ( "id" integer NOT NULL PRIMARY KEY AUTOINCREMENT, "byte_size" bigint NOT NULL, "checksum" varchar, "content_type" varchar, "filename" varchar NOT NULL, "key" varchar NOT NULL UNIQUE, "metadata" varchar, "service_name" varchar NOT NULL );

CREATE TABLE "active_storage_attachments" ( "id" integer NOT NULL PRIMARY KEY AUTOINCREMENT, "blob_id" integer NOT NULL, "name" varchar NOT NULL, "record_id" bigint NOT NULL, "record_type" varchar NOT NULL, FOREIGN KEY ("blob_id") REFERENCES "active_storage_blobs" ("id") ON DELETE NO ACTION ON UPDATE NO ACTION );

CREATE TABLE "active_storage_variant_records" ( "id" integer NOT NULL PRIMARY KEY AUTOINCREMENT, "blob_id" integer NOT NULL, "variation_digest" varchar NOT NULL, FOREIGN KEY ("blob_id") REFERENCES "active_storage_blobs" ("id") ON DELETE NO ACTION ON UPDATE NO ACTION );

CREATE TABLE "ar_internal_metadata" ( "key" varchar NOT NULL PRIMARY KEY, "value" varchar );

CREATE TABLE "users" ( "id" integer NOT NULL PRIMARY KEY AUTOINCREMENT, "bio" varchar, "bot_token" varchar UNIQUE, "email_address" varchar UNIQUE, "name" varchar NOT NULL, "password_digest" varchar, "role" integer NOT NULL DEFAULT 0, "status" integer NOT NULL DEFAULT 0 );

CREATE TABLE "bans" ( "id" integer NOT NULL PRIMARY KEY AUTOINCREMENT, "ip_address" varchar NOT NULL, "user_id" integer NOT NULL, FOREIGN KEY ("user_id") REFERENCES "users" ("id") ON DELETE NO ACTION ON UPDATE NO ACTION );

CREATE TABLE "rooms" ( "id" integer NOT NULL PRIMARY KEY AUTOINCREMENT, "creator_id" bigint NOT NULL, "name" varchar, "type" varchar NOT NULL, "messages_count" integer NOT NULL DEFAULT 0 );

CREATE TABLE "messages" ( "id" integer NOT NULL PRIMARY KEY AUTOINCREMENT, "client_message_id" varchar NOT NULL, "creator_id" integer NOT NULL, "room_id" integer NOT NULL, FOREIGN KEY ("creator_id") REFERENCES "users" ("id") ON DELETE NO ACTION ON UPDATE NO ACTION, FOREIGN KEY ("room_id") REFERENCES "rooms" ("id") ON DELETE NO ACTION ON UPDATE NO ACTION );

CREATE TABLE "boosts" ( "id" integer NOT NULL PRIMARY KEY AUTOINCREMENT, "booster_id" integer NOT NULL, "content" varchar NOT NULL, "message_id" integer NOT NULL, FOREIGN KEY ("message_id") REFERENCES "messages" ("id") ON DELETE NO ACTION ON UPDATE NO ACTION );

CREATE TABLE "memberships" ( "id" integer NOT NULL PRIMARY KEY AUTOINCREMENT, "connections" integer NOT NULL DEFAULT 0, "involvement" varchar DEFAULT 'mentions', "room_id" integer NOT NULL, "user_id" integer NOT NULL );

CREATE TABLE "message_search_index_content" ( "id" integer NOT NULL PRIMARY KEY AUTOINCREMENT );

CREATE TABLE "message_search_index_data" ( "id" integer NOT NULL PRIMARY KEY AUTOINCREMENT, "block" varbinary_blob );

CREATE TABLE "message_search_index_docsize" ( "id" integer NOT NULL PRIMARY KEY AUTOINCREMENT, "sz" varbinary_blob );

CREATE TABLE "push_subscriptions" ( "id" integer NOT NULL PRIMARY KEY AUTOINCREMENT, "auth_key" varchar, "endpoint" varchar, "p256dh_key" varchar, "user_agent" varchar, "user_id" integer NOT NULL, FOREIGN KEY ("user_id") REFERENCES "users" ("id") ON DELETE NO ACTION ON UPDATE NO ACTION );

CREATE TABLE "schema_migrations" ( "version" varchar NOT NULL PRIMARY KEY );

CREATE TABLE "searches" ( "id" integer NOT NULL PRIMARY KEY AUTOINCREMENT, "query" varchar NOT NULL, "user_id" integer NOT NULL, FOREIGN KEY ("user_id") REFERENCES "users" ("id") ON DELETE NO ACTION ON UPDATE NO ACTION );

CREATE TABLE "sessions" ( "id" integer NOT NULL PRIMARY KEY AUTOINCREMENT, "ip_address" varchar, "token" varchar NOT NULL UNIQUE, "user_agent" varchar, "user_id" integer NOT NULL, FOREIGN KEY ("user_id") REFERENCES "users" ("id") ON DELETE NO ACTION ON UPDATE NO ACTION );

CREATE TABLE "webhooks" ( "id" integer NOT NULL PRIMARY KEY AUTOINCREMENT, "url" varchar, "user_id" integer NOT NULL, FOREIGN KEY ("user_id") REFERENCES "users" ("id") ON DELETE NO ACTION ON UPDATE NO ACTION );
