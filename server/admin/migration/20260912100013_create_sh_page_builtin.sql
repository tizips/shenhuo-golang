-- +goose Up
-- +goose StatementBegin
create table `sh_page_builtin`
(
    `id`         int unsigned not null auto_increment,
    `key`        varchar(64)  not null default '' comment '标识',
    `page_id`    int unsigned not null default 0 comment '页面ID',
    `created_at` timestamp    not null default CURRENT_TIMESTAMP,
    `updated_at` timestamp    not null default CURRENT_TIMESTAMP,
    `deleted_at` timestamp             default null,
    primary key (`id`),
    key `idx_key` (`key`),
    key `idx_page_id` (`page_id`),
    key `idx_deleted_at` (`deleted_at`)
) collate = utf8mb4_unicode_ci comment ='神火-内置页面';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists `sh_page_builtin`;
-- +goose StatementEnd
