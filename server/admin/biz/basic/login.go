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

// DoLoginOfAccount
// @Summary 账号登录
// @Description 管理员账号登录接口，登录成功后返回双令牌（JWT），后续请求按以下方式使用：
// @Description 1. access_token：访问令牌，放入 `Authorization` 请求头（直接放令牌本身，不带 Bearer 前缀）访问所有需要登录的接口；有效时长为 `access_lifetime` 秒（默认 86400 秒，即 24 小时）。
// @Description 2. refresh_token：刷新令牌，平时不随请求发送，仅在访问令牌过期后用于换取新令牌对（见下）。
// @Description 刷新机制（无专门的刷新接口，由鉴权中间件隐式完成）：
// @Description - 访问令牌过期或无效时接口返回 401；此时在任意需要登录的接口上携带 `Refresh-Token: <refresh_token>` 请求头重发请求（注意：不要同时携带 Authorization 头，两者同时存在会直接返回 401）；
// @Description - 服务端校验刷新令牌成功后自动轮换令牌并正常处理本次业务请求；新的令牌对通过响应头 `Token-Pair` 返回（JSON 字符串，字段与登录响应中的令牌字段一致：session_id / access_token / refresh_token / issued_at / access_lifetime / refresh_lifetime / grace_lifetime），前端必须解析并覆盖本地保存的旧令牌；
// @Description - 同一刷新令牌只能轮换一次：并发场景下多个请求同时携带同一旧刷新令牌时，在 `grace_lifetime` 秒（默认 3 秒）容错窗口内会返回同一份新令牌，不会被判为重放；前端只需保存最先收到的一份即可；
// @Description - 超过容错窗口后重复使用旧刷新令牌（黑名单模式下判为重放）、刷新令牌已过期（有效时长 `refresh_lifetime` 秒，默认 30 天）或已被注销时，返回 401，此时必须引导用户重新登录；
// @Description - 轮换后的新令牌对：session_id 保持不变，访问令牌有效时长不变，刷新令牌有效期重新起算（滑动过期）；
// @Description - 令牌相关响应（登录与刷新）均携带 `Cache-Control: no-store`，前端与代理层均不得缓存。
// @Description 退出登录（POST /basic/account/logout）会将当前访问令牌加入黑名单；前端退出时应同时丢弃本地保存的 access_token 与 refresh_token。
// @Tags 基础-登录
// @Accept json
// @Produce json
// @Param request body req.DoLoginOfAccount true "登录信息"
// @Success 200 {object} res.DoLogin "登录成功"
// @Router /basic/login/account [post]
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

// DoLoginOfOut
// @Summary 退出登录
// @Description 退出当前登录
// @Tags 基础-账户
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {object} nil "退出成功"
// @Router /basic/account/logout [post]
func DoLoginOfOut(c context.Context, ctx *app.RequestContext) {

	if ok, _ := auth.BlacklistOfJwtValue(c, ctx); !ok {
		http.Fail(ctx, "退出失败，请稍后重试")
		return
	}

	http.Success[any](ctx)
}
