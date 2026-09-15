package site

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/herhe-com/framework/contracts/http/response"
	"github.com/herhe-com/framework/facades"
	"github.com/herhe-com/framework/http"
	"github.com/tizips/shenhuo/model"
	req "github.com/tizips/shenhuo/server/admin/http/request/site"
	res "github.com/tizips/shenhuo/server/admin/http/response/site"
)

// ToPageOfPaginate
// @Summary 获取页面列表
// @Description Permissions: site.page.paginate
// @Tags 站点-页面
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param page query int false "页码"
// @Param size query int false "每页数量"
// @Param keyword query string false "关键词"
// @Success 200 {object} response.Paginate[res.ToPageOfPaginate] "页面列表"
// @Router /site/pages [get]
func ToPageOfPaginate(c context.Context, ctx *app.RequestContext) {

	var request req.ToPageOfPaginate

	if err := ctx.BindAndValidate(&request); err != nil {
		http.BadRequest(ctx, err)
		return
	}

	responses := response.Paginate[res.ToPageOfPaginate]{
		Page: request.GetPage(),
		Size: request.GetSize(),
	}

	tx := facades.Database().Default().WithContext(c).Model(&model.ShPage{})

	if request.Keyword != "" {
		tx = tx.Where("`title` LIKE ?", "%"+request.Keyword+"%")
	}

	tx.Count(&responses.Total)

	if responses.Total > 0 {

		var pages []model.ShPage

		tx.Order("`id` desc").Offset(request.GetOffset()).Limit(request.GetLimit()).Find(&pages)

		responses.Data = make([]res.ToPageOfPaginate, len(pages))

		for index, item := range pages {
			responses.Data[index] = res.ToPageOfPaginate{
				ID:        item.ID,
				Title:     item.Title,
				CreatedAt: item.CreatedAt.ToDateTimeString(),
			}
		}
	}

	http.Success(ctx, responses)
}

// ToPageOfInformation
// @Summary 获取页面详情
// @Description 获取指定页面详情
// @Tags 站点-页面
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path int true "页面ID"
// @Success 200 {object} res.ToPageOfInformation "页面详情"
// @Router /site/pages/{id} [get]
func ToPageOfInformation(c context.Context, ctx *app.RequestContext) {

	var request req.ToPageOfInformation

	if err := ctx.BindAndValidate(&request); err != nil {
		http.BadRequest(ctx, err)
		return
	}

	var page model.ShPage

	if err := firstByID(c, &page, request.ID); err != nil {
		writeFindError(ctx, err)
		return
	}

	http.Success(ctx, res.ToPageOfInformation{
		ID:        page.ID,
		Title:     page.Title,
		Content:   page.Content,
		CreatedAt: page.CreatedAt.ToDateTimeString(),
	})
}

// ToPageOfOpening
// @Summary 获取可用页面列表
// @Description 获取所有可用的页面列表，默认返回最新的 20 个
// @Tags 站点-页面
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {array} response.Opening[uint] "页面列表"
// @Router /site/page/opening [get]
func ToPageOfOpening(c context.Context, ctx *app.RequestContext) {

	var pages []model.ShPage

	facades.Database().Default().WithContext(c).Order("`id` desc").Limit(20).Find(&pages)

	responses := make([]response.Opening[uint], len(pages))

	for index, item := range pages {
		responses[index] = response.Opening[uint]{
			ID:   item.ID,
			Name: item.Title,
		}
	}

	http.Success(ctx, responses)
}

// DoPageOfCreate
// @Summary 创建页面
// @Description Permissions: site.page.create
// @Tags 站点-页面
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param request body req.DoPageOfCreate true "页面信息"
// @Success 200 {object} nil "创建成功"
// @Router /site/page [post]
func DoPageOfCreate(c context.Context, ctx *app.RequestContext) {

	var request req.DoPageOfCreate

	if err := ctx.BindAndValidate(&request); err != nil {
		http.BadRequest(ctx, err)
		return
	}

	page := model.ShPage{
		Title:   request.Title,
		Content: request.Content,
	}

	if result := facades.Database().Default().WithContext(c).Create(&page); result.Error != nil {
		http.Fail(ctx, "创建失败：%v", result.Error)
		return
	}

	http.Success[any](ctx)
}

// DoPageOfUpdate
// @Summary 更新页面
// @Description Permissions: site.page.update
// @Tags 站点-页面
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path int true "页面ID"
// @Param request body req.DoPageOfUpdate true "页面信息"
// @Success 200 {object} nil "更新成功"
// @Router /site/pages/{id} [put]
func DoPageOfUpdate(c context.Context, ctx *app.RequestContext) {

	var request req.DoPageOfUpdate

	if err := ctx.BindAndValidate(&request); err != nil {
		http.BadRequest(ctx, err)
		return
	}

	var page model.ShPage

	if err := firstByID(c, &page, request.ID); err != nil {
		writeFindError(ctx, err)
		return
	}

	page.Title = request.Title
	page.Content = request.Content

	if result := facades.Database().Default().WithContext(c).Save(&page); result.Error != nil {
		http.Fail(ctx, "修改失败：%v", result.Error)
		return
	}

	http.Success[any](ctx)
}

// DoPageOfDelete
// @Summary 删除页面
// @Description Permissions: site.page.delete
// @Tags 站点-页面
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path int true "页面ID"
// @Success 200 {object} nil "删除成功"
// @Router /site/pages/{id} [delete]
func DoPageOfDelete(c context.Context, ctx *app.RequestContext) {

	var request req.DoPageOfDelete

	if err := ctx.BindAndValidate(&request); err != nil {
		http.BadRequest(ctx, err)
		return
	}

	var page model.ShPage

	if err := firstByID(c, &page, request.ID); err != nil {
		writeFindError(ctx, err)
		return
	}

	if result := facades.Database().Default().WithContext(c).Delete(&page); result.Error != nil {
		http.Fail(ctx, "删除失败：%v", result.Error)
		return
	}

	http.Success[any](ctx)
}
