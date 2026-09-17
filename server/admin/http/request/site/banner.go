package site

import "github.com/herhe-com/framework/contracts/http/request"

// ToBannerOfPaginate 轮播图分页查询请求
type ToBannerOfPaginate struct {
	request.Paginate
}

// DoBannerOfCreate 创建轮播图请求
type DoBannerOfCreate struct {
	Title string `json:"title" form:"title" validate:"required,max=64" label:"标题"`        // 轮播图标题
	Image string `json:"image" form:"image" validate:"required,max=255,url" label:"图片"`   // 图片地址
	Link  string `json:"link" form:"link" validate:"omitempty,max=255" label:"跳转链接"`      // 点击跳转链接；可为空
	Order uint8  `json:"order" form:"order" validate:"omitempty,min=1,max=99" label:"序号"` // 排序序号；数值越小越靠前
}

// DoBannerOfUpdate 更新轮播图请求
type DoBannerOfUpdate struct {
	DoBannerOfCreate
	request.IDOfUint
}

// DoBannerOfDelete 删除轮播图请求
type DoBannerOfDelete struct {
	request.IDOfUint
}
