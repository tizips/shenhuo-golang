-- +goose Up
-- +goose StatementBegin
alter table `sh_schedule`
    add column `category_id` int unsigned not null default 0 comment '日程分类ID' after `id`,
    add column `items` json default null comment '子项目（名称/时间）' after `time`;
-- +goose StatementEnd

-- +goose StatementBegin
insert into `sh_schedule_category` (`name`, `order`) values ('赛事安排', 50);
-- +goose StatementEnd

-- +goose StatementBegin
update `sh_schedule` set `category_id` = last_insert_id();
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
alter table `sh_schedule`
    drop column `category_id`,
    drop column `items`;
-- +goose StatementEnd
