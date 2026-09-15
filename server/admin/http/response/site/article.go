package site

// ToArticleOfPaginate 文章分页项响应
type ToArticleOfPaginate struct {
	ID          uint   `json:"id"`           // 文章ID
	Title       string `json:"title"`        // 文章标题
	Thumb       string `json:"thumb"`        // 缩略图地址
	PublishedAt string `json:"published_at"` // 发布时间
	IsTop       uint8  `json:"is_top"`       // 是否置顶；枚举：1=是，2=否
	IsRecommend uint8  `json:"is_recommend"` // 是否首页推荐；枚举：1=是，2=否
	CreatedAt   string `json:"created_at"`   // 创建时间
}

// ToArticleOfInformation 文章详情响应
type ToArticleOfInformation struct {
	ID          uint   `json:"id"`           // 文章ID
	Title       string `json:"title"`        // 文章标题
	Thumb       string `json:"thumb"`        // 缩略图地址
	Content     string `json:"content"`      // 正文内容；富文本
	PublishedAt string `json:"published_at"` // 发布时间
	IsTop       uint8  `json:"is_top"`       // 是否置顶；枚举：1=是，2=否
	IsRecommend uint8  `json:"is_recommend"` // 是否首页推荐；枚举：1=是，2=否
	CreatedAt   string `json:"created_at"`   // 创建时间
}
