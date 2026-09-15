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

// ToSceneOfPaginate
// @Summary 获取场景列表
// @Description Permissions: site.scene.paginate
// @Tags 站点-场景
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param page query int false "页码"
// @Param size query int false "每页数量"
// @Success 200 {object} response.Paginate[res.ToSceneOfPaginate] "场景列表"
// @Router /site/scenes [get]
func ToSceneOfPaginate(c context.Context, ctx *app.RequestContext) {

	var request req.ToSceneOfPaginate

	if err := ctx.BindAndValidate(&request); err != nil {
		http.BadRequest(ctx, err)
		return
	}

	responses := response.Paginate[res.ToSceneOfPaginate]{
		Page: request.GetPage(),
		Size: request.GetSize(),
	}

	tx := facades.Database().Default().WithContext(c).Model(&model.ShScene{})

	tx.Count(&responses.Total)

	if responses.Total > 0 {

		var scenes []model.ShScene

		tx.Order("`order` asc, `id` asc").Offset(request.GetOffset()).Limit(request.GetLimit()).Find(&scenes)

		responses.Data = make([]res.ToSceneOfPaginate, len(scenes))

		for index, item := range scenes {
			responses.Data[index] = res.ToSceneOfPaginate{
				ID:        item.ID,
				Name:      item.Name,
				Order:     item.Order,
				CreatedAt: item.CreatedAt.ToDateTimeString(),
			}
		}
	}

	http.Success(ctx, responses)
}

// ToSceneOfOpening
// @Summary 获取可用场景列表
// @Description 获取所有可用的场景列表，默认返回最新的 20 个
// @Tags 站点-场景
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {array} response.Opening[uint] "场景列表"
// @Router /site/scene/opening [get]
func ToSceneOfOpening(c context.Context, ctx *app.RequestContext) {

	var scenes []model.ShScene

	facades.Database().Default().WithContext(c).Order("`id` desc").Limit(20).Find(&scenes)

	responses := make([]response.Opening[uint], len(scenes))

	for index, item := range scenes {
		responses[index] = response.Opening[uint]{
			ID:   item.ID,
			Name: item.Name,
		}
	}

	http.Success(ctx, responses)
}

// DoSceneOfCreate
// @Summary 创建场景
// @Description Permissions: site.scene.create
// @Tags 站点-场景
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param request body req.DoSceneOfCreate true "场景信息"
// @Success 200 {object} nil "创建成功"
// @Router /site/scene [post]
func DoSceneOfCreate(c context.Context, ctx *app.RequestContext) {

	var request req.DoSceneOfCreate

	if err := ctx.BindAndValidate(&request); err != nil {
		http.BadRequest(ctx, err)
		return
	}

	scene := model.ShScene{Name: request.Name, Order: helper.Order(request.Order)}

	if result := facades.Database().Default().WithContext(c).Create(&scene); result.Error != nil {
		http.Fail(ctx, "创建失败：%v", result.Error)
		return
	}

	http.Success[any](ctx)
}

// DoSceneOfUpdate
// @Summary 更新场景
// @Description Permissions: site.scene.update
// @Tags 站点-场景
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path int true "场景ID"
// @Param request body req.DoSceneOfUpdate true "场景信息"
// @Success 200 {object} nil "更新成功"
// @Router /site/scenes/{id} [put]
func DoSceneOfUpdate(c context.Context, ctx *app.RequestContext) {

	var request req.DoSceneOfUpdate

	if err := ctx.BindAndValidate(&request); err != nil {
		http.BadRequest(ctx, err)
		return
	}

	var scene model.ShScene

	if err := firstByID(c, &scene, request.ID); err != nil {
		writeFindError(ctx, err)
		return
	}

	scene.Name = request.Name
	scene.Order = helper.Order(request.Order)

	if result := facades.Database().Default().WithContext(c).Save(&scene); result.Error != nil {
		http.Fail(ctx, "修改失败：%v", result.Error)
		return
	}

	http.Success[any](ctx)
}

// DoSceneOfDelete
// @Summary 删除场景
// @Description Permissions: site.scene.delete
// @Tags 站点-场景
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path int true "场景ID"
// @Success 200 {object} nil "删除成功"
// @Router /site/scenes/{id} [delete]
func DoSceneOfDelete(c context.Context, ctx *app.RequestContext) {

	var request req.DoSceneOfDelete

	if err := ctx.BindAndValidate(&request); err != nil {
		http.BadRequest(ctx, err)
		return
	}

	var scene model.ShScene

	if err := firstByID(c, &scene, request.ID); err != nil {
		writeFindError(ctx, err)
		return
	}

	tx := facades.Database().Default().WithContext(c).Begin()

	if result := tx.Where("`scene_id`=?", scene.ID).Delete(&model.ShMedia{}); result.Error != nil {
		tx.Rollback()
		http.Fail(ctx, "删除失败：%v", result.Error)
		return
	}

	if result := tx.Delete(&scene); result.Error != nil {
		tx.Rollback()
		http.Fail(ctx, "删除失败：%v", result.Error)
		return
	}

	tx.Commit()

	http.Success[any](ctx)
}
