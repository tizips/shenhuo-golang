package site

import (
	"context"
	"net/url"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/herhe-com/framework/http"
	"github.com/xuri/excelize/v2"
)

// ToPersonOfTemplate
// @Summary 下载人员导入模板
// @Description Permissions: site.person.import
// @Tags 站点-人员
// @Produce application/octet-stream
// @Security ApiKeyAuth
// @Success 200 {file} file "人员导入模板"
// @Router /site/person/template [get]
func ToPersonOfTemplate(c context.Context, ctx *app.RequestContext) {

	file := excelize.NewFile()
	defer file.Close()

	sheet := file.GetSheetName(0)

	headers := []string{"姓名", "单位", "手机号", "参赛号", "小组名称"}
	if err := file.SetSheetRow(sheet, "A1", &headers); err != nil {
		http.Fail(ctx, "生成模板失败：%v", err)
		return
	}

	example := []string{"张三", "XX单位", "13800000000", "001", "第一组"}
	if err := file.SetSheetRow(sheet, "A2", &example); err != nil {
		http.Fail(ctx, "生成模板失败：%v", err)
		return
	}

	writeTemplate(ctx, file, "人员导入模板.xlsx")
}

// ToScoreOfTemplate
// @Summary 下载成绩导入模板
// @Description Permissions: site.score.import
// @Tags 站点-成绩
// @Produce application/octet-stream
// @Security ApiKeyAuth
// @Success 200 {file} file "成绩导入模板"
// @Router /site/score/template [get]
func ToScoreOfTemplate(c context.Context, ctx *app.RequestContext) {

	file := excelize.NewFile()
	defer file.Close()

	sheet := file.GetSheetName(0)

	headers := []string{"参赛号", "姓名", "理论考试", "管理能力考试", "井下紧急避险考核", "自救器操作考核", "心肺复苏考核"}
	if err := file.SetSheetRow(sheet, "A1", &headers); err != nil {
		http.Fail(ctx, "生成模板失败：%v", err)
		return
	}

	example := []string{"001", "张三", "90", "85", "80", "95", "88"}
	if err := file.SetSheetRow(sheet, "A2", &example); err != nil {
		http.Fail(ctx, "生成模板失败：%v", err)
		return
	}

	writeTemplate(ctx, file, "成绩导入模板.xlsx")
}

// writeTemplate 输出 Excel 模板文件。
func writeTemplate(ctx *app.RequestContext, file *excelize.File, filename string) {

	buffer, err := file.WriteToBuffer()
	if err != nil {
		http.Fail(ctx, "生成模板失败：%v", err)
		return
	}

	ctx.Header("Content-Type", "application/octet-stream")
	ctx.Header("Content-Disposition", `attachment; filename*=UTF-8''`+url.PathEscape(filename))
	ctx.Data(200, "application/octet-stream", buffer.Bytes())
}
