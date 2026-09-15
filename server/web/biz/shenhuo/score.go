package shenhuo

import (
	"context"
	"errors"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/herhe-com/framework/auth"
	"github.com/herhe-com/framework/constants/global"
	"github.com/herhe-com/framework/facades"
	"github.com/herhe-com/framework/http"
	"github.com/tizips/shenhuo/model"
	req "github.com/tizips/shenhuo/server/web/http/request/shenhuo"
	res "github.com/tizips/shenhuo/server/web/http/response/shenhuo"
	"gorm.io/gorm"
)

// ToScoreOfQuery
// @Summary 查询成绩
// @Description 按照手机号和密码验证身份后查询对应用户的成绩
// @Tags 成绩
// @Accept json
// @Produce json
// @Param request body req.DoScoreOfQuery true "查询信息"
// @Success 200 {object} res.ToScore "成绩信息"
// @Router /score/query [post]
func ToScoreOfQuery(c context.Context, ctx *app.RequestContext) {

	var request req.DoScoreOfQuery

	if err := ctx.BindAndValidate(&request); err != nil {
		http.BadRequest(ctx, err)
		return
	}

	var person model.ShPerson

	fp := facades.Database().Default().WithContext(c).First(&person, "`mobile`=?", request.Mobile)
	if fp.Error != nil || !auth.CheckPassword(request.Password, person.Password) {
		http.Fail(ctx, "手机号或密码错误")
		return
	}

	var score model.ShScore

	fs := facades.Database().Default().WithContext(c).Preload("Person").First(&score, "`person_id`=?", person.ID)
	if errors.Is(fs.Error, gorm.ErrRecordNotFound) {
		http.NotFound(ctx, "暂无成绩")
		return
	} else if fs.Error != nil {
		http.Fail(ctx, "查询失败：%v", fs.Error)
		return
	}

	http.Success(ctx, scoreToResponse(score))
}

// ToScoreOfMine
// @Summary 获取我的成绩
// @Description 参赛人员查询自己的成绩
// @Tags 成绩
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {object} res.ToScore "我的成绩"
// @Router /score/mine [get]
func ToScoreOfMine(c context.Context, ctx *app.RequestContext) {

	var person model.ShPerson

	fp := facades.Database().Default().WithContext(c).First(&person, "`id`=?", auth.ID(ctx))
	if errors.Is(fp.Error, gorm.ErrRecordNotFound) {
		http.Unauthorized(ctx)
		return
	} else if fp.Error != nil {
		http.Fail(ctx, "查询失败：%v", fp.Error)
		return
	}

	if person.MustChangePassword == global.YES {
		http.Fail(ctx, "请先修改密码")
		return
	}

	var score model.ShScore

	fs := facades.Database().Default().WithContext(c).Preload("Person").First(&score, "`person_id`=?", person.ID)
	if errors.Is(fs.Error, gorm.ErrRecordNotFound) {
		http.NotFound(ctx, "暂无成绩")
		return
	} else if fs.Error != nil {
		http.Fail(ctx, "查询失败：%v", fs.Error)
		return
	}

	http.Success(ctx, scoreToResponse(score))
}

func scoreToResponse(item model.ShScore) res.ToScore {
	row := res.ToScore{
		ID:       item.ID,
		PersonID: item.PersonID,
		Total:    item.Total,
		Items:    item.ItemList(),
		Order:    item.Order,
	}

	if item.Person != nil {
		row.Name = item.Person.Name
		row.Number = item.Person.Number
		row.Unit = item.Person.Unit
		row.GroupName = item.Person.GroupName
	}

	return row
}
