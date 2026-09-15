-- +goose Up
-- +goose StatementBegin
create table `sh_media`
(
    `id`         int unsigned not null auto_increment,
    `scene_id`   int unsigned not null default 0 comment '场景ID',
    `type`       varchar(16)  not null default '' comment '类型：image=图片；video=视频',
    `title`      varchar(128) not null default '' comment '标题',
    `url`        varchar(255) not null default '' comment '链接',
    `is_top`     tinyint unsigned not null default 2 comment '是否置顶：1=是；2=否',
    `created_at` timestamp    not null default CURRENT_TIMESTAMP,
    `updated_at` timestamp    not null default CURRENT_TIMESTAMP,
    `deleted_at` timestamp             default null,
    primary key (`id`),
    key `idx_scene_id` (`scene_id`),
    key `idx_is_top` (`is_top`),
    key `idx_deleted_at` (`deleted_at`)
) collate = utf8mb4_unicode_ci comment ='神火-精彩媒体';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists `sh_media`;
-- +goose StatementEnd
