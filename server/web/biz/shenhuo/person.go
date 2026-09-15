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

// ToPersonOfList
// @Summary 获取人员列表
// @Description 获取人员列表；不分页
// @Tags 人员
// @Produce json
// @Success 200 {array} res.ToPerson "人员列表"
// @Router /persons [get]
func ToPersonOfList(c context.Context, ctx *app.RequestContext) {

	var people []model.ShPerson

	facades.Database().Default().WithContext(c).Select("`name`", "`unit`").Order("`name` asc, `id` asc").Find(&people)

	responses := make([]res.ToPerson, len(people))
	for index, item := range people {
		responses[index] = res.ToPerson{
			Name: item.Name,
			Unit: item.Unit,
		}
	}

	http.Success(ctx, responses)
}

// ToPersonOfGroupList
// @Summary 获取参赛人员分组列表
// @Description 获取参赛人员分组摘要列表；不含人员明细，明细请调用 /persons/groups/{id}
// @Tags 人员
// @Produce json
// @Success 200 {array} res.ToPersonGroup "参赛人员分组摘要列表"
// @Router /persons/groups [get]
func ToPersonOfGroupList(c context.Context, ctx *app.RequestContext) {

	var categories []model.ShDrawCategory
	facades.Database().Default().WithContext(c).Order("`order` asc, `id` asc").Find(&categories)

	type countOfCategory struct {
		CategoryID uint
		Total      int
	}

	var counts []countOfCategory
	facades.Database().Default().WithContext(c).
		Model(&model.ShDraw{}).
		Select("`category_id`, COUNT(*) AS `total`").
		Group("`category_id`").
		Scan(&counts)

	totals := make(map[uint]int, len(counts))
	for _, item := range counts {
		totals[item.CategoryID] = item.Total
	}

	responses := make([]res.ToPersonGroup, len(categories))
	for index, item := range categories {
		responses[index] = res.ToPersonGroup{
			ID:    item.ID,
			Name:  item.Name,
			Icon:  item.Icon,
			Order: item.Order,
			Total: totals[item.ID],
		}
	}

	http.Success(ctx, responses)
}

// ToPersonOfGroupInformation
// @Summary 获取单个分组的人员列表
// @Description 按照抽签分组获取单个分组的参赛人员列表；不分页
// @Tags 人员
// @Produce json
// @Param id path int true "分组ID（抽签类别ID）"
// @Success 200 {object} res.ToPersonGroupOfPersons "分组参赛人员列表"
// @Router /persons/groups/{id} [get]
func ToPersonOfGroupInformation(c context.Context, ctx *app.RequestContext) {

	var request req.ToPersonOfGroupInformation

	if err := ctx.BindAndValidate(&request); err != nil {
		http.BadRequest(ctx, err)
		return
	}

	var category model.ShDrawCategory

	fu := facades.Database().Default().WithContext(c).First(&category, "`id`=?", request.ID)
	if errors.Is(fu.Error, gorm.ErrRecordNotFound) {
		http.NotFound(ctx, "未找到该数据")
		return
	} else if fu.Error != nil {
		http.Fail(ctx, "查询失败：%v", fu.Error)
		return
	}

	var draws []model.ShDraw
	facades.Database().Default().WithContext(c).Preload("Person").Where("`category_id`=?", request.ID).Find(&draws)

	persons := make([]res.ToGroupedPerson, 0, len(draws))
	for _, draw := range draws {
		if draw.Person == nil {
			continue
		}
		persons = append(persons, res.ToGroupedPerson{
			ID:        draw.Person.ID,
			Name:      draw.Person.Name,
			Number:    draw.Person.Number,
			Unit:      draw.Person.Unit,
			GroupName: draw.Person.GroupName,
		})
	}

	http.Success(ctx, res.ToPersonGroupOfPersons{
		ID:      category.ID,
		Name:    category.Name,
		Icon:    category.Icon,
		Total:   len(persons),
		Persons: persons,
	})
}
