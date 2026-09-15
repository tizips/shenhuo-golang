package site

// ToManagerOfPaginate 管理员分页项响应
type ToManagerOfPaginate struct {
	ID        string `json:"id"`         // 管理员ID；雪花算法生成
	Name      string `json:"name"`       // 管理员姓名
	Mobile    string `json:"mobile"`     // 手机号
	IsEnable  uint8  `json:"is_enable"`  // 是否启用；枚举：1=启用，2=停用
	CreatedAt string `json:"created_at"` // 创建时间
}
