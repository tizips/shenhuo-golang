package basic

// DoLogin 登录响应
// 双令牌（JWT）机制：
//   - access_token 用于日常请求；过期后（接口返回 401）改用 refresh_token 换取新令牌对；
//   - 刷新无专门接口：在任意需要登录的接口上携带 `Refresh-Token: <refresh_token>` 请求头重发请求，
//     服务端轮换令牌后通过响应头 `Token-Pair` 返回新令牌对（JSON，字段与登录响应中的令牌字段一致），
//     前端必须解析并覆盖本地旧令牌；
//   - 同一刷新令牌只能轮换一次；并发刷新在 grace_lifetime 秒容错窗口内返回同一份新令牌；
//   - 超过容错窗口重复使用旧刷新令牌、刷新令牌过期（refresh_lifetime 秒，默认 30 天）或已注销时返回 401，
//     此时必须重新登录；
//   - 轮换后的新令牌对 session_id 不变，刷新令牌有效期重新起算（滑动过期）。
type DoLogin struct {
	SessionID          string `json:"session_id"`                     // 登录会话ID；同一轮登录及后续令牌轮换保持不变，不作为请求凭证
	AccessToken        string `json:"access_token"`                   // 访问令牌；放入 Authorization 请求头（直接放令牌本身，不带 Bearer 前缀）访问需要登录的接口；有效时长见 access_lifetime
	RefreshToken       string `json:"refresh_token"`                  // 刷新令牌；平时不随请求发送，仅在访问令牌过期后用于 Refresh-Token 请求头换取新令牌对；过期或重复使用后返回 401，需重新登录
	IssuedAt           int64  `json:"issued_at"`                      // 签发时间；Unix 秒（UTC）；access_token/refresh_token 的过期时间均为 issued_at + 各自有效时长
	AccessLifetime     int64  `json:"access_lifetime"`                // 访问令牌有效时长；单位：秒；过期时间为 issued_at + access_lifetime；轮换后新令牌的有效时长与本次相同
	RefreshLifetime    int64  `json:"refresh_lifetime"`               // 刷新令牌有效时长；单位：秒；过期时间为 issued_at + refresh_lifetime（默认 30 天）；轮换后重新起算
	GraceLifetime      int64  `json:"grace_lifetime,omitempty"`       // 并发刷新容错时长；单位：秒；登录响应通常省略，刷新令牌对时返回；同一旧刷新令牌在此窗口内重复使用会返回同一份新令牌
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
