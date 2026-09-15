package main

import (
	"github.com/tizips/shenhuo/server/web/bootstrap"
)

// @title Web API
// @version 1.0
// @description 前台API文档
// @BasePath /
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

func main() {

	bootstrap.Boot()
}
