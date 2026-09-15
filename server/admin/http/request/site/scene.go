package site

import "github.com/herhe-com/framework/contracts/http/request"

// ToSceneOfPaginate 场景分页查询请求
type ToSceneOfPaginate struct {
	request.Paginate
}

// DoSceneOfCreate 创建场景请求
type DoSceneOfCreate struct {
	Name  string `json:"name" form:"name" validate:"required,max=64" label:"名称"`          // 场景名称
	Order uint8  `json:"order" form:"order" validate:"omitempty,min=1,max=99" label:"序号"` // 排序序号；数值越小越靠前
}

// DoSceneOfUpdate 更新场景请求
type DoSceneOfUpdate struct {
	DoSceneOfCreate
	request.IDOfUint
}

// DoSceneOfDelete 删除场景请求
type DoSceneOfDelete struct {
	request.IDOfUint
}
