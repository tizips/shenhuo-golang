package main

import (
	"github.com/tizips/shenhuo/server/admin/bootstrap"
)

// @title Admin API
// @version 1.0
// @description 管理后台API文档
// @BasePath /
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

func main() {

	bootstrap.Boot()
}
