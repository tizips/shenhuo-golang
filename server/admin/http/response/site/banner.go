package site

// ToBannerOfPaginate 轮播图分页项响应
type ToBannerOfPaginate struct {
	ID        uint   `json:"id"`         // 轮播图ID
	Title     string `json:"title"`      // 轮播图标题
	Image     string `json:"image"`      // 图片地址
	Link      string `json:"link"`       // 点击跳转链接；可为空
	Order     uint8  `json:"order"`      // 排序序号；数值越小越靠前
	CreatedAt string `json:"created_at"` // 创建时间
}
