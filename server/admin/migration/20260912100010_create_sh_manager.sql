-- +goose Up
-- +goose StatementBegin
create table `sh_manager`
(
    `id`         varchar(64)      not null,
    `name`       varchar(32)      not null default '' comment '姓名',
    `mobile`     varchar(20)      not null default '' comment '手机号',
    `password`   varchar(64)      not null default '' comment '密码',
    `is_enable`  tinyint unsigned not null default 0 comment '是否启用：1=是；2=否',
    `created_at` timestamp        not null default CURRENT_TIMESTAMP,
    `updated_at` timestamp        not null default CURRENT_TIMESTAMP,
    `deleted_at` timestamp                 default null,
    primary key (`id`),
    key `idx_mobile` (`mobile`),
    key `idx_deleted_at` (`deleted_at`)
) collate = utf8mb4_unicode_ci comment ='神火-小程序管理人员';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists `sh_manager`;
-- +goose StatementEnd
