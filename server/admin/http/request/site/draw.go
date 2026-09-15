package site

import "github.com/herhe-com/framework/contracts/http/request"

// ToDrawOfPaginate 抽签结果分页查询请求
type ToDrawOfPaginate struct {
	CategoryID uint `json:"category_id" form:"category_id" query:"category_id" validate:"omitempty,gt=0" label:"抽签类别"` // 抽签类别ID；留空表示查全部类别
	request.Paginate
}

// DoDrawOfCreate 登记抽签结果请求
type DoDrawOfCreate struct {
	CategoryID uint `json:"category_id" form:"category_id" validate:"required,gt=0" label:"抽签类别"` // 抽签类别ID
}

// DoDrawOfDelete 删除抽签结果请求
type DoDrawOfDelete struct {
	request.IDOfUint
}
