package shenhuo

import (
	"context"
	"strings"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/herhe-com/framework/constants/global"
	"github.com/herhe-com/framework/contracts/http/response"
	"github.com/herhe-com/framework/facades"
	"github.com/herhe-com/framework/http"
	"github.com/tizips/shenhuo/model"
	req "github.com/tizips/shenhuo/server/web/http/request/shenhuo"
	res "github.com/tizips/shenhuo/server/web/http/response/shenhuo"
)

// ToMediaOfPaginate
// @Summary 获取媒体列表
// @Description 获取精彩图片或视频分页列表
// @Tags 媒体
// @Accept json
// @Produce json
// @Param page query int false "页码"
// @Param size query int false "每页数量"
// @Param scene_id query int false "场景ID"
// @Param type query string false "媒体类型；枚举：image=图片，video=视频"
// @Success 200 {object} response.Paginate[res.ToMedia] "媒体列表"
// @Router /medias [get]
func ToMediaOfPaginate(c context.Context, ctx *app.RequestContext) {

	var request req.ToMediaOfPaginate

	if err := ctx.BindAndValidate(&request); err != nil {
		http.BadRequest(ctx, err)
		return
	}

	responses := response.Paginate[res.ToMedia]{
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

		tx.Order("`is_top` desc, `id` desc").Offset(request.GetOffset()).Limit(request.GetLimit()).Find(&medias)

		responses.Data = make([]res.ToMedia, len(medias))
		for index, item := range medias {
			responses.Data[index] = mediaToResponse(item)
		}
	}

	http.Success(ctx, responses)
}

// DoMediaOfCreateByImage
// @Summary 批量添加图片
// @Description 管理人员批量添加精彩图片
// @Tags 媒体
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param request body req.DoMediaOfCreateByImage true "图片信息"
// @Success 200 {array} res.ToMedia "创建成功"
// @Router /media/images [post]
func DoMediaOfCreateByImage(c context.Context, ctx *app.RequestContext) {

	var request req.DoMediaOfCreateByImage

	if err := ctx.BindAndValidate(&request); err != nil {
		http.BadRequest(ctx, err)
		return
	}

	var total int64
	facades.Database().Default().WithContext(c).Model(&model.ShScene{}).Where("`id`=?", request.SceneID).Count(&total)
	if total == 0 {
		http.Fail(ctx, "场景不存在")
		return
	}

	medias := make([]model.ShMedia, len(request.URLs))
	for index, url := range request.URLs {
		medias[index] = model.ShMedia{
			SceneID: request.SceneID,
			Type:    model.ShMediaOfTypeImage,
			URL:     url,
			IsTop:   global.NO,
		}
	}

	if result := facades.Database().Default().WithContext(c).Create(&medias); result.Error != nil {
		http.Fail(ctx, "创建失败：%v", result.Error)
		return
	}

	responses := make([]res.ToMedia, len(medias))
	for index, item := range medias {
		responses[index] = mediaToResponse(item)
	}

	http.Success(ctx, responses)
}

// DoMediaOfCreateByVideo
// @Summary 添加视频
// @Description 管理人员添加精彩视频，标题必填
// @Tags 媒体
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param request body req.DoMediaOfCreateByVideo true "视频信息"
// @Success 200 {object} res.ToMedia "创建成功"
// @Router /media/video [post]
func DoMediaOfCreateByVideo(c context.Context, ctx *app.RequestContext) {

	var request req.DoMediaOfCreateByVideo

	if err := ctx.BindAndValidate(&request); err != nil {
		http.BadRequest(ctx, err)
		return
	}

	if strings.TrimSpace(request.Title) == "" {
		http.Fail(ctx, "视频标题不能为空")
		return
	}

	var total int64
	facades.Database().Default().WithContext(c).Model(&model.ShScene{}).Where("`id`=?", request.SceneID).Count(&total)
	if total == 0 {
		http.Fail(ctx, "场景不存在")
		return
	}

	media := model.ShMedia{
		SceneID: request.SceneID,
		Type:    model.ShMediaOfTypeVideo,
		Title:   strings.TrimSpace(request.Title),
		URL:     request.URL,
		IsTop:   global.NO,
	}

	if result := facades.Database().Default().WithContext(c).Create(&media); result.Error != nil {
		http.Fail(ctx, "创建失败：%v", result.Error)
		return
	}

	http.Success(ctx, mediaToResponse(media))
}

// ToMediaOfPinned
// @Summary 获取置顶媒体
// @Description 获取置顶的精彩图片或视频
// @Tags 媒体
// @Accept json
// @Produce json
// @Success 200 {array} res.ToMedia "置顶媒体"
// @Router /media/pinned [get]
func ToMediaOfPinned(c context.Context, ctx *app.RequestContext) {

	var medias []model.ShMedia

	facades.Database().Default().WithContext(c).Where("`is_top`=?", global.YES).Order("`id` desc").Find(&medias)

	responses := make([]res.ToMedia, len(medias))
	for index, item := range medias {
		responses[index] = mediaToResponse(item)
	}

	http.Success(ctx, responses)
}

func mediaToResponse(item model.ShMedia) res.ToMedia {
	return res.ToMedia{
		ID:      item.ID,
		SceneID: item.SceneID,
		Type:    item.Type,
		Title:   item.Title,
		URL:     item.URL,
		IsTop:   item.IsTop,
	}
}
