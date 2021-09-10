DROP DATABASE IF EXISTS `clean-architecture-sample`;
CREATE DATABASE `clean-architecture-sample`;
USE `clean-architecture-sample`;

CREATE TABLE IF NOT EXISTS `users` (
  `id` char(36) PRIMARY KEY NOT NULL,
  `name` varchar(32) NOT NULL UNIQUE,
  `email` varchar(32) NOT NULL,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- sample data
INSERT INTO `users` (`id`, `name`, `email`) VALUES (UUID(), 'John Doe', 'jonny@example.com');
INSERT INTO `users` (`id`, `name`, `email`) VALUES (UUID(), 'Mike Smith', 'mike@example.com');
