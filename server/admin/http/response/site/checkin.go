package site

// ToCheckinOfPaginate 签到分页项响应
type ToCheckinOfPaginate struct {
	ID        uint   `json:"id"`         // 签到ID
	Name      string `json:"name"`       // 参赛人员姓名
	CreatedAt string `json:"created_at"` // 签到时间
}

// ToCheckinOfInformation 签到详情响应
type ToCheckinOfInformation struct {
	ToCheckinOfPaginate
}
