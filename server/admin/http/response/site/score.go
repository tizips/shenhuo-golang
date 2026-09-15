package site

import "github.com/tizips/shenhuo/model"

// ToScoreOfPaginate 成绩分页项响应
type ToScoreOfPaginate struct {
	ID        uint                `json:"id"`         // 成绩ID
	PersonID  string              `json:"person_id"`  // 参赛人员ID
	Name      string              `json:"name"`       // 人员姓名
	Number    string              `json:"number"`     // 参赛号
	Unit      string              `json:"unit"`       // 所属单位
	GroupName string              `json:"group_name"` // 所属小组名称
	Total     string              `json:"total"`      // 总分
	Items     []model.ShScoreItem `json:"items"`      // 各小项成绩列表
	Order     uint8               `json:"order"`      // 名次；按总分自动计算，数值越小名次越靠前
	CreatedAt string              `json:"created_at"` // 创建时间
}

// ToScoreOfInformation 成绩详情响应
type ToScoreOfInformation struct {
	ToScoreOfPaginate
}

// DoScoreOfImport 批量导入成绩结果响应
type DoScoreOfImport struct {
	Total   int `json:"total"`   // 导入总行数
	Created int `json:"created"` // 成功创建数
	Updated int `json:"updated"` // 成功更新数
	Skipped int `json:"skipped"` // 跳过数；通常为无效数据
}
