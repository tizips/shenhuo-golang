package site

// ToPageBuiltinOfPaginate 内置页面分页项响应
type ToPageBuiltinOfPaginate struct {
	ID        uint   `json:"id"`         // 内置页面ID
	Key       string `json:"key"`        // 内置页面标识
	PageID    uint   `json:"page_id"`    // 关联的单页ID
	PageTitle string `json:"page_title"` // 关联的单页标题
	CreatedAt string `json:"created_at"` // 创建时间
}
