package site

// ToAppealOfPaginate 仲裁申诉分页项响应
type ToAppealOfPaginate struct {
	ID        uint   `json:"id"`         // 申诉ID
	PersonID  string `json:"person_id"`  // 参赛人员ID
	Name      string `json:"name"`       // 人员姓名
	Number    string `json:"number"`     // 参赛号
	Unit      string `json:"unit"`       // 所属单位
	GroupName string `json:"group_name"` // 所属小组名称
	IDCard    string `json:"id_card"`    // 身份证号
	Reason    string `json:"reason"`     // 申诉原因
	CreatedAt string `json:"created_at"` // 提交时间
}

// ToAppealOfInformation 仲裁申诉详情响应
type ToAppealOfInformation struct {
	ToAppealOfPaginate
}
