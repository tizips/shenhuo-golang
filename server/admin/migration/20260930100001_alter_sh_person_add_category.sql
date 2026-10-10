-- +goose Up
-- +goose StatementBegin
alter table `sh_person`
    add column `category_id` int unsigned not null default 0 comment '所属抽签类别；小组直接绑定抽签类别' after `number`;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
alter table `sh_person`
    drop column `category_id`;
-- +goose StatementEnd
