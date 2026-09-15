package site

import "github.com/herhe-com/framework/contracts/http/request"

// ToPageOfPaginate 单页分页查询请求
type ToPageOfPaginate struct {
	Keyword string `json:"keyword" form:"keyword" query:"keyword" validate:"omitempty,max=120" label:"关键词"` // 搜索关键词；按标题模糊匹配
	request.Paginate
}

// DoPageOfCreate 创建单页请求
type DoPageOfCreate struct {
	Title   string `json:"title" form:"title" validate:"required,max=120" label:"标题"` // 单页标题
	Content string `json:"content" form:"content" validate:"required" label:"内容"`     // 单页内容；富文本
}

// DoPageOfUpdate 更新单页请求
type DoPageOfUpdate struct {
	DoPageOfCreate
	request.IDOfUint
}

// DoPageOfDelete 删除单页请求
type DoPageOfDelete struct {
	request.IDOfUint
}

// ToPageOfInformation 单页详情查询请求
type ToPageOfInformation struct {
	request.IDOfUint
}
