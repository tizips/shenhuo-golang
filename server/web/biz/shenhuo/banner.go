package shenhuo

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/herhe-com/framework/facades"
	"github.com/herhe-com/framework/http"
	"github.com/tizips/shenhuo/model"
	res "github.com/tizips/shenhuo/server/web/http/response/shenhuo"
)

// ToBannerOfList
// @Summary 获取轮播列表
// @Description 获取轮播列表
// @Tags 轮播
// @Accept json
// @Produce json
// @Success 200 {array} res.ToBanner "轮播列表"
// @Router /banners [get]
func ToBannerOfList(c context.Context, ctx *app.RequestContext) {

	var banners []model.ShBanner

	facades.Database().Default().WithContext(c).Order("`order` asc, `id` asc").Find(&banners)

	responses := make([]res.ToBanner, len(banners))
	for index, item := range banners {
		responses[index] = res.ToBanner{
			ID:    item.ID,
			Title: item.Title,
			Image: item.Image,
			Link:  item.Link,
			Order: item.Order,
		}
	}

	http.Success(ctx, responses)
}
