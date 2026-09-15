package basic

// DoLoginOfAccount 管理端账号密码登录请求
type DoLoginOfAccount struct {
	Username string `json:"username" form:"username" validate:"required,username" label:"用户名"` // 登录用户名
	Password string `json:"password" form:"password" validate:"required,password" label:"密码"`  // 登录密码
}
