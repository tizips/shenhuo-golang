-- +goose Up
-- +goose StatementBegin
create table `sh_article`
(
    `id`           int unsigned     not null auto_increment,
    `title`        varchar(120)     not null default '' comment '标题',
    `thumb`        varchar(255)     not null default '' comment '缩略图',
    `content`      longtext         not null comment '富文本内容',
    `published_at` timestamp        not null default CURRENT_TIMESTAMP comment '发布时间',
    `is_top`       tinyint unsigned not null default 2 comment '是否置顶：1=是；2=否',
    `is_recommend` tinyint unsigned not null default 2 comment '是否首页推荐：1=是；2=否',
    `created_at`   timestamp        not null default CURRENT_TIMESTAMP,
    `updated_at`   timestamp        not null default CURRENT_TIMESTAMP,
    `deleted_at`   timestamp                 default null,
    primary key (`id`),
    key `idx_is_top` (`is_top`),
    key `idx_published_at` (`published_at`),
    key `idx_deleted_at` (`deleted_at`)
) collate = utf8mb4_unicode_ci comment ='神火-热点资讯';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists `sh_article`;
-- +goose StatementEnd
