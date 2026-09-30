package shenhuo

// DoScoreOfQuery 成绩查询请求
type DoScoreOfQuery struct {
	IDCard   string `json:"id_card" form:"id_card" validate:"required,idCard" label:"身份证号"` // 身份证号
	Password string `json:"password" form:"password" validate:"required" label:"密码"`        // 登录密码
}
