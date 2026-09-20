package config

import (
	"github.com/herhe-com/framework/facades"
)

func init() {

	cfg := facades.Config()
	cfg.Add("system", map[string]any{
		"video": cfg.Env("system.video", true),
	})
}
