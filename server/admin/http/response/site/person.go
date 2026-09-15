package site

// ToPersonOfPaginate 参赛人员分页项响应
type ToPersonOfPaginate struct {
	ID                 string `json:"id"`                   // 人员ID；雪花算法生成
	Name               string `json:"name"`                 // 姓名
	Unit               string `json:"unit"`                 // 所属单位
	Mobile             string `json:"mobile"`               // 手机号
	Number             string `json:"number"`               // 参赛号
	GroupName          string `json:"group_name"`           // 所属小组名称
	MustChangePassword uint8  `json:"must_change_password"` // 是否需强制修改密码；枚举：1=是，2=否
	CreatedAt          string `json:"created_at"`           // 创建时间
}

// ToPersonOfInformation 参赛人员详情响应
type ToPersonOfInformation struct {
	ToPersonOfPaginate
}

// DoPersonOfImport 批量导入参赛人员结果响应
type DoPersonOfImport struct {
	Total   int `json:"total"`   // 导入总行数
	Created int `json:"created"` // 成功创建数
	Skipped int `json:"skipped"` // 跳过数；通常为重复数据
}
