package basic

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/herhe-com/framework/auth"
	"github.com/herhe-com/framework/constants/global"
	"github.com/herhe-com/framework/facades"
	"github.com/herhe-com/framework/http"
	"github.com/tizips/shenhuo/model"
	"github.com/tizips/shenhuo/server/web/constants"
	req "github.com/tizips/shenhuo/server/web/http/request/basic"
	res "github.com/tizips/shenhuo/server/web/http/response/basic"
	"github.com/tizips/shenhuo/wechat"
)

// DoLoginOfManager
// @Summary 管理人员登录
// @Description 小程序管理人员登录，登录成功后返回双令牌（JWT），后续请求按以下方式使用：
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
// @Param request body req.DoLoginOfMobile true "登录信息"
// @Success 200 {object} res.DoLogin "登录成功"
// @Router /basic/login/manager [post]
func DoLoginOfManager(c context.Context, ctx *app.RequestContext) {

	var request req.DoLoginOfMobile

	if err := ctx.BindAndValidate(&request); err != nil {
		http.BadRequest(ctx, err)
		return
	}

	var manager model.ShManager

	fu := facades.Database().Default().WithContext(c).First(&manager, "`mobile`=? and `is_enable`=?", request.Mobile, global.YES)
	if fu.Error != nil || !auth.CheckPassword(request.Password, manager.Password) {
		http.Fail(ctx, "手机号或密码错误")
		return
	}

	writeLogin(ctx, manager.ID, constants.JwtKindManager, manager.Name, manager.Mobile, 0)
}

// DoLoginOfPerson
// @Summary 参赛人员登录
// @Description 参赛人员登录，登录成功后返回双令牌（JWT），后续请求按以下方式使用：
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
// @Description 登录成功后若 `must_change_password=1`，前端应引导参赛人员先修改密码（PUT /basic/account/password）。
// @Tags 基础-登录
// @Accept json
// @Produce json
// @Param request body req.DoLoginOfMobile true "登录信息"
// @Success 200 {object} res.DoLogin "登录成功"
// @Router /basic/login/person [post]
func DoLoginOfPerson(c context.Context, ctx *app.RequestContext) {

	var request req.DoLoginOfMobile

	if err := ctx.BindAndValidate(&request); err != nil {
		http.BadRequest(ctx, err)
		return
	}

	var person model.ShPerson

	fu := facades.Database().Default().WithContext(c).First(&person, "`mobile`=?", request.Mobile)
	if fu.Error != nil || !auth.CheckPassword(request.Password, person.Password) {
		http.Fail(ctx, "手机号或密码错误")
		return
	}

	writeLogin(ctx, person.ID, constants.JwtKindPerson, person.Name, person.Mobile, person.MustChangePassword)
}

// DoLoginOfWeChat
// @Summary 参赛人员微信登录
// @Description 参赛人员使用手机号、密码和微信服务号网页授权 Code 登录，登录成功后将 OpenID 绑定到当前用户。
// @Description 登录成功后返回双令牌（JWT），后续请求按以下方式使用：
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
// @Param request body req.DoLoginOfWeChat true "登录信息"
// @Success 200 {object} res.DoLogin "登录成功"
// @Router /basic/login/wechat [post]
func DoLoginOfWeChat(c context.Context, ctx *app.RequestContext) {

	var request req.DoLoginOfWeChat

	if err := ctx.BindAndValidate(&request); err != nil {
		http.BadRequest(ctx, err)
		return
	}

	// 先校验账号密码：Code 一次性目只能用 5 分钟，避免账号密码错误时白白消耗 Code
	var person model.ShPerson

	fu := facades.Database().Default().WithContext(c).First(&person, "`mobile`=?", request.Mobile)
	if fu.Error != nil || !auth.CheckPassword(request.Password, person.Password) {
		http.Fail(ctx, "手机号或密码错误")
		return
	}

	openID, err := wechat.OpenIDOfCode(c, request.Code)
	if err != nil {
		http.Fail(ctx, "微信授权失败：%v", err)
		return
	}

	// 绑定 OpenID 到当前用户；绑定失败不阻断登录
	if result := facades.Database().Default().WithContext(c).Model(&model.ShPerson{}).
		Where("`id`=?", person.ID).
		Update("openid", openID); result.Error == nil {
		person.OpenID = openID
	}

	writeLogin(ctx, person.ID, constants.JwtKindPerson, person.Name, person.Mobile, person.MustChangePassword)
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

// ToAccountOfInformation
// @Summary 获取账户信息
// @Description 获取当前登录用户的账户信息
// @Tags 基础-账户
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {object} res.ToAccountOfInformation "账户信息"
// @Router /basic/account/information [get]
func ToAccountOfInformation(c context.Context, ctx *app.RequestContext) {

	kind := jwtKind(ctx)
	id := auth.ID(ctx)

	if kind == constants.JwtKindManager {
		var manager model.ShManager
		if err := facades.Database().Default().WithContext(c).First(&manager, "`id`=?", id).Error; err != nil {
			http.Unauthorized(ctx)
			return
		}

		http.Success(ctx, res.ToAccountOfInformation{
			ID:     manager.ID,
			Name:   manager.Name,
			Mobile: manager.Mobile,
			Kind:   kind,
		})
		return
	}

	var person model.ShPerson
	if err := facades.Database().Default().WithContext(c).First(&person, "`id`=?", id).Error; err != nil {
		http.Unauthorized(ctx)
		return
	}

	http.Success(ctx, res.ToAccountOfInformation{
		ID:                 person.ID,
		Name:               person.Name,
		Mobile:             person.Mobile,
		Number:             person.Number,
		Unit:               person.Unit,
		GroupName:          person.GroupName,
		MustChangePassword: person.MustChangePassword,
		Kind:               kind,
	})
}

// DoAccountOfPassword
// @Summary 修改密码
// @Description 参赛人员修改密码
// @Tags 基础-账户
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param request body req.DoAccountOfPassword true "新密码"
// @Success 200 {object} nil "修改成功"
// @Router /basic/account/password [put]
func DoAccountOfPassword(c context.Context, ctx *app.RequestContext) {

	var request req.DoAccountOfPassword

	if err := ctx.BindAndValidate(&request); err != nil {
		http.BadRequest(ctx, err)
		return
	}

	if jwtKind(ctx) != constants.JwtKindPerson {
		http.Unauthorized(ctx)
		return
	}

	result := facades.Database().Default().WithContext(c).Model(&model.ShPerson{}).Where("`id`=?", auth.ID(ctx)).Updates(map[string]any{
		"password":             auth.Password(request.Password),
		"must_change_password": global.NO,
	})
	if result.Error != nil {
		http.Fail(ctx, "修改失败：%v", result.Error)
		return
	}

	http.Success[any](ctx)
}

func writeLogin(ctx *app.RequestContext, id, kind, name, mobile string, mustChange uint8) {

	pair, err := auth.NewLoginJWToken(id, true, map[string]any{constants.JwtKindKey: kind})
	if err != nil {
		http.Login(ctx)
		return
	}

	ctx.Header("Cache-Control", "no-store")
	ctx.Header("Pragma", "no-cache")

	http.Success(ctx, res.DoLogin{
		SessionID:          pair.SessionID,
		AccessToken:        pair.AccessToken,
		RefreshToken:       pair.RefreshToken,
		IssuedAt:           pair.IssuedAt,
		AccessLifetime:     pair.AccessLifetime,
		RefreshLifetime:    pair.RefreshLifetime,
		GraceLifetime:      pair.GraceLifetime,
		MustChangePassword: mustChange,
		Name:               name,
		Mobile:             mobile,
	})
}

func jwtKind(ctx *app.RequestContext) string {
	claims := auth.Claims(ctx)
	if claims == nil || claims.Ext == nil {
		return ""
	}
	value, _ := claims.Ext[constants.JwtKindKey].(string)
	return value
}
