-- +goose Up
-- +goose StatementBegin
alter table `sh_person`
    add column `id_card` varchar(18) not null default '' comment '身份证号' after `mobile`,
    add key `idx_id_card` (`id_card`);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
alter table `sh_person`
    drop key `idx_id_card`,
    drop column `id_card`;
-- +goose StatementEnd
