package config

import (
	"github.com/herhe-com/framework/facades"
)

func init() {

	cfg := facades.Config()
	cfg.Add("wechat", map[string]any{
		"official_account": map[string]any{
			"app_id": cfg.Env("wechat.official_account.app_id", ""),
			"secret": cfg.Env("wechat.official_account.secret", ""),
			"oauth": map[string]any{
				"redirect": cfg.Env("wechat.official_account.oauth.redirect", ""),
				"scope":    cfg.Env("wechat.official_account.oauth.scope", "snsapi_base"),
			},
		},
		"templates": map[string]any{
			"score": map[string]any{
				"id":  cfg.Env("wechat.templates.score.id", ""),
				"url": cfg.Env("wechat.templates.score.url", ""),
				"keys": map[string]any{
					"name":   cfg.Env("wechat.templates.score.keys.name", ""),
					"number": cfg.Env("wechat.templates.score.keys.number", ""),
				},
			},
			"draw": map[string]any{
				"id":  cfg.Env("wechat.templates.draw.id", ""),
				"url": cfg.Env("wechat.templates.draw.url", ""),
				"keys": map[string]any{
					"name": cfg.Env("wechat.templates.draw.keys.name", ""),
					"unit": cfg.Env("wechat.templates.draw.keys.unit", ""),
				},
			},
		},
	})
}
