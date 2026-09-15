package config

import (
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/herhe-com/framework/facades"
	"github.com/herhe-com/framework/http/middleware"
	"github.com/tizips/shenhuo/server/web/route"
)

func init() {

	cfg := facades.Config()
	cfg.Add("server", map[string]any{
		"name":    cfg.Env("server.name", "UPER"),
		"address": cfg.Env("server.address", "0.0.0.0"),
		"port":    cfg.Env("server.port", "9600"),
		"route":   route.Router,
		//"handle":  func(server *server.Hertz) {},
		//"options": []config.Option{},
		"middlewares": []app.HandlerFunc{
			middleware.Access(),
		},
	})
}
