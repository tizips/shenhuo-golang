package site

import "github.com/herhe-com/framework/contracts/http/request"

// ToManagerOfPaginate 管理员分页查询请求
type ToManagerOfPaginate struct {
	Keyword string `json:"keyword" form:"keyword" query:"keyword" validate:"omitempty,max=32" label:"关键词"` // 搜索关键词；按姓名或手机号匹配
	request.Paginate
}

// DoManagerOfCreate 创建管理员请求
type DoManagerOfCreate struct {
	Name     string `json:"name" form:"name" validate:"required,max=32" label:"姓名"`           // 管理员姓名
	Mobile   string `json:"mobile" form:"mobile" validate:"required,mobile" label:"手机号"`      // 手机号；用于登录
	Password string `json:"password" form:"password" validate:"required,password" label:"密码"` // 登录密码
	request.Enable
}

// DoManagerOfUpdate 更新管理员请求
type DoManagerOfUpdate struct {
	Name     string `json:"name" form:"name" validate:"required,max=32" label:"姓名"`            // 管理员姓名
	Mobile   string `json:"mobile" form:"mobile" validate:"required,mobile" label:"手机号"`       // 手机号；用于登录
	Password string `json:"password" form:"password" validate:"omitempty,password" label:"密码"` // 新密码；留空表示不修改
	request.IDOfSnowflake
	request.Enable
}

// DoManagerOfDelete 删除管理员请求
type DoManagerOfDelete struct {
	request.IDOfSnowflake
}

// DoManagerOfEnable 启用/停用管理员请求
type DoManagerOfEnable struct {
	request.IDOfSnowflake
	request.Enable
}
