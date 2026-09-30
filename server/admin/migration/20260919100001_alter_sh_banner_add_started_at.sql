-- +goose Up
-- +goose StatementBegin
alter table `sh_banner`
    add column `started_at` timestamp null default null comment '生效开始时间' after `order`,
    add column `ended_at`   timestamp null default null comment '生效结束时间' after `started_at`;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
alter table `sh_banner`
    drop column `started_at`,
    drop column `ended_at`;
-- +goose StatementEnd
