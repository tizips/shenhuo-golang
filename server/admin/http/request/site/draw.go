package site

import "github.com/herhe-com/framework/contracts/http/request"

// ToDrawOfList 抽签结果查询请求
type ToDrawOfList struct {
	CategoryID uint `json:"category_id" form:"category_id" query:"category_id" validate:"required,gt=0" label:"抽签类别"` // 抽签类别ID
}

// DoDrawOfCreate 登记抽签结果请求
type DoDrawOfCreate struct {
	CategoryID uint `json:"category_id" form:"category_id" validate:"required,gt=0" label:"抽签类别"` // 抽签类别ID
}

// DoDrawOfDelete 删除抽签结果请求
type DoDrawOfDelete struct {
	request.IDOfUint
}
