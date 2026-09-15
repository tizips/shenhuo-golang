-- +goose Up
-- +goose StatementBegin
create table `sh_draw`
(
    `id`          int unsigned not null auto_increment,
    `category_id` int unsigned not null default 0 comment '抽签类别',
    `person_id`   varchar(64)  not null default '' comment '人员ID',
    `created_at`  timestamp    not null default CURRENT_TIMESTAMP,
    `updated_at`  timestamp    not null default CURRENT_TIMESTAMP,
    primary key (`id`),
    unique key `uk_category_person` (`category_id`, `person_id`),
    key `idx_person_id` (`person_id`)
) collate = utf8mb4_unicode_ci comment ='神火-抽签结果';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists `sh_draw`;
-- +goose StatementEnd
