package site

import "github.com/herhe-com/framework/contracts/http/request"

// ToScheduleCategoryOfPaginate 日程分类分页查询请求
type ToScheduleCategoryOfPaginate struct {
	request.Paginate
}

// DoScheduleCategoryOfCreate 创建日程分类请求
type DoScheduleCategoryOfCreate struct {
	Name  string `json:"name" form:"name" validate:"required,max=64" label:"名称"`          // 分类名称
	Order uint8  `json:"order" form:"order" validate:"omitempty,min=1,max=99" label:"序号"` // 排序序号；数值越小越靠前
}

// DoScheduleCategoryOfUpdate 更新日程分类请求
type DoScheduleCategoryOfUpdate struct {
	DoScheduleCategoryOfCreate
	request.IDOfUint
}

// DoScheduleCategoryOfDelete 删除日程分类请求
type DoScheduleCategoryOfDelete struct {
	request.IDOfUint
}
