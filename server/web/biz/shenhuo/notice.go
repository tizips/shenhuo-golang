package shenhuo

import (
	"context"
	"errors"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/herhe-com/framework/facades"
	"github.com/herhe-com/framework/http"
	"github.com/tizips/shenhuo/model"
	req "github.com/tizips/shenhuo/server/web/http/request/shenhuo"
	res "github.com/tizips/shenhuo/server/web/http/response/shenhuo"
	"github.com/tizips/shenhuo/wechat"
	"gorm.io/gorm"
)

// DoNoticeOfTest
// @Summary 测试推送微信通知
// @Description 给绑定了微信的参赛人员直接发送模板消息，用于测试模板配置与文案
// @Tags 通知
// @Accept json
// @Produce json
// @Param request body req.DoNoticeOfTest true "测试信息"
// @Success 200 {array} res.DoNoticeOfTest "各类型通知发送结果"
// @Router /notice/test [post]
func DoNoticeOfTest(c context.Context, ctx *app.RequestContext) {

	var request req.DoNoticeOfTest

	if err := ctx.BindAndValidate(&request); err != nil {
		http.BadRequest(ctx, err)
		return
	}

	var person model.ShPerson

	fp := facades.Database().Default().WithContext(c).First(&person, "`id`=?", request.PersonID)
	if errors.Is(fp.Error, gorm.ErrRecordNotFound) {
		http.NotFound(ctx, "未找到该人员")
		return
	} else if fp.Error != nil {
		http.Fail(ctx, "查询失败：%v", fp.Error)
		return
	}

	if person.OpenID == "" {
		http.Fail(ctx, "该参赛人员未绑定微信，无法发送通知")
		return
	}

	types := []string{request.Type}
	if request.Type == "all" {
		types = []string{"score", "draw"}
	}

	responses := make([]res.DoNoticeOfTest, 0, len(types))
	for _, kind := range types {
		err := sendTestNotice(c, kind, &person)
		responses = append(responses, res.DoNoticeOfTest{
			Type:  kind,
			Ok:    err == nil,
			Error: errorText(err),
		})
	}

	http.Success(ctx, responses)
}

// sendTestNotice 按类型发送测试模板消息。
func sendTestNotice(c context.Context, kind string, person *model.ShPerson) error {

	switch kind {

	case "score":
		var score model.ShScore
		fs := facades.Database().Default().WithContext(c).First(&score, "`person_id`=?", person.ID)
		if fs.Error != nil && !errors.Is(fs.Error, gorm.ErrRecordNotFound) {
			return fs.Error
		}
		return wechat.SendScoreNotice(c, person.OpenID, person.Name, person.Number, score.Total)

	case "draw":
		group := person.GroupName
		var draw model.ShDraw
		fd := facades.Database().Default().WithContext(c).Preload("Category").
			Order("`id` desc").First(&draw, "`person_id`=?", person.ID)
		if fd.Error == nil && draw.Category != nil && draw.Category.Name != "" {
			group = draw.Category.Name
		}
		return wechat.SendDrawNotice(c, person.OpenID, person.Name, person.Unit, group)
	}

	return errors.New("未知的通知类型")
}

func errorText(err error) string {
	if err == nil {
		return ""
	}
	return err.Error()
}
