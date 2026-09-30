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

// ToCheckinOfPaginate
// @Summary 获取签到列表
// @Description Permissions: site.checkin.paginate
// @Tags 站点-签到
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param name query string false "姓名"
// @Param page query int false "页码"
// @Param size query int false "每页数量"
// @Success 200 {object} response.Paginate[res.ToCheckinOfPaginate] "签到列表"
// @Router /site/checkins [get]
func ToCheckinOfPaginate(c context.Context, ctx *app.RequestContext) {

	var request req.ToCheckinOfPaginate

	if err := ctx.BindAndValidate(&request); err != nil {
		http.BadRequest(ctx, err)
		return
	}

	responses := response.Paginate[res.ToCheckinOfPaginate]{
		Page: request.GetPage(),
		Size: request.GetSize(),
	}

	tx := facades.Database().Default().WithContext(c).Model(&model.ShCheckin{})

	if request.Name != "" {
		tx = tx.Where("`name` LIKE ?", "%"+request.Name+"%")
	}

	tx.Count(&responses.Total)

	if responses.Total > 0 {

		var checkins []model.ShCheckin

		tx.Order("`id` desc").Offset(request.GetOffset()).Limit(request.GetLimit()).Find(&checkins)

		responses.Data = make([]res.ToCheckinOfPaginate, len(checkins))

		for index, item := range checkins {
			responses.Data[index] = checkinToResponse(item)
		}
	}

	http.Success(ctx, responses)
}

// ToCheckinOfInformation
// @Summary 获取签到详情
// @Description 获取指定签到详情
// @Tags 站点-签到
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path int true "签到ID"
// @Success 200 {object} res.ToCheckinOfInformation "签到详情"
// @Router /site/checkins/{id} [get]
func ToCheckinOfInformation(c context.Context, ctx *app.RequestContext) {

	var request req.ToCheckinOfInformation

	if err := ctx.BindAndValidate(&request); err != nil {
		http.BadRequest(ctx, err)
		return
	}

	var checkin model.ShCheckin

	if err := firstByID(c, &checkin, request.ID); err != nil {
		writeFindError(ctx, err)
		return
	}

	http.Success(ctx, res.ToCheckinOfInformation{ToCheckinOfPaginate: checkinToResponse(checkin)})
}

func checkinToResponse(item model.ShCheckin) res.ToCheckinOfPaginate {
	return res.ToCheckinOfPaginate{
		ID:        item.ID,
		Name:      item.Name,
		CreatedAt: item.CreatedAt.ToDateTimeString(),
	}
}
