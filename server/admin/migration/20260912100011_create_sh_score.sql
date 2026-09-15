-- +goose Up
-- +goose StatementBegin
create table `sh_score`
(
    `id`         int unsigned not null auto_increment,
    `person_id`  varchar(64)  not null default '' comment '人员ID',
    `total`      varchar(32)  not null default '' comment '总分',
    `items`      json                  default null comment '小项成绩',
    `order`      tinyint unsigned not null default 50 comment '序号',
    `created_at` timestamp    not null default CURRENT_TIMESTAMP,
    `updated_at` timestamp    not null default CURRENT_TIMESTAMP,
    `deleted_at` timestamp             default null,
    primary key (`id`),
    unique key `uk_person_id` (`person_id`),
    key `idx_deleted_at` (`deleted_at`)
) collate = utf8mb4_unicode_ci comment ='神火-成绩';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists `sh_score`;
-- +goose StatementEnd
