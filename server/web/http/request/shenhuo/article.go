package shenhuo

import "github.com/herhe-com/framework/contracts/http/request"

// ToArticleOfPaginate 文章分页查询请求
type ToArticleOfPaginate struct {
	request.Paginate
}

// ToArticleOfInformation 文章详情查询请求
type ToArticleOfInformation struct {
	request.IDOfUint
}
