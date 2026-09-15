package site

// ToNavOfPaginate 导航分页项响应
type ToNavOfPaginate struct {
	ID        uint   `json:"id"`         // 导航ID
	Title     string `json:"title"`      // 导航标题
	Icon      string `json:"icon"`       // 图标地址
	Type      string `json:"type"`       // 导航类型；枚举：link=外链，page=单页
	Value     string `json:"value"`      // 导航值；link 为跳转地址，page 为单页ID
	PageName  string `json:"page_name"`  // 页面名称；type=page 时为关联单页的标题
	Order     uint8  `json:"order"`      // 排序序号；数值越小越靠前
	CreatedAt string `json:"created_at"` // 创建时间
}
