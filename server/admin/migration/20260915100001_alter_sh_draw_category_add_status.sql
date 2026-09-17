-- +goose Up
-- +goose StatementBegin
alter table `sh_draw_category`
    add column `status` tinyint unsigned not null default 1 comment '抽签进度；枚举：1=未开始，2=已结束' after `quota`;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
alter table `sh_draw_category`
    drop column `status`;
-- +goose StatementEnd
