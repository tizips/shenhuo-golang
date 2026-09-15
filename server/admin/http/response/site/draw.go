package site

// ToDrawOfPaginate 抽签结果分页项响应
type ToDrawOfPaginate struct {
	ID         uint   `json:"id"`          // 抽签结果ID
	CategoryID uint   `json:"category_id"` // 抽签类别ID
	Category   string `json:"category"`    // 抽签类别名称
	PersonID   string `json:"person_id"`   // 参赛人员ID
	Name       string `json:"name"`        // 人员姓名
	Number     string `json:"number"`      // 参赛号
	Unit       string `json:"unit"`        // 所属单位
	GroupName  string `json:"group_name"`  // 所属小组名称
	CreatedAt  string `json:"created_at"`  // 中签时间
}

// DoDrawOfCreate 登记抽签结果响应
type DoDrawOfCreate struct {
	ToDrawOfPaginate
}
