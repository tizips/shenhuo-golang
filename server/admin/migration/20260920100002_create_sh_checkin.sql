-- +goose Up
-- +goose StatementBegin
create table `sh_checkin`
(
    `id`         int unsigned not null auto_increment,
    `name`       varchar(32)  not null default '' comment '姓名',
    `created_at` timestamp    not null default CURRENT_TIMESTAMP,
    `updated_at` timestamp    not null default CURRENT_TIMESTAMP,
    `deleted_at` timestamp             default null,
    primary key (`id`),
    key `idx_name` (`name`),
    key `idx_created_at` (`created_at`),
    key `idx_deleted_at` (`deleted_at`)
) collate = utf8mb4_unicode_ci comment ='神火-签到';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists `sh_checkin`;
-- +goose StatementEnd
