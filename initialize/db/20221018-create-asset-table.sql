-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied

CREATE TABLE IF NOT EXISTS `tb_sys_asset` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT 'auto increment id',
  `created_at` datetime(3) DEFAULT NULL COMMENT 'create time',
  `updated_at` datetime(3) DEFAULT NULL COMMENT 'update time',
  `deleted_at` datetime(3) DEFAULT NULL COMMENT 'soft delete time',
  `modified_user_id` bigint unsigned DEFAULT NULL COMMENT 'modified_user_id',
  `asset_id` varchar(100) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT 'asset id',
  `name` varchar(191) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT 'name',
  `type` varchar(80) COLLATE utf8mb4_general_ci COMMENT 'type',
  `child_type` varchar(80) COLLATE utf8mb4_general_ci COMMENT 'child_type',
  `project_name` longtext COLLATE utf8mb4_general_ci COMMENT 'project_name',
  `origin` longtext COLLATE utf8mb4_general_ci COMMENT 'origin',
  `model` longtext COLLATE utf8mb4_general_ci COMMENT 'model',
  `value` DECIMAL COLLATE utf8mb4_general_ci COMMENT 'value',
  `unit` varchar(80) COLLATE utf8mb4_general_ci COMMENT 'unit',
  `in_date` datetime(3) COLLATE utf8mb4_general_ci COMMENT 'in_date',
  `op_date` datetime(3) COLLATE utf8mb4_general_ci COMMENT 'op_date',
  `dep_year` tinyint(1) DEFAULT '0' COMMENT 'dep_year',
  `location` longtext COLLATE utf8mb4_general_ci COMMENT 'location',
  `image` longtext COLLATE utf8mb4_general_ci COMMENT 'image_url',
  `status` tinyint(1) DEFAULT '1' COMMENT 'status(0: disabled, 1: enable)',
  `assign_user_id` bigint unsigned DEFAULT NULL COMMENT 'assign_user_id',
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_asset_id` (`asset_id`),
  KEY `idx_deleted_at` (`deleted_at`)
  ) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;