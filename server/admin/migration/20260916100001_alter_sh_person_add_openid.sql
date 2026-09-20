-- +goose Up
-- +goose StatementBegin
alter table `sh_person`
    add column `openid` varchar(64) not null default '' comment '微信公众号 OpenID' after `group_name`,
    add key `idx_openid` (`openid`);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
alter table `sh_person`
    drop key `idx_openid`,
    drop column `openid`;
-- +goose StatementEnd
