package site

import "github.com/herhe-com/framework/contracts/http/request"

// ToScoreOfPaginate 成绩分页查询请求
type ToScoreOfPaginate struct {
	Keyword string `json:"keyword" form:"keyword" query:"keyword" validate:"omitempty,max=64" label:"关键词"` // 搜索关键词；按姓名、参赛号或单位匹配
	request.Paginate
}

// DoScoreOfCreate 创建成绩请求；五项小项固定，总分与名次由系统自动计算
type DoScoreOfCreate struct {
	PersonID   string `json:"person_id" form:"person_id" validate:"required,snowflake" label:"人员"`      // 参赛人员ID
	Theory     string `json:"theory" form:"theory" validate:"omitempty,max=32" label:"理论考试"`            // 理论考试成绩
	Management string `json:"management" form:"management" validate:"omitempty,max=32" label:"管理能力考试"`  // 管理能力考试成绩
	Escape     string `json:"escape" form:"escape" validate:"omitempty,max=32" label:"井下紧急避险考核"`        // 井下紧急避险考核成绩
	Respirator string `json:"respirator" form:"respirator" validate:"omitempty,max=32" label:"自救器操作考核"` // 自救器操作考核成绩
	CPR        string `json:"cpr" form:"cpr" validate:"omitempty,max=32" label:"心肺复苏考核"`                // 心肺复苏考核成绩
}

// DoScoreOfUpdate 更新成绩请求；五项小项固定，总分与名次由系统自动计算
type DoScoreOfUpdate struct {
	Theory     string `json:"theory" form:"theory" validate:"omitempty,max=32" label:"理论考试"`            // 理论考试成绩
	Management string `json:"management" form:"management" validate:"omitempty,max=32" label:"管理能力考试"`  // 管理能力考试成绩
	Escape     string `json:"escape" form:"escape" validate:"omitempty,max=32" label:"井下紧急避险考核"`        // 井下紧急避险考核成绩
	Respirator string `json:"respirator" form:"respirator" validate:"omitempty,max=32" label:"自救器操作考核"` // 自救器操作考核成绩
	CPR        string `json:"cpr" form:"cpr" validate:"omitempty,max=32" label:"心肺复苏考核"`                // 心肺复苏考核成绩
	request.IDOfUint
}

// DoScoreOfDelete 删除成绩请求
type DoScoreOfDelete struct {
	request.IDOfUint
}

// ToScoreOfInformation 成绩详情查询请求
type ToScoreOfInformation struct {
	request.IDOfUint
}

// DoScoreOfNotify 推送成绩通知请求；ID 为抽签分组（抽签类别）ID，向该分组下所有已出成绩的人员推送
type DoScoreOfNotify struct {
	CategoryID uint `json:"category_id" form:"category_id" validate:"required,gt=0" label:"抽签分组"` // 抽签分组ID
}
