package shenhuo

import (
	"context"
	"errors"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/go-redsync/redsync/v4"
	"github.com/herhe-com/framework/facades"
	"github.com/herhe-com/framework/http"
	"github.com/herhe-com/framework/microservice/locker"
	"github.com/tizips/shenhuo/helper"
	"github.com/tizips/shenhuo/model"
	req "github.com/tizips/shenhuo/server/web/http/request/shenhuo"
	res "github.com/tizips/shenhuo/server/web/http/response/shenhuo"
	"gorm.io/gorm"
)

// ToDrawOfList
// @Summary 获取抽签结果
// @Description 按类别展示抽签结果
// @Tags 抽签
// @Accept json
// @Produce json
// @Success 200 {array} res.ToDrawCategory "抽签结果"
// @Router /draws [get]
func ToDrawOfList(c context.Context, ctx *app.RequestContext) {

	var categories []model.ShDrawCategory

	facades.Database().Default().WithContext(c).Order("`order` asc, `id` asc").Find(&categories)

	responses := make([]res.ToDrawCategory, len(categories))

	for index, item := range categories {

		var draws []model.ShDraw
		facades.Database().Default().WithContext(c).Preload("Person").Where("`category_id`=?", item.ID).Order("`id` asc").Find(&draws)

		winners := make([]res.ToDraw, len(draws))
		for i, draw := range draws {
			winners[i] = drawToResponse(draw)
		}

		responses[index] = res.ToDrawCategory{
			ID:      item.ID,
			Name:    item.Name,
			Icon:    item.Icon,
			Quota:   item.Quota,
			Order:   item.Order,
			Drawn:   len(draws),
			Winners: winners,
		}
	}

	http.Success(ctx, responses)
}

// DoDrawOfCreate
// @Summary 抽签
// @Description 管理人员抽签
// @Tags 抽签
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param request body req.DoDrawOfCreate true "抽签信息"
// @Success 200 {array} res.ToDraw "中签结果"
// @Router /draw [post]
func DoDrawOfCreate(c context.Context, ctx *app.RequestContext) {

	var request req.DoDrawOfCreate

	if err := ctx.BindAndValidate(&request); err != nil {
		http.BadRequest(ctx, err)
		return
	}

	mutex := facades.Locker().NewMutex(locker.Keys("draw", request.CategoryID))
	if err := mutex.Lock(); err != nil {
		http.Fail(ctx, "正在抽签，请稍后重试")
		return
	}
	defer func(lock *redsync.Mutex) {
		_, _ = lock.Unlock()
	}(mutex)

	tx := facades.Database().Default().WithContext(c).Begin()

	draws, err := helper.DrawCategory(tx, request.CategoryID)
	if err != nil {
		tx.Rollback()
		if errors.Is(err, gorm.ErrRecordNotFound) {
			http.NotFound(ctx, "未找到该抽签类别")
			return
		}
		http.Fail(ctx, "%s", err.Error())
		return
	}

	tx.Commit()

	responses := make([]res.ToDraw, len(draws))
	for index, item := range draws {
		responses[index] = drawToResponse(item)
	}

	http.Success(ctx, responses)
}

func drawToResponse(item model.ShDraw) res.ToDraw {
	row := res.ToDraw{
		ID:        item.ID,
		PersonID:  item.PersonID,
		CreatedAt: item.CreatedAt.ToDateTimeString(),
	}

	if item.Person != nil {
		row.Name = item.Person.Name
		row.Number = item.Person.Number
		row.Unit = item.Person.Unit
		row.GroupName = item.Person.GroupName
	}

	return row
}
