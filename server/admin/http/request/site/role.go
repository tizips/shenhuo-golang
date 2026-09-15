package site

import "github.com/herhe-com/framework/contracts/http/request"

// ToRoleByPaginate 角色分页查询请求
type ToRoleByPaginate struct {
	request.Paginate
}

// DoRoleByCreate 创建角色请求
type DoRoleByCreate struct {
	Name        string   `json:"name" form:"name" validate:"required,max=32" label:"名称"`                       // 角色名称
	Permissions []string `json:"permissions" form:"permissions[]" validate:"required,min=1,unique" label:"权限"` // 权限标识列表
	Summary     string   `json:"summary" form:"summary" validate:"omitempty,max=255" label:"简介"`               // 角色简介；可为空
}

// DoRoleByUpdate 更新角色请求
type DoRoleByUpdate struct {
	Name        string   `json:"name" form:"name" validate:"required,max=32" label:"名称"`                     // 角色名称
	Permissions []string `json:"permissions" form:"permissions" validate:"required,min=1,unique" label:"权限"` // 权限标识列表
	Summary     string   `json:"summary" form:"summary" validate:"omitempty,max=255" label:"简介"`             // 角色简介；可为空

	request.IDOfUint
}
