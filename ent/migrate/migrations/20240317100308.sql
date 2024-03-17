-- Modify "blogs" table
ALTER TABLE `blogs` ADD COLUMN `user_id` char(36) NOT NULL, ADD INDEX `blogs_users_user` (`user_id`), ADD CONSTRAINT `blogs_users_user` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON UPDATE NO ACTION ON DELETE NO ACTION;
