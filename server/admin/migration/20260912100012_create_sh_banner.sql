-- +goose Up
-- +goose StatementBegin
create table `sh_banner`
(
    `id`         int unsigned     not null auto_increment,
    `title`      varchar(64)      not null default '' comment '标题',
    `image`      varchar(255)     not null default '' comment '图片链接',
    `link`       varchar(255)     not null default '' comment '跳转链接',
    `order`      tinyint unsigned not null default 50 comment '序号',
    `created_at` timestamp        not null default CURRENT_TIMESTAMP,
    `updated_at` timestamp        not null default CURRENT_TIMESTAMP,
    `deleted_at` timestamp                 default null,
    primary key (`id`),
    key `idx_deleted_at` (`deleted_at`)
) collate = utf8mb4_unicode_ci comment ='神火-轮播';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists `sh_banner`;
-- +goose StatementEnd
