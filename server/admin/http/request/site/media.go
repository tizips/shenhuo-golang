package site

import "github.com/herhe-com/framework/contracts/http/request"

// ToMediaOfPaginate 媒体分页查询请求
type ToMediaOfPaginate struct {
	SceneID uint   `json:"scene_id" form:"scene_id" query:"scene_id" validate:"omitempty,gt=0" label:"场景"`  // 场景ID；留空表示查全部场景
	Type    string `json:"type" form:"type" query:"type" validate:"omitempty,oneof=image video" label:"类型"` // 媒体类型；枚举：image=图片，video=视频；留空表示查全部类型
	request.Paginate
}

// DoMediaOfCreate 创建媒体请求
type DoMediaOfCreate struct {
	SceneID uint   `json:"scene_id" form:"scene_id" validate:"required,gt=0" label:"场景"`      // 所属场景ID
	Type    string `json:"type" form:"type" validate:"required,oneof=image video" label:"类型"` // 媒体类型；枚举：image=图片，video=视频
	Title   string `json:"title" form:"title" validate:"omitempty,max=128" label:"标题"`        // 媒体标题；可为空
	URL     string `json:"url" form:"url" validate:"required,max=255,url" label:"链接"`         // 媒体资源地址
	IsTop   uint8  `json:"is_top" form:"is_top" validate:"required,oneof=1 2" label:"置顶"`     // 是否置顶；枚举：1=是，2=否
}

// DoMediaOfUpdate 更新媒体请求
type DoMediaOfUpdate struct {
	DoMediaOfCreate
	request.IDOfUint
}

// DoMediaOfDelete 删除媒体请求
type DoMediaOfDelete struct {
	request.IDOfUint
}
