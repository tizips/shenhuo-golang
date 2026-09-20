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

		articles := route.Group("articles")
		{
			articles.GET(":id", site.ToArticleOfInformation)
			articles.GET("", middleware.Permission("site.article.paginate"), site.ToArticleOfPaginate)
			articles.PUT(":id", middleware.Permission("site.article.update"), site.DoArticleOfUpdate)
			articles.DELETE(":id", middleware.Permission("site.article.delete"), site.DoArticleOfDelete)
		}

		article := route.Group("article")
		{
			article.POST("", middleware.Permission("site.article.create"), site.DoArticleOfCreate)
		}

		pages := route.Group("pages")
		{
			pages.GET(":id", site.ToPageOfInformation)
			pages.GET("", middleware.Permission("site.page.paginate"), site.ToPageOfPaginate)
			pages.PUT(":id", middleware.Permission("site.page.update"), site.DoPageOfUpdate)
			pages.DELETE(":id", middleware.Permission("site.page.delete"), site.DoPageOfDelete)
		}

		page := route.Group("page")
		{
			page.POST("", middleware.Permission("site.page.create"), site.DoPageOfCreate)
			page.GET("opening", site.ToPageOfOpening)
		}

		pageBuiltins := route.Group("page-builtins")
		{
			pageBuiltins.GET("", middleware.Permission("site.page_builtin.paginate"), site.ToPageBuiltinOfPaginate)
			pageBuiltins.PUT(":id", middleware.Permission("site.page_builtin.update"), site.DoPageBuiltinOfUpdate)
			pageBuiltins.DELETE(":id", middleware.Permission("site.page_builtin.delete"), site.DoPageBuiltinOfDelete)
		}

		pageBuiltin := route.Group("page-builtin")
		{
			pageBuiltin.POST("", middleware.Permission("site.page_builtin.create"), site.DoPageBuiltinOfCreate)
		}

		navs := route.Group("navs")
		{
			navs.GET("", middleware.Permission("site.nav.paginate"), site.ToNavOfPaginate)
			navs.PUT(":id", middleware.Permission("site.nav.update"), site.DoNavOfUpdate)
			navs.DELETE(":id", middleware.Permission("site.nav.delete"), site.DoNavOfDelete)
		}

		nav := route.Group("nav")
		{
			nav.POST("", middleware.Permission("site.nav.create"), site.DoNavOfCreate)
		}

		banners := route.Group("banners")
		{
			banners.GET("", middleware.Permission("site.banner.paginate"), site.ToBannerOfPaginate)
			banners.PUT(":id", middleware.Permission("site.banner.update"), site.DoBannerOfUpdate)
			banners.DELETE(":id", middleware.Permission("site.banner.delete"), site.DoBannerOfDelete)
		}

		banner := route.Group("banner")
		{
			banner.POST("", middleware.Permission("site.banner.create"), site.DoBannerOfCreate)
		}

		schedules := route.Group("schedules")
		{
			schedules.GET("", middleware.Permission("site.schedule.paginate"), site.ToScheduleOfPaginate)
			schedules.PUT(":id", middleware.Permission("site.schedule.update"), site.DoScheduleOfUpdate)
			schedules.DELETE(":id", middleware.Permission("site.schedule.delete"), site.DoScheduleOfDelete)
		}

		schedule := route.Group("schedule")
		{
			schedule.POST("", middleware.Permission("site.schedule.create"), site.DoScheduleOfCreate)
		}

		scheduleCategories := route.Group("schedule-categories")
		{
			scheduleCategories.GET("", middleware.Permission("site.schedule_category.paginate"), site.ToScheduleCategoryOfPaginate)
			scheduleCategories.PUT(":id", middleware.Permission("site.schedule_category.update"), site.DoScheduleCategoryOfUpdate)
			scheduleCategories.DELETE(":id", middleware.Permission("site.schedule_category.delete"), site.DoScheduleCategoryOfDelete)
		}

		scheduleCategory := route.Group("schedule-category")
		{
			scheduleCategory.POST("", middleware.Permission("site.schedule_category.create"), site.DoScheduleCategoryOfCreate)
			scheduleCategory.GET("opening", site.ToScheduleCategoryOfOpening)
		}

		persons := route.Group("persons")
		{
			persons.GET(":id", site.ToPersonOfInformation)
			persons.GET("", middleware.Permission("site.person.paginate"), site.ToPersonOfPaginate)
			persons.PUT(":id", middleware.Permission("site.person.update"), site.DoPersonOfUpdate)
			persons.DELETE(":id", middleware.Permission("site.person.delete"), site.DoPersonOfDelete)
		}

		person := route.Group("person")
		{
			person.POST("", middleware.Permission("site.person.create"), site.DoPersonOfCreate)
			person.POST("import", middleware.Permission("site.person.import"), site.DoPersonOfImport)
			person.GET("template", middleware.Permission("site.person.import"), site.ToPersonOfTemplate)
			person.GET("opening", site.ToPersonOfOpening)
		}

		drawCategories := route.Group("draw-categories")
		{
			drawCategories.GET("", middleware.Permission("site.draw_category.paginate"), site.ToDrawCategoryOfPaginate)
			drawCategories.PUT(":id", middleware.Permission("site.draw_category.update"), site.DoDrawCategoryOfUpdate)
			drawCategories.DELETE(":id", middleware.Permission("site.draw_category.delete"), site.DoDrawCategoryOfDelete)
		}

		drawCategory := route.Group("draw-category")
		{
			drawCategory.POST("", middleware.Permission("site.draw_category.create"), site.DoDrawCategoryOfCreate)
			drawCategory.GET("opening", site.ToDrawCategoryOfOpening)
		}

		draws := route.Group("draws")
		{
			draws.GET("", middleware.Permission("site.draw.paginate"), site.ToDrawOfPaginate)
			draws.DELETE(":id", middleware.Permission("site.draw.delete"), site.DoDrawOfDelete)
		}

		draw := route.Group("draw")
		{
			draw.POST("", middleware.Permission("site.draw.create"), site.DoDrawOfCreate)
		}

		scenes := route.Group("scenes")
		{
			scenes.GET("", middleware.Permission("site.scene.paginate"), site.ToSceneOfPaginate)
			scenes.PUT(":id", middleware.Permission("site.scene.update"), site.DoSceneOfUpdate)
			scenes.DELETE(":id", middleware.Permission("site.scene.delete"), site.DoSceneOfDelete)
		}

		scene := route.Group("scene")
		{
			scene.POST("", middleware.Permission("site.scene.create"), site.DoSceneOfCreate)
			scene.GET("opening", site.ToSceneOfOpening)
		}

		medias := route.Group("medias")
		{
			medias.GET("", middleware.Permission("site.media.paginate"), site.ToMediaOfPaginate)
			medias.PUT(":id", middleware.Permission("site.media.update"), site.DoMediaOfUpdate)
			medias.DELETE(":id", middleware.Permission("site.media.delete"), site.DoMediaOfDelete)
		}

		media := route.Group("media")
		{
			media.POST("images", middleware.Permission("site.media.create"), site.DoMediaOfCreateByImage)
			media.POST("video", middleware.Permission("site.media.create"), site.DoMediaOfCreateByVideo)
		}

		managers := route.Group("managers")
		{
			managers.GET("", middleware.Permission("site.manager.paginate"), site.ToManagerOfPaginate)
			managers.PUT(":id", middleware.Permission("site.manager.update"), site.DoManagerOfUpdate)
			managers.DELETE(":id", middleware.Permission("site.manager.delete"), site.DoManagerOfDelete)
		}

		manager := route.Group("manager")
		{
			manager.POST("", middleware.Permission("site.manager.create"), site.DoManagerOfCreate)
			manager.PUT("enable", middleware.Permission("site.manager.enable"), site.DoManagerOfEnable)
		}

		scores := route.Group("scores")
		{
			scores.GET(":id", site.ToScoreOfInformation)
			scores.GET("", middleware.Permission("site.score.paginate"), site.ToScoreOfPaginate)
			scores.PUT(":id", middleware.Permission("site.score.update"), site.DoScoreOfUpdate)
			scores.DELETE(":id", middleware.Permission("site.score.delete"), site.DoScoreOfDelete)
		}

		score := route.Group("score")
		{
			score.POST("", middleware.Permission("site.score.create"), site.DoScoreOfCreate)
			score.POST("import", middleware.Permission("site.score.import"), site.DoScoreOfImport)
			score.POST("notify", middleware.Permission("site.score.notify"), site.DoScoreOfNotify)
			score.GET("template", middleware.Permission("site.score.import"), site.ToScoreOfTemplate)
		}
	}
}
