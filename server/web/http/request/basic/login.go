package basic

// DoLoginOfMobile 大屏端手机号密码登录请求
type DoLoginOfMobile struct {
	Mobile   string `json:"mobile" form:"mobile" validate:"required,mobile" label:"手机号"`      // 登录手机号
	Password string `json:"password" form:"password" validate:"required,password" label:"密码"` // 登录密码
}

// DoLoginOfWeChat 参赛人员微信网页授权登录请求
type DoLoginOfWeChat struct {
	Mobile   string `json:"mobile" form:"mobile" validate:"required,mobile" label:"手机号"`      // 登录手机号
	Password string `json:"password" form:"password" validate:"required,password" label:"密码"` // 登录密码
	Code     string `json:"code" form:"code" validate:"required" label:"网页授权Code"`            // 微信服务号网页授权 Code
}

// DoAccountOfPassword 修改当前账号密码请求
type DoAccountOfPassword struct {
	Password string `json:"password" form:"password" validate:"required,password" label:"密码"` // 新密码
}
