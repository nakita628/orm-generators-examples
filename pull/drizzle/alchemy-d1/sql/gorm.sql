CREATE TABLE `users` (`id` integer PRIMARY KEY AUTOINCREMENT,`email` text NOT NULL,`name` text NOT NULL,`created_at` integer NOT NULL DEFAULT (unixepoch()));

CREATE UNIQUE INDEX `users_email_key` ON `users`(`email`);

CREATE TABLE `posts` (`id` integer PRIMARY KEY AUTOINCREMENT,`user_id` integer NOT NULL,`title` text NOT NULL,`body` text NOT NULL,`created_at` integer NOT NULL DEFAULT (unixepoch()),CONSTRAINT `posts_user_id_fkey` FOREIGN KEY (`user_id`) REFERENCES `users`(`id`) ON DELETE CASCADE ON UPDATE NO ACTION);
