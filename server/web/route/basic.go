package route

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/herhe-com/framework/http/middleware"
	"github.com/tizips/shenhuo/server/web/biz/basic"
)

func BasicRouter(router *server.Hertz) {

	route := router.Group("basic")
	{
		upload := route.Group("upload").Use(middleware.Auth())
		{
			upload.POST("file", basic.DoUploadOfFile)
		}
	}
}
