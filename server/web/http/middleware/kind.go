package middleware

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/herhe-com/framework/auth"
	"github.com/herhe-com/framework/http"
	"github.com/tizips/shenhuo/server/web/constants"
)

func Kind(kind string) app.HandlerFunc {

	return func(c context.Context, ctx *app.RequestContext) {

		if !auth.Check(ctx) {
			ctx.Abort()
			http.Unauthorized(ctx)
			return
		}

		claims := auth.Claims(ctx)
		if claims == nil || claims.Ext == nil {
			ctx.Abort()
			http.Unauthorized(ctx)
			return
		}

		value, _ := claims.Ext[constants.JwtKindKey].(string)
		if value != kind {
			ctx.Abort()
			http.Unauthorized(ctx)
			return
		}

		ctx.Next(c)
	}
}

func Manager() app.HandlerFunc {
	return Kind(constants.JwtKindManager)
}

func Person() app.HandlerFunc {
	return Kind(constants.JwtKindPerson)
}
