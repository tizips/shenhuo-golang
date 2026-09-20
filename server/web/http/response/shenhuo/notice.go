package shenhuo

// DoNoticeOfTest 测试推送结果
type DoNoticeOfTest struct {
	Type  string `json:"type"`            // 通知类型；枚举：score=成绩查询通知，draw=活动中签通知
	Ok    bool   `json:"ok"`              // 是否发送成功
	Error string `json:"error,omitempty"` // 失败原因
}
