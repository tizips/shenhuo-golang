-- +goose Up
-- +goose StatementBegin
create table `sh_page`
(
    `id`         int unsigned not null auto_increment,
    `title`      varchar(120) not null default '' comment '标题',
    `content`    longtext     not null comment '富文本内容',
    `created_at` timestamp    not null default CURRENT_TIMESTAMP,
    `updated_at` timestamp    not null default CURRENT_TIMESTAMP,
    `deleted_at` timestamp             default null,
    primary key (`id`),
    key `idx_deleted_at` (`deleted_at`)
) collate = utf8mb4_unicode_ci comment ='神火-页面';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists `sh_page`;
-- +goose StatementEnd
