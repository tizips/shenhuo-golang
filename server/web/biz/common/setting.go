package common

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/go-playground/validator/v10"
	"github.com/gookit/goutil/strutil"
	"github.com/herhe-com/framework/facades"
	"github.com/herhe-com/framework/http"
	"github.com/tizips/shenhuo/model"
	"github.com/tizips/shenhuo/server/web/http/response/common"
)

func ToSetting(c context.Context, ctx *app.RequestContext, module string) {

	var settings []model.ComSetting

	facades.Database().Default().Order("`order` asc, `id` asc").Find(&settings, "`module`=?", module)

	responses := make([]common.ToSetting, len(settings))

	for index, item := range settings {
		responses[index] = common.ToSetting{
			ID:         item.ID,
			Type:       item.Type,
			Label:      item.Label,
			Key:        item.Key,
			Val:        item.Val,
			IsRequired: item.IsRequired,
			CreatedAt:  item.CreatedAt.ToDateTimeString(),
		}
	}

	http.Success(ctx, responses)
}

func DoSetting(c context.Context, ctx *app.RequestContext, module string) {

	var request map[string]string

	if err := ctx.Bind(&request); err != nil {
		http.BadRequest(ctx, err)
		return
	}

	var settings []model.ComSetting

	facades.Database().Default().Find(&settings, "`module`=?", module)

	updates := make(map[uint]string)

	valid, ok := facades.Optional[*validator.Validate]()
	if !ok {
		http.Fail(ctx, "请先开启验证器")
		return
	}

	for _, item := range settings {

		req, ok := request[item.Key]

		if item.IsRequired == model.ComSettingForIsRequiredOfYes && (!ok || strutil.IsEmpty(req)) {
			http.BadRequest(ctx, item.Label+"不能为空")
			return
		}

		if !strutil.IsEmpty(req) {

			var err error

			if item.Type == model.ComSettingForTypeOfEnable {
				err = valid.Var(req, "oneof=1 2")
			} else if item.Type == model.ComSettingForTypeOfURL || item.Type == model.ComSettingForTypeOfPicture {
				err = valid.Var(req, "url")
			} else if item.Type == model.ComSettingForTypeOfEmail {
				err = valid.Var(req, "email")
			}

			if err != nil {
				http.BadRequest(ctx, err)
				return
			}
		}

		if req != item.Val {
			updates[item.ID] = req
		}
	}

	if len(updates) > 0 {

		tx := facades.Database().Default().Begin()

		for index, item := range updates {
			if us := tx.Model(model.ComSetting{}).Where("`id`=?", index).Update("val", item); us.Error != nil {
				tx.Rollback()
				http.Fail(ctx, "保存失败：%v", us.Error)
				return
			}
		}

		tx.Commit()
	}

	http.Success[any](ctx)
}
