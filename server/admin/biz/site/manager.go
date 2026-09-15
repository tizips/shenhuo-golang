package site

import (
	"context"
	"errors"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/herhe-com/framework/auth"
	"github.com/herhe-com/framework/contracts/http/response"
	"github.com/herhe-com/framework/facades"
	"github.com/herhe-com/framework/http"
	"github.com/tizips/shenhuo/model"
	req "github.com/tizips/shenhuo/server/admin/http/request/site"
	res "github.com/tizips/shenhuo/server/admin/http/response/site"
)

// ToManagerOfPaginate
// @Summary 获取管理人员列表
// @Description Permissions: site.manager.paginate
// @Tags 站点-管理人员
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param page query int false "页码"
// @Param size query int false "每页数量"
// @Param keyword query string false "关键词"
// @Success 200 {object} response.Paginate[res.ToManagerOfPaginate] "管理人员列表"
// @Router /site/managers [get]
func ToManagerOfPaginate(c context.Context, ctx *app.RequestContext) {

	var request req.ToManagerOfPaginate

	if err := ctx.BindAndValidate(&request); err != nil {
		http.BadRequest(ctx, err)
		return
	}

	responses := response.Paginate[res.ToManagerOfPaginate]{
		Page: request.GetPage(),
		Size: request.GetSize(),
	}

	tx := facades.Database().Default().WithContext(c).Model(&model.ShManager{})

	if request.Keyword != "" {
		like := "%" + request.Keyword + "%"
		tx = tx.Where("`name` LIKE ? OR `mobile` LIKE ?", like, like)
	}

	tx.Count(&responses.Total)

	if responses.Total > 0 {

		var managers []model.ShManager

		tx.Order("`id` desc").Offset(request.GetOffset()).Limit(request.GetLimit()).Find(&managers)

		responses.Data = make([]res.ToManagerOfPaginate, len(managers))

		for index, item := range managers {
			responses.Data[index] = res.ToManagerOfPaginate{
				ID:        item.ID,
				Name:      item.Name,
				Mobile:    item.Mobile,
				IsEnable:  item.IsEnable,
				CreatedAt: item.CreatedAt.ToDateTimeString(),
			}
		}
	}

	http.Success(ctx, responses)
}

// DoManagerOfCreate
// @Summary 创建管理人员
// @Description Permissions: site.manager.create
// @Tags 站点-管理人员
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param request body req.DoManagerOfCreate true "管理人员信息"
// @Success 200 {object} nil "创建成功"
// @Router /site/manager [post]
func DoManagerOfCreate(c context.Context, ctx *app.RequestContext) {

	var request req.DoManagerOfCreate

	if err := ctx.BindAndValidate(&request); err != nil {
		http.BadRequest(ctx, err)
		return
	}

	if err := assertManagerUnique(c, request.Mobile, ""); err != nil {
		http.Fail(ctx, "%s", err.Error())
		return
	}

	manager := model.ShManager{
		ID:       facades.Snowflake().Generate().String(),
		Name:     request.Name,
		Mobile:   request.Mobile,
		Password: auth.Password(request.Password),
		IsEnable: request.IsEnable,
	}

	if result := facades.Database().Default().WithContext(c).Create(&manager); result.Error != nil {
		http.Fail(ctx, "创建失败：%v", result.Error)
		return
	}

	http.Success[any](ctx)
}

// DoManagerOfUpdate
// @Summary 更新管理人员
// @Description Permissions: site.manager.update
// @Tags 站点-管理人员
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path string true "管理人员ID"
// @Param request body req.DoManagerOfUpdate true "管理人员信息"
// @Success 200 {object} nil "更新成功"
// @Router /site/managers/{id} [put]
func DoManagerOfUpdate(c context.Context, ctx *app.RequestContext) {

	var request req.DoManagerOfUpdate

	if err := ctx.BindAndValidate(&request); err != nil {
		http.BadRequest(ctx, err)
		return
	}

	var manager model.ShManager

	if err := firstBySID(c, &manager, request.ID); err != nil {
		writeFindError(ctx, err)
		return
	}

	if err := assertManagerUnique(c, request.Mobile, manager.ID); err != nil {
		http.Fail(ctx, "%s", err.Error())
		return
	}

	manager.Name = request.Name
	manager.Mobile = request.Mobile
	manager.IsEnable = request.IsEnable

	if request.Password != "" {
		manager.Password = auth.Password(request.Password)
	}

	if result := facades.Database().Default().WithContext(c).Save(&manager); result.Error != nil {
		http.Fail(ctx, "修改失败：%v", result.Error)
		return
	}

	http.Success[any](ctx)
}

// DoManagerOfDelete
// @Summary 删除管理人员
// @Description Permissions: site.manager.delete
// @Tags 站点-管理人员
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path string true "管理人员ID"
// @Success 200 {object} nil "删除成功"
// @Router /site/managers/{id} [delete]
func DoManagerOfDelete(c context.Context, ctx *app.RequestContext) {

	var request req.DoManagerOfDelete

	if err := ctx.BindAndValidate(&request); err != nil {
		http.BadRequest(ctx, err)
		return
	}

	var manager model.ShManager

	if err := firstBySID(c, &manager, request.ID); err != nil {
		writeFindError(ctx, err)
		return
	}

	if result := facades.Database().Default().WithContext(c).Delete(&manager); result.Error != nil {
		http.Fail(ctx, "删除失败：%v", result.Error)
		return
	}

	http.Success[any](ctx)
}

// DoManagerOfEnable
// @Summary 启用/禁用管理人员
// @Description Permissions: site.manager.enable
// @Tags 站点-管理人员
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param request body req.DoManagerOfEnable true "启用信息"
// @Success 200 {object} nil "操作成功"
// @Router /site/manager/enable [put]
func DoManagerOfEnable(c context.Context, ctx *app.RequestContext) {

	var request req.DoManagerOfEnable

	if err := ctx.BindAndValidate(&request); err != nil {
		http.BadRequest(ctx, err)
		return
	}

	var manager model.ShManager

	if err := firstBySID(c, &manager, request.ID); err != nil {
		writeFindError(ctx, err)
		return
	}

	if result := facades.Database().Default().WithContext(c).Model(&manager).Update("is_enable", request.IsEnable); result.Error != nil {
		http.Fail(ctx, "启禁失败：%v", result.Error)
		return
	}

	http.Success[any](ctx)
}

func assertManagerUnique(c context.Context, mobile, exclude string) error {
	var total int64

	q := facades.Database().Default().WithContext(c).Model(&model.ShManager{}).Where("`mobile`=?", mobile)
	if exclude != "" {
		q = q.Where("`id`<>?", exclude)
	}
	q.Count(&total)

	if total > 0 {
		return errors.New("手机号已存在")
	}

	return nil
}
