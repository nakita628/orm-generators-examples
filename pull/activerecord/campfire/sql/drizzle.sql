CREATE TABLE `accounts` (
	`id` integer PRIMARY KEY AUTOINCREMENT,
	`custom_styles` text,
	`join_code` text NOT NULL,
	`name` text NOT NULL,
	`singleton_guard` integer DEFAULT 0 NOT NULL
);

CREATE TABLE `action_text_rich_texts` (
	`id` integer PRIMARY KEY AUTOINCREMENT,
	`body` text,
	`name` text NOT NULL,
	`record_id` bigint NOT NULL,
	`record_type` text NOT NULL
);

CREATE TABLE `active_storage_attachments` (
	`id` integer PRIMARY KEY AUTOINCREMENT,
	`blob_id` integer NOT NULL,
	`name` text NOT NULL,
	`record_id` bigint NOT NULL,
	`record_type` text NOT NULL,
	CONSTRAINT `active_storage_attachments_blob_id_fkey` FOREIGN KEY (`blob_id`) REFERENCES `active_storage_blobs`(`id`)
);

CREATE TABLE `active_storage_blobs` (
	`id` integer PRIMARY KEY AUTOINCREMENT,
	`byte_size` bigint NOT NULL,
	`checksum` text,
	`content_type` text,
	`filename` text NOT NULL,
	`key` text NOT NULL,
	`metadata` text,
	`service_name` text NOT NULL
);

CREATE TABLE `active_storage_variant_records` (
	`id` integer PRIMARY KEY AUTOINCREMENT,
	`blob_id` integer NOT NULL,
	`variation_digest` text NOT NULL,
	CONSTRAINT `active_storage_variant_records_blob_id_fkey` FOREIGN KEY (`blob_id`) REFERENCES `active_storage_blobs`(`id`)
);

CREATE TABLE `ar_internal_metadata` (
	`key` text PRIMARY KEY,
	`value` text
);

CREATE TABLE `bans` (
	`id` integer PRIMARY KEY AUTOINCREMENT,
	`ip_address` text NOT NULL,
	`user_id` integer NOT NULL,
	CONSTRAINT `bans_user_id_fkey` FOREIGN KEY (`user_id`) REFERENCES `users`(`id`)
);

CREATE TABLE `boosts` (
	`id` integer PRIMARY KEY AUTOINCREMENT,
	`booster_id` integer NOT NULL,
	`content` text NOT NULL,
	`message_id` integer NOT NULL,
	CONSTRAINT `boosts_message_id_fkey` FOREIGN KEY (`message_id`) REFERENCES `messages`(`id`)
);

CREATE TABLE `memberships` (
	`id` integer PRIMARY KEY AUTOINCREMENT,
	`connections` integer DEFAULT 0 NOT NULL,
	`involvement` text DEFAULT 'mentions',
	`room_id` integer NOT NULL,
	`user_id` integer NOT NULL
);

CREATE TABLE `message_search_index_content` (
	`id` integer PRIMARY KEY AUTOINCREMENT
);

CREATE TABLE `message_search_index_data` (
	`id` integer PRIMARY KEY AUTOINCREMENT,
	`block` blob
);

CREATE TABLE `message_search_index_docsize` (
	`id` integer PRIMARY KEY AUTOINCREMENT,
	`sz` blob
);

CREATE TABLE `messages` (
	`id` integer PRIMARY KEY AUTOINCREMENT,
	`client_message_id` text NOT NULL,
	`creator_id` integer NOT NULL,
	`room_id` integer NOT NULL,
	CONSTRAINT `messages_creator_id_fkey` FOREIGN KEY (`creator_id`) REFERENCES `users`(`id`),
	CONSTRAINT `messages_room_id_fkey` FOREIGN KEY (`room_id`) REFERENCES `rooms`(`id`)
);

CREATE TABLE `push_subscriptions` (
	`id` integer PRIMARY KEY AUTOINCREMENT,
	`auth_key` text,
	`endpoint` text,
	`p256dh_key` text,
	`user_agent` text,
	`user_id` integer NOT NULL,
	CONSTRAINT `push_subscriptions_user_id_fkey` FOREIGN KEY (`user_id`) REFERENCES `users`(`id`)
);

CREATE TABLE `rooms` (
	`id` integer PRIMARY KEY AUTOINCREMENT,
	`creator_id` bigint NOT NULL,
	`name` text,
	`type` text NOT NULL,
	`messages_count` integer DEFAULT 0 NOT NULL
);

CREATE TABLE `schema_migrations` (
	`version` text PRIMARY KEY
);

CREATE TABLE `searches` (
	`id` integer PRIMARY KEY AUTOINCREMENT,
	`query` text NOT NULL,
	`user_id` integer NOT NULL,
	CONSTRAINT `searches_user_id_fkey` FOREIGN KEY (`user_id`) REFERENCES `users`(`id`)
);

CREATE TABLE `sessions` (
	`id` integer PRIMARY KEY AUTOINCREMENT,
	`ip_address` text,
	`token` text NOT NULL,
	`user_agent` text,
	`user_id` integer NOT NULL,
	CONSTRAINT `sessions_user_id_fkey` FOREIGN KEY (`user_id`) REFERENCES `users`(`id`)
);

CREATE TABLE `users` (
	`id` integer PRIMARY KEY AUTOINCREMENT,
	`bio` text,
	`bot_token` text,
	`email_address` text,
	`name` text NOT NULL,
	`password_digest` text,
	`role` integer DEFAULT 0 NOT NULL,
	`status` integer DEFAULT 0 NOT NULL
);

CREATE TABLE `webhooks` (
	`id` integer PRIMARY KEY AUTOINCREMENT,
	`url` text,
	`user_id` integer NOT NULL,
	CONSTRAINT `webhooks_user_id_fkey` FOREIGN KEY (`user_id`) REFERENCES `users`(`id`)
);

CREATE UNIQUE INDEX `index_accounts_on_singleton_guard` ON `accounts` (`singleton_guard`);

CREATE UNIQUE INDEX `index_action_text_rich_texts_uniqueness` ON `action_text_rich_texts` (`record_type`,`record_id`,`name`);

CREATE INDEX `index_active_storage_attachments_on_blob_id` ON `active_storage_attachments` (`blob_id`);

CREATE UNIQUE INDEX `index_active_storage_attachments_uniqueness` ON `active_storage_attachments` (`record_type`,`record_id`,`name`,`blob_id`);

CREATE UNIQUE INDEX `index_active_storage_blobs_on_key` ON `active_storage_blobs` (`key`);

CREATE UNIQUE INDEX `index_active_storage_variant_records_uniqueness` ON `active_storage_variant_records` (`blob_id`,`variation_digest`);

CREATE INDEX `index_bans_on_user_id` ON `bans` (`user_id`);

CREATE INDEX `index_bans_on_ip_address` ON `bans` (`ip_address`);

CREATE INDEX `index_boosts_on_message_id` ON `boosts` (`message_id`);

CREATE INDEX `index_boosts_on_booster_id` ON `boosts` (`booster_id`);

CREATE INDEX `index_memberships_on_user_id` ON `memberships` (`user_id`);

CREATE INDEX `index_memberships_on_room_id` ON `memberships` (`room_id`);

CREATE UNIQUE INDEX `index_memberships_on_room_id_and_user_id` ON `memberships` (`room_id`,`user_id`);

CREATE INDEX `index_messages_on_room_id` ON `messages` (`room_id`);

CREATE INDEX `index_messages_on_creator_id` ON `messages` (`creator_id`);

CREATE INDEX `index_push_subscriptions_on_user_id` ON `push_subscriptions` (`user_id`);

CREATE INDEX `idx_on_endpoint_p256dh_key_auth_key_7553014576` ON `push_subscriptions` (`endpoint`,`p256dh_key`,`auth_key`);

CREATE INDEX `index_searches_on_user_id` ON `searches` (`user_id`);

CREATE UNIQUE INDEX `index_sessions_on_token` ON `sessions` (`token`);

CREATE INDEX `index_sessions_on_user_id` ON `sessions` (`user_id`);

CREATE UNIQUE INDEX `index_users_on_bot_token` ON `users` (`bot_token`);

CREATE UNIQUE INDEX `index_users_on_email_address` ON `users` (`email_address`);

CREATE INDEX `index_webhooks_on_user_id` ON `webhooks` (`user_id`);
