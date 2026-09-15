package shenhuo

// DoDrawOfCreate 抽签请求
type DoDrawOfCreate struct {
	CategoryID uint `json:"category_id" form:"category_id" validate:"required,gt=0" label:"抽签类别"` // 抽签类别ID
}
