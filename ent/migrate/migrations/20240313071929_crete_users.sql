-- Create "users" table
CREATE TABLE `users` (`id` bigint NOT NULL AUTO_INCREMENT, `age` bigint NOT NULL, `name` varchar(255) NOT NULL DEFAULT "unknown", PRIMARY KEY (`id`)) CHARSET utf8mb4 COLLATE utf8mb4_bin;
