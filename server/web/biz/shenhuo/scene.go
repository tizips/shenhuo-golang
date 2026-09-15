package shenhuo

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/herhe-com/framework/facades"
	"github.com/herhe-com/framework/http"
	"github.com/tizips/shenhuo/model"
	res "github.com/tizips/shenhuo/server/web/http/response/shenhuo"
)

// ToSceneOfList
// @Summary 获取场景列表
// @Description 获取场景及媒体列表
// @Tags 场景
// @Accept json
// @Produce json
// @Success 200 {array} res.ToScene "场景列表"
// @Router /scenes [get]
func ToSceneOfList(c context.Context, ctx *app.RequestContext) {

	var scenes []model.ShScene

	facades.Database().Default().WithContext(c).Preload("Medias").Order("`order` asc, `id` asc").Find(&scenes)

	responses := make([]res.ToScene, len(scenes))
	for index, item := range scenes {
		medias := make([]res.ToMedia, len(item.Medias))
		for i, media := range item.Medias {
			medias[i] = res.ToMedia{
				ID:      media.ID,
				SceneID: media.SceneID,
				Type:    media.Type,
				URL:     media.URL,
				IsTop:   media.IsTop,
			}
		}
		responses[index] = res.ToScene{
			ID:     item.ID,
			Name:   item.Name,
			Order:  item.Order,
			Medias: medias,
		}
	}

	http.Success(ctx, responses)
}
