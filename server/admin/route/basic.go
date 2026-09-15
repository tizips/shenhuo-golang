package route

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/herhe-com/framework/http/middleware"
	"github.com/tizips/shenhuo/server/admin/biz/basic"
)

func BasicRouter(router *server.Hertz) {

	route := router.Group("basic")
	{

		login := route.Group("login")
		{
			login.POST("account", middleware.LoginLimiter(), basic.DoLoginOfAccount)
		}

		account := route.Group("account").Use(middleware.Auth())
		{
			account.GET("information", basic.ToAccountOfInformation)
			account.GET("modules", basic.ToAccountOfModules)
			account.GET("permissions", basic.ToAccountOfPermissions)
			account.POST("logout", basic.DoLoginOfOut)
			account.PUT("", basic.DoAccount)
		}

		upload := route.Group("upload").Use(middleware.Auth())
		{
			upload.POST("file", basic.DoUploadOfFile)
		}
	}
}
