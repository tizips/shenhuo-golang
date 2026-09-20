package shenhuo

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/herhe-com/framework/facades"
	"github.com/herhe-com/framework/http"
	res "github.com/tizips/shenhuo/server/web/http/response/shenhuo"
)

// ToSystemOfConfig
// @Summary 获取系统配置
// @Description 获取系统配置
// @Tags 系统
// @Accept json
// @Produce json
// @Success 200 {object} res.ToSystemOfConfig "系统配置"
// @Router /system/config [get]
func ToSystemOfConfig(c context.Context, ctx *app.RequestContext) {

	http.Success(ctx, res.ToSystemOfConfig{
		Video: facades.Config().GetBool("system.video"),
	})
}
