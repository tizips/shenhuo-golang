package site

// ToBannerOfPaginate 轮播图分页项响应
type ToBannerOfPaginate struct {
	ID        uint   `json:"id"`                                       // 轮播图ID
	Title     string `json:"title"`                                    // 轮播图标题
	Image     string `json:"image"`                                    // 图片地址
	Link      string `json:"link"`                                     // 点击跳转链接；可为空
	Order     uint8  `json:"order"`                                    // 排序序号；数值越小越靠前
	StartedAt string `json:"started_at" example:"2026-09-30 00:00:00"` // 生效开始时间
	EndedAt   string `json:"ended_at" example:"2026-10-31 23:59:59"`   // 生效结束时间
	CreatedAt string `json:"created_at"`                               // 创建时间
}
