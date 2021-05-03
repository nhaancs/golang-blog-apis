CREATE TABLE `users` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `created_at` datetime DEFAULT (now()),
  `updated_at` datetime,
  `deleted_at` datetime,
  `is_enabled` boolean DEFAULT true,
  `first_name` varchar(200) NOT NULL,
  `last_name` varchar(200) NOT NULL,
  `avatar` json NOT NULL,
  `bio` text,
  `email` varchar(255) NOT NULL,
  `password` varchar(255) NOT NULL,
  `salt` varchar(255) NOT NULL,
  `token` varchar(255),
  `role` varchar(50) NOT NULL
);

CREATE TABLE `categories` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `created_at` datetime DEFAULT (now()),
  `updated_at` datetime,
  `deleted_at` datetime,
  `is_enabled` boolean DEFAULT true,
  `name` varchar(200) NOT NULL,
  `slug` varchar(255) UNIQUE NOT NULL,
  `desc` text
);

CREATE TABLE `posts` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `created_at` datetime DEFAULT (now()),
  `updated_at` datetime,
  `deleted_at` datetime,
  `is_enabled` boolean DEFAULT true,
  `title` varchar(200) NOT NULL,
  `slug` varchar(255) UNIQUE NOT NULL,
  `short_desc` varchar(255) NOT NULL,
  `body` text NOT NULL,
  `image` json NOT NULL,
  `published_at` datetime DEFAULT (now()),
  `keywords` varchar(255),
  `category_id` int,
  `user_id` int
);

CREATE TABLE `comments` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `created_at` datetime DEFAULT (now()),
  `updated_at` datetime,
  `deleted_at` datetime,
  `is_enabled` boolean DEFAULT true,
  `post_id` int NOT NULL,
  `user_id` int NOT NULL,
  `body` text NOT NULL
);

CREATE TABLE `favorites` (
  `created_at` datetime DEFAULT (now()),
  `user_id` int NOT NULL,
  `post_id` int NOT NULL,
  PRIMARY KEY (`user_id`, `post_id`)
);

ALTER TABLE `posts` ADD FOREIGN KEY (`category_id`) REFERENCES `categories` (`id`);

ALTER TABLE `posts` ADD FOREIGN KEY (`user_id`) REFERENCES `users` (`id`);

ALTER TABLE `comments` ADD FOREIGN KEY (`post_id`) REFERENCES `posts` (`id`);

ALTER TABLE `comments` ADD FOREIGN KEY (`user_id`) REFERENCES `users` (`id`);

ALTER TABLE `favorites` ADD FOREIGN KEY (`user_id`) REFERENCES `users` (`id`);

ALTER TABLE `favorites` ADD FOREIGN KEY (`post_id`) REFERENCES `posts` (`id`);

CREATE INDEX `users_index_0` ON `users` (`deleted_at`);

CREATE INDEX `users_index_1` ON `users` (`created_at`);

CREATE INDEX `users_index_2` ON `users` (`is_enabled`);

CREATE INDEX `users_index_3` ON `users` (`role`);

CREATE INDEX `users_index_4` ON `users` (`email`);

CREATE INDEX `categories_index_5` ON `categories` (`deleted_at`);

CREATE INDEX `categories_index_6` ON `categories` (`created_at`);

CREATE INDEX `categories_index_7` ON `categories` (`is_enabled`);

CREATE INDEX `categories_index_8` ON `categories` (`slug`);

CREATE INDEX `posts_index_9` ON `posts` (`deleted_at`);

CREATE INDEX `posts_index_10` ON `posts` (`created_at`);

CREATE INDEX `posts_index_11` ON `posts` (`is_enabled`);

CREATE INDEX `posts_index_12` ON `posts` (`slug`);

CREATE INDEX `posts_index_13` ON `posts` (`title`);

CREATE INDEX `posts_index_14` ON `posts` (`category_id`);

CREATE INDEX `comments_index_15` ON `comments` (`deleted_at`);

CREATE INDEX `comments_index_16` ON `comments` (`created_at`);

CREATE INDEX `comments_index_17` ON `comments` (`is_enabled`);

CREATE INDEX `comments_index_18` ON `comments` (`post_id`);

CREATE INDEX `favorites_index_19` ON `favorites` (`created_at`);

CREATE INDEX `favorites_index_20` ON `favorites` (`post_id`);
