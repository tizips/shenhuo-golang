package basic

// ToAccountOfPermissions 获取当前账号在指定模块下的权限列表请求
type ToAccountOfPermissions struct {
	Module string `json:"module" query:"module" form:"module" validate:"required" label:"模块"` // 模块标识；用于定位权限所属模块
}

// DoAccount 更新当前账号资料请求
type DoAccount struct {
	Mobile   string `json:"mobile" form:"mobile" validate:"omitempty,mobile" label:"手机号"`      // 手机号；留空表示不修改
	Email    string `json:"email" form:"email" validate:"omitempty,email" label:"邮箱"`          // 邮箱；留空表示不修改
	Password string `json:"password" form:"password" validate:"omitempty,password" label:"密码"` // 新密码；留空表示不修改
}
