package shenhuo

// ToSystemOfConfig 系统配置响应
type ToSystemOfConfig struct {
	Video bool `json:"video"` // 是否开启视频模块；枚举：true=开启，false=关闭
}
