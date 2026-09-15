package route

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/herhe-com/framework/http/middleware"
	"github.com/tizips/shenhuo/server/admin/biz/site"
)

func SiteRouter(router *server.Hertz) {

	route := router.Group("site")
	route.Use(middleware.Auth())
	{
		permissions := route.Group("permissions")
		{
			permissions.GET("", site.ToPermissions)
		}

		roles := route.Group("roles")
		{
			roles.GET(":id", site.ToRoleByInformation)
			roles.GET("", middleware.Permission("site.role.paginate"), site.ToRoleByPaginate)
			roles.PUT(":id", middleware.Permission("site.role.update"), site.DoRoleByUpdate)
			roles.DELETE(":id", middleware.Permission("site.role.delete"), site.DoRoleByDelete)
		}

		role := route.Group("role")
		{
			role.POST("", middleware.Permission("site.role.create"), site.DoRoleByCreate)
			role.GET("opening", site.ToRoleByOpening)
		}

		users := route.Group("users")
		{
			users.GET(":id", site.ToRoleByInformation)
			users.GET("", middleware.Permission("site.user.paginate"), site.ToUserByPaginate)
			users.PUT(":id", middleware.Permission("site.user.update"), site.DoUserByUpdate)
			users.DELETE(":id", middleware.Permission("site.user.delete"), site.DoUserByDelete)
		}

		user := route.Group("user")
		{
			user.POST("", middleware.Permission("site.user.create"), site.DoUserByCreate)
			user.PUT("enable", middleware.Permission("site.user.enable"), site.DoUserByEnable)
		}
	}
}
