package site

// ToScheduleCategoryOfPaginate 日程分类分页项响应
type ToScheduleCategoryOfPaginate struct {
	ID        uint   `json:"id"`         // 分类ID
	Name      string `json:"name"`       // 分类名称
	Order     uint8  `json:"order"`      // 排序序号；数值越小越靠前
	Total     int64  `json:"total"`      // 分类下的日程数量
	CreatedAt string `json:"created_at"` // 创建时间
}
