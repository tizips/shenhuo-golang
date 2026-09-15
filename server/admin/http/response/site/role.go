package site

// ToRoleByPaginate 角色分页项响应
type ToRoleByPaginate struct {
	ID        uint   `json:"id"`         // 角色ID
	Name      string `json:"name"`       // 角色名称
	Summary   string `json:"summary"`    // 角色简介
	CreatedAt string `json:"created_at"` // 创建时间
}

// ToRoleByInformation 角色详情响应
type ToRoleByInformation struct {
	ID          uint     `json:"id"`          // 角色ID
	Name        string   `json:"name"`        // 角色名称
	Permissions []string `json:"permissions"` // 权限标识列表
	Summary     string   `json:"summary"`     // 角色简介
	CreatedAt   string   `json:"created_at"`  // 创建时间
}
