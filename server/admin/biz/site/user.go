package site

import (
	"context"
	"errors"
	"fmt"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/herhe-com/framework/auth"
	"github.com/herhe-com/framework/constants/global"
	"github.com/herhe-com/framework/contracts/http/response"
	"github.com/herhe-com/framework/facades"
	"github.com/herhe-com/framework/http"
	"github.com/tizips/shenhuo/model"
	req "github.com/tizips/shenhuo/server/admin/http/request/site"
	res "github.com/tizips/shenhuo/server/admin/http/response/site"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// ToUserByPaginate
// @Summary 获取用户列表
// @Description Permissions: site.user.paginate
// @Tags 站点-用户
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param page query int false "页码"
// @Param size query int false "每页数量"
// @Success 200 {object} response.Paginate[res.ToUserByPaginate] "用户列表"
// @Router /site/users [get]
func ToUserByPaginate(c context.Context, ctx *app.RequestContext) {

	var request req.ToUserByPaginate

	if err := ctx.BindAndValidate(&request); err != nil {
		http.BadRequest(ctx, err)
		return
	}

	responses := response.Paginate[res.ToUserByPaginate]{
		Page: request.GetPage(),
		Size: request.GetSize(),
	}

	ter := facades.Database().Default().WithContext(context.Background()).
		Select("1").
		Model(&model.SysUserBindRole{}).
		Where(fmt.Sprintf("`%s`.`id`=`%s`.`user_id`", model.TableSysUser, model.TableSysUserBindRole))

	if ok, _ := facades.Casbin().HasRoleForUser(auth.NameOfUser(auth.ID(ctx)), auth.NameOfDeveloper()); !ok {
		ter = ter.Where("`role_id`<>?", auth.CodeOfDeveloper)
	}

	tx := facades.Database().Default().WithContext(c).Where("exists (?)", ter)

	tx.Model(&model.SysUser{}).Count(&responses.Total)

	if responses.Total > 0 {

		var roles []model.SysUser

		tx.
			Preload("BindRoles.Role").
			Offset(request.GetOffset()).
			Limit(request.GetLimit()).
			Order("`id` desc").
			Find(&roles)

		responses.Data = make([]res.ToUserByPaginate, len(roles))

		for index, item := range roles {

			responses.Data[index] = res.ToUserByPaginate{
				ID:        item.ID,
				Nickname:  item.Nickname,
				Roles:     make([]res.ToUserByPaginateOfRoles, 0),
				IsEnable:  item.IsEnable,
				CreatedAt: item.CreatedAt.ToDateTimeString(),
			}

			if item.Username != nil {
				responses.Data[index].Username = *item.Username
			}

			if item.Mobile != nil {
				responses.Data[index].Mobile = *item.Mobile
			}

			if item.Email != nil {
				responses.Data[index].Email = *item.Email
			}

			for _, value := range item.BindRoles {
				if value.Role != nil {
					responses.Data[index].Roles = append(responses.Data[index].Roles, res.ToUserByPaginateOfRoles{
						ID:   value.Role.ID,
						Name: value.Role.Name,
					})
				}
			}
		}
	}

	http.Success(ctx, responses)

}

// DoUserByCreate
// @Summary 创建用户
// @Description Permissions: site.user.create
// @Tags 站点-用户
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param request body req.DoUserByCreate true "用户信息"
// @Success 200 {object} nil "创建成功"
// @Router /site/user [post]
func DoUserByCreate(c context.Context, ctx *app.RequestContext) {

	var request req.DoUserByCreate

	if err := ctx.BindAndValidate(&request); err != nil {
		http.BadRequest(ctx, err)
		return
	}

	var total int64 = 0

	vr := facades.Database().Default().WithContext(context.Background()).Where("`id` IN (?)", request.Roles)

	if ok, _ := facades.Casbin().HasRoleForUser(auth.NameOfUser(auth.ID(ctx)), auth.NameOfDeveloper()); !ok {
		vr = vr.Where("`id`<>?", auth.CodeOfDeveloper)
	}

	vr.Model(&model.SysRole{}).Count(&total)

	if int(total) != len(request.Roles) {
		http.Fail(ctx, "部分角色未找到")
		return
	}

	tx := facades.Database().Default().Begin()

	user := model.SysUser{
		ID:       facades.Snowflake().Generate().String(),
		Nickname: request.Nickname,
		Username: &request.Username,
		Password: auth.Password(request.Password),
		IsEnable: request.IsEnable,
	}

	if request.Mobile != "" {
		user.Mobile = &request.Mobile
	}

	if request.Email != "" {
		user.Email = &request.Email
	}

	if result := tx.Create(&user); result.Error != nil {
		tx.Rollback()
		http.Fail(ctx, "创建失败：%v", result.Error)
		return
	}

	bindings := make([]model.SysUserBindRole, len(request.Roles))

	for index, item := range request.Roles {
		bindings[index] = model.SysUserBindRole{
			Platform:       auth.Platform(ctx),
			OrganizationID: auth.Organization(ctx),
			UserID:         user.ID,
			RoleID:         item,
		}
	}

	if result := tx.Create(&bindings); result.Error != nil {
		tx.Rollback()
		http.Fail(ctx, "创建失败：%v", result.Error)
		return
	}

	items := make([]string, len(bindings))

	for index, item := range bindings {
		items[index] = auth.NameOfRole(item.RoleID)
	}

	if _, err := facades.Casbin().AddRolesForUser(auth.NameOfUser(user.ID), items, auth.SPlatform(ctx)); err != nil {
		tx.Rollback()
		http.Fail(ctx, "创建失败：%v", err)
		return
	}

	tx.Commit()

	http.Success[any](ctx)
}

// DoUserByUpdate
// @Summary 更新用户
// @Description Permissions: site.user.update
// @Tags 站点-用户
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path string true "用户ID"
// @Param request body req.DoUserByUpdate true "用户信息"
// @Success 200 {object} nil "更新成功"
// @Router /site/users/{id} [put]
func DoUserByUpdate(c context.Context, ctx *app.RequestContext) {

	var request req.DoUserByUpdate

	if err := ctx.BindAndValidate(&request); err != nil {
		http.BadRequest(ctx, err)
		return
	}

	var user model.SysUser

	fu := facades.Database().Default().
		Preload("BindRoles.Role").
		Where("exists (?)", facades.Database().Default().
			Select("1").
			Model(model.SysUserBindRole{}).
			Where(fmt.Sprintf("`%s`.`id`=`%s`.`user_id`", model.TableSysUser, model.TableSysUserBindRole)),
		).
		First(&user, "`id`=?", request.ID)
	if errors.Is(fu.Error, gorm.ErrRecordNotFound) {
		http.NotFound(ctx, "未找到该数据")
		return
	} else if fu.Error != nil {
		http.Fail(ctx, "查询失败：%v", fu.Error)
		return
	}

	//	现可用的角色
	rids := make([]uint, 0)

	for _, item := range user.BindRoles {
		if item.Role != nil {
			rids = append(rids, item.RoleID)
		}
	}

	creates := make([]model.SysUserBindRole, 0)
	deletes := make([]uint, 0)

	for _, item := range request.Roles {
		mark := true
		for _, value := range rids {
			if item == value {
				mark = false
			}
		}
		if mark {
			creates = append(creates, model.SysUserBindRole{
				UserID: user.ID,
				RoleID: item,
			})
		}
	}

	for _, item := range rids {
		mark := true
		for _, value := range request.Roles {
			if item == value {
				mark = false
			}
		}
		if mark {
			deletes = append(deletes, item)
		}
	}

	if len(creates) > 0 {

		var total int64 = 0

		vr := facades.Database().Default().WithContext(context.Background()).
			Where("`id` IN (?)", request.Roles)

		if ok, _ := facades.Casbin().HasRoleForUser(auth.NameOfUser(auth.ID(ctx)), auth.NameOfDeveloper()); !ok {
			vr = vr.Where("`id`<>?", auth.CodeOfDeveloper)
		}

		vr.Model(&model.SysRole{}).Count(&total)

		if int(total) != len(request.Roles) {
			http.Fail(ctx, "部分角色未找到")
			return
		}
	}

	tx := facades.Database().Default().Begin()

	user.Nickname = request.Nickname
	user.IsEnable = request.IsEnable

	if request.Password != "" {
		user.Password = auth.Password(request.Password)
	}

	if request.Mobile != "" {
		user.Mobile = &request.Mobile
	}

	if request.Email != "" {
		user.Email = &request.Email
	}

	if uu := tx.Omit(clause.Associations).Save(&user); uu.Error != nil {
		tx.Rollback()
		http.Fail(ctx, "修改失败：%v", uu.Error)
		return
	}

	if len(creates) > 0 {

		if cb := tx.Create(&creates); cb.Error != nil {
			tx.Rollback()
			http.Fail(ctx, "修改失败：%v", cb.Error)
			return
		}

		items := make([]string, len(creates))

		for index, item := range creates {
			items[index] = auth.NameOfRole(item.RoleID)
		}

		if _, err := facades.Casbin().AddRolesForUser(auth.NameOfUser(user.ID), items); err != nil {
			tx.Rollback()
			http.Fail(ctx, "修改失败：%v", err)
			return
		}
	}

	if len(deletes) > 0 {

		if db := tx.Delete(&model.SysUserBindRole{}, "`user_id`=? and `role_id` IN (?)", user.ID, deletes); db.Error != nil {
			tx.Rollback()
			http.Fail(ctx, "修改失败：%v", db.Error)
			return
		}

		for _, item := range deletes {

			if _, err := facades.Casbin().DeleteRoleForUser(auth.NameOfUser(user.ID), auth.NameOfRole(item)); err != nil {
				tx.Rollback()
				http.Fail(ctx, "修改失败：%v", err)
				return
			}
		}
	}

	if request.IsEnable != user.IsEnable {

		var bindings []model.SysUserBindRole

		_ = facades.Database().Default().Model(&user).Association("BindRoles").Find(&bindings)

		if request.IsEnable == global.YES && user.IsEnable == global.NO {
			//	启用用户角色

			items := make([]string, len(bindings))

			for index, item := range bindings {
				items[index] = auth.NameOfRole(item.RoleID)
			}

			if _, err := facades.Casbin().AddRolesForUser(auth.NameOfUser(user.ID), items); err != nil {
				tx.Rollback()
				http.Fail(ctx, "修改失败：%v", err)
				return
			}

		} else if request.IsEnable == global.NO && user.IsEnable == global.YES {
			//	禁用用户角色

			for _, item := range bindings {

				if _, err := facades.Casbin().DeleteRoleForUser(auth.NameOfUser(user.ID), auth.NameOfRole(item.RoleID)); err != nil {
					tx.Rollback()
					http.Fail(ctx, "修改失败：%v", err)
					return
				}
			}
		}
	}

	tx.Commit()

	http.Success[any](ctx)
}

// DoUserByDelete
// @Summary 删除用户
// @Description Permissions: site.user.delete
// @Tags 站点-用户
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path string true "用户ID"
// @Success 200 {object} nil "删除成功"
// @Router /site/users/{id} [delete]
func DoUserByDelete(c context.Context, ctx *app.RequestContext) {

	id := ctx.Param("id")

	var user model.SysUser

	fu := facades.Database().Default().
		Preload("BindRoles", func(t *gorm.DB) *gorm.DB {

			return t.
				Where("exists (?)", facades.Database().Default().
					Select("1").
					Model(&model.SysRole{}).
					Where(fmt.Sprintf("`%s`.`role_id`=`%s`.`id`", model.TableSysUserBindRole, model.TableSysRole)),
				)
		}).
		First(&user, "`id`=?", id)
	if errors.Is(fu.Error, gorm.ErrRecordNotFound) {
		http.NotFound(ctx, "未找到该数据")
		return
	} else if fu.Error != nil {
		http.Fail(ctx, "查询失败：%v", fu.Error)
		return
	}

	tx := facades.Database().Default().Begin()

	if du := tx.Delete(&user); du.Error != nil {
		tx.Rollback()
		http.Fail(ctx, "删除失败：%v", du.Error)
		return
	}

	for _, item := range user.BindRoles {

		if _, err := facades.Casbin().DeleteRoleForUser(auth.NameOfUser(user.ID), auth.NameOfRole(item.RoleID)); err != nil {
			tx.Rollback()
			http.Fail(ctx, "删除失败：%v", err)
			return
		}
	}

	tx.Commit()

	http.Success[any](ctx)
}

// DoUserByEnable
// @Summary 启用/禁用用户
// @Description Permissions: site.user.enable
// @Tags 站点-用户
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param request body req.DoUserByEnable true "启用信息"
// @Success 200 {object} nil "操作成功"
// @Router /site/user/enable [put]
func DoUserByEnable(c context.Context, ctx *app.RequestContext) {

	var request req.DoUserByEnable

	if err := ctx.BindAndValidate(&request); err != nil {
		http.BadRequest(ctx, err)
		return
	}

	var user model.SysUser

	fu := facades.Database().Default().
		Preload("BindRoles").
		Where("exists (?)", facades.Database().Default().
			Select("1").
			Model(model.SysUserBindRole{}).
			Where(fmt.Sprintf("`%s`.`id`=`%s`.`user_id`", model.TableSysUser, model.TableSysUserBindRole)),
		).
		First(&user, "`id`=?", request.ID)
	if errors.Is(fu.Error, gorm.ErrRecordNotFound) {
		http.NotFound(ctx, "未找到该数据")
		return
	} else if fu.Error != nil {
		http.Fail(ctx, "查询失败：%v", fu.Error)
		return
	}

	tx := facades.Database().Default().Begin()

	if uu := tx.Omit(clause.Associations).Model(user).Update("is_enable", request.IsEnable); uu.Error != nil {
		tx.Rollback()
		http.Fail(ctx, "删除失败：%v", uu.Error)
		return
	}

	if request.IsEnable != user.IsEnable {

		if request.IsEnable == global.YES && user.IsEnable == global.NO {
			//	启用用户角色

			items := make([]string, len(user.BindRoles))

			for index, item := range user.BindRoles {
				items[index] = auth.NameOfRole(item.RoleID)
			}

			if _, err := facades.Casbin().AddRolesForUser(auth.NameOfUser(user.ID), items); err != nil {
				tx.Rollback()
				http.Fail(ctx, "修改失败：%v", err)
				return
			}

		} else if request.IsEnable == global.NO && user.IsEnable == global.YES {
			//	禁用用户角色

			for _, item := range user.BindRoles {

				if _, err := facades.Casbin().DeleteRoleForUser(auth.NameOfUser(user.ID), auth.NameOfRole(item.RoleID)); err != nil {
					tx.Rollback()
					http.Fail(ctx, "修改失败：%v", err)
					return
				}
			}
		}
	}

	tx.Commit()

	http.Success[any](ctx)
}
