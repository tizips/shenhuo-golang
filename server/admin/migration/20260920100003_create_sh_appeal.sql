-- +goose Up
-- +goose StatementBegin
create table `sh_appeal`
(
    `id`         int unsigned not null auto_increment,
    `person_id`  varchar(64)  not null default '' comment '人员ID',
    `reason`     varchar(500) not null comment '申诉原因',
    `created_at` timestamp    not null default CURRENT_TIMESTAMP,
    `updated_at` timestamp    not null default CURRENT_TIMESTAMP,
    `deleted_at` timestamp             default null,
    primary key (`id`),
    key `idx_person_id` (`person_id`),
    key `idx_deleted_at` (`deleted_at`)
) collate = utf8mb4_unicode_ci comment ='神火-仲裁申诉';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists `sh_appeal`;
-- +goose StatementEnd
