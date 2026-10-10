package site

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/go-redsync/redsync/v4"
	"github.com/herhe-com/framework/facades"
	"github.com/herhe-com/framework/http"
	"github.com/herhe-com/framework/microservice/locker"
	"github.com/tizips/shenhuo/helper"
	"github.com/tizips/shenhuo/model"
	req "github.com/tizips/shenhuo/server/admin/http/request/site"
	res "github.com/tizips/shenhuo/server/admin/http/response/site"
	"github.com/tizips/shenhuo/wechat"
	"github.com/xuri/excelize/v2"
	"gorm.io/gorm"
)

// ToDrawOfList
// @Summary 获取抽签结果列表
// @Description Permissions: site.draw.paginate
// @Tags 站点-抽签
// @Produce json
// @Security ApiKeyAuth
// @Param category_id query int true "抽签类别"
// @Success 200 {object} []res.ToDrawOfList "抽签结果"
// @Router /site/draws [get]
func ToDrawOfList(c context.Context, ctx *app.RequestContext) {

	var request req.ToDrawOfList

	if err := ctx.BindAndValidate(&request); err != nil {
		http.BadRequest(ctx, err)
		return
	}

	var draws []model.ShDraw

	// 按 id 正序返回；序号按行索引实时生成，不入库
	tx := facades.Database().Default().WithContext(c).
		Preload("Person").
		Preload("Category").
		Where("`category_id`=?", request.CategoryID).
		Order("`id` asc").
		Find(&draws)
	if tx.Error != nil {
		http.Fail(ctx, "获取抽签结果失败：%v", tx.Error)
		return
	}

	responses := make([]res.ToDrawOfList, len(draws))

	for index, item := range draws {
		responses[index] = drawToResponse(item)
		responses[index].Sequence = index + 1
	}

	http.Success(ctx, responses)
}

// DoDrawOfCreate
// @Summary 抽签
// @Description Permissions: site.draw.create
// @Tags 站点-抽签
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param request body req.DoDrawOfCreate true "抽签信息"
// @Success 200 {array} res.DoDrawOfCreate "中签结果"
// @Router /site/draw [post]
func DoDrawOfCreate(c context.Context, ctx *app.RequestContext) {

	var request req.DoDrawOfCreate

	if err := ctx.BindAndValidate(&request); err != nil {
		http.BadRequest(ctx, err)
		return
	}

	mutex := facades.Locker().NewMutex(locker.Keys("draw", request.CategoryID))
	if err := mutex.Lock(); err != nil {
		http.Fail(ctx, "正在抽签，请稍后重试")
		return
	}
	defer func(lock *redsync.Mutex) {
		_, _ = lock.Unlock()
	}(mutex)

	tx := facades.Database().Default().WithContext(c).Begin()

	draws, err := helper.DrawCategory(tx, request.CategoryID)
	if err != nil {
		tx.Rollback()
		if errors.Is(err, gorm.ErrRecordNotFound) {
			http.NotFound(ctx, "未找到该抽签类别")
			return
		}
		http.Fail(ctx, "%s", err.Error())
		return
	}

	tx.Commit()

	// 抽签结束后向中签人员发送微信通知；尽力而为，不影响抽签结果
	wechat.NotifyDraws(c, facades.Database().Default(), draws)

	responses := make([]res.DoDrawOfCreate, len(draws))
	for index, item := range draws {
		responses[index] = res.DoDrawOfCreate{ToDrawOfList: drawToResponse(item)}
	}

	http.Success(ctx, responses)
}

// DoDrawOfDelete
// @Summary 删除抽签结果
// @Description Permissions: site.draw.delete
// @Tags 站点-抽签
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path int true "抽签结果ID"
// @Success 200 {object} nil "删除成功"
// @Router /site/draws/{id} [delete]
func DoDrawOfDelete(c context.Context, ctx *app.RequestContext) {

	var request req.DoDrawOfDelete

	if err := ctx.BindAndValidate(&request); err != nil {
		http.BadRequest(ctx, err)
		return
	}

	var item model.ShDraw

	if err := firstByID(c, &item, request.ID); err != nil {
		writeFindError(ctx, err)
		return
	}

	if result := facades.Database().Default().WithContext(c).Delete(&item); result.Error != nil {
		http.Fail(ctx, "删除失败：%v", result.Error)
		return
	}

	http.Success[any](ctx)
}

// ToDrawOfExport
// @Summary 导出抽签结果 Excel
// @Description Permissions: site.draw.paginate
// @Tags 站点-抽签
// @Produce application/octet-stream
// @Security ApiKeyAuth
// @Router /site/draws/export [get]
func ToDrawOfExport(c context.Context, ctx *app.RequestContext) {

	file := excelize.NewFile()
	defer file.Close()

	var categories []model.ShDrawCategory

	tx := facades.Database().Default().WithContext(c).Order("`order` asc, `id` asc").Find(&categories)
	if tx.Error != nil {
		http.Fail(ctx, "查询抽签类别失败：%v", tx.Error)
		return
	}

	sheetIndex := 0

	for _, category := range categories {

		var draws []model.ShDraw

		tx := facades.Database().Default().WithContext(c).Preload("Person").Where("`category_id`=?", category.ID).Order("`id` asc").Find(&draws)
		if tx.Error != nil {
			http.Fail(ctx, "查询抽签结果失败：%v", tx.Error)
			return
		}

		if len(draws) == 0 {
			continue
		}

		sheet := drawSheetName(category.Name, sheetIndex)
		if sheetIndex == 0 {
			file.SetSheetName(file.GetSheetName(0), sheet)
		} else {
			if _, err := file.NewSheet(sheet); err != nil {
				http.Fail(ctx, "创建工作表失败：%v", err)
				return
			}
		}
		sheetIndex++

		headers := []string{"序号", "参赛号", "姓名", "身份证", "手机号", "单位", "小组名称", "中签时间"}
		if err := file.SetSheetRow(sheet, "A1", &headers); err != nil {
			http.Fail(ctx, "写入表头失败：%v", err)
			return
		}

		// 身份证/手机号/单位/小组名称加宽一倍（默认列宽约 9.14）
		for _, col := range []string{"D", "E", "F", "G"} {
			if err := file.SetColWidth(sheet, col, col, 20); err != nil {
				http.Fail(ctx, "设置列宽失败：%v", err)
				return
			}
		}

		for index, item := range draws {

			row := []string{
				fmt.Sprintf("%d", index+1),
				"", "", "", "", "", "", item.CreatedAt.ToDateTimeString(),
			}

			if item.Person != nil {
				row[1] = item.Person.Number
				row[2] = item.Person.Name
				row[3] = item.Person.IDCard
				row[4] = item.Person.Mobile
				row[5] = item.Person.Unit
				row[6] = item.Person.GroupName
			}

			cell, _ := excelize.CoordinatesToCellName(1, index+2)
			if err := file.SetSheetRow(sheet, cell, &row); err != nil {
				http.Fail(ctx, "写入数据失败：%v", err)
				return
			}
		}
	}

	writeTemplate(ctx, file, "抽签结果.xlsx")
}

// drawSheetName 将抽签类别名称转换为合法的工作表名称。
func drawSheetName(name string, index int) string {

	name = strings.Map(func(r rune) rune {
		switch r {
		case ':', '\\', '/', '?', '*', '[', ']':
			return '_'
		}
		return r
	}, name)

	if runes := []rune(name); len(runes) > 31 {
		name = string(runes[:31])
	}

	if name == "" {
		return fmt.Sprintf("Sheet%d", index+1)
	}

	return name
}

func drawToResponse(item model.ShDraw) res.ToDrawOfList {
	row := res.ToDrawOfList{
		ID:         item.ID,
		CategoryID: item.CategoryID,
		PersonID:   item.PersonID,
		CreatedAt:  item.CreatedAt.ToDateTimeString(),
	}

	if item.Category != nil {
		row.Category = item.Category.Name
	}

	if item.Person != nil {
		row.Name = item.Person.Name
		row.Number = item.Person.Number
		row.Unit = item.Person.Unit
		row.GroupName = item.Person.GroupName
	}

	return row
}
