-- Modify "blogs" table
ALTER TABLE `blogs` ADD COLUMN `create_time` timestamp NOT NULL, ADD COLUMN `update_time` timestamp NOT NULL;
-- Modify "projects" table
ALTER TABLE `projects` ADD COLUMN `create_time` timestamp NOT NULL, ADD COLUMN `update_time` timestamp NOT NULL;
-- Modify "users" table
ALTER TABLE `users` MODIFY COLUMN `id` char(36) NOT NULL, ADD COLUMN `create_time` timestamp NOT NULL, ADD COLUMN `update_time` timestamp NOT NULL;
