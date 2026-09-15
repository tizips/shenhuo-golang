package site

import "github.com/herhe-com/framework/contracts/http/request"

// ScheduleOfItem 日程子项目
type ScheduleOfItem struct {
	Name string `json:"name" form:"name" validate:"required,max=120" label:"子项名称"` // 子项名称
	Time string `json:"time" form:"time" validate:"omitempty,max=64" label:"子项时间"` // 子项时间；展示用文本；可为空
}

// ToScheduleOfPaginate 日程分页查询请求
type ToScheduleOfPaginate struct {
	CategoryID uint `json:"category_id" form:"category_id" query:"category_id" validate:"omitempty,gt=0" label:"日程分类"` // 日程分类ID；留空表示查全部分类
	request.Paginate
}

// DoScheduleOfCreate 创建日程请求
type DoScheduleOfCreate struct {
	CategoryID  uint             `json:"category_id" form:"category_id" validate:"required,gt=0" label:"日程分类"`   // 日程分类ID
	Title       string           `json:"title" form:"title" validate:"required,max=120" label:"标题"`              // 日程标题
	Subtitle    string           `json:"subtitle" form:"subtitle" validate:"omitempty,max=120" label:"副标题"`      // 副标题；可为空
	Description string           `json:"description" form:"description" validate:"omitempty,max=255" label:"描述"` // 日程描述；可为空
	Time        string           `json:"time" form:"time" validate:"omitempty,max=64" label:"时间"`                // 日程时间；展示用文本
	Items       []ScheduleOfItem `json:"items" form:"items" validate:"omitempty,max=20,dive" label:"子项目"`        // 子项目（一天多场安排）；可为空
	Order       uint8            `json:"order" form:"order" validate:"omitempty,min=1,max=99" label:"序号"`        // 排序序号；数值越小越靠前
}

// DoScheduleOfUpdate 更新日程请求
type DoScheduleOfUpdate struct {
	DoScheduleOfCreate
	request.IDOfUint
}

// DoScheduleOfDelete 删除日程请求
type DoScheduleOfDelete struct {
	request.IDOfUint
}
