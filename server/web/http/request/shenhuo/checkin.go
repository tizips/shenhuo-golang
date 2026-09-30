package shenhuo

// DoCheckinOfCreate 签到请求
type DoCheckinOfCreate struct {
	Name string `json:"name" form:"name" validate:"required,max=32" label:"姓名"` // 签到姓名
}
