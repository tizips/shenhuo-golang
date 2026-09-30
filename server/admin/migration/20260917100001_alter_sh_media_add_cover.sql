-- +goose Up
-- +goose StatementBegin
alter table `sh_media`
    add column `cover` varchar(255) not null default '' comment '视频封面' after `url`;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
alter table `sh_media`
    drop column `cover`;
-- +goose StatementEnd
