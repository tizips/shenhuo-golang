-- +goose Up
-- +goose StatementBegin
create table `sh_schedule`
(
    `id`          int unsigned not null auto_increment,
    `title`       varchar(120) not null default '' comment '标题',
    `subtitle`    varchar(120) not null default '' comment '副标题',
    `description` varchar(255) not null default '' comment '描述',
    `time`        varchar(64)  not null default '' comment '时间（自由文本）',
    `order`       tinyint unsigned not null default 50 comment '序号',
    `created_at`  timestamp    not null default CURRENT_TIMESTAMP,
    `updated_at`  timestamp    not null default CURRENT_TIMESTAMP,
    `deleted_at`  timestamp             default null,
    primary key (`id`),
    key `idx_deleted_at` (`deleted_at`)
) collate = utf8mb4_unicode_ci comment ='神火-日期安排';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists `sh_schedule`;
-- +goose StatementEnd
