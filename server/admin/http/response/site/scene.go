package site

// ToSceneOfPaginate 场景分页项响应
type ToSceneOfPaginate struct {
	ID        uint   `json:"id"`         // 场景ID
	Name      string `json:"name"`       // 场景名称
	Order     uint8  `json:"order"`      // 排序序号；数值越小越靠前
	CreatedAt string `json:"created_at"` // 创建时间
}
