package shenhuo

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/herhe-com/framework/auth"
	"github.com/herhe-com/framework/facades"
	"github.com/herhe-com/framework/http"
	"github.com/tizips/shenhuo/model"
	req "github.com/tizips/shenhuo/server/web/http/request/shenhuo"
)

// DoAppealOfCreate
// @Summary 仲裁申诉
// @Description 按照身份证号和密码验证身份后提交仲裁申诉
// @Tags 仲裁申诉
// @Accept json
// @Produce json
// @Param request body req.DoAppealOfCreate true "申诉信息"
// @Success 200 {object} nil "提交成功"
// @Router /appeal [post]
func DoAppealOfCreate(c context.Context, ctx *app.RequestContext) {

	var request req.DoAppealOfCreate

	if err := ctx.BindAndValidate(&request); err != nil {
		http.BadRequest(ctx, err)
		return
	}

	var person model.ShPerson

	fp := facades.Database().Default().WithContext(c).First(&person, "`id_card`=?", request.IDCard)
	if fp.Error != nil || !auth.CheckPassword(request.Password, person.Password) {
		http.Fail(ctx, "身份证号或密码错误")
		return
	}

	appeal := model.ShAppeal{
		PersonID: person.ID,
		Reason:   request.Reason,
	}

	if result := facades.Database().Default().WithContext(c).Create(&appeal); result.Error != nil {
		http.Fail(ctx, "提交失败：%v", result.Error)
		return
	}

	http.Success[any](ctx)
}
