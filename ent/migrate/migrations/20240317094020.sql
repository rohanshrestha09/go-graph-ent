-- Modify "blogs" table
ALTER TABLE `blogs` ADD COLUMN `status` enum('PUBLISHED','UNPUBLISHED') NOT NULL DEFAULT "PUBLISHED";
