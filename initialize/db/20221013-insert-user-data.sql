-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied

INSERT INTO `tb_sys_user` (`id`, `created_at`, `updated_at`, `deleted_at`, `username`, `password`, `mobile`, `avatar`,
                           `name`, `status`, `role_id`, `dept_id`, `last_login`, `locked`)
VALUES (1, '2022-10-13 11:01:40.848', '2022-10-13 11:01:40.848', NULL, 'admin',
        '$2a$10$rZ6DhktF36hPsubzSVzkDO/YBP8zPs09Ka565hrbKFgfaNiVz.vy.', '19999999999',
        'https://wpimg.wallstcn.com/f778738c-e4f8-4870-b634-56703b4acafe.gif', '连辰',
        1, 1, 1, NULL, 0);

