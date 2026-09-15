-- +goose Up
-- +goose StatementBegin
create table `sh_schedule_category`
(
    `id`         int unsigned not null auto_increment,
    `name`       varchar(64)  not null default '' comment '名称',
    `order`      tinyint unsigned not null default 50 comment '序号',
    `created_at` timestamp    not null default CURRENT_TIMESTAMP,
    `updated_at` timestamp    not null default CURRENT_TIMESTAMP,
    `deleted_at` timestamp             default null,
    primary key (`id`),
    key `idx_deleted_at` (`deleted_at`)
) collate = utf8mb4_unicode_ci comment ='神火-日程分类';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists `sh_schedule_category`;
-- +goose StatementEnd
