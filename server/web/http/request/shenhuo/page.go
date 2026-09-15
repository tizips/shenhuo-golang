package shenhuo

import "github.com/herhe-com/framework/contracts/http/request"

// ToPageOfInformation 单页详情查询请求
type ToPageOfInformation struct {
	request.IDOfUint
}

// ToPageOfInformationByKey 按标识查询单页详情请求
type ToPageOfInformationByKey struct {
	Key string `json:"key" form:"key" path:"key" query:"key" validate:"required,max=64" label:"页面标识"` // 内置页面标识
}
