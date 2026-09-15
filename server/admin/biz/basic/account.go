package basic

import (
	"context"
	"errors"
	"fmt"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/herhe-com/framework/auth"
	"github.com/herhe-com/framework/facades"
	"github.com/herhe-com/framework/http"
	"github.com/samber/lo"
	"github.com/tizips/shenhuo/model"
	req "github.com/tizips/shenhuo/server/admin/http/request/basic"
	res "github.com/tizips/shenhuo/server/admin/http/response/basic"
	"gorm.io/gorm"
)

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

	var user model.SysUser

	fu := facades.Database().Default().First(&user, "`id`=?", auth.ID(ctx))

	if errors.Is(fu.Error, gorm.ErrRecordNotFound) {
		http.Unauthorized(ctx)
		return
	}

	responses := res.ToAccountOfInformation{
		Nickname: user.Nickname,
		Avatar:   user.Avatar,
	}

	if user.Username != nil {
		responses.Username = *user.Username
	}

	if user.Mobile != nil {
		responses.Mobile = *user.Mobile
	}

	if user.Email != nil {
		responses.Username = *user.Email
	}

	responses.Platform.Code = auth.Platform(ctx)
	responses.Platform.Name = facades.Config().GetString("app.title")

	http.Success(ctx, responses)
}

// ToAccountOfModules
// @Summary 获取模块列表
// @Description 获取当前用户可访问的模块列表
// @Tags 基础-账户
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {array} res.ToAccountOfModules "模块列表"
// @Router /basic/account/modules [get]
func ToAccountOfModules(c context.Context, ctx *app.RequestContext) {

	responses := make([]res.ToAccountOfModules, 0)

	modules := auth.Modules(auth.Platform(ctx))

	if ok, _ := facades.Casbin().HasRoleForUser(auth.NameOfUser(auth.ID(ctx)), auth.NameOfDeveloper()); ok {

		for _, item := range modules {
			responses = append(responses, res.ToAccountOfModules{
				Code: item.Code,
				Name: item.Name,
			})
		}

	} else {

		var moduleCodes []string

		facades.Database().Default().
			Model(&model.SysRoleBindPermission{}).
			Distinct("module").
			Where("exists (?)", facades.Database().Default().
				Model(&model.SysUserBindRole{}).
				Select("1").
				Where(fmt.Sprintf("`%s`.`role_id`=`%s`.`role_id` and `%s`.`user_id`=?", model.TableSysRoleBindPermission, model.TableSysUserBindRole, model.TableSysUserBindRole), auth.ID(ctx)),
			).
			Where("exists (?)", facades.Database().Default().
				Model(&model.SysRole{}).
				Select("1").
				Where(fmt.Sprintf("`%s`.`role_id`=`%s`.`id`", model.TableSysRoleBindPermission, model.TableSysRole)),
			).
			Pluck("module", &moduleCodes)

		for _, item := range modules {

			if lo.Contains(moduleCodes, item.Code) {

				responses = append(responses, res.ToAccountOfModules{
					Code: item.Code,
					Name: item.Name,
				})
			}
		}
	}

	http.Success(ctx, responses)
}

// ToAccountOfPermissions
// @Summary 获取权限列表
// @Description 获取当前用户在指定模块下的权限列表
// @Tags 基础-账户
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param module query string true "模块代码"
// @Success 200 {array} string "权限列表"
// @Router /basic/account/permissions [get]
func ToAccountOfPermissions(c context.Context, ctx *app.RequestContext) {

	var request req.ToAccountOfPermissions

	if err := ctx.BindAndValidate(&request); err != nil {
		http.BadRequest(ctx, err)
		return
	}

	responses := make([]string, 0)

	if ok, _ := facades.Casbin().HasRoleForUser(auth.NameOfUser(auth.ID(ctx)), auth.NameOfDeveloper()); ok {

		modules := auth.Modules(auth.Platform(ctx))

		for _, item := range modules {
			if item.Code == request.Module {
				responses = item.Permissions
				break
			}
		}

	} else {

		facades.Database().Default().
			Model(&model.SysRoleBindPermission{}).
			Where("module = ?", request.Module).
			Where("exists (?)", facades.Database().Default().
				Model(&model.SysUserBindRole{}).
				Select("1").
				Where(fmt.Sprintf("`%s`.`role_id`=`%s`.`role_id` and `%s`.`user_id`=?", model.TableSysRoleBindPermission, model.TableSysUserBindRole, model.TableSysUserBindRole), auth.ID(ctx)),
			).
			Where("exists (?)", facades.Database().Default().
				Model(&model.SysRole{}).
				Select("1").
				Where(fmt.Sprintf("`%s`.`role_id`=`%s`.`id`", model.TableSysRoleBindPermission, model.TableSysRole)),
			).
			Pluck("permission", &responses)
	}

	http.Success(ctx, responses)
}

// DoAccount
// @Summary 更新账户信息
// @Description 更新当前登录用户的账户信息
// @Tags 基础-账户
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param request body req.DoAccount true "账户信息"
// @Success 200 {object} nil "更新成功"
// @Router /basic/account [put]
func DoAccount(c context.Context, ctx *app.RequestContext) {

	var request req.DoAccount

	if err := ctx.BindAndValidate(&request); err != nil {
		http.BadRequest(ctx, err)
		return
	}

	updates := make(map[string]any)

	if request.Mobile != "" {
		updates["mobile"] = request.Mobile
	}

	if request.Email != "" {
		updates["email"] = request.Email
	}

	if request.Password != "" {
		updates["password"] = auth.Password(request.Password)
	}

	if len(updates) > 0 {

		if result := facades.Database().Default().Model(&model.SysUser{}).Where("`id` = ?", auth.ID(ctx)).Updates(updates); result.Error != nil {
			http.Fail(ctx, "修改失败：%v", result.Error)
			return
		}
	}

	http.Success[any](ctx)
}
