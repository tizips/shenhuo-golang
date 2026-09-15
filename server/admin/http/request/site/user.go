package site

import "github.com/herhe-com/framework/contracts/http/request"

// ToUserByPaginate 用户分页查询请求
type ToUserByPaginate struct {
	request.Paginate
}

// DoUserByCreate 创建用户请求
type DoUserByCreate struct {
	Username string `json:"username" form:"username" validate:"required,username" label:"用户名"` // 登录用户名
	Nickname string `json:"nickname" form:"nickname" validate:"required,max=32" label:"昵称"`    // 用户昵称
	Mobile   string `json:"mobile" form:"mobile" validate:"omitempty,mobile" label:"手机号"`      // 手机号；可为空
	Email    string `json:"email" form:"email" validate:"omitempty,email,max=64" label:"邮箱"`   // 邮箱；可为空
	Password string `json:"password" form:"password" validate:"required,password" label:"密码"`  // 登录密码
	Roles    []uint `json:"roles" form:"roles" validate:"required,min=1,unique" label:"角色"`    // 绑定的角色ID列表

	request.Enable
}

// DoUserByUpdate 更新用户请求
type DoUserByUpdate struct {
	request.IDOfSnowflake

	Nickname string `json:"nickname" form:"nickname" validate:"required,max=32" label:"昵称"`    // 用户昵称
	Mobile   string `json:"mobile" form:"mobile" validate:"omitempty,mobile" label:"手机号"`      // 手机号；可为空
	Email    string `json:"email" form:"email" validate:"omitempty,email,max=64" label:"邮箱"`   // 邮箱；可为空
	Password string `json:"password" form:"password" validate:"omitempty,password" label:"密码"` // 新密码；留空表示不修改
	Roles    []uint `json:"roles" form:"roles" validate:"required,min=1,unique" label:"角色"`    // 绑定的角色ID列表

	request.Enable
}

// DoUserByEnable 启用/停用用户请求
type DoUserByEnable struct {
	request.IDOfSnowflake
	request.Enable
}
