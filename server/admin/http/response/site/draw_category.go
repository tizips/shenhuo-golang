package site

// ToDrawCategoryOfPaginate 抽签类别分页项响应
type ToDrawCategoryOfPaginate struct {
	ID        uint   `json:"id"`         // 类别ID
	Name      string `json:"name"`       // 类别名称
	Icon      string `json:"icon"`       // 图标地址
	Quota     uint   `json:"quota"`      // 中签名额人数
	Order     uint8  `json:"order"`      // 排序序号；数值越小越靠前
	Drawn     int64  `json:"drawn"`      // 已中签人数
	CreatedAt string `json:"created_at"` // 创建时间
}
