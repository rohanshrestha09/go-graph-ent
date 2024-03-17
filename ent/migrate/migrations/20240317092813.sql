-- Modify "blogs" table
ALTER TABLE `blogs` ADD COLUMN `title` varchar(255) NOT NULL, ADD COLUMN `slug` varchar(255) NOT NULL, ADD COLUMN `content` longtext NOT NULL, ADD COLUMN `image` varchar(255) NULL, ADD UNIQUE INDEX `slug` (`slug`);
-- Modify "users" table
ALTER TABLE `users` DROP COLUMN `age`, MODIFY COLUMN `name` varchar(255) NOT NULL, ADD COLUMN `password` varchar(255) NOT NULL, ADD COLUMN `image` varchar(255) NULL;
-- Create "user_blogs" table
CREATE TABLE `user_blogs` (`user_id` char(36) NOT NULL, `blog_id` bigint NOT NULL, PRIMARY KEY (`user_id`, `blog_id`), INDEX `user_blogs_blog_id` (`blog_id`), CONSTRAINT `user_blogs_blog_id` FOREIGN KEY (`blog_id`) REFERENCES `blogs` (`id`) ON UPDATE NO ACTION ON DELETE CASCADE, CONSTRAINT `user_blogs_user_id` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON UPDATE NO ACTION ON DELETE CASCADE) CHARSET utf8mb4 COLLATE utf8mb4_bin;
-- Drop "projects" table
DROP TABLE `projects`;
