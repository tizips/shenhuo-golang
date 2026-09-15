package shenhuo

// DoScoreOfQuery 成绩查询请求
type DoScoreOfQuery struct {
	Mobile   string `json:"mobile" form:"mobile" validate:"required,mobile" label:"手机号"` // 登录手机号
	Password string `json:"password" form:"password" validate:"required" label:"密码"`     // 登录密码
}
