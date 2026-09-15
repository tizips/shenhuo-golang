-- +goose Up
-- +goose StatementBegin
create table `sh_nav`
(
    `id`         int unsigned not null auto_increment,
    `title`      varchar(64)  not null default '' comment '标题',
    `icon`       varchar(255) not null default '' comment '图标链接',
    `type`       varchar(16)  not null default '' comment '类型：link=链接；page=页面',
    `value`      varchar(255) not null default '' comment '链接或页面ID',
    `order`      tinyint unsigned not null default 50 comment '序号',
    `created_at` timestamp    not null default CURRENT_TIMESTAMP,
    `updated_at` timestamp    not null default CURRENT_TIMESTAMP,
    `deleted_at` timestamp             default null,
    primary key (`id`),
    key `idx_deleted_at` (`deleted_at`)
) collate = utf8mb4_unicode_ci comment ='神火-导航';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists `sh_nav`;
-- +goose StatementEnd
