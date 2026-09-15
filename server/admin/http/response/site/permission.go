package site

// ToPermissions 权限树节点响应
type ToPermissions struct {
	Code     string          `json:"code"`               // 权限标识
	Name     string          `json:"name"`               // 权限名称
	Children []ToPermissions `json:"children,omitempty"` // 子权限列表；无子级时省略
}
