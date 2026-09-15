package site

import "github.com/herhe-com/framework/contracts/http/request"

// ToPersonOfPaginate 参赛人员分页查询请求
type ToPersonOfPaginate struct {
	Keyword string `json:"keyword" form:"keyword" query:"keyword" validate:"omitempty,max=64" label:"关键词"` // 搜索关键词；按姓名、参赛号或手机号匹配
	request.Paginate
}

// DoPersonOfCreate 创建参赛人员请求
type DoPersonOfCreate struct {
	Name      string `json:"name" form:"name" validate:"required,max=32" label:"姓名"`               // 人员姓名
	Unit      string `json:"unit" form:"unit" validate:"required,max=64" label:"单位"`               // 所属单位
	Mobile    string `json:"mobile" form:"mobile" validate:"required,mobile" label:"手机号"`          // 手机号；用于登录
	Password  string `json:"password" form:"password" validate:"required,password" label:"密码"`     // 登录密码
	Number    string `json:"number" form:"number" validate:"required,max=32" label:"参赛号"`          // 参赛号
	GroupName string `json:"group_name" form:"group_name" validate:"required,max=64" label:"小组名称"` // 所属小组名称
}

// DoPersonOfUpdate 更新参赛人员请求
type DoPersonOfUpdate struct {
	Name      string `json:"name" form:"name" validate:"required,max=32" label:"姓名"`               // 人员姓名
	Unit      string `json:"unit" form:"unit" validate:"required,max=64" label:"单位"`               // 所属单位
	Mobile    string `json:"mobile" form:"mobile" validate:"required,mobile" label:"手机号"`          // 手机号；用于登录
	Password  string `json:"password" form:"password" validate:"omitempty,password" label:"密码"`    // 新密码；留空表示不修改
	Number    string `json:"number" form:"number" validate:"required,max=32" label:"参赛号"`          // 参赛号
	GroupName string `json:"group_name" form:"group_name" validate:"required,max=64" label:"小组名称"` // 所属小组名称
	request.IDOfSnowflake
}

// DoPersonOfDelete 删除参赛人员请求
type DoPersonOfDelete struct {
	request.IDOfSnowflake
}

// ToPersonOfInformation 参赛人员详情查询请求
type ToPersonOfInformation struct {
	request.IDOfSnowflake
}
