package route

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/herhe-com/framework/http/middleware"
	"github.com/tizips/shenhuo/server/web/biz/basic"
	middle "github.com/tizips/shenhuo/server/web/http/middleware"
)

func BasicRouter(router *server.Hertz) {

	route := router.Group("basic")
	{
		login := route.Group("login")
		{
			login.POST("manager", middleware.LoginLimiter(), basic.DoLoginOfManager)
			login.POST("person", middleware.LoginLimiter(), basic.DoLoginOfPerson)
		}

		account := route.Group("account").Use(middleware.Auth())
		{
			account.GET("information", basic.ToAccountOfInformation)
			account.POST("logout", basic.DoLoginOfOut)
			account.PUT("password", middle.Person(), basic.DoAccountOfPassword)
		}

		upload := route.Group("upload").Use(middleware.Auth(), middle.Manager())
		{
			upload.POST("file", basic.DoUploadOfFile)
		}
	}
}
