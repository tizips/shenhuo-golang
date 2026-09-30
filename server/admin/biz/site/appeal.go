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

// ToAppealOfPaginate
// @Summary 获取仲裁申诉列表
// @Description Permissions: site.appeal.paginate
// @Tags 站点-仲裁申诉
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param keyword query string false "关键词"
// @Param page query int false "页码"
// @Param size query int false "每页数量"
// @Success 200 {object} response.Paginate[res.ToAppealOfPaginate] "申诉列表"
// @Router /site/appeals [get]
func ToAppealOfPaginate(c context.Context, ctx *app.RequestContext) {

	var request req.ToAppealOfPaginate

	if err := ctx.BindAndValidate(&request); err != nil {
		http.BadRequest(ctx, err)
		return
	}

	responses := response.Paginate[res.ToAppealOfPaginate]{
		Page: request.GetPage(),
		Size: request.GetSize(),
	}

	tx := facades.Database().Default().WithContext(c).Model(&model.ShAppeal{})

	if request.Keyword != "" {
		like := "%" + request.Keyword + "%"
		tx = tx.Where("exists (?)", facades.Database().Default().WithContext(c).
			Model(&model.ShPerson{}).
			Select("1").
			Where(model.TableShPerson+".id="+model.TableShAppeal+".person_id").
			Where("`name` LIKE ? OR `number` LIKE ? OR `id_card` LIKE ?", like, like, like),
		)
	}

	tx.Count(&responses.Total)

	if responses.Total > 0 {

		var appeals []model.ShAppeal

		tx.Preload("Person").Order("`id` desc").Offset(request.GetOffset()).Limit(request.GetLimit()).Find(&appeals)

		responses.Data = make([]res.ToAppealOfPaginate, len(appeals))

		for index, item := range appeals {
			responses.Data[index] = appealToResponse(item)
		}
	}

	http.Success(ctx, responses)
}

// ToAppealOfInformation
// @Summary 获取仲裁申诉详情
// @Description 获取指定仲裁申诉详情
// @Tags 站点-仲裁申诉
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path int true "申诉ID"
// @Success 200 {object} res.ToAppealOfInformation "申诉详情"
// @Router /site/appeals/{id} [get]
func ToAppealOfInformation(c context.Context, ctx *app.RequestContext) {

	var request req.ToAppealOfInformation

	if err := ctx.BindAndValidate(&request); err != nil {
		http.BadRequest(ctx, err)
		return
	}

	var appeal model.ShAppeal

	fu := facades.Database().Default().WithContext(c).Preload("Person").First(&appeal, "`id`=?", request.ID)
	if fu.Error != nil {
		writeFindError(ctx, fu.Error)
		return
	}

	http.Success(ctx, res.ToAppealOfInformation{ToAppealOfPaginate: appealToResponse(appeal)})
}

func appealToResponse(item model.ShAppeal) res.ToAppealOfPaginate {
	row := res.ToAppealOfPaginate{
		ID:        item.ID,
		PersonID:  item.PersonID,
		Reason:    item.Reason,
		CreatedAt: item.CreatedAt.ToDateTimeString(),
	}

	if item.Person != nil {
		row.Name = item.Person.Name
		row.Number = item.Person.Number
		row.Unit = item.Person.Unit
		row.GroupName = item.Person.GroupName
		row.IDCard = item.Person.IDCard
	}

	return row
}
