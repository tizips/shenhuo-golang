package site

import (
	"context"
	"strings"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/herhe-com/framework/contracts/http/response"
	"github.com/herhe-com/framework/facades"
	"github.com/herhe-com/framework/http"
	"github.com/tizips/shenhuo/model"
	req "github.com/tizips/shenhuo/server/admin/http/request/site"
	res "github.com/tizips/shenhuo/server/admin/http/response/site"
)

// ToMediaOfPaginate
// @Summary 获取媒体列表
// @Description Permissions: site.media.paginate
// @Tags 站点-媒体
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param page query int false "页码"
// @Param size query int false "每页数量"
// @Param scene_id query int false "场景"
// @Param type query string false "类型：image=图片，video=视频" Enums(image,video)
// @Success 200 {object} response.Paginate[res.ToMediaOfPaginate] "媒体列表"
// @Router /site/medias [get]
func ToMediaOfPaginate(c context.Context, ctx *app.RequestContext) {

	var request req.ToMediaOfPaginate

	if err := ctx.BindAndValidate(&request); err != nil {
		http.BadRequest(ctx, err)
		return
	}

	responses := response.Paginate[res.ToMediaOfPaginate]{
		Page: request.GetPage(),
		Size: request.GetSize(),
	}

	tx := facades.Database().Default().WithContext(c).Model(&model.ShMedia{})

	if request.SceneID > 0 {
		tx = tx.Where("`scene_id`=?", request.SceneID)
	}

	if request.Type != "" {
		tx = tx.Where("`type`=?", request.Type)
	}

	tx.Count(&responses.Total)

	if responses.Total > 0 {

		var medias []model.ShMedia

		tx.Preload("Scene").Order("`id` desc").Offset(request.GetOffset()).Limit(request.GetLimit()).Find(&medias)

		responses.Data = make([]res.ToMediaOfPaginate, len(medias))

		for index, item := range medias {
			responses.Data[index] = mediaToResponse(item)
		}
	}

	http.Success(ctx, responses)
}

// DoMediaOfCreate
// @Summary 创建媒体
// @Description Permissions: site.media.create
// @Tags 站点-媒体
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param request body req.DoMediaOfCreate true "媒体信息"
// @Success 200 {object} nil "创建成功"
// @Router /site/media [post]
func DoMediaOfCreate(c context.Context, ctx *app.RequestContext) {

	var request req.DoMediaOfCreate

	if err := ctx.BindAndValidate(&request); err != nil {
		http.BadRequest(ctx, err)
		return
	}

	if request.Type == model.ShMediaOfTypeVideo && strings.TrimSpace(request.Title) == "" {
		http.Fail(ctx, "视频标题不能为空")
		return
	}

	if !sceneExists(c, request.SceneID) {
		http.Fail(ctx, "场景不存在")
		return
	}

	media := model.ShMedia{
		SceneID: request.SceneID,
		Type:    request.Type,
		Title:   strings.TrimSpace(request.Title),
		URL:     request.URL,
		IsTop:   request.IsTop,
	}

	if result := facades.Database().Default().WithContext(c).Create(&media); result.Error != nil {
		http.Fail(ctx, "创建失败：%v", result.Error)
		return
	}

	http.Success[any](ctx)
}

// DoMediaOfUpdate
// @Summary 更新媒体
// @Description Permissions: site.media.update
// @Tags 站点-媒体
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path int true "媒体ID"
// @Param request body req.DoMediaOfUpdate true "媒体信息"
// @Success 200 {object} nil "更新成功"
// @Router /site/medias/{id} [put]
func DoMediaOfUpdate(c context.Context, ctx *app.RequestContext) {

	var request req.DoMediaOfUpdate

	if err := ctx.BindAndValidate(&request); err != nil {
		http.BadRequest(ctx, err)
		return
	}

	if !sceneExists(c, request.SceneID) {
		http.Fail(ctx, "场景不存在")
		return
	}

	var media model.ShMedia

	if err := firstByID(c, &media, request.ID); err != nil {
		writeFindError(ctx, err)
		return
	}

	if request.Type == model.ShMediaOfTypeVideo && strings.TrimSpace(request.Title) == "" {
		http.Fail(ctx, "视频标题不能为空")
		return
	}

	media.SceneID = request.SceneID
	media.Type = request.Type
	media.Title = strings.TrimSpace(request.Title)
	media.URL = request.URL
	media.IsTop = request.IsTop

	if result := facades.Database().Default().WithContext(c).Save(&media); result.Error != nil {
		http.Fail(ctx, "修改失败：%v", result.Error)
		return
	}

	http.Success[any](ctx)
}

// DoMediaOfDelete
// @Summary 删除媒体
// @Description Permissions: site.media.delete
// @Tags 站点-媒体
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path int true "媒体ID"
// @Success 200 {object} nil "删除成功"
// @Router /site/medias/{id} [delete]
func DoMediaOfDelete(c context.Context, ctx *app.RequestContext) {

	var request req.DoMediaOfDelete

	if err := ctx.BindAndValidate(&request); err != nil {
		http.BadRequest(ctx, err)
		return
	}

	var media model.ShMedia

	if err := firstByID(c, &media, request.ID); err != nil {
		writeFindError(ctx, err)
		return
	}

	if result := facades.Database().Default().WithContext(c).Delete(&media); result.Error != nil {
		http.Fail(ctx, "删除失败：%v", result.Error)
		return
	}

	http.Success[any](ctx)
}

func sceneExists(c context.Context, id uint) bool {
	var total int64
	facades.Database().Default().WithContext(c).Model(&model.ShScene{}).Where("`id`=?", id).Count(&total)
	return total > 0
}

func mediaToResponse(item model.ShMedia) res.ToMediaOfPaginate {
	row := res.ToMediaOfPaginate{
		ID:        item.ID,
		SceneID:   item.SceneID,
		Type:      item.Type,
		Title:     item.Title,
		URL:       item.URL,
		IsTop:     item.IsTop,
		CreatedAt: item.CreatedAt.ToDateTimeString(),
	}

	if item.Scene != nil {
		row.Scene = item.Scene.Name
	}

	return row
}
