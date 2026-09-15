package shenhuo

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/herhe-com/framework/facades"
	"github.com/herhe-com/framework/http"
	"github.com/tizips/shenhuo/model"
	res "github.com/tizips/shenhuo/server/web/http/response/shenhuo"
)

// ToScheduleOfList
// @Summary 获取日程列表
// @Description 按日程分类分组返回日程列表
// @Tags 日程
// @Accept json
// @Produce json
// @Success 200 {array} res.ToScheduleCategory "日程分类列表"
// @Router /schedules [get]
func ToScheduleOfList(c context.Context, ctx *app.RequestContext) {

	var schedules []model.ShSchedule

	facades.Database().Default().WithContext(c).Order("`order` asc, `id` asc").Find(&schedules)

	groups := make(map[uint][]res.ToSchedule, len(schedules))

	for _, item := range schedules {

		subitems := make([]res.ToScheduleItem, 0, len(item.Items))
		for _, subitem := range item.Items {
			subitems = append(subitems, res.ToScheduleItem{
				Name: subitem.Name,
				Time: subitem.Time,
			})
		}

		groups[item.CategoryID] = append(groups[item.CategoryID], res.ToSchedule{
			ID:          item.ID,
			Title:       item.Title,
			Subtitle:    item.Subtitle,
			Description: item.Description,
			Time:        item.Time,
			Items:       subitems,
			Order:       item.Order,
		})
	}

	var categories []model.ShScheduleCategory

	facades.Database().Default().WithContext(c).Order("`order` asc, `id` asc").Find(&categories)

	responses := make([]res.ToScheduleCategory, 0, len(categories))

	for _, category := range categories {
		responses = append(responses, res.ToScheduleCategory{
			ID:        category.ID,
			Name:      category.Name,
			Order:     category.Order,
			Schedules: groups[category.ID],
		})
	}

	http.Success(ctx, responses)
}
