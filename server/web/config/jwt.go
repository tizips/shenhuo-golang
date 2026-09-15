package config

import (
	"github.com/herhe-com/framework/facades"
)

func init() {

	cfg := facades.Config()
	cfg.Add("jwt", map[string]any{
		"secret":   cfg.Env("jwt.secret", ""),
		"sub":      cfg.Env("jwt.sub", "web"),
		"leeway":   cfg.Env("jwt.leeway", 3),
		"lifetime": cfg.Env("jwt.lifetime", 60*24),
		"refresh": map[string]any{
			"mode":     cfg.Env("jwt.refresh.mode", "blacklist"),
			"lifetime": cfg.Env("jwt.refresh.lifetime", 30),
			"leeway":   cfg.Env("jwt.refresh.leeway", 3),
		},
	})
}
