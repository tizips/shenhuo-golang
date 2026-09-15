package shenhuo

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/herhe-com/framework/facades"
	"github.com/herhe-com/framework/http"
	"github.com/tizips/shenhuo/model"
	res "github.com/tizips/shenhuo/server/web/http/response/shenhuo"
)

// ToNavOfList
// @Summary 获取导航列表
// @Description 获取导航列表
// @Tags 导航
// @Accept json
// @Produce json
// @Success 200 {array} res.ToNav "导航列表"
// @Router /navs [get]
func ToNavOfList(c context.Context, ctx *app.RequestContext) {

	var navs []model.ShNav

	facades.Database().Default().WithContext(c).Order("`order` asc, `id` asc").Find(&navs)

	responses := make([]res.ToNav, len(navs))
	for index, item := range navs {
		responses[index] = res.ToNav{
			ID:    item.ID,
			Title: item.Title,
			Icon:  item.Icon,
			Type:  item.Type,
			Value: item.Value,
			Order: item.Order,
		}
	}

	http.Success(ctx, responses)
}
