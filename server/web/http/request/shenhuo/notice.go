package shenhuo

// DoNoticeOfTest 测试推送通知请求
type DoNoticeOfTest struct {
	PersonID string `json:"person_id" form:"person_id" validate:"required,snowflake" label:"人员"`    // 参赛人员ID
	Type     string `json:"type" form:"type" validate:"required,oneof=score draw all" label:"通知类型"` // 通知类型；枚举：score=成绩查询通知，draw=活动中签通知，all=全部
}
