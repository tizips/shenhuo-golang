package shenhuo

import "github.com/herhe-com/framework/contracts/http/request"

// ToMediaOfPaginate 媒体分页查询请求
type ToMediaOfPaginate struct {
	request.Paginate
	SceneID uint   `json:"scene_id" form:"scene_id" validate:"omitempty,gt=0" label:"场景"`      // 所属场景ID；可为空
	Type    string `json:"type" form:"type" validate:"omitempty,oneof=image video" label:"类型"` // 媒体类型；枚举：image=图片，video=视频
}

// DoMediaOfCreateByImage 批量添加图片请求
type DoMediaOfCreateByImage struct {
	SceneID uint     `json:"scene_id" form:"scene_id" validate:"required,gt=0" label:"场景"`                             // 所属场景ID
	URLs    []string `json:"urls" form:"urls" validate:"required,min=1,max=20,dive,required,url,max=255" label:"图片链接"` // 图片资源地址列表
}

// DoMediaOfCreateByVideo 添加视频请求
type DoMediaOfCreateByVideo struct {
	SceneID uint   `json:"scene_id" form:"scene_id" validate:"required,gt=0" label:"场景"` // 所属场景ID
	Title   string `json:"title" form:"title" validate:"required,max=128" label:"标题"`    // 视频标题
	URL     string `json:"url" form:"url" validate:"required,max=255,url" label:"链接"`    // 视频资源地址
}
