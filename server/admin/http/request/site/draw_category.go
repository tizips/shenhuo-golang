package site

import "github.com/herhe-com/framework/contracts/http/request"

// ToDrawCategoryOfPaginate 抽签类别分页查询请求
type ToDrawCategoryOfPaginate struct {
	request.Paginate
}

// DoDrawCategoryOfCreate 创建抽签类别请求
type DoDrawCategoryOfCreate struct {
	Name  string `json:"name" form:"name" validate:"required,max=64" label:"名称"`          // 类别名称
	Icon  string `json:"icon" form:"icon" validate:"required,max=255,url" label:"图标"`     // 图标地址
	Quota uint   `json:"quota" form:"quota" validate:"required,gt=0" label:"中签人数"`        // 中签名额人数
	Order uint8  `json:"order" form:"order" validate:"omitempty,min=1,max=99" label:"序号"` // 排序序号；数值越小越靠前
}

// DoDrawCategoryOfUpdate 更新抽签类别请求
type DoDrawCategoryOfUpdate struct {
	DoDrawCategoryOfCreate
	request.IDOfUint
}

// DoDrawCategoryOfDelete 删除抽签类别请求
type DoDrawCategoryOfDelete struct {
	request.IDOfUint
}
