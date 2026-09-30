package site

import "github.com/herhe-com/framework/contracts/http/request"

// ToAppealOfPaginate 仲裁申诉分页查询请求
type ToAppealOfPaginate struct {
	Keyword string `json:"keyword" form:"keyword" query:"keyword" validate:"omitempty,max=64" label:"关键词"` // 搜索关键词；按姓名、参赛号或身份证号匹配
	request.Paginate
}

// ToAppealOfInformation 仲裁申诉详情查询请求
type ToAppealOfInformation struct {
	request.IDOfUint
}
