package site

import (
	"context"
	"errors"
	"time"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/herhe-com/framework/auth"
	"github.com/herhe-com/framework/constants/global"
	"github.com/herhe-com/framework/contracts/http/response"
	"github.com/herhe-com/framework/facades"
	"github.com/herhe-com/framework/http"
	"github.com/tizips/shenhuo/helper"
	"github.com/tizips/shenhuo/model"
	req "github.com/tizips/shenhuo/server/admin/http/request/site"
	res "github.com/tizips/shenhuo/server/admin/http/response/site"
	"github.com/xuri/excelize/v2"
	"gorm.io/gorm"
)

// ToPersonOfPaginate
// @Summary 获取人员列表
// @Description Permissions: site.person.paginate
// @Tags 站点-人员
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param page query int false "页码"
// @Param size query int false "每页数量"
// @Param keyword query string false "关键词"
// @Success 200 {object} response.Paginate[res.ToPersonOfPaginate] "人员列表"
// @Router /site/persons [get]
func ToPersonOfPaginate(c context.Context, ctx *app.RequestContext) {

	var request req.ToPersonOfPaginate

	if err := ctx.BindAndValidate(&request); err != nil {
		http.BadRequest(ctx, err)
		return
	}

	responses := response.Paginate[res.ToPersonOfPaginate]{
		Page: request.GetPage(),
		Size: request.GetSize(),
	}

	tx := facades.Database().Default().WithContext(c).Model(&model.ShPerson{})

	if request.Keyword != "" {
		like := "%" + request.Keyword + "%"
		tx = tx.Where("`name` LIKE ? OR `mobile` LIKE ? OR `number` LIKE ? OR `group_name` LIKE ? OR `unit` LIKE ?", like, like, like, like, like)
	}

	tx.Count(&responses.Total)

	if responses.Total > 0 {

		var people []model.ShPerson

		tx.Order("`id` desc").Offset(request.GetOffset()).Limit(request.GetLimit()).Find(&people)

		responses.Data = make([]res.ToPersonOfPaginate, len(people))

		for index, item := range people {
			responses.Data[index] = personToPaginate(item)
		}
	}

	http.Success(ctx, responses)
}

// ToPersonOfInformation
// @Summary 获取人员详情
// @Description 获取指定人员详情
// @Tags 站点-人员
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path string true "人员ID"
// @Success 200 {object} res.ToPersonOfInformation "人员详情"
// @Router /site/persons/{id} [get]
func ToPersonOfInformation(c context.Context, ctx *app.RequestContext) {

	var request req.ToPersonOfInformation

	if err := ctx.BindAndValidate(&request); err != nil {
		http.BadRequest(ctx, err)
		return
	}

	var person model.ShPerson

	if err := firstBySID(c, &person, request.ID); err != nil {
		writeFindError(ctx, err)
		return
	}

	http.Success(ctx, res.ToPersonOfInformation{ToPersonOfPaginate: personToPaginate(person)})
}

// ToPersonOfOpening
// @Summary 获取可用人员列表
// @Description 获取所有可用的人员列表
// @Tags 站点-人员
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {array} response.Opening[string] "人员列表"
// @Router /site/person/opening [get]
func ToPersonOfOpening(c context.Context, ctx *app.RequestContext) {

	var people []model.ShPerson

	facades.Database().Default().WithContext(c).Order("`id` desc").Find(&people)

	responses := make([]response.Opening[string], len(people))

	for index, item := range people {
		responses[index] = response.Opening[string]{
			ID:   item.ID,
			Name: item.Name + "（" + item.Number + "）",
		}
	}

	http.Success(ctx, responses)
}

// DoPersonOfCreate
// @Summary 创建人员
// @Description Permissions: site.person.create
// @Tags 站点-人员
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param request body req.DoPersonOfCreate true "人员信息"
// @Success 200 {object} nil "创建成功"
// @Router /site/person [post]
func DoPersonOfCreate(c context.Context, ctx *app.RequestContext) {

	var request req.DoPersonOfCreate

	if err := ctx.BindAndValidate(&request); err != nil {
		http.BadRequest(ctx, err)
		return
	}

	if err := assertPersonUnique(c, request.Mobile, request.Number, ""); err != nil {
		http.Fail(ctx, "%s", err.Error())
		return
	}

	person := model.ShPerson{
		ID:                 facades.Snowflake().Generate().String(),
		Name:               request.Name,
		Unit:               request.Unit,
		Mobile:             request.Mobile,
		Password:           auth.Password(request.Password),
		Number:             request.Number,
		GroupName:          request.GroupName,
		MustChangePassword: global.NO,
	}

	if result := facades.Database().Default().WithContext(c).Create(&person); result.Error != nil {
		http.Fail(ctx, "创建失败：%v", result.Error)
		return
	}

	http.Success[any](ctx)
}

// DoPersonOfUpdate
// @Summary 更新人员
// @Description Permissions: site.person.update
// @Tags 站点-人员
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path string true "人员ID"
// @Param request body req.DoPersonOfUpdate true "人员信息"
// @Success 200 {object} nil "更新成功"
// @Router /site/persons/{id} [put]
func DoPersonOfUpdate(c context.Context, ctx *app.RequestContext) {

	var request req.DoPersonOfUpdate

	if err := ctx.BindAndValidate(&request); err != nil {
		http.BadRequest(ctx, err)
		return
	}

	var person model.ShPerson

	if err := firstBySID(c, &person, request.ID); err != nil {
		writeFindError(ctx, err)
		return
	}

	if err := assertPersonUnique(c, request.Mobile, request.Number, person.ID); err != nil {
		http.Fail(ctx, "%s", err.Error())
		return
	}

	person.Name = request.Name
	person.Unit = request.Unit
	person.Mobile = request.Mobile
	person.Number = request.Number
	person.GroupName = request.GroupName

	if request.Password != "" {
		person.Password = auth.Password(request.Password)
		person.MustChangePassword = global.NO
	}

	if result := facades.Database().Default().WithContext(c).Save(&person); result.Error != nil {
		http.Fail(ctx, "修改失败：%v", result.Error)
		return
	}

	http.Success[any](ctx)
}

// DoPersonOfDelete
// @Summary 删除人员
// @Description Permissions: site.person.delete
// @Tags 站点-人员
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path string true "人员ID"
// @Success 200 {object} nil "删除成功"
// @Router /site/persons/{id} [delete]
func DoPersonOfDelete(c context.Context, ctx *app.RequestContext) {

	var request req.DoPersonOfDelete

	if err := ctx.BindAndValidate(&request); err != nil {
		http.BadRequest(ctx, err)
		return
	}

	var person model.ShPerson

	if err := firstBySID(c, &person, request.ID); err != nil {
		writeFindError(ctx, err)
		return
	}

	if result := facades.Database().Default().WithContext(c).Delete(&person); result.Error != nil {
		http.Fail(ctx, "删除失败：%v", result.Error)
		return
	}

	http.Success[any](ctx)
}

// DoPersonOfImport
// @Summary 导入人员
// @Description Permissions: site.person.import
// @Tags 站点-人员
// @Accept multipart/form-data
// @Produce json
// @Security ApiKeyAuth
// @Param file formData file true "Excel 文件"
// @Success 200 {object} res.DoPersonOfImport "导入结果"
// @Router /site/person/import [post]
func DoPersonOfImport(c context.Context, ctx *app.RequestContext) {

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

	headers := helper.HeaderIndex(rows[0])

	now := time.Now()
	responses := res.DoPersonOfImport{Total: len(rows) - 1}

	tx := facades.Database().Default().WithContext(c).Begin()

	for _, row := range rows[1:] {
		name := helper.CellByName(row, headers, "姓名")
		unit := helper.CellByName(row, headers, "单位")
		mobile := helper.CellByName(row, headers, "手机号", "手机")
		number := helper.CellByName(row, headers, "参赛号", "编号")
		group := helper.CellByName(row, headers, "小组名称", "小组")

		if name == "" && unit == "" && mobile == "" && number == "" && group == "" {
			responses.Skipped++
			continue
		}

		if name == "" || unit == "" || mobile == "" || number == "" || group == "" {
			tx.Rollback()
			http.Fail(ctx, "导入失败：存在必填字段为空的行")
			return
		}

		if err = assertPersonUniqueTx(tx, mobile, number, ""); err != nil {
			responses.Skipped++
			continue
		}

		person := model.ShPerson{
			ID:                 facades.Snowflake().Generate().String(),
			Name:               name,
			Unit:               unit,
			Mobile:             mobile,
			Password:           auth.Password(helper.PersonImportPassword(mobile, now)),
			Number:             number,
			GroupName:          group,
			MustChangePassword: global.YES,
		}

		if result := tx.Create(&person); result.Error != nil {
			tx.Rollback()
			http.Fail(ctx, "导入失败：%v", result.Error)
			return
		}

		responses.Created++
	}

	tx.Commit()

	http.Success(ctx, responses)
}

func personToPaginate(item model.ShPerson) res.ToPersonOfPaginate {
	return res.ToPersonOfPaginate{
		ID:                 item.ID,
		Name:               item.Name,
		Unit:               item.Unit,
		Mobile:             item.Mobile,
		Number:             item.Number,
		GroupName:          item.GroupName,
		MustChangePassword: item.MustChangePassword,
		CreatedAt:          item.CreatedAt.ToDateTimeString(),
	}
}

func firstBySID(c context.Context, dest any, id string) error {
	return facades.Database().Default().WithContext(c).First(dest, "`id`=?", id).Error
}

func assertPersonUnique(c context.Context, mobile, number, exclude string) error {
	return assertPersonUniqueTx(facades.Database().Default().WithContext(c), mobile, number, exclude)
}

func assertPersonUniqueTx(tx *gorm.DB, mobile, number, exclude string) error {
	var total int64

	q := tx.Model(&model.ShPerson{}).Where("`mobile`=?", mobile)
	if exclude != "" {
		q = q.Where("`id`<>?", exclude)
	}
	q.Count(&total)
	if total > 0 {
		return errors.New("手机号已存在")
	}

	q = tx.Model(&model.ShPerson{}).Where("`number`=?", number)
	if exclude != "" {
		q = q.Where("`id`<>?", exclude)
	}
	q.Count(&total)
	if total > 0 {
		return errors.New("参赛号已存在")
	}

	return nil
}
