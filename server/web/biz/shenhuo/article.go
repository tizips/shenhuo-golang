package shenhuo

import (
	"context"
	"errors"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/herhe-com/framework/constants/global"
	"github.com/herhe-com/framework/contracts/http/response"
	"github.com/herhe-com/framework/facades"
	"github.com/herhe-com/framework/http"
	"github.com/tizips/shenhuo/model"
	req "github.com/tizips/shenhuo/server/web/http/request/shenhuo"
	res "github.com/tizips/shenhuo/server/web/http/response/shenhuo"
	"gorm.io/gorm"
)

// ToArticleOfPinned
// @Summary 获取置顶资讯
// @Description 获取已发布的置顶资讯
// @Tags 资讯
// @Accept json
// @Produce json
// @Success 200 {array} res.ToArticle "置顶资讯"
// @Router /article/pinned [get]
func ToArticleOfPinned(c context.Context, ctx *app.RequestContext) {

	var articles []model.ShArticle

	database().Where("`is_top`=?", global.YES).Order("`published_at` desc, `id` desc").Find(&articles)

	responses := make([]res.ToArticle, len(articles))
	for index, item := range articles {
		responses[index] = articleToList(item)
	}

	http.Success(ctx, responses)
}

// ToArticleOfRecommend
// @Summary 获取首页推荐资讯
// @Description 获取已推荐的资讯列表
// @Tags 资讯
// @Accept json
// @Produce json
// @Success 200 {array} res.ToArticle "推荐资讯"
// @Router /article/recommend [get]
func ToArticleOfRecommend(c context.Context, ctx *app.RequestContext) {

	var articles []model.ShArticle

	database().Where("`is_recommend`=?", global.YES).Order("`published_at` desc, `id` desc").Find(&articles)

	responses := make([]res.ToArticle, len(articles))
	for index, item := range articles {
		responses[index] = articleToList(item)
	}

	http.Success(ctx, responses)
}

// ToArticleOfPaginate
// @Summary 获取资讯列表
// @Description 获取已发布资讯分页列表
// @Tags 资讯
// @Accept json
// @Produce json
// @Param page query int false "页码"
// @Param size query int false "每页数量"
// @Success 200 {object} response.Paginate[res.ToArticle] "资讯列表"
// @Router /articles [get]
func ToArticleOfPaginate(c context.Context, ctx *app.RequestContext) {

	var request req.ToArticleOfPaginate

	if err := ctx.BindAndValidate(&request); err != nil {
		http.BadRequest(ctx, err)
		return
	}

	responses := response.Paginate[res.ToArticle]{
		Page: request.GetPage(),
		Size: request.GetSize(),
	}

	tx := database().Where("`is_top`=?", global.NO)

	tx.Count(&responses.Total)

	if responses.Total > 0 {

		var articles []model.ShArticle

		tx.Order("`published_at` desc, `id` desc").Offset(request.GetOffset()).Limit(request.GetLimit()).Find(&articles)

		responses.Data = make([]res.ToArticle, len(articles))
		for index, item := range articles {
			responses.Data[index] = articleToList(item)
		}
	}

	http.Success(ctx, responses)
}

// ToArticleOfInformation
// @Summary 获取资讯详情
// @Description 获取已发布资讯详情
// @Tags 资讯
// @Accept json
// @Produce json
// @Param id path int true "资讯ID"
// @Success 200 {object} res.ToArticle "资讯详情"
// @Router /articles/{id} [get]
func ToArticleOfInformation(c context.Context, ctx *app.RequestContext) {

	var request req.ToArticleOfInformation

	if err := ctx.BindAndValidate(&request); err != nil {
		http.BadRequest(ctx, err)
		return
	}

	var article model.ShArticle

	fu := database().First(&article, "`id`=?", request.ID)
	if errors.Is(fu.Error, gorm.ErrRecordNotFound) {
		http.NotFound(ctx, "未找到该数据")
		return
	} else if fu.Error != nil {
		http.Fail(ctx, "查询失败：%v", fu.Error)
		return
	}

	item := articleToList(article)
	item.Content = article.Content

	http.Success(ctx, item)
}

func database() *gorm.DB {
	return facades.Database().Default().Model(&model.ShArticle{})
}

func articleToList(item model.ShArticle) res.ToArticle {
	return res.ToArticle{
		ID:          item.ID,
		Title:       item.Title,
		Thumb:       item.Thumb,
		PublishedAt: item.PublishedAt.ToDateTimeString(),
		IsTop:       item.IsTop,
		IsRecommend: item.IsRecommend,
	}
}
