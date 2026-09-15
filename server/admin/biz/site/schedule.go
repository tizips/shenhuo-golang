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

// ToScheduleOfPaginate
// @Summary 获取日程列表
// @Description Permissions: site.schedule.paginate
// @Tags 站点-日程
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param page query int false "页码"
// @Param size query int false "每页数量"
// @Success 200 {object} response.Paginate[res.ToScheduleOfPaginate] "日程列表"
// @Router /site/schedules [get]
func ToScheduleOfPaginate(c context.Context, ctx *app.RequestContext) {

	var request req.ToScheduleOfPaginate

	if err := ctx.BindAndValidate(&request); err != nil {
		http.BadRequest(ctx, err)
		return
	}

	responses := response.Paginate[res.ToScheduleOfPaginate]{
		Page: request.GetPage(),
		Size: request.GetSize(),
	}

	tx := facades.Database().Default().WithContext(c).Model(&model.ShSchedule{})

	if request.CategoryID > 0 {
		tx.Where("`category_id`=?", request.CategoryID)
	}

	tx.Count(&responses.Total)

	if responses.Total > 0 {

		var items []model.ShSchedule

		tx.Order("`order` asc, `id` asc").Offset(request.GetOffset()).Limit(request.GetLimit()).Find(&items)

		var categories []model.ShScheduleCategory
		facades.Database().Default().WithContext(c).Find(&categories)

		names := make(map[uint]string, len(categories))
		for _, category := range categories {
			names[category.ID] = category.Name
		}

		responses.Data = make([]res.ToScheduleOfPaginate, len(items))

		for index, item := range items {

			subitems := make([]res.ToScheduleOfItem, 0, len(item.Items))
			for _, subitem := range item.Items {
				subitems = append(subitems, res.ToScheduleOfItem{
					Name: subitem.Name,
					Time: subitem.Time,
				})
			}

			responses.Data[index] = res.ToScheduleOfPaginate{
				ID:           item.ID,
				CategoryID:   item.CategoryID,
				CategoryName: names[item.CategoryID],
				Title:        item.Title,
				Subtitle:     item.Subtitle,
				Description:  item.Description,
				Time:         item.Time,
				Items:        subitems,
				Order:        item.Order,
				CreatedAt:    item.CreatedAt.ToDateTimeString(),
			}
		}
	}

	http.Success(ctx, responses)
}

// DoScheduleOfCreate
// @Summary 创建日程
// @Description Permissions: site.schedule.create
// @Tags 站点-日程
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param request body req.DoScheduleOfCreate true "日程信息"
// @Success 200 {object} nil "创建成功"
// @Router /site/schedule [post]
func DoScheduleOfCreate(c context.Context, ctx *app.RequestContext) {

	var request req.DoScheduleOfCreate

	if err := ctx.BindAndValidate(&request); err != nil {
		http.BadRequest(ctx, err)
		return
	}

	subitems := make([]model.ShScheduleItem, 0, len(request.Items))
	for _, subitem := range request.Items {
		subitems = append(subitems, model.ShScheduleItem{
			Name: subitem.Name,
			Time: subitem.Time,
		})
	}

	item := model.ShSchedule{
		CategoryID:  request.CategoryID,
		Title:       request.Title,
		Subtitle:    request.Subtitle,
		Description: request.Description,
		Time:        request.Time,
		Items:       subitems,
		Order:       helper.Order(request.Order),
	}

	if result := facades.Database().Default().WithContext(c).Create(&item); result.Error != nil {
		http.Fail(ctx, "创建失败：%v", result.Error)
		return
	}

	http.Success[any](ctx)
}

// DoScheduleOfUpdate
// @Summary 更新日程
// @Description Permissions: site.schedule.update
// @Tags 站点-日程
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path int true "日程ID"
// @Param request body req.DoScheduleOfUpdate true "日程信息"
// @Success 200 {object} nil "更新成功"
// @Router /site/schedules/{id} [put]
func DoScheduleOfUpdate(c context.Context, ctx *app.RequestContext) {

	var request req.DoScheduleOfUpdate

	if err := ctx.BindAndValidate(&request); err != nil {
		http.BadRequest(ctx, err)
		return
	}

	var item model.ShSchedule

	if err := firstByID(c, &item, request.ID); err != nil {
		writeFindError(ctx, err)
		return
	}

	subitems := make([]model.ShScheduleItem, 0, len(request.Items))
	for _, subitem := range request.Items {
		subitems = append(subitems, model.ShScheduleItem{
			Name: subitem.Name,
			Time: subitem.Time,
		})
	}

	item.CategoryID = request.CategoryID
	item.Title = request.Title
	item.Subtitle = request.Subtitle
	item.Description = request.Description
	item.Time = request.Time
	item.Items = subitems
	item.Order = helper.Order(request.Order)

	if result := facades.Database().Default().WithContext(c).Save(&item); result.Error != nil {
		http.Fail(ctx, "修改失败：%v", result.Error)
		return
	}

	http.Success[any](ctx)
}

// DoScheduleOfDelete
// @Summary 删除日程
// @Description Permissions: site.schedule.delete
// @Tags 站点-日程
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path int true "日程ID"
// @Success 200 {object} nil "删除成功"
// @Router /site/schedules/{id} [delete]
func DoScheduleOfDelete(c context.Context, ctx *app.RequestContext) {

	var request req.DoScheduleOfDelete

	if err := ctx.BindAndValidate(&request); err != nil {
		http.BadRequest(ctx, err)
		return
	}

	var item model.ShSchedule

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
