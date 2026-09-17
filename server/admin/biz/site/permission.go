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

// ToPermissions
// @Summary 获取权限树
// @Description 获取当前用户可分配的权限树
// @Tags 站点-权限
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {array} authConstants.Tree "权限树"
// @Router /site/permissions [get]
func ToPermissions(c context.Context, ctx *app.RequestContext) {

	var responses []authConstants.Tree

	//	开发者返回完整权限树

	if ok, _ := facades.Casbin().HasRoleForUser(auth.NameOfUser(auth.ID(ctx)), auth.NameOfDeveloper()); ok {

		responses = auth.Trees(auth.Platform(ctx), true)

	} else {

		var codes []string

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

		responses = auth.Trees(auth.Platform(ctx), false, codes)
	}

	http.Success(ctx, responses)
}
