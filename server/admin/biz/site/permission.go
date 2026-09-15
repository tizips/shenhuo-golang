package site

import (
	"context"
	"fmt"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/herhe-com/framework/auth"
	authConstants "github.com/herhe-com/framework/contracts/auth"
	"github.com/herhe-com/framework/database/orm/scope"
	"github.com/herhe-com/framework/facades"
	"github.com/herhe-com/framework/http"
	"github.com/tizips/shenhuo/model"
)

func ToPermissions(c context.Context, ctx *app.RequestContext) {

	var responses []authConstants.Tree

	var codes []string

	if ok, _ := facades.Casbin().HasRoleForUser(auth.NameOfUser(auth.ID(ctx)), auth.NameOfDeveloper()); !ok {

		facades.Database().Default().
			Scopes(scope.Platform(ctx)).
			Model(&model.SysRoleBindPermission{}).
			Where("exists (?)", facades.Database().Default().
				Model(&model.SysUserBindRole{}).
				Select("1").
				Where(fmt.Sprintf("`%s`.`role_id`=`%s`.`role_id` and `%s`.`user_id`=?", model.TableSysRoleBindPermission, model.TableSysUserBindRole, model.TableSysUserBindRole), auth.ID(ctx)),
			).
			Where("exists (?)", facades.Database().Default().
				Model(&model.SysRole{}).
				Select("1").Where(fmt.Sprintf("`%s`.`role_id`=`%s`.`id`", model.TableSysRoleBindPermission, model.TableSysRole)),
			).
			Pluck("permission", &codes)
	}

	responses = auth.Trees(auth.Platform(ctx), false, codes)

	http.Success(ctx, responses)
}
