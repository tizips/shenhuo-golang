package shenhuo

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/golang-module/carbon/v2"
	"github.com/herhe-com/framework/facades"
	"github.com/herhe-com/framework/http"
	"github.com/tizips/shenhuo/model"
	req "github.com/tizips/shenhuo/server/web/http/request/shenhuo"
)

// DoCheckinOfCreate
// @Summary 签到
// @Description 参赛人员按照姓名进行签到；同一天同一姓名不可重复签到
// @Tags 签到
// @Accept json
// @Produce json
// @Param request body req.DoCheckinOfCreate true "签到信息"
// @Success 200 {object} nil "签到成功"
// @Router /checkin [post]
func DoCheckinOfCreate(c context.Context, ctx *app.RequestContext) {

	var request req.DoCheckinOfCreate

	if err := ctx.BindAndValidate(&request); err != nil {
		http.BadRequest(ctx, err)
		return
	}

	var total int64

	facades.Database().Default().WithContext(c).Model(&model.ShCheckin{}).
		Where("`name`=? and `created_at`>=?", request.Name, carbon.Now().StartOfDay()).
		Count(&total)

	if total > 0 {
		http.Fail(ctx, "今日已签到，请勿重复签到")
		return
	}

	checkin := model.ShCheckin{Name: request.Name}

	if result := facades.Database().Default().WithContext(c).Create(&checkin); result.Error != nil {
		http.Fail(ctx, "签到失败：%v", result.Error)
		return
	}

	http.Success[any](ctx)
}
