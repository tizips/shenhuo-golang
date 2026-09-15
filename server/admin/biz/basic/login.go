package basic

import (
	"context"
	"errors"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/herhe-com/framework/auth"
	"github.com/herhe-com/framework/constants/global"
	"github.com/herhe-com/framework/facades"
	"github.com/herhe-com/framework/http"
	"github.com/tizips/shenhuo/model"
	req "github.com/tizips/shenhuo/server/admin/http/request/basic"
	res "github.com/tizips/shenhuo/server/admin/http/response/basic"
	"gorm.io/gorm"
)

func DoLoginOfAccount(c context.Context, ctx *app.RequestContext) {

	var request req.DoLoginOfAccount

	if err := ctx.BindAndValidate(&request); err != nil {
		http.BadRequest(ctx, err)
		return
	}

	var user model.SysUser

	fu := facades.Database().Default().WithContext(c).First(&user, "`username`=? and `is_enable`=?", request.Username, global.YES)

	if fu.Error != nil {
		http.Fail(ctx, "用户名或密码错误")
		return
	}

	if !auth.CheckPassword(request.Password, user.Password) {
		http.Fail(ctx, "用户名或密码错误")
		return
	}

	var bind model.SysUserBindRole

	fb := facades.Database().Default().WithContext(c).Order("`platform` asc").First(&bind, "`user_id`=?", user.ID)

	if errors.Is(fb.Error, gorm.ErrRecordNotFound) {
		http.NotFound(ctx, "未查询到被授权的角色")
		return
	} else if fb.Error != nil {
		http.Fail(ctx, "登陆失败：%v", fb.Error)
		return
	}

	pair, err := auth.NewLoginJWToken(user.ID, true, nil)
	if err != nil {
		http.Login(ctx)
		return
	}

	responses := res.DoLogin{
		SessionID:       pair.SessionID,
		AccessToken:     pair.AccessToken,
		RefreshToken:    pair.RefreshToken,
		IssuedAt:        pair.IssuedAt,
		AccessLifetime:  pair.AccessLifetime,
		RefreshLifetime: pair.RefreshLifetime,
		GraceLifetime:   pair.GraceLifetime,
	}

	ctx.Header("Cache-Control", "no-store")
	ctx.Header("Pragma", "no-cache")

	http.Success(ctx, responses)
}

func DoLoginOfOut(c context.Context, ctx *app.RequestContext) {

	if ok, _ := auth.BlacklistOfJwtValue(c, ctx); !ok {
		http.Fail(ctx, "退出失败，请稍后重试")
		return
	}

	http.Success[any](ctx)
}
