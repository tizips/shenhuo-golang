package site

import (
	"context"
	"errors"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/dromara/carbon/v2"
	"github.com/herhe-com/framework/contracts/http/response"
	"github.com/herhe-com/framework/facades"
	"github.com/herhe-com/framework/http"
	"github.com/tizips/shenhuo/helper"
	"github.com/tizips/shenhuo/model"
	req "github.com/tizips/shenhuo/server/admin/http/request/site"
	res "github.com/tizips/shenhuo/server/admin/http/response/site"
)

// ToBannerOfPaginate
// @Summary 获取轮播列表
// @Description Permissions: site.banner.paginate
// @Tags 站点-轮播
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param page query int false "页码"
// @Param size query int false "每页数量"
// @Success 200 {object} response.Paginate[res.ToBannerOfPaginate] "轮播列表"
// @Router /site/banners [get]
func ToBannerOfPaginate(c context.Context, ctx *app.RequestContext) {

	var request req.ToBannerOfPaginate

	if err := ctx.BindAndValidate(&request); err != nil {
		http.BadRequest(ctx, err)
		return
	}

	responses := response.Paginate[res.ToBannerOfPaginate]{
		Page: request.GetPage(),
		Size: request.GetSize(),
	}

	tx := facades.Database().Default().WithContext(c).Model(&model.ShBanner{})

	tx.Count(&responses.Total)

	if responses.Total > 0 {

		var banners []model.ShBanner

		tx.Order("`order` asc, `id` asc").Offset(request.GetOffset()).Limit(request.GetLimit()).Find(&banners)

		responses.Data = make([]res.ToBannerOfPaginate, len(banners))

		for index, item := range banners {
			responses.Data[index] = res.ToBannerOfPaginate{
				ID:        item.ID,
				Title:     item.Title,
				Image:     item.Image,
				Link:      item.Link,
				Order:     item.Order,
				StartedAt: formatNullableDateTime(&item.StartedAt),
				EndedAt:   formatNullableDateTime(&item.EndedAt),
				CreatedAt: item.CreatedAt.ToDateTimeString(),
			}
		}
	}

	http.Success(ctx, responses)
}

// DoBannerOfCreate
// @Summary 创建轮播
// @Description Permissions: site.banner.create
// @Tags 站点-轮播
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param request body req.DoBannerOfCreate true "轮播信息"
// @Success 200 {object} nil "创建成功"
// @Router /site/banner [post]
func DoBannerOfCreate(c context.Context, ctx *app.RequestContext) {

	var request req.DoBannerOfCreate

	if err := ctx.BindAndValidate(&request); err != nil {
		http.BadRequest(ctx, err)
		return
	}

	started, ended, err := parsePeriodOfBanner(request.StartedAt, request.EndedAt)
	if err != nil {
		http.BadRequest(ctx, err)
		return
	}

	banner := model.ShBanner{
		Title:     request.Title,
		Image:     request.Image,
		Link:      request.Link,
		Order:     helper.Order(request.Order),
		StartedAt: started,
		EndedAt:   ended,
	}

	if result := facades.Database().Default().WithContext(c).Create(&banner); result.Error != nil {
		http.Fail(ctx, "创建失败：%v", result.Error)
		return
	}

	http.Success[any](ctx)
}

// DoBannerOfUpdate
// @Summary 更新轮播
// @Description Permissions: site.banner.update
// @Tags 站点-轮播
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path int true "轮播ID"
// @Param request body req.DoBannerOfUpdate true "轮播信息"
// @Success 200 {object} nil "更新成功"
// @Router /site/banners/{id} [put]
func DoBannerOfUpdate(c context.Context, ctx *app.RequestContext) {

	var request req.DoBannerOfUpdate

	if err := ctx.BindAndValidate(&request); err != nil {
		http.BadRequest(ctx, err)
		return
	}

	var banner model.ShBanner

	if err := firstByID(c, &banner, request.ID); err != nil {
		writeFindError(ctx, err)
		return
	}

	started, ended, err := parsePeriodOfBanner(request.StartedAt, request.EndedAt)
	if err != nil {
		http.BadRequest(ctx, err)
		return
	}

	banner.Title = request.Title
	banner.Image = request.Image
	banner.Link = request.Link
	banner.Order = helper.Order(request.Order)
	banner.StartedAt = started
	banner.EndedAt = ended

	if result := facades.Database().Default().WithContext(c).Save(&banner); result.Error != nil {
		http.Fail(ctx, "修改失败：%v", result.Error)
		return
	}

	http.Success[any](ctx)
}

// DoBannerOfDelete
// @Summary 删除轮播
// @Description Permissions: site.banner.delete
// @Tags 站点-轮播
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path int true "轮播ID"
// @Success 200 {object} nil "删除成功"
// @Router /site/banners/{id} [delete]
func DoBannerOfDelete(c context.Context, ctx *app.RequestContext) {

	var request req.DoBannerOfDelete

	if err := ctx.BindAndValidate(&request); err != nil {
		http.BadRequest(ctx, err)
		return
	}

	var banner model.ShBanner

	if err := firstByID(c, &banner, request.ID); err != nil {
		writeFindError(ctx, err)
		return
	}

	if result := facades.Database().Default().WithContext(c).Delete(&banner); result.Error != nil {
		http.Fail(ctx, "删除失败：%v", result.Error)
		return
	}

	http.Success[any](ctx)
}

// parsePeriodOfBanner 解析轮播有效期；开始/结束时间均为必填
func parsePeriodOfBanner(started, ended string) (carbon.Carbon, carbon.Carbon, error) {

	if started == "" {
		return carbon.Carbon{}, carbon.Carbon{}, errors.New("生效开始时间不能为空")
	}

	if ended == "" {
		return carbon.Carbon{}, carbon.Carbon{}, errors.New("生效结束时间不能为空")
	}

	start := carbon.Parse(started)
	if start.Error != nil {
		return carbon.Carbon{}, carbon.Carbon{}, errors.New("生效开始时间格式错误")
	}

	end := carbon.Parse(ended)
	if end.Error != nil {
		return carbon.Carbon{}, carbon.Carbon{}, errors.New("生效结束时间格式错误")
	}

	if end.Lt(start) {
		return carbon.Carbon{}, carbon.Carbon{}, errors.New("生效结束时间不能早于生效开始时间")
	}

	return *start, *end, nil
}

// formatNullableDateTime 可空时间格式化为展示字符串；为空返回空串
func formatNullableDateTime(value *carbon.Carbon) string {
	if value == nil || value.IsNil() || value.IsInvalid() {
		return ""
	}
	return value.ToDateTimeString()
}
