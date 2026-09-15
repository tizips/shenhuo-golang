package basic

// DoLogin 登录响应
type DoLogin struct {
	SessionID          string `json:"session_id"`                     // 登录会话ID；同一轮登录及后续令牌轮换保持不变，不作为请求凭证
	AccessToken        string `json:"access_token"`                   // 访问令牌；仅用于 Authorization 请求头
	RefreshToken       string `json:"refresh_token"`                  // 刷新令牌；仅在访问令牌失效后用于 Refresh-Token 请求头
	IssuedAt           int64  `json:"issued_at"`                      // 签发时间；Unix 秒（UTC）
	AccessLifetime     int64  `json:"access_lifetime"`                // 访问令牌有效时长；单位：秒；过期时间为 issued_at + access_lifetime
	RefreshLifetime    int64  `json:"refresh_lifetime"`               // 刷新令牌有效时长；单位：秒；过期时间为 issued_at + refresh_lifetime
	GraceLifetime      int64  `json:"grace_lifetime,omitempty"`       // 并发刷新容错时长；单位：秒；登录响应通常省略，刷新令牌对时返回
	MustChangePassword uint8  `json:"must_change_password,omitempty"` // 是否需强制修改密码；枚举：1=是，2=否；1 时前端应引导修改密码
	Name               string `json:"name"`                           // 姓名
	Mobile             string `json:"mobile"`                         // 手机号
}

// ToAccountOfInformation 当前账号信息响应
type ToAccountOfInformation struct {
	ID                 string `json:"id"`                             // 账号ID；雪花算法生成
	Name               string `json:"name"`                           // 姓名
	Mobile             string `json:"mobile"`                         // 手机号
	Number             string `json:"number,omitempty"`               // 参赛号；仅参赛人员有值
	Unit               string `json:"unit,omitempty"`                 // 所属单位；仅参赛人员有值
	GroupName          string `json:"group_name,omitempty"`           // 所属小组名称；仅参赛人员有值
	MustChangePassword uint8  `json:"must_change_password,omitempty"` // 是否需强制修改密码；枚举：1=是，2=否
	Kind               string `json:"kind"`                           // 账号类型；枚举：manager=管理员，person=参赛人员
}
