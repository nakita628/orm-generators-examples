CREATE TABLE `accounts` (`id` integer PRIMARY KEY AUTOINCREMENT,`custom_styles` text,`join_code` text NOT NULL,`name` text NOT NULL,`singleton_guard` integer NOT NULL DEFAULT 0);

CREATE UNIQUE INDEX `idx_accounts_singleton_guard` ON `accounts`(`singleton_guard`);

CREATE TABLE `action_text_rich_texts` (`id` integer PRIMARY KEY AUTOINCREMENT,`body` text,`name` text NOT NULL,`record_id` integer NOT NULL,`record_type` text NOT NULL);

CREATE UNIQUE INDEX `action_text_rich_texts_record_type_record_id_name_key` ON `action_text_rich_texts`(`name`,`record_id`,`record_type`);

CREATE TABLE `active_storage_blobs` (`id` integer PRIMARY KEY AUTOINCREMENT,`byte_size` integer NOT NULL,`checksum` text,`content_type` text,`filename` text NOT NULL,`key` text NOT NULL,`metadata` text,`service_name` text NOT NULL);

CREATE UNIQUE INDEX `idx_active_storage_blobs_key` ON `active_storage_blobs`(`key`);

CREATE TABLE `active_storage_attachments` (`id` integer PRIMARY KEY AUTOINCREMENT,`blob_id` integer NOT NULL,`name` text NOT NULL,`record_id` integer NOT NULL,`record_type` text NOT NULL,CONSTRAINT `fk_active_storage_blobs_active_storage_attachments` FOREIGN KEY (`blob_id`) REFERENCES `active_storage_blobs`(`id`) ON DELETE NO ACTION ON UPDATE NO ACTION);

CREATE INDEX `index_active_storage_attachments_on_blob_id` ON `active_storage_attachments`(`blob_id`);

CREATE UNIQUE INDEX `active_storage_attachments_record_type_record_id_name_blob_id_key` ON `active_storage_attachments`(`blob_id`,`name`,`record_id`,`record_type`);

CREATE TABLE `active_storage_variant_records` (`id` integer PRIMARY KEY AUTOINCREMENT,`blob_id` integer NOT NULL,`variation_digest` text NOT NULL,CONSTRAINT `fk_active_storage_blobs_active_storage_variant_records` FOREIGN KEY (`blob_id`) REFERENCES `active_storage_blobs`(`id`) ON DELETE NO ACTION ON UPDATE NO ACTION);

CREATE UNIQUE INDEX `active_storage_variant_records_blob_id_variation_digest_key` ON `active_storage_variant_records`(`blob_id`,`variation_digest`);

CREATE TABLE `ar_internal_metadata` (`key` text,`value` text,PRIMARY KEY (`key`));

CREATE TABLE `users` (`id` integer PRIMARY KEY AUTOINCREMENT,`bio` text,`bot_token` text,`email_address` text,`name` text NOT NULL,`password_digest` text,`role` integer NOT NULL DEFAULT 0,`status` integer NOT NULL DEFAULT 0);

CREATE UNIQUE INDEX `idx_users_email_address` ON `users`(`email_address`);

CREATE UNIQUE INDEX `idx_users_bot_token` ON `users`(`bot_token`);

CREATE TABLE `bans` (`id` integer PRIMARY KEY AUTOINCREMENT,`ip_address` text NOT NULL,`user_id` integer NOT NULL,CONSTRAINT `fk_users_bans` FOREIGN KEY (`user_id`) REFERENCES `users`(`id`) ON DELETE NO ACTION ON UPDATE NO ACTION);

CREATE INDEX `index_bans_on_user_id` ON `bans`(`user_id`);

CREATE INDEX `index_bans_on_ip_address` ON `bans`(`ip_address`);

CREATE TABLE `rooms` (`id` integer PRIMARY KEY AUTOINCREMENT,`creator_id` integer NOT NULL,`name` text,`type` text NOT NULL,`messages_count` integer NOT NULL DEFAULT 0);

CREATE TABLE `messages` (`id` integer PRIMARY KEY AUTOINCREMENT,`client_message_id` text NOT NULL,`creator_id` integer NOT NULL,`room_id` integer NOT NULL,CONSTRAINT `fk_users_messages` FOREIGN KEY (`creator_id`) REFERENCES `users`(`id`) ON DELETE NO ACTION ON UPDATE NO ACTION,CONSTRAINT `fk_rooms_messages` FOREIGN KEY (`room_id`) REFERENCES `rooms`(`id`) ON DELETE NO ACTION ON UPDATE NO ACTION);

CREATE INDEX `index_messages_on_room_id_and_created_at` ON `messages`(`room_id`);

CREATE INDEX `index_messages_on_room_id_and_updated_at` ON `messages`(`room_id`);

CREATE INDEX `index_messages_on_room_id` ON `messages`(`room_id`);

CREATE INDEX `index_messages_on_creator_id` ON `messages`(`creator_id`);

CREATE TABLE `boosts` (`id` integer PRIMARY KEY AUTOINCREMENT,`booster_id` integer NOT NULL,`content` text NOT NULL,`message_id` integer NOT NULL,CONSTRAINT `fk_messages_boosts` FOREIGN KEY (`message_id`) REFERENCES `messages`(`id`) ON DELETE NO ACTION ON UPDATE NO ACTION);

CREATE INDEX `index_boosts_on_message_id` ON `boosts`(`message_id`);

CREATE INDEX `index_boosts_on_booster_id` ON `boosts`(`booster_id`);

CREATE TABLE `memberships` (`id` integer PRIMARY KEY AUTOINCREMENT,`connections` integer NOT NULL DEFAULT 0,`involvement` text DEFAULT "mentions",`room_id` integer NOT NULL,`user_id` integer NOT NULL);

CREATE INDEX `index_memberships_on_user_id` ON `memberships`(`user_id`);

CREATE INDEX `index_memberships_on_room_id_and_created_at` ON `memberships`(`room_id`);

CREATE INDEX `index_memberships_on_room_id` ON `memberships`(`room_id`);

CREATE UNIQUE INDEX `memberships_room_id_user_id_key` ON `memberships`(`room_id`,`user_id`);

CREATE TABLE `message_search_index_content` (`id` integer PRIMARY KEY AUTOINCREMENT);

CREATE TABLE `message_search_index_data` (`id` integer PRIMARY KEY AUTOINCREMENT,`block` blob);

CREATE TABLE `message_search_index_docsize` (`id` integer PRIMARY KEY AUTOINCREMENT,`sz` blob);

CREATE TABLE `push_subscriptions` (`id` integer PRIMARY KEY AUTOINCREMENT,`auth_key` text,`endpoint` text,`p256dh_key` text,`user_agent` text,`user_id` integer NOT NULL,CONSTRAINT `fk_users_push_subscriptions` FOREIGN KEY (`user_id`) REFERENCES `users`(`id`) ON DELETE NO ACTION ON UPDATE NO ACTION);

CREATE INDEX `index_push_subscriptions_on_user_id` ON `push_subscriptions`(`user_id`);

CREATE INDEX `idx_on_endpoint_p256dh_key_auth_key_7553014576` ON `push_subscriptions`(`auth_key`,`endpoint`,`p256dh_key`);

CREATE TABLE `schema_migrations` (`version` text,PRIMARY KEY (`version`));

CREATE TABLE `searches` (`id` integer PRIMARY KEY AUTOINCREMENT,`query` text NOT NULL,`user_id` integer NOT NULL,CONSTRAINT `fk_users_searches` FOREIGN KEY (`user_id`) REFERENCES `users`(`id`) ON DELETE NO ACTION ON UPDATE NO ACTION);

CREATE INDEX `index_searches_on_user_id` ON `searches`(`user_id`);

CREATE TABLE `sessions` (`id` integer PRIMARY KEY AUTOINCREMENT,`ip_address` text,`token` text NOT NULL,`user_agent` text,`user_id` integer NOT NULL,CONSTRAINT `fk_users_sessions` FOREIGN KEY (`user_id`) REFERENCES `users`(`id`) ON DELETE NO ACTION ON UPDATE NO ACTION);

CREATE INDEX `index_sessions_on_user_id` ON `sessions`(`user_id`);

CREATE UNIQUE INDEX `idx_sessions_token` ON `sessions`(`token`);

CREATE TABLE `webhooks` (`id` integer PRIMARY KEY AUTOINCREMENT,`url` text,`user_id` integer NOT NULL,CONSTRAINT `fk_users_webhooks` FOREIGN KEY (`user_id`) REFERENCES `users`(`id`) ON DELETE NO ACTION ON UPDATE NO ACTION);

CREATE INDEX `index_webhooks_on_user_id` ON `webhooks`(`user_id`);
