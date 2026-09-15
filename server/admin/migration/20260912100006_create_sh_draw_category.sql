-- +goose Up
-- +goose StatementBegin
create table `sh_draw_category`
(
    `id`         int unsigned not null auto_increment,
    `name`       varchar(64)  not null default '' comment '名称',
    `icon`       varchar(255) not null default '' comment '图标链接',
    `quota`      int unsigned not null default 0 comment '中签人数',
    `order`      tinyint unsigned not null default 50 comment '序号',
    `created_at` timestamp    not null default CURRENT_TIMESTAMP,
    `updated_at` timestamp    not null default CURRENT_TIMESTAMP,
    `deleted_at` timestamp             default null,
    primary key (`id`),
    key `idx_deleted_at` (`deleted_at`)
) collate = utf8mb4_unicode_ci comment ='神火-抽签类别';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists `sh_draw_category`;
-- +goose StatementEnd
