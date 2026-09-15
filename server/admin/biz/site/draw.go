package site

import (
	"context"
	"errors"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/go-redsync/redsync/v4"
	"github.com/herhe-com/framework/contracts/http/response"
	"github.com/herhe-com/framework/facades"
	"github.com/herhe-com/framework/http"
	"github.com/herhe-com/framework/microservice/locker"
	"github.com/tizips/shenhuo/helper"
	"github.com/tizips/shenhuo/model"
	req "github.com/tizips/shenhuo/server/admin/http/request/site"
	res "github.com/tizips/shenhuo/server/admin/http/response/site"
	"gorm.io/gorm"
)

// ToDrawOfPaginate
// @Summary 获取抽签结果列表
// @Description Permissions: site.draw.paginate
// @Tags 站点-抽签
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param page query int false "页码"
// @Param size query int false "每页数量"
// @Param category_id query int false "抽签类别"
// @Success 200 {object} response.Paginate[res.ToDrawOfPaginate] "抽签结果"
// @Router /site/draws [get]
func ToDrawOfPaginate(c context.Context, ctx *app.RequestContext) {

	var request req.ToDrawOfPaginate

	if err := ctx.BindAndValidate(&request); err != nil {
		http.BadRequest(ctx, err)
		return
	}

	responses := response.Paginate[res.ToDrawOfPaginate]{
		Page: request.GetPage(),
		Size: request.GetSize(),
	}

	tx := facades.Database().Default().WithContext(c).Model(&model.ShDraw{})

	if request.CategoryID > 0 {
		tx = tx.Where("`category_id`=?", request.CategoryID)
	}

	tx.Count(&responses.Total)

	if responses.Total > 0 {

		var draws []model.ShDraw

		tx.Preload("Person").Preload("Category").Order("`id` desc").Offset(request.GetOffset()).Limit(request.GetLimit()).Find(&draws)

		responses.Data = make([]res.ToDrawOfPaginate, len(draws))

		for index, item := range draws {
			responses.Data[index] = drawToResponse(item)
		}
	}

	http.Success(ctx, responses)
}

// DoDrawOfCreate
// @Summary 抽签
// @Description Permissions: site.draw.create
// @Tags 站点-抽签
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param request body req.DoDrawOfCreate true "抽签信息"
// @Success 200 {array} res.DoDrawOfCreate "中签结果"
// @Router /site/draw [post]
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

	responses := make([]res.DoDrawOfCreate, len(draws))
	for index, item := range draws {
		responses[index] = res.DoDrawOfCreate{ToDrawOfPaginate: drawToResponse(item)}
	}

	http.Success(ctx, responses)
}

// DoDrawOfDelete
// @Summary 删除抽签结果
// @Description Permissions: site.draw.delete
// @Tags 站点-抽签
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path int true "抽签结果ID"
// @Success 200 {object} nil "删除成功"
// @Router /site/draws/{id} [delete]
func DoDrawOfDelete(c context.Context, ctx *app.RequestContext) {

	var request req.DoDrawOfDelete

	if err := ctx.BindAndValidate(&request); err != nil {
		http.BadRequest(ctx, err)
		return
	}

	var item model.ShDraw

	if err := firstByID(c, &item, request.ID); err != nil {
		writeFindError(ctx, err)
		return
	}

	if result := facades.Database().Default().WithContext(c).Delete(&item); result.Error != nil {
		http.Fail(ctx, "删除失败：%v", result.Error)
		return
	}

	http.Success[any](ctx)
}

func drawToResponse(item model.ShDraw) res.ToDrawOfPaginate {
	row := res.ToDrawOfPaginate{
		ID:         item.ID,
		CategoryID: item.CategoryID,
		PersonID:   item.PersonID,
		CreatedAt:  item.CreatedAt.ToDateTimeString(),
	}

	if item.Category != nil {
		row.Category = item.Category.Name
	}

	if item.Person != nil {
		row.Name = item.Person.Name
		row.Number = item.Person.Number
		row.Unit = item.Person.Unit
		row.GroupName = item.Person.GroupName
	}

	return row
}
