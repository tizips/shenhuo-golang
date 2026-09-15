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

// ToPageBuiltinOfPaginate
// @Summary 获取内置页面列表
// @Description Permissions: site.page_builtin.paginate
// @Tags 站点-内置页面
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param page query int false "页码"
// @Param size query int false "每页数量"
// @Param keyword query string false "关键词"
// @Success 200 {object} response.Paginate[res.ToPageBuiltinOfPaginate] "内置页面列表"
// @Router /site/page-builtins [get]
func ToPageBuiltinOfPaginate(c context.Context, ctx *app.RequestContext) {

	var request req.ToPageBuiltinOfPaginate

	if err := ctx.BindAndValidate(&request); err != nil {
		http.BadRequest(ctx, err)
		return
	}

	responses := response.Paginate[res.ToPageBuiltinOfPaginate]{
		Page: request.GetPage(),
		Size: request.GetSize(),
	}

	tx := facades.Database().Default().WithContext(c).Model(&model.ShPageBuiltin{})

	if request.Keyword != "" {
		tx = tx.Where("`key` LIKE ?", "%"+request.Keyword+"%")
	}

	tx.Count(&responses.Total)

	if responses.Total > 0 {

		var builtins []model.ShPageBuiltin

		tx.Order("`id` desc").Offset(request.GetOffset()).Limit(request.GetLimit()).Find(&builtins)

		ids := make([]uint, 0, len(builtins))
		for _, item := range builtins {
			ids = append(ids, item.PageID)
		}

		titles := make(map[uint]string, len(builtins))
		if len(ids) > 0 {
			var pages []model.ShPage
			facades.Database().Default().WithContext(c).Where("`id` IN ?", ids).Find(&pages)
			for _, item := range pages {
				titles[item.ID] = item.Title
			}
		}

		responses.Data = make([]res.ToPageBuiltinOfPaginate, len(builtins))

		for index, item := range builtins {
			responses.Data[index] = res.ToPageBuiltinOfPaginate{
				ID:        item.ID,
				Key:       item.Key,
				PageID:    item.PageID,
				PageTitle: titles[item.PageID],
				CreatedAt: item.CreatedAt.ToDateTimeString(),
			}
		}
	}

	http.Success(ctx, responses)
}

// DoPageBuiltinOfCreate
// @Summary 创建内置页面
// @Description Permissions: site.page_builtin.create
// @Tags 站点-内置页面
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param request body req.DoPageBuiltinOfCreate true "内置页面信息"
// @Success 200 {object} nil "创建成功"
// @Router /site/page-builtin [post]
func DoPageBuiltinOfCreate(c context.Context, ctx *app.RequestContext) {

	var request req.DoPageBuiltinOfCreate

	if err := ctx.BindAndValidate(&request); err != nil {
		http.BadRequest(ctx, err)
		return
	}

	if existsPageBuiltinOfKey(c, request.Key, 0) {
		http.Fail(ctx, "标识已存在")
		return
	}

	if !existsPageByID(c, request.PageID) {
		http.Fail(ctx, "页面不存在")
		return
	}

	builtin := model.ShPageBuiltin{
		Key:    request.Key,
		PageID: request.PageID,
	}

	if result := facades.Database().Default().WithContext(c).Create(&builtin); result.Error != nil {
		http.Fail(ctx, "创建失败：%v", result.Error)
		return
	}

	http.Success[any](ctx)
}

// DoPageBuiltinOfUpdate
// @Summary 更新内置页面
// @Description Permissions: site.page_builtin.update
// @Tags 站点-内置页面
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path int true "内置页面ID"
// @Param request body req.DoPageBuiltinOfUpdate true "内置页面信息"
// @Success 200 {object} nil "更新成功"
// @Router /site/page-builtins/{id} [put]
func DoPageBuiltinOfUpdate(c context.Context, ctx *app.RequestContext) {

	var request req.DoPageBuiltinOfUpdate

	if err := ctx.BindAndValidate(&request); err != nil {
		http.BadRequest(ctx, err)
		return
	}

	var builtin model.ShPageBuiltin

	if err := firstByID(c, &builtin, request.ID); err != nil {
		writeFindError(ctx, err)
		return
	}

	if existsPageBuiltinOfKey(c, request.Key, request.ID) {
		http.Fail(ctx, "标识已存在")
		return
	}

	if !existsPageByID(c, request.PageID) {
		http.Fail(ctx, "页面不存在")
		return
	}

	builtin.Key = request.Key
	builtin.PageID = request.PageID

	if result := facades.Database().Default().WithContext(c).Save(&builtin); result.Error != nil {
		http.Fail(ctx, "修改失败：%v", result.Error)
		return
	}

	http.Success[any](ctx)
}

// DoPageBuiltinOfDelete
// @Summary 删除内置页面
// @Description Permissions: site.page_builtin.delete
// @Tags 站点-内置页面
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path int true "内置页面ID"
// @Success 200 {object} nil "删除成功"
// @Router /site/page-builtins/{id} [delete]
func DoPageBuiltinOfDelete(c context.Context, ctx *app.RequestContext) {

	var request req.DoPageBuiltinOfDelete

	if err := ctx.BindAndValidate(&request); err != nil {
		http.BadRequest(ctx, err)
		return
	}

	var builtin model.ShPageBuiltin

	if err := firstByID(c, &builtin, request.ID); err != nil {
		writeFindError(ctx, err)
		return
	}

	if result := facades.Database().Default().WithContext(c).Delete(&builtin); result.Error != nil {
		http.Fail(ctx, "删除失败：%v", result.Error)
		return
	}

	http.Success[any](ctx)
}

func existsPageBuiltinOfKey(c context.Context, key string, excludeID uint) bool {

	tx := facades.Database().Default().WithContext(c).Model(&model.ShPageBuiltin{}).Where("`key`=?", key)

	if excludeID > 0 {
		tx = tx.Where("`id`<>?", excludeID)
	}

	var total int64
	tx.Count(&total)

	return total > 0
}

func existsPageByID(c context.Context, id uint) bool {

	var total int64
	facades.Database().Default().WithContext(c).Model(&model.ShPage{}).Where("`id`=?", id).Count(&total)

	return total > 0
}
