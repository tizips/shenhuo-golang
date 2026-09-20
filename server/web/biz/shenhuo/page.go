package shenhuo

import (
	"context"
	"errors"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/herhe-com/framework/facades"
	"github.com/herhe-com/framework/http"
	"github.com/tizips/shenhuo/model"
	req "github.com/tizips/shenhuo/server/web/http/request/shenhuo"
	res "github.com/tizips/shenhuo/server/web/http/response/shenhuo"
	"gorm.io/gorm"
)

// ToPageOfList
// @Summary 获取页面列表
// @Description 获取页面列表
// @Tags 页面
// @Accept json
// @Produce json
// @Success 200 {array} res.ToPage "页面列表"
// @Router /pages [get]
func ToPageOfList(c context.Context, ctx *app.RequestContext) {

	var pages []model.ShPage

	facades.Database().Default().WithContext(c).Order("`id` desc").Find(&pages)

	responses := make([]res.ToPage, len(pages))
	for index, item := range pages {
		responses[index] = res.ToPage{
			ID:    item.ID,
			Title: item.Title,
		}
	}

	http.Success(ctx, responses)
}

// ToPageOfInformation
// @Summary 获取页面详情
// @Description 获取指定页面详情
// @Tags 页面
// @Accept json
// @Produce json
// @Param id path int true "页面ID"
// @Success 200 {object} res.ToPage "页面详情"
// @Router /pages/{id} [get]
func ToPageOfInformation(c context.Context, ctx *app.RequestContext) {

	var request req.ToPageOfInformation

	if err := ctx.BindAndValidate(&request); err != nil {
		http.BadRequest(ctx, err)
		return
	}

	var page model.ShPage

	fp := facades.Database().Default().WithContext(c).First(&page, "`id`=?", request.ID)
	if errors.Is(fp.Error, gorm.ErrRecordNotFound) {
		http.NotFound(ctx, "未找到该数据")
		return
	} else if fp.Error != nil {
		http.Fail(ctx, "查询失败：%v", fp.Error)
		return
	}

	http.Success(ctx, res.ToPage{
		ID:      page.ID,
		Title:   page.Title,
		Content: page.Content,
	})
}

// ToPageOfInformationByKey
// @Summary 按标识获取页面详情
// @Description 根据内置页面标识获取页面详情
// @Tags 页面
// @Accept json
// @Produce json
// @Param key path string true "页面标识"
// @Success 200 {object} res.ToPage "页面详情"
// @Router /page/keys/{key} [get]
func ToPageOfInformationByKey(c context.Context, ctx *app.RequestContext) {

	var request req.ToPageOfInformationByKey

	if err := ctx.BindAndValidate(&request); err != nil {
		http.BadRequest(ctx, err)
		return
	}

	var builtin model.ShPageBuiltin

	fp := facades.Database().Default().WithContext(c).First(&builtin, "`key`=?", request.Key)
	if errors.Is(fp.Error, gorm.ErrRecordNotFound) {
		http.NotFound(ctx, "未找到该数据")
		return
	} else if fp.Error != nil {
		http.Fail(ctx, "查询失败：%v", fp.Error)
		return
	}

	var page model.ShPage

	fp = facades.Database().Default().WithContext(c).First(&page, "`id`=?", builtin.PageID)
	if errors.Is(fp.Error, gorm.ErrRecordNotFound) {
		http.NotFound(ctx, "未找到该数据")
		return
	} else if fp.Error != nil {
		http.Fail(ctx, "查询失败：%v", fp.Error)
		return
	}

	http.Success(ctx, res.ToPage{
		ID:      page.ID,
		Title:   page.Title,
		Content: page.Content,
	})
}
