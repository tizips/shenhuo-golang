package site

import (
	"context"
	"strconv"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/herhe-com/framework/contracts/http/response"
	"github.com/herhe-com/framework/facades"
	"github.com/herhe-com/framework/http"
	"github.com/tizips/shenhuo/helper"
	"github.com/tizips/shenhuo/model"
	req "github.com/tizips/shenhuo/server/admin/http/request/site"
	res "github.com/tizips/shenhuo/server/admin/http/response/site"
)

// ToNavOfPaginate
// @Summary 获取导航列表
// @Description Permissions: site.nav.paginate
// @Tags 站点-导航
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param page query int false "页码"
// @Param size query int false "每页数量"
// @Success 200 {object} response.Paginate[res.ToNavOfPaginate] "导航列表"
// @Router /site/navs [get]
func ToNavOfPaginate(c context.Context, ctx *app.RequestContext) {

	var request req.ToNavOfPaginate

	if err := ctx.BindAndValidate(&request); err != nil {
		http.BadRequest(ctx, err)
		return
	}

	responses := response.Paginate[res.ToNavOfPaginate]{
		Page: request.GetPage(),
		Size: request.GetSize(),
	}

	tx := facades.Database().Default().WithContext(c).Model(&model.ShNav{})

	tx.Count(&responses.Total)

	if responses.Total > 0 {

		var navs []model.ShNav

		tx.Order("`order` asc, `id` asc").Offset(request.GetOffset()).Limit(request.GetLimit()).Find(&navs)

		names := make(map[uint]string, len(navs))
		ids := make([]uint, 0, len(navs))
		for _, item := range navs {
			if item.Type != model.ShNavOfTypePage {
				continue
			}
			if id, err := strconv.ParseUint(item.Value, 10, 64); err == nil && id > 0 {
				ids = append(ids, uint(id))
			}
		}
		if len(ids) > 0 {
			var pages []model.ShPage
			facades.Database().Default().WithContext(c).Where("`id` IN ?", ids).Find(&pages)
			for _, item := range pages {
				names[item.ID] = item.Title
			}
		}

		responses.Data = make([]res.ToNavOfPaginate, len(navs))

		for index, item := range navs {
			var pageName string
			if item.Type == model.ShNavOfTypePage {
				if id, err := strconv.ParseUint(item.Value, 10, 64); err == nil {
					pageName = names[uint(id)]
				}
			}
			responses.Data[index] = res.ToNavOfPaginate{
				ID:        item.ID,
				Title:     item.Title,
				Icon:      item.Icon,
				Type:      item.Type,
				Value:     item.Value,
				PageName:  pageName,
				Order:     item.Order,
				CreatedAt: item.CreatedAt.ToDateTimeString(),
			}
		}
	}

	http.Success(ctx, responses)
}

// DoNavOfCreate
// @Summary 创建导航
// @Description Permissions: site.nav.create
// @Tags 站点-导航
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param request body req.DoNavOfCreate true "导航信息"
// @Success 200 {object} nil "创建成功"
// @Router /site/nav [post]
func DoNavOfCreate(c context.Context, ctx *app.RequestContext) {

	var request req.DoNavOfCreate

	if err := ctx.BindAndValidate(&request); err != nil {
		http.BadRequest(ctx, err)
		return
	}

	if !validNavValue(c, request.Type, request.Value) {
		http.Fail(ctx, "页面不存在")
		return
	}

	nav := model.ShNav{
		Title: request.Title,
		Icon:  request.Icon,
		Type:  request.Type,
		Value: request.Value,
		Order: helper.Order(request.Order),
	}

	if result := facades.Database().Default().WithContext(c).Create(&nav); result.Error != nil {
		http.Fail(ctx, "创建失败：%v", result.Error)
		return
	}

	http.Success[any](ctx)
}

// DoNavOfUpdate
// @Summary 更新导航
// @Description Permissions: site.nav.update
// @Tags 站点-导航
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path int true "导航ID"
// @Param request body req.DoNavOfUpdate true "导航信息"
// @Success 200 {object} nil "更新成功"
// @Router /site/navs/{id} [put]
func DoNavOfUpdate(c context.Context, ctx *app.RequestContext) {

	var request req.DoNavOfUpdate

	if err := ctx.BindAndValidate(&request); err != nil {
		http.BadRequest(ctx, err)
		return
	}

	if !validNavValue(c, request.Type, request.Value) {
		http.Fail(ctx, "页面不存在")
		return
	}

	var nav model.ShNav

	if err := firstByID(c, &nav, request.ID); err != nil {
		writeFindError(ctx, err)
		return
	}

	nav.Title = request.Title
	nav.Icon = request.Icon
	nav.Type = request.Type
	nav.Value = request.Value
	nav.Order = helper.Order(request.Order)

	if result := facades.Database().Default().WithContext(c).Save(&nav); result.Error != nil {
		http.Fail(ctx, "修改失败：%v", result.Error)
		return
	}

	http.Success[any](ctx)
}

// DoNavOfDelete
// @Summary 删除导航
// @Description Permissions: site.nav.delete
// @Tags 站点-导航
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path int true "导航ID"
// @Success 200 {object} nil "删除成功"
// @Router /site/navs/{id} [delete]
func DoNavOfDelete(c context.Context, ctx *app.RequestContext) {

	var request req.DoNavOfDelete

	if err := ctx.BindAndValidate(&request); err != nil {
		http.BadRequest(ctx, err)
		return
	}

	var nav model.ShNav

	if err := firstByID(c, &nav, request.ID); err != nil {
		writeFindError(ctx, err)
		return
	}

	if result := facades.Database().Default().WithContext(c).Delete(&nav); result.Error != nil {
		http.Fail(ctx, "删除失败：%v", result.Error)
		return
	}

	http.Success[any](ctx)
}

func validNavValue(c context.Context, typ, value string) bool {
	if typ != model.ShNavOfTypePage {
		return true
	}

	id, err := strconv.ParseUint(value, 10, 64)
	if err != nil || id == 0 {
		return false
	}

	var total int64
	facades.Database().Default().WithContext(c).Model(&model.ShPage{}).Where("`id`=?", id).Count(&total)

	return total > 0
}
