package common

// ToSetting 系统设置项响应
type ToSetting struct {
	ID         uint   `json:"id"`          // 设置项ID
	Type       string `json:"type"`        // 表单类型；枚举：input=文本框，enable=开关，url=链接，email=邮箱，picture=图片，textarea=多行文本
	Label      string `json:"label"`       // 显示名称
	Key        string `json:"key"`         // 设置项键名
	Val        string `json:"val"`         // 设置项值
	IsRequired uint8  `json:"is_required"` // 是否必填；枚举：1=是，2=否
	CreatedAt  string `json:"created_at"`  // 创建时间
}
