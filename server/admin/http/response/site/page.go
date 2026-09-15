package site

// ToPageOfPaginate 单页分页项响应
type ToPageOfPaginate struct {
	ID        uint   `json:"id"`         // 单页ID
	Title     string `json:"title"`      // 单页标题
	CreatedAt string `json:"created_at"` // 创建时间
}

// ToPageOfInformation 单页详情响应
type ToPageOfInformation struct {
	ID        uint   `json:"id"`         // 单页ID
	Title     string `json:"title"`      // 单页标题
	Content   string `json:"content"`    // 单页内容；富文本
	CreatedAt string `json:"created_at"` // 创建时间
}
