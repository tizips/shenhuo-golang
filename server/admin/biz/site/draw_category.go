package site

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/herhe-com/framework/contracts/http/response"
	"github.com/herhe-com/framework/facades"
	"github.com/herhe-com/framework/http"
	"github.com/tizips/shenhuo/helper"
	"github.com/tizips/shenhuo/model"
	req "github.com/tizips/shenhuo/server/admin/http/request/site"
	res "github.com/tizips/shenhuo/server/admin/http/response/site"
)

// ToDrawCategoryOfPaginate
// @Summary 获取抽签类别列表
// @Description Permissions: site.draw_category.paginate
// @Tags 站点-抽签
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param page query int false "页码"
// @Param size query int false "每页数量"
// @Success 200 {object} response.Paginate[res.ToDrawCategoryOfPaginate] "抽签类别列表"
// @Router /site/draw-categories [get]
func ToDrawCategoryOfPaginate(c context.Context, ctx *app.RequestContext) {

	var request req.ToDrawCategoryOfPaginate

	if err := ctx.BindAndValidate(&request); err != nil {
		http.BadRequest(ctx, err)
		return
	}

	responses := response.Paginate[res.ToDrawCategoryOfPaginate]{
		Page: request.GetPage(),
		Size: request.GetSize(),
	}

	tx := facades.Database().Default().WithContext(c).Model(&model.ShDrawCategory{})

	tx.Count(&responses.Total)

	if responses.Total > 0 {

		var categories []model.ShDrawCategory

		tx.Order("`order` asc, `id` asc").Offset(request.GetOffset()).Limit(request.GetLimit()).Find(&categories)

		responses.Data = make([]res.ToDrawCategoryOfPaginate, len(categories))

		for index, item := range categories {
			var drawn int64
			facades.Database().Default().WithContext(c).Model(&model.ShDraw{}).Where("`category_id`=?", item.ID).Count(&drawn)

			responses.Data[index] = res.ToDrawCategoryOfPaginate{
				ID:        item.ID,
				Name:      item.Name,
				Icon:      item.Icon,
				Quota:     item.Quota,
				Order:     item.Order,
				Drawn:     drawn,
				CreatedAt: item.CreatedAt.ToDateTimeString(),
			}
		}
	}

	http.Success(ctx, responses)
}

// ToDrawCategoryOfOpening
// @Summary 获取可用抽签类别
// @Description 获取所有可用的抽签类别，默认返回最新的 20 个
// @Tags 站点-抽签
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {array} response.Opening[uint] "抽签类别列表"
// @Router /site/draw-category/opening [get]
func ToDrawCategoryOfOpening(c context.Context, ctx *app.RequestContext) {

	var categories []model.ShDrawCategory

	facades.Database().Default().WithContext(c).Order("`id` desc").Limit(20).Find(&categories)

	responses := make([]response.Opening[uint], len(categories))

	for index, item := range categories {
		responses[index] = response.Opening[uint]{
			ID:   item.ID,
			Name: item.Name,
		}
	}

	http.Success(ctx, responses)
}

// DoDrawCategoryOfCreate
// @Summary 创建抽签类别
// @Description Permissions: site.draw_category.create
// @Tags 站点-抽签
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param request body req.DoDrawCategoryOfCreate true "抽签类别信息"
// @Success 200 {object} nil "创建成功"
// @Router /site/draw-category [post]
func DoDrawCategoryOfCreate(c context.Context, ctx *app.RequestContext) {

	var request req.DoDrawCategoryOfCreate

	if err := ctx.BindAndValidate(&request); err != nil {
		http.BadRequest(ctx, err)
		return
	}

	item := model.ShDrawCategory{
		Name:  request.Name,
		Icon:  request.Icon,
		Quota: request.Quota,
		Order: helper.Order(request.Order),
	}

	if result := facades.Database().Default().WithContext(c).Create(&item); result.Error != nil {
		http.Fail(ctx, "创建失败：%v", result.Error)
		return
	}

	http.Success[any](ctx)
}

// DoDrawCategoryOfUpdate
// @Summary 更新抽签类别
// @Description Permissions: site.draw_category.update
// @Tags 站点-抽签
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path int true "抽签类别ID"
// @Param request body req.DoDrawCategoryOfUpdate true "抽签类别信息"
// @Success 200 {object} nil "更新成功"
// @Router /site/draw-categories/{id} [put]
func DoDrawCategoryOfUpdate(c context.Context, ctx *app.RequestContext) {

	var request req.DoDrawCategoryOfUpdate

	if err := ctx.BindAndValidate(&request); err != nil {
		http.BadRequest(ctx, err)
		return
	}

	var item model.ShDrawCategory

	if err := firstByID(c, &item, request.ID); err != nil {
		writeFindError(ctx, err)
		return
	}

	item.Name = request.Name
	item.Icon = request.Icon
	item.Quota = request.Quota
	item.Order = helper.Order(request.Order)

	if result := facades.Database().Default().WithContext(c).Save(&item); result.Error != nil {
		http.Fail(ctx, "修改失败：%v", result.Error)
		return
	}

	http.Success[any](ctx)
}

// DoDrawCategoryOfDelete
// @Summary 删除抽签类别
// @Description Permissions: site.draw_category.delete
// @Tags 站点-抽签
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path int true "抽签类别ID"
// @Success 200 {object} nil "删除成功"
// @Router /site/draw-categories/{id} [delete]
func DoDrawCategoryOfDelete(c context.Context, ctx *app.RequestContext) {

	var request req.DoDrawCategoryOfDelete

	if err := ctx.BindAndValidate(&request); err != nil {
		http.BadRequest(ctx, err)
		return
	}

	var item model.ShDrawCategory

	if err := firstByID(c, &item, request.ID); err != nil {
		writeFindError(ctx, err)
		return
	}

	tx := facades.Database().Default().WithContext(c).Begin()

	if result := tx.Where("`category_id`=?", item.ID).Delete(&model.ShDraw{}); result.Error != nil {
		tx.Rollback()
		http.Fail(ctx, "删除失败：%v", result.Error)
		return
	}

	if result := tx.Delete(&item); result.Error != nil {
		tx.Rollback()
		http.Fail(ctx, "删除失败：%v", result.Error)
		return
	}

	tx.Commit()

	http.Success[any](ctx)
}
