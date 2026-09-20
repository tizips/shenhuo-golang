package route

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/herhe-com/framework/http/middleware"
	"github.com/tizips/shenhuo/server/web/biz/shenhuo"
	middle "github.com/tizips/shenhuo/server/web/http/middleware"
)

func ShenhuoRouter(router *server.Hertz) {

	article := router.Group("article")
	{
		article.GET("pinned", shenhuo.ToArticleOfPinned)
		article.GET("recommend", shenhuo.ToArticleOfRecommend)
	}

	articles := router.Group("articles")
	{
		articles.GET(":id", shenhuo.ToArticleOfInformation)
		articles.GET("", shenhuo.ToArticleOfPaginate)
	}

	pages := router.Group("pages")
	{
		pages.GET(":id", shenhuo.ToPageOfInformation)
		pages.GET("", shenhuo.ToPageOfList)
	}

	page := router.Group("page")
	{
		page.GET("keys/:key", shenhuo.ToPageOfInformationByKey)
	}

	persons := router.Group("persons")
	{
		persons.GET("", shenhuo.ToPersonOfList)

		groups := persons.Group("groups")
		{
			groups.GET("", shenhuo.ToPersonOfGroupList)
			groups.GET(":id", shenhuo.ToPersonOfGroupInformation)
		}
	}

	navs := router.Group("navs")
	{
		navs.GET("", shenhuo.ToNavOfList)
	}

	banners := router.Group("banners")
	{
		banners.GET("", shenhuo.ToBannerOfList)
	}

	schedules := router.Group("schedules")
	{
		schedules.GET("", shenhuo.ToScheduleOfList)
	}

	draws := router.Group("draws")
	{
		draws.GET("", shenhuo.ToDrawOfList)
	}

	scenes := router.Group("scenes")
	{
		scenes.GET("", shenhuo.ToSceneOfList)
	}

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

	system := router.Group("system")
	{
		system.GET("config", shenhuo.ToSystemOfConfig)
	}

	wechat := router.Group("wechat")
	{
		wechat.GET("oauth", shenhuo.ToWechatOfOAuth)
	}

	score := router.Group("score")
	{
		score.POST("query", shenhuo.ToScoreOfQuery)
		score.GET("mine", middleware.Auth(), middle.Person(), shenhuo.ToScoreOfMine)
	}

	draw := router.Group("draw").Use(middleware.Auth(), middle.Manager())
	{
		draw.POST("", shenhuo.DoDrawOfCreate)
	}
}
