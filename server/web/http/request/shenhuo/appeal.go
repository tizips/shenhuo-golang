package shenhuo

// DoAppealOfCreate 仲裁申诉请求
type DoAppealOfCreate struct {
	IDCard   string `json:"id_card" form:"id_card" validate:"required,idCard" label:"身份证号"` // 身份证号；用于验证身份
	Password string `json:"password" form:"password" validate:"required" label:"密码"`        // 登录密码；用于验证身份
	Reason   string `json:"reason" form:"reason" validate:"required,max=500" label:"申诉原因"`  // 申诉原因
}
