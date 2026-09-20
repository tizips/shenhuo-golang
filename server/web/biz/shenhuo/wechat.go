package shenhuo

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/herhe-com/framework/http"
	"github.com/tizips/shenhuo/wechat"
)

// ToWechatOfOAuth
// @Summary 微信服务号网页授权跳转
// @Description 跳转到微信服务号网页授权页面，授权后微信将重定向到配置的回调地址（wechat.official_account.oauth.redirect）并附带 code 与 state 参数
// @Tags 微信
// @Produce json
// @Param state query string false "透传给回调地址的 state 参数"
// @Success 302 {string} string "跳转到微信授权页面"
// @Router /wechat/oauth [get]
func ToWechatOfOAuth(c context.Context, ctx *app.RequestContext) {

	location, err := wechat.OAuthURL(string(ctx.QueryArgs().Peek("state")))
	if err != nil {
		http.Fail(ctx, "微信授权跳转失败：%v", err)
		return
	}

	ctx.Redirect(consts.StatusFound, []byte(location))
}
