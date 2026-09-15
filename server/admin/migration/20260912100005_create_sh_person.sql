-- +goose Up
-- +goose StatementBegin
create table `sh_person`
(
    `id`                   varchar(64)      not null,
    `name`                 varchar(32)      not null default '' comment '姓名',
    `unit`                 varchar(64)      not null default '' comment '单位',
    `mobile`               varchar(20)      not null default '' comment '手机号',
    `password`             varchar(64)      not null default '' comment '密码',
    `number`               varchar(32)      not null default '' comment '参赛号',
    `group_name`           varchar(64)      not null default '' comment '小组名称',
    `must_change_password` tinyint unsigned not null default 2 comment '是否首次登录改密：1=是；2=否',
    `created_at`           timestamp        not null default CURRENT_TIMESTAMP,
    `updated_at`           timestamp        not null default CURRENT_TIMESTAMP,
    `deleted_at`           timestamp                 default null,
    primary key (`id`),
    key `idx_mobile` (`mobile`),
    key `idx_number` (`number`),
    key `idx_deleted_at` (`deleted_at`)
) collate = utf8mb4_unicode_ci comment ='神火-人员';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists `sh_person`;
-- +goose StatementEnd
