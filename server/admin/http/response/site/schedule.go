package site

// ToScheduleOfItem 日程子项目响应
type ToScheduleOfItem struct {
	Name string `json:"name"` // 子项名称
	Time string `json:"time"` // 子项时间；展示用文本
}

// ToScheduleOfPaginate 日程分页项响应
type ToScheduleOfPaginate struct {
	ID           uint               `json:"id"`            // 日程ID
	CategoryID   uint               `json:"category_id"`   // 日程分类ID
	CategoryName string             `json:"category_name"` // 日程分类名称
	Title        string             `json:"title"`         // 日程标题
	Subtitle     string             `json:"subtitle"`      // 副标题
	Description  string             `json:"description"`   // 日程描述
	Time         string             `json:"time"`          // 日程时间；展示用文本
	Items        []ToScheduleOfItem `json:"items"`         // 子项目（一天多场安排）
	Order        uint8              `json:"order"`         // 排序序号；数值越小越靠前
	CreatedAt    string             `json:"created_at"`    // 创建时间
}
