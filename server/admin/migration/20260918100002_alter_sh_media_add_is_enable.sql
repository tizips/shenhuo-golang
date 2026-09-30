-- +goose Up
-- +goose StatementBegin
alter table `sh_media`
    add column `is_enable` tinyint unsigned not null default 1 comment '是否启用：1=是；2=否' after `is_top`;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
alter table `sh_media`
    drop column `is_enable`;
-- +goose StatementEnd
