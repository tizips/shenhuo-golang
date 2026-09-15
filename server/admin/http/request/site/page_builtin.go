package site

import "github.com/herhe-com/framework/contracts/http/request"

// ToPageBuiltinOfPaginate 内置页面分页查询请求
type ToPageBuiltinOfPaginate struct {
	Keyword string `json:"keyword" form:"keyword" query:"keyword" validate:"omitempty,max=120" label:"关键词"` // 搜索关键词；按标识模糊匹配
	request.Paginate
}

// DoPageBuiltinOfCreate 创建内置页面请求
type DoPageBuiltinOfCreate struct {
	Key    string `json:"key" form:"key" validate:"required,max=64" label:"标识"`         // 内置页面标识；全局唯一
	PageID uint   `json:"page_id" form:"page_id" validate:"required,gt=0" label:"页面ID"` // 关联的单页ID
}

// DoPageBuiltinOfUpdate 更新内置页面请求
type DoPageBuiltinOfUpdate struct {
	DoPageBuiltinOfCreate
	request.IDOfUint
}

// DoPageBuiltinOfDelete 删除内置页面请求
type DoPageBuiltinOfDelete struct {
	request.IDOfUint
}
