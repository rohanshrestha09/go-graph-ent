-- Modify "blogs" table
ALTER TABLE `blogs` ADD COLUMN `user_blogs` char(36) NULL, ADD INDEX `blogs_users_blogs` (`user_blogs`), ADD CONSTRAINT `blogs_users_blogs` FOREIGN KEY (`user_blogs`) REFERENCES `users` (`id`) ON UPDATE NO ACTION ON DELETE SET NULL;
-- Drop "user_blogs" table
DROP TABLE `user_blogs`;
