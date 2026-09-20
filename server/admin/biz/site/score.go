package site

import (
	"context"
	"errors"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/herhe-com/framework/contracts/http/response"
	"github.com/herhe-com/framework/facades"
	"github.com/herhe-com/framework/http"
	"github.com/tizips/shenhuo/helper"
	"github.com/tizips/shenhuo/model"
	req "github.com/tizips/shenhuo/server/admin/http/request/site"
	res "github.com/tizips/shenhuo/server/admin/http/response/site"
	"github.com/tizips/shenhuo/wechat"
	"github.com/xuri/excelize/v2"
	"gorm.io/gorm"
)

// ToScoreOfPaginate
// @Summary 获取成绩列表
// @Description Permissions: site.score.paginate
// @Tags 站点-成绩
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param page query int false "页码"
// @Param size query int false "每页数量"
// @Param keyword query string false "关键词"
// @Success 200 {object} response.Paginate[res.ToScoreOfPaginate] "成绩列表"
// @Router /site/scores [get]
func ToScoreOfPaginate(c context.Context, ctx *app.RequestContext) {

	var request req.ToScoreOfPaginate

	if err := ctx.BindAndValidate(&request); err != nil {
		http.BadRequest(ctx, err)
		return
	}

	responses := response.Paginate[res.ToScoreOfPaginate]{
		Page: request.GetPage(),
		Size: request.GetSize(),
	}

	tx := facades.Database().Default().WithContext(c).Model(&model.ShScore{})

	if request.Keyword != "" {
		like := "%" + request.Keyword + "%"
		tx = tx.Where("exists (?)", facades.Database().Default().WithContext(c).
			Model(&model.ShPerson{}).
			Select("1").
			Where(model.TableShPerson+".id="+model.TableShScore+".person_id").
			Where("`name` LIKE ? OR `number` LIKE ? OR `mobile` LIKE ? OR `group_name` LIKE ?", like, like, like, like),
		)
	}

	tx.Count(&responses.Total)

	if responses.Total > 0 {

		var scores []model.ShScore

		tx.Preload("Person").Order("`order` asc, `id` asc").Offset(request.GetOffset()).Limit(request.GetLimit()).Find(&scores)

		responses.Data = make([]res.ToScoreOfPaginate, len(scores))

		for index, item := range scores {
			responses.Data[index] = scoreToResponse(item)
		}
	}

	http.Success(ctx, responses)
}

// ToScoreOfInformation
// @Summary 获取成绩详情
// @Description 获取指定成绩详情
// @Tags 站点-成绩
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path int true "成绩ID"
// @Success 200 {object} res.ToScoreOfInformation "成绩详情"
// @Router /site/scores/{id} [get]
func ToScoreOfInformation(c context.Context, ctx *app.RequestContext) {

	var request req.ToScoreOfInformation

	if err := ctx.BindAndValidate(&request); err != nil {
		http.BadRequest(ctx, err)
		return
	}

	var score model.ShScore

	fu := facades.Database().Default().WithContext(c).Preload("Person").First(&score, "`id`=?", request.ID)
	if fu.Error != nil {
		writeFindError(ctx, fu.Error)
		return
	}

	http.Success(ctx, res.ToScoreOfInformation{ToScoreOfPaginate: scoreToResponse(score)})
}

// DoScoreOfCreate 创建成绩
// @Summary 创建成绩
// @Description Permissions: site.score.create；总分与名次由五项小项成绩自动计算
// @Tags 站点-成绩
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param request body req.DoScoreOfCreate true "成绩信息"
// @Success 200 {object} nil "创建成功"
// @Router /site/score [post]
func DoScoreOfCreate(c context.Context, ctx *app.RequestContext) {

	var request req.DoScoreOfCreate

	if err := ctx.BindAndValidate(&request); err != nil {
		http.BadRequest(ctx, err)
		return
	}

	if !personExists(c, request.PersonID) {
		http.Fail(ctx, "人员不存在")
		return
	}

	items := buildScoreItems(request.Theory, request.Management, request.Escape, request.Respirator, request.CPR)

	total, err := helper.CalculateScoreTotal(scoreItemValues(items))
	if err != nil {
		http.Fail(ctx, "%s", err.Error())
		return
	}

	var existed model.ShScore
	fs := facades.Database().Default().WithContext(c).Unscoped().First(&existed, "`person_id`=?", request.PersonID)
	if fs.Error == nil {
		if !existed.DeletedAt.Valid {
			http.Fail(ctx, "该人员已有成绩")
			return
		}

		existed.Total = total
		existed.Items = model.MarshalScoreItems(items)
		existed.DeletedAt = gorm.DeletedAt{}

		if result := facades.Database().Default().WithContext(c).Unscoped().Save(&existed); result.Error != nil {
			http.Fail(ctx, "创建失败：%v", result.Error)
			return
		}

		if err := refreshScoreRanks(facades.Database().Default().WithContext(c)); err != nil {
			http.Fail(ctx, "重新计算名次失败：%v", err)
			return
		}

		http.Success[any](ctx)
		return
	} else if !errors.Is(fs.Error, gorm.ErrRecordNotFound) {
		http.Fail(ctx, "查询失败：%v", fs.Error)
		return
	}

	score := model.ShScore{
		PersonID: request.PersonID,
		Total:    total,
		Items:    model.MarshalScoreItems(items),
		Order:    helper.DefaultOrder,
	}

	if result := facades.Database().Default().WithContext(c).Create(&score); result.Error != nil {
		http.Fail(ctx, "创建失败：%v", result.Error)
		return
	}

	if err := refreshScoreRanks(facades.Database().Default().WithContext(c)); err != nil {
		http.Fail(ctx, "重新计算名次失败：%v", err)
		return
	}

	http.Success[any](ctx)
}

// DoScoreOfUpdate 更新成绩
// @Summary 更新成绩
// @Description Permissions: site.score.update；总分与名次由五项小项成绩自动计算
// @Tags 站点-成绩
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path int true "成绩ID"
// @Param request body req.DoScoreOfUpdate true "成绩信息"
// @Success 200 {object} nil "更新成功"
// @Router /site/scores/{id} [put]
func DoScoreOfUpdate(c context.Context, ctx *app.RequestContext) {

	var request req.DoScoreOfUpdate

	if err := ctx.BindAndValidate(&request); err != nil {
		http.BadRequest(ctx, err)
		return
	}

	var score model.ShScore

	if err := firstByID(c, &score, request.ID); err != nil {
		writeFindError(ctx, err)
		return
	}

	items := buildScoreItems(request.Theory, request.Management, request.Escape, request.Respirator, request.CPR)

	total, err := helper.CalculateScoreTotal(scoreItemValues(items))
	if err != nil {
		http.Fail(ctx, "%s", err.Error())
		return
	}

	score.Total = total
	score.Items = model.MarshalScoreItems(items)

	if result := facades.Database().Default().WithContext(c).Save(&score); result.Error != nil {
		http.Fail(ctx, "修改失败：%v", result.Error)
		return
	}

	if err := refreshScoreRanks(facades.Database().Default().WithContext(c)); err != nil {
		http.Fail(ctx, "重新计算名次失败：%v", err)
		return
	}

	http.Success[any](ctx)
}

// DoScoreOfDelete
// @Summary 删除成绩
// @Description Permissions: site.score.delete
// @Tags 站点-成绩
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path int true "成绩ID"
// @Success 200 {object} nil "删除成功"
// @Router /site/scores/{id} [delete]
func DoScoreOfDelete(c context.Context, ctx *app.RequestContext) {

	var request req.DoScoreOfDelete

	if err := ctx.BindAndValidate(&request); err != nil {
		http.BadRequest(ctx, err)
		return
	}

	var score model.ShScore

	if err := firstByID(c, &score, request.ID); err != nil {
		writeFindError(ctx, err)
		return
	}

	if result := facades.Database().Default().WithContext(c).Delete(&score); result.Error != nil {
		http.Fail(ctx, "删除失败：%v", result.Error)
		return
	}

	if err := refreshScoreRanks(facades.Database().Default().WithContext(c)); err != nil {
		http.Fail(ctx, "重新计算名次失败：%v", err)
		return
	}

	http.Success[any](ctx)
}

// DoScoreOfImport 导入成绩
// @Summary 导入成绩
// @Description Permissions: site.score.import；模板包含参赛号与五个固定小项，总分与名次由系统自动计算
// @Tags 站点-成绩
// @Accept multipart/form-data
// @Produce json
// @Security ApiKeyAuth
// @Param file formData file true "Excel 文件"
// @Success 200 {object} res.DoScoreOfImport "导入结果"
// @Router /site/score/import [post]
func DoScoreOfImport(c context.Context, ctx *app.RequestContext) {

	header, err := ctx.FormFile("file")
	if err != nil {
		http.BadRequest(ctx, errors.New("请上传 Excel 文件"))
		return
	}

	src, err := header.Open()
	if err != nil {
		http.Fail(ctx, "文件读取失败：%v", err)
		return
	}
	defer src.Close()

	file, err := excelize.OpenReader(src)
	if err != nil {
		http.Fail(ctx, "Excel 解析失败：%v", err)
		return
	}
	defer file.Close()

	rows, err := helper.ReadExcelRows(file)
	if err != nil || len(rows) < 2 {
		http.Fail(ctx, "Excel 内容为空")
		return
	}

	numberCol, itemCols, err := helper.ParseScoreHeader(rows[0])
	if err != nil {
		http.Fail(ctx, "%s", err.Error())
		return
	}

	responses := res.DoScoreOfImport{Total: len(rows) - 1}

	tx := facades.Database().Default().WithContext(c).Begin()

	for index, row := range rows[1:] {
		number := helper.Cell(row, numberCol)

		if number == "" {
			responses.Skipped++
			continue
		}

		items := helper.ScoreItemsFromRow(row, itemCols)

		total, err := helper.CalculateScoreTotal(scoreItemValues(items))
		if err != nil {
			tx.Rollback()
			http.Fail(ctx, "导入失败：第 %d 行参赛号 %s 的%s", index+2, number, err.Error())
			return
		}

		var person model.ShPerson
		fp := tx.First(&person, "`number`=?", number)
		if errors.Is(fp.Error, gorm.ErrRecordNotFound) {
			responses.Skipped++
			continue
		} else if fp.Error != nil {
			tx.Rollback()
			http.Fail(ctx, "导入失败：%v", fp.Error)
			return
		}

		payload := model.ShScore{
			PersonID: person.ID,
			Total:    total,
			Items:    model.MarshalScoreItems(items),
			Order:    helper.DefaultOrder,
		}

		var existed model.ShScore
		fs := tx.Unscoped().First(&existed, "`person_id`=?", person.ID)
		if errors.Is(fs.Error, gorm.ErrRecordNotFound) {
			if result := tx.Create(&payload); result.Error != nil {
				tx.Rollback()
				http.Fail(ctx, "导入失败：%v", result.Error)
				return
			}
			responses.Created++
			continue
		} else if fs.Error != nil {
			tx.Rollback()
			http.Fail(ctx, "导入失败：%v", fs.Error)
			return
		}

		existed.Total = payload.Total
		existed.Items = payload.Items
		existed.DeletedAt = gorm.DeletedAt{}

		if result := tx.Unscoped().Save(&existed); result.Error != nil {
			tx.Rollback()
			http.Fail(ctx, "导入失败：%v", result.Error)
			return
		}

		responses.Updated++
	}

	if !refreshScoreRanksInTx(ctx, tx) {
		return
	}

	tx.Commit()

	http.Success(ctx, responses)
}

// DoScoreOfNotify 推送成绩通知
// @Summary 推送成绩通知
// @Description Permissions: site.score.notify；按抽签分组推送，向分组下所有已出成绩、且绑定了微信的参赛人员发送成绩发布通知
// @Tags 站点-成绩
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param request body req.DoScoreOfNotify true "请求参数（抽签分组ID）"
// @Success 200 {object} res.DoScoreOfNotify "推送结果"
// @Router /site/score/notify [post]
func DoScoreOfNotify(c context.Context, ctx *app.RequestContext) {

	var request req.DoScoreOfNotify

	if err := ctx.BindAndValidate(&request); err != nil {
		http.BadRequest(ctx, err)
		return
	}

	var category model.ShDrawCategory
	if err := facades.Database().Default().WithContext(c).First(&category, "`id`=?", request.CategoryID).Error; err != nil {
		writeFindError(ctx, err)
		return
	}

	var scores []model.ShScore

	facades.Database().Default().WithContext(c).
		Joins("JOIN `"+model.TableShDraw+"` ON `"+model.TableShDraw+"`.`person_id` = `"+model.TableShScore+"`.`person_id`").
		Where("`"+model.TableShDraw+"`.`category_id`=?", request.CategoryID).
		Preload("Person").
		Order("`" + model.TableShScore + "`.`order` asc").
		Find(&scores)

	if len(scores) == 0 {
		http.Fail(ctx, "该抽签分组下暂无成绩记录")
		return
	}

	response := res.DoScoreOfNotify{Group: category.Name, Total: len(scores)}

	for _, score := range scores {

		if score.Person == nil || score.Person.OpenID == "" {
			response.Skipped++
			continue
		}

		if err := wechat.SendScoreNotice(c, score.Person.OpenID, score.Person.Name, score.Person.Number, score.Total); err != nil {
			response.Failed++
			continue
		}

		response.Sent++
	}

	http.Success(ctx, response)
}

// buildScoreItems 按固定小项顺序组装成绩小项。
func buildScoreItems(theory, management, escape, respirator, cpr string) []model.ShScoreItem {
	return []model.ShScoreItem{
		{Name: helper.ScoreItemNames[0], Value: theory},
		{Name: helper.ScoreItemNames[1], Value: management},
		{Name: helper.ScoreItemNames[2], Value: escape},
		{Name: helper.ScoreItemNames[3], Value: respirator},
		{Name: helper.ScoreItemNames[4], Value: cpr},
	}
}

// scoreItemValues 提取各小项成绩值，用于计算总分。
func scoreItemValues(items []model.ShScoreItem) []string {
	values := make([]string, len(items))
	for index, item := range items {
		values[index] = item.Value
	}
	return values
}

// refreshScoreRanks 按总分重新计算名次；失败时返回错误。
func refreshScoreRanks(db *gorm.DB) error {
	return helper.RefreshScoreRanks(db)
}

// refreshScoreRanksInTx 在事务内重新计算名次；失败时回滚事务并返回错误响应。
func refreshScoreRanksInTx(ctx *app.RequestContext, tx *gorm.DB) bool {
	if err := helper.RefreshScoreRanks(tx); err == nil {
		return true
	} else {
		tx.Rollback()
		http.Fail(ctx, "导入失败：重新计算名次失败：%v", err)
		return false
	}
}

func personExists(c context.Context, id string) bool {
	var total int64
	facades.Database().Default().WithContext(c).Model(&model.ShPerson{}).Where("`id`=?", id).Count(&total)
	return total > 0
}

func scoreToResponse(item model.ShScore) res.ToScoreOfPaginate {
	row := res.ToScoreOfPaginate{
		ID:        item.ID,
		PersonID:  item.PersonID,
		Total:     item.Total,
		Items:     item.ItemList(),
		Order:     item.Order,
		CreatedAt: item.CreatedAt.ToDateTimeString(),
	}

	if item.Person != nil {
		row.Name = item.Person.Name
		row.Number = item.Person.Number
		row.Unit = item.Person.Unit
		row.GroupName = item.Person.GroupName
	}

	return row
}
