package site

import "github.com/herhe-com/framework/contracts/http/request"

// ToArticleOfPaginate 文章分页查询请求
type ToArticleOfPaginate struct {
	Keyword string `json:"keyword" form:"keyword" query:"keyword" validate:"omitempty,max=120" label:"关键词"` // 搜索关键词；按标题模糊匹配
	request.Paginate
}

// DoArticleOfCreate 创建文章请求
type DoArticleOfCreate struct {
	Title       string `json:"title" form:"title" validate:"required,max=120" label:"标题"`                   // 文章标题
	Thumb       string `json:"thumb" form:"thumb" validate:"required,max=255,url" label:"缩略图"`              // 缩略图地址
	Content     string `json:"content" form:"content" validate:"required" label:"内容"`                       // 正文内容；富文本
	PublishedAt string `json:"published_at" form:"published_at" validate:"required,max=32" label:"发布时间"`    // 发布时间
	IsTop       uint8  `json:"is_top" form:"is_top" validate:"required,oneof=1 2" label:"置顶"`               // 是否置顶；枚举：1=是，2=否
	IsRecommend uint8  `json:"is_recommend" form:"is_recommend" validate:"required,oneof=1 2" label:"首页推荐"` // 是否首页推荐；枚举：1=是，2=否
}

// DoArticleOfUpdate 更新文章请求
type DoArticleOfUpdate struct {
	DoArticleOfCreate
	request.IDOfUint
}

// DoArticleOfDelete 删除文章请求
type DoArticleOfDelete struct {
	request.IDOfUint
}

// ToArticleOfInformation 文章详情查询请求
type ToArticleOfInformation struct {
	request.IDOfUint
}
