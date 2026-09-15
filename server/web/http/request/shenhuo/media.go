package shenhuo

import "github.com/herhe-com/framework/contracts/http/request"

// ToMediaOfPaginate 媒体分页查询请求
type ToMediaOfPaginate struct {
	request.Paginate
	SceneID uint   `json:"scene_id" form:"scene_id" validate:"omitempty,gt=0" label:"场景"`      // 所属场景ID；可为空
	Type    string `json:"type" form:"type" validate:"omitempty,oneof=image video" label:"类型"` // 媒体类型；枚举：image=图片，video=视频
}

// DoMediaOfCreate 创建媒体请求
type DoMediaOfCreate struct {
	SceneID uint   `json:"scene_id" form:"scene_id" validate:"required,gt=0" label:"场景"`      // 所属场景ID
	Type    string `json:"type" form:"type" validate:"required,oneof=image video" label:"类型"` // 媒体类型；枚举：image=图片，video=视频
	Title   string `json:"title" form:"title" validate:"omitempty,max=128" label:"标题"`        // 媒体标题；可为空
	URL     string `json:"url" form:"url" validate:"required,max=255,url" label:"链接"`         // 媒体资源地址
}
