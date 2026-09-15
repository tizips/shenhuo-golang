package site

import "github.com/herhe-com/framework/contracts/http/request"

// ToNavOfPaginate 导航分页查询请求
type ToNavOfPaginate struct {
	request.Paginate
}

// DoNavOfCreate 创建导航请求
type DoNavOfCreate struct {
	Title string `json:"title" form:"title" validate:"required,max=64" label:"标题"`        // 导航标题
	Icon  string `json:"icon" form:"icon" validate:"required,max=255,url" label:"图标"`     // 图标地址
	Type  string `json:"type" form:"type" validate:"required,oneof=link page" label:"类型"` // 导航类型；枚举：link=外链，page=单页
	Value string `json:"value" form:"value" validate:"required,max=255" label:"值"`        // 导航值；link 为跳转地址，page 为单页ID
	Order uint8  `json:"order" form:"order" validate:"omitempty,min=1,max=99" label:"序号"` // 排序序号；数值越小越靠前
}

// DoNavOfUpdate 更新导航请求
type DoNavOfUpdate struct {
	DoNavOfCreate
	request.IDOfUint
}

// DoNavOfDelete 删除导航请求
type DoNavOfDelete struct {
	request.IDOfUint
}
