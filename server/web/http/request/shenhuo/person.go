package shenhuo

import "github.com/herhe-com/framework/contracts/http/request"

// ToPersonOfList 人员列表查询请求
type ToPersonOfList struct {
	CategoryID uint `json:"category_id" form:"category_id" query:"category_id" validate:"required,gt=0" label:"抽签组别"` // 抽签组别；抽签类别ID
}

// ToPersonOfGroupInformation 分组人员查询请求
type ToPersonOfGroupInformation struct {
	request.IDOfUint
}
