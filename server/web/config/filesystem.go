package config

import (
	"github.com/herhe-com/framework/facades"
	"github.com/herhe-com/framework/filesystem"
)

func init() {

	cfg := facades.Config()
	cfg.Add("filesystem", map[string]any{
		"default": cfg.Env("filesystem.default", "default"),
		"disks": map[string]any{
			"default": map[string]any{
				"driver":   cfg.Env("filesystem.disks.default.driver", filesystem.DriverQiniu),
				"access":   cfg.Env("filesystem.disks.default.access"),
				"secret":   cfg.Env("filesystem.disks.default.secret"),
				"region":   cfg.Env("filesystem.disks.default.region", "us-east-1"),
				"bucket":   cfg.Env("filesystem.disks.default.bucket"),
				"domain":   cfg.Env("filesystem.disks.default.domain"),
				"endpoint": cfg.Env("filesystem.disks.default.endpoint"),
				"prefix":   cfg.Env("filesystem.disks.default.prefix"),
				"ssl":      cfg.Env("filesystem.disks.default.ssl", false),
			},
		},
	})
}
