package basic

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/herhe-com/framework/facades"
	"github.com/herhe-com/framework/http"
	req "github.com/tizips/shenhuo/server/admin/http/request/basic"
	"github.com/tizips/shenhuo/server/admin/http/response/basic"
	"path/filepath"
)

// DoUploadOfFile
// @Summary 上传文件
// @Description 上传文件到指定目录
// @Tags 基础-上传
// @Accept multipart/form-data
// @Produce json
// @Security ApiKeyAuth
// @Param file formData file true "文件"
// @Param dir formData string true "目录"
// @Success 200 {object} basic.DoUploadOfFile "上传结果"
// @Router /basic/upload/file [post]
func DoUploadOfFile(c context.Context, ctx *app.RequestContext) {

	var request req.DoUploadOfFile

	if err := ctx.BindAndValidate(&request); err != nil {
		http.BadRequest(ctx, err)
		return
	}

	file, err := ctx.FormFile("file")

	if err != nil {
		http.Fail(ctx, "上传失败：%v", err)
		return
	}

	fp, err := file.Open()

	if err != nil {
		http.Fail(ctx, "文件读取失败：%v", err)
		return
	}

	filename := facades.Snowflake().Generate().String() + filepath.Ext(file.Filename)
	uri := request.Dir + "/" + filename

	if err = facades.Storage().Put(uri, fp, file.Size); err != nil {
		http.Fail(ctx, "上传失败：%v", err)
		return
	}

	responses := basic.DoUploadOfFile{
		Name: filename,
		Uri:  uri,
		Url:  facades.Storage().Url(uri),
	}

	http.Success(ctx, responses)
}
