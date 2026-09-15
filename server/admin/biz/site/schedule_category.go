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

// ToScheduleCategoryOfPaginate
// @Summary 获取日程分类列表
// @Description Permissions: site.schedule_category.paginate
// @Tags 站点-日程
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param page query int false "页码"
// @Param size query int false "每页数量"
// @Success 200 {object} response.Paginate[res.ToScheduleCategoryOfPaginate] "日程分类列表"
// @Router /site/schedule-categories [get]
func ToScheduleCategoryOfPaginate(c context.Context, ctx *app.RequestContext) {

	var request req.ToScheduleCategoryOfPaginate

	if err := ctx.BindAndValidate(&request); err != nil {
		http.BadRequest(ctx, err)
		return
	}

	responses := response.Paginate[res.ToScheduleCategoryOfPaginate]{
		Page: request.GetPage(),
		Size: request.GetSize(),
	}

	tx := facades.Database().Default().WithContext(c).Model(&model.ShScheduleCategory{})

	tx.Count(&responses.Total)

	if responses.Total > 0 {

		var categories []model.ShScheduleCategory

		tx.Order("`order` asc, `id` asc").Offset(request.GetOffset()).Limit(request.GetLimit()).Find(&categories)

		responses.Data = make([]res.ToScheduleCategoryOfPaginate, len(categories))

		for index, item := range categories {

			var total int64
			facades.Database().Default().WithContext(c).Model(&model.ShSchedule{}).Where("`category_id`=?", item.ID).Count(&total)

			responses.Data[index] = res.ToScheduleCategoryOfPaginate{
				ID:        item.ID,
				Name:      item.Name,
				Order:     item.Order,
				Total:     total,
				CreatedAt: item.CreatedAt.ToDateTimeString(),
			}
		}
	}

	http.Success(ctx, responses)
}

// ToScheduleCategoryOfOpening
// @Summary 获取可用日程分类
// @Description 获取所有可用的日程分类，默认返回最新的 20 个
// @Tags 站点-日程
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {array} response.Opening[uint] "日程分类列表"
// @Router /site/schedule-category/opening [get]
func ToScheduleCategoryOfOpening(c context.Context, ctx *app.RequestContext) {

	var categories []model.ShScheduleCategory

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

// DoScheduleCategoryOfCreate
// @Summary 创建日程分类
// @Description Permissions: site.schedule_category.create
// @Tags 站点-日程
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param request body req.DoScheduleCategoryOfCreate true "日程分类信息"
// @Success 200 {object} nil "创建成功"
// @Router /site/schedule-category [post]
func DoScheduleCategoryOfCreate(c context.Context, ctx *app.RequestContext) {

	var request req.DoScheduleCategoryOfCreate

	if err := ctx.BindAndValidate(&request); err != nil {
		http.BadRequest(ctx, err)
		return
	}

	item := model.ShScheduleCategory{
		Name:  request.Name,
		Order: helper.Order(request.Order),
	}

	if result := facades.Database().Default().WithContext(c).Create(&item); result.Error != nil {
		http.Fail(ctx, "创建失败：%v", result.Error)
		return
	}

	http.Success[any](ctx)
}

// DoScheduleCategoryOfUpdate
// @Summary 更新日程分类
// @Description Permissions: site.schedule_category.update
// @Tags 站点-日程
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path int true "日程分类ID"
// @Param request body req.DoScheduleCategoryOfUpdate true "日程分类信息"
// @Success 200 {object} nil "更新成功"
// @Router /site/schedule-categories/{id} [put]
func DoScheduleCategoryOfUpdate(c context.Context, ctx *app.RequestContext) {

	var request req.DoScheduleCategoryOfUpdate

	if err := ctx.BindAndValidate(&request); err != nil {
		http.BadRequest(ctx, err)
		return
	}

	var item model.ShScheduleCategory

	if err := firstByID(c, &item, request.ID); err != nil {
		writeFindError(ctx, err)
		return
	}

	item.Name = request.Name
	item.Order = helper.Order(request.Order)

	if result := facades.Database().Default().WithContext(c).Save(&item); result.Error != nil {
		http.Fail(ctx, "修改失败：%v", result.Error)
		return
	}

	http.Success[any](ctx)
}

// DoScheduleCategoryOfDelete
// @Summary 删除日程分类
// @Description Permissions: site.schedule_category.delete
// @Tags 站点-日程
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path int true "日程分类ID"
// @Success 200 {object} nil "删除成功"
// @Router /site/schedule-categories/{id} [delete]
func DoScheduleCategoryOfDelete(c context.Context, ctx *app.RequestContext) {

	var request req.DoScheduleCategoryOfDelete

	if err := ctx.BindAndValidate(&request); err != nil {
		http.BadRequest(ctx, err)
		return
	}

	var item model.ShScheduleCategory

	if err := firstByID(c, &item, request.ID); err != nil {
		writeFindError(ctx, err)
		return
	}

	tx := facades.Database().Default().WithContext(c).Begin()

	if result := tx.Where("`category_id`=?", item.ID).Delete(&model.ShSchedule{}); result.Error != nil {
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
