package site

import (
	"context"
	"errors"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/golang-module/carbon/v2"
	"github.com/herhe-com/framework/contracts/http/response"
	"github.com/herhe-com/framework/facades"
	"github.com/herhe-com/framework/http"
	"github.com/tizips/shenhuo/model"
	req "github.com/tizips/shenhuo/server/admin/http/request/site"
	res "github.com/tizips/shenhuo/server/admin/http/response/site"
	"gorm.io/gorm"
)

// ToArticleOfPaginate
// @Summary 获取资讯列表
// @Description Permissions: site.article.paginate
// @Tags 站点-资讯
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param page query int false "页码"
// @Param size query int false "每页数量"
// @Param keyword query string false "关键词"
// @Success 200 {object} response.Paginate[res.ToArticleOfPaginate] "资讯列表"
// @Router /site/articles [get]
func ToArticleOfPaginate(c context.Context, ctx *app.RequestContext) {

	var request req.ToArticleOfPaginate

	if err := ctx.BindAndValidate(&request); err != nil {
		http.BadRequest(ctx, err)
		return
	}

	responses := response.Paginate[res.ToArticleOfPaginate]{
		Page: request.GetPage(),
		Size: request.GetSize(),
	}

	tx := facades.Database().Default().WithContext(c).Model(&model.ShArticle{})

	if request.Keyword != "" {
		tx = tx.Where("`title` LIKE ?", "%"+request.Keyword+"%")
	}

	tx.Count(&responses.Total)

	if responses.Total > 0 {

		var articles []model.ShArticle

		tx.Order("`id` desc").Offset(request.GetOffset()).Limit(request.GetLimit()).Find(&articles)

		responses.Data = make([]res.ToArticleOfPaginate, len(articles))

		for index, item := range articles {
			responses.Data[index] = res.ToArticleOfPaginate{
				ID:          item.ID,
				Title:       item.Title,
				Thumb:       item.Thumb,
				PublishedAt: item.PublishedAt.ToDateTimeString(),
				IsTop:       item.IsTop,
				IsRecommend: item.IsRecommend,
				CreatedAt:   item.CreatedAt.ToDateTimeString(),
			}
		}
	}

	http.Success(ctx, responses)
}

// ToArticleOfInformation
// @Summary 获取资讯详情
// @Description 获取指定资讯详情
// @Tags 站点-资讯
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path int true "资讯ID"
// @Success 200 {object} res.ToArticleOfInformation "资讯详情"
// @Router /site/articles/{id} [get]
func ToArticleOfInformation(c context.Context, ctx *app.RequestContext) {

	var request req.ToArticleOfInformation

	if err := ctx.BindAndValidate(&request); err != nil {
		http.BadRequest(ctx, err)
		return
	}

	var article model.ShArticle

	if err := firstByID(c, &article, request.ID); err != nil {
		writeFindError(ctx, err)
		return
	}

	http.Success(ctx, res.ToArticleOfInformation{
		ID:          article.ID,
		Title:       article.Title,
		Thumb:       article.Thumb,
		Content:     article.Content,
		PublishedAt: article.PublishedAt.ToDateTimeString(),
		IsTop:       article.IsTop,
		IsRecommend: article.IsRecommend,
		CreatedAt:   article.CreatedAt.ToDateTimeString(),
	})
}

// DoArticleOfCreate
// @Summary 创建资讯
// @Description Permissions: site.article.create
// @Tags 站点-资讯
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param request body req.DoArticleOfCreate true "资讯信息"
// @Success 200 {object} nil "创建成功"
// @Router /site/article [post]
func DoArticleOfCreate(c context.Context, ctx *app.RequestContext) {

	var request req.DoArticleOfCreate

	if err := ctx.BindAndValidate(&request); err != nil {
		http.BadRequest(ctx, err)
		return
	}

	published, err := parseDateTime(request.PublishedAt)
	if err != nil {
		http.BadRequest(ctx, err)
		return
	}

	article := model.ShArticle{
		Title:       request.Title,
		Thumb:       request.Thumb,
		Content:     request.Content,
		PublishedAt: published,
		IsTop:       request.IsTop,
		IsRecommend: request.IsRecommend,
	}

	if result := facades.Database().Default().WithContext(c).Create(&article); result.Error != nil {
		http.Fail(ctx, "创建失败：%v", result.Error)
		return
	}

	http.Success[any](ctx)
}

// DoArticleOfUpdate
// @Summary 更新资讯
// @Description Permissions: site.article.update
// @Tags 站点-资讯
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path int true "资讯ID"
// @Param request body req.DoArticleOfUpdate true "资讯信息"
// @Success 200 {object} nil "更新成功"
// @Router /site/articles/{id} [put]
func DoArticleOfUpdate(c context.Context, ctx *app.RequestContext) {

	var request req.DoArticleOfUpdate

	if err := ctx.BindAndValidate(&request); err != nil {
		http.BadRequest(ctx, err)
		return
	}

	published, err := parseDateTime(request.PublishedAt)
	if err != nil {
		http.BadRequest(ctx, err)
		return
	}

	var article model.ShArticle

	if err = firstByID(c, &article, request.ID); err != nil {
		writeFindError(ctx, err)
		return
	}

	article.Title = request.Title
	article.Thumb = request.Thumb
	article.Content = request.Content
	article.PublishedAt = published
	article.IsTop = request.IsTop
	article.IsRecommend = request.IsRecommend

	if result := facades.Database().Default().WithContext(c).Save(&article); result.Error != nil {
		http.Fail(ctx, "修改失败：%v", result.Error)
		return
	}

	http.Success[any](ctx)
}

// DoArticleOfDelete
// @Summary 删除资讯
// @Description Permissions: site.article.delete
// @Tags 站点-资讯
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path int true "资讯ID"
// @Success 200 {object} nil "删除成功"
// @Router /site/articles/{id} [delete]
func DoArticleOfDelete(c context.Context, ctx *app.RequestContext) {

	var request req.DoArticleOfDelete

	if err := ctx.BindAndValidate(&request); err != nil {
		http.BadRequest(ctx, err)
		return
	}

	var article model.ShArticle

	if err := firstByID(c, &article, request.ID); err != nil {
		writeFindError(ctx, err)
		return
	}

	if result := facades.Database().Default().WithContext(c).Delete(&article); result.Error != nil {
		http.Fail(ctx, "删除失败：%v", result.Error)
		return
	}

	http.Success[any](ctx)
}

func parseDateTime(value string) (carbon.Carbon, error) {
	parsed := carbon.Parse(value)
	if parsed.Error != nil {
		return carbon.Carbon{}, errors.New("发布时间格式错误")
	}
	return parsed, nil
}

func firstByID(c context.Context, dest any, id uint) error {
	return facades.Database().Default().WithContext(c).First(dest, "`id`=?", id).Error
}

func writeFindError(ctx *app.RequestContext, err error) {
	if errors.Is(err, gorm.ErrRecordNotFound) {
		http.NotFound(ctx, "未找到该数据")
		return
	}
	http.Fail(ctx, "查询失败：%v", err)
}
