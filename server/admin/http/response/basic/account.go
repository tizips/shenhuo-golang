package basic

// ToAccountOfInformation 当前账号信息响应
type ToAccountOfInformation struct {
	Nickname string `json:"nickname"`           // 昵称
	Username string `json:"username,omitempty"` // 登录用户名；未设置时省略
	Mobile   string `json:"mobile,omitempty"`   // 手机号；未设置时省略
	Email    string `json:"email,omitempty"`    // 邮箱；未设置时省略
	Avatar   string `json:"avatar"`             // 头像地址
	Platform struct {
		Code uint16 `json:"code"` // 平台编码
		Name string `json:"name"` // 平台名称
	} `json:"platform"`
}

// ToAccountOfModules 当前账号可见模块响应
type ToAccountOfModules struct {
	Code string `json:"code"` // 模块标识
	Name string `json:"name"` // 模块名称
}
