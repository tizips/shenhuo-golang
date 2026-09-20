package route

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/herhe-com/framework/http/middleware"
	"github.com/tizips/shenhuo/server/web/biz/shenhuo"
	middle "github.com/tizips/shenhuo/server/web/http/middleware"
)

func ShenhuoRouter(router *server.Hertz) {

	articles := router.Group("articles")
	{
		articles.GET("pinned", shenhuo.ToArticleOfPinned)
		articles.GET(":id", shenhuo.ToArticleOfInformation)
		articles.GET("", shenhuo.ToArticleOfPaginate)
	}

	pages := router.Group("pages")
	{
		pages.GET(":id", shenhuo.ToPageOfInformation)
		pages.GET("key/:key", shenhuo.ToPageOfInformationByKey)
		pages.GET("", shenhuo.ToPageOfList)
	}

	router.GET("persons", shenhuo.ToPersonOfList)
	router.GET("persons/groups", shenhuo.ToPersonOfGroupList)
	router.GET("persons/groups/:id", shenhuo.ToPersonOfGroupInformation)
	router.GET("navs", shenhuo.ToNavOfList)
	router.GET("banners", shenhuo.ToBannerOfList)
	router.GET("schedules", shenhuo.ToScheduleOfList)
	router.GET("draws", shenhuo.ToDrawOfList)
	router.GET("scenes", shenhuo.ToSceneOfList)

	media := router.Group("media")
	{
		media.GET("pinned", shenhuo.ToMediaOfPinned)
		media.POST("images", middleware.Auth(), middle.Manager(), shenhuo.DoMediaOfCreateByImage)
		media.POST("video", middleware.Auth(), middle.Manager(), shenhuo.DoMediaOfCreateByVideo)
	}

	medias := router.Group("medias")
	{
		medias.GET("", shenhuo.ToMediaOfPaginate)
	}

	router.POST("score/query", shenhuo.ToScoreOfQuery)

	score := router.Group("score").Use(middleware.Auth(), middle.Person())
	{
		score.GET("mine", shenhuo.ToScoreOfMine)
	}

	draw := router.Group("draw").Use(middleware.Auth(), middle.Manager())
	{
		draw.POST("", shenhuo.DoDrawOfCreate)
	}
}
