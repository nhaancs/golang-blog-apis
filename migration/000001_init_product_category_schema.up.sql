CREATE TABLE `products` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `created_at` datetime DEFAULT (now()),
  `updated_at` datetime,
  `deleted_at` datetime,
  `is_enabled` boolean DEFAULT true,
  `name` varchar(200) NOT NULL,
  `slug` varchar(255) UNIQUE NOT NULL,
  `short_desc` varchar(255) NOT NULL,
  `long_desc` text,
  `images` json NOT NULL,
  `unit_key` varchar(50) NOT NULL,
  `unit_name` varchar(255) NOT NULL,
  `price` decimal(13,2) DEFAULT 0,
  `quantity` decimal(13,2) DEFAULT 0,
  `is_unlimited` boolean DEFAULT false
);

CREATE INDEX `products_index_0` ON `products` (`deleted_at`);

CREATE INDEX `products_index_1` ON `products` (`created_at`);

CREATE INDEX `products_index_2` ON `products` (`is_enabled`);

CREATE INDEX `products_index_3` ON `products` (`name`);

CREATE INDEX `products_index_4` ON `products` (`slug`);
