-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied

CREATE TABLE IF NOT EXISTS `tb_sys_user` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT 'auto increment id',
  `created_at` datetime(3) DEFAULT NULL COMMENT 'create time',
  `updated_at` datetime(3) DEFAULT NULL COMMENT 'update time',
  `deleted_at` datetime(3) DEFAULT NULL COMMENT 'soft delete time',
  `username` varchar(191) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT 'user login name',
  `password` longtext COLLATE utf8mb4_general_ci COMMENT 'password',
  `mobile` longtext COLLATE utf8mb4_general_ci COMMENT 'mobile number',
  `avatar` longtext COLLATE utf8mb4_general_ci COMMENT 'avatar url',
  `name` longtext COLLATE utf8mb4_general_ci COMMENT 'name',
  `status` tinyint(1) DEFAULT '1' COMMENT 'status(0: disabled, 1: enable)',
  `dept_id` bigint unsigned DEFAULT NULL COMMENT 'department id',
  `role_id` bigint unsigned DEFAULT NULL COMMENT 'role id',
  `last_login` datetime(3) DEFAULT NULL COMMENT 'last login time',
  `locked` tinyint(1) DEFAULT '0' COMMENT 'locked(0: unlock, 1: locked)',
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_username` (`username`),
  KEY `idx_deleted_at` (`deleted_at`)
  ) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;