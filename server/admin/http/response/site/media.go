package site

// ToMediaOfPaginate 媒体分页项响应
type ToMediaOfPaginate struct {
	ID        uint   `json:"id"`         // 媒体ID
	SceneID   uint   `json:"scene_id"`   // 所属场景ID
	Scene     string `json:"scene"`      // 所属场景名称
	Type      string `json:"type"`       // 媒体类型；枚举：image=图片，video=视频
	Title     string `json:"title"`      // 媒体标题
	URL       string `json:"url"`        // 媒体资源地址
	IsTop     uint8  `json:"is_top"`     // 是否置顶；枚举：1=是，2=否
	CreatedAt string `json:"created_at"` // 创建时间
}
