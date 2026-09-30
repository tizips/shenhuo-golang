package site

import "github.com/herhe-com/framework/contracts/http/request"

// ToCheckinOfPaginate 签到分页查询请求
type ToCheckinOfPaginate struct {
	Name string `json:"name" form:"name" query:"name" validate:"omitempty,max=32" label:"姓名"` // 参赛人员姓名
	request.Paginate
}

// ToCheckinOfInformation 签到详情查询请求
type ToCheckinOfInformation struct {
	request.IDOfUint
}
