-- Modify "blogs" table
ALTER TABLE `blogs` DROP COLUMN `create_time`, DROP COLUMN `update_time`, ADD COLUMN `created_at` timestamp NOT NULL, ADD COLUMN `updated_at` timestamp NOT NULL;
-- Modify "projects" table
ALTER TABLE `projects` DROP COLUMN `create_time`, DROP COLUMN `update_time`, ADD COLUMN `created_at` timestamp NOT NULL, ADD COLUMN `updated_at` timestamp NOT NULL;
-- Modify "users" table
ALTER TABLE `users` DROP COLUMN `create_time`, DROP COLUMN `update_time`, ADD COLUMN `created_at` timestamp NOT NULL, ADD COLUMN `updated_at` timestamp NOT NULL;
