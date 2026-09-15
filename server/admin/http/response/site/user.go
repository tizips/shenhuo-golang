package site

// ToUserByPaginate 用户分页项响应
type ToUserByPaginate struct {
	ID        string                    `json:"id"`                 // 用户ID；雪花算法生成
	Nickname  string                    `json:"nickname"`           // 用户昵称
	Username  string                    `json:"username,omitempty"` // 登录用户名；未设置时省略
	Mobile    string                    `json:"mobile,omitempty"`   // 手机号；未设置时省略
	Email     string                    `json:"email,omitempty"`    // 邮箱；未设置时省略
	Roles     []ToUserByPaginateOfRoles `json:"roles"`              // 绑定的角色列表
	IsEnable  uint8                     `json:"is_enable"`          // 是否启用；枚举：1=启用，2=停用
	CreatedAt string                    `json:"created_at"`         // 创建时间
}

// ToUserByPaginateOfRoles 用户分页项中的角色信息
type ToUserByPaginateOfRoles struct {
	ID   uint   `json:"id"`   // 角色ID
	Name string `json:"name"` // 角色名称
}
