package config

import (
	"github.com/herhe-com/framework/auth"
	contractauth "github.com/herhe-com/framework/contracts/auth"
	"github.com/herhe-com/framework/facades"
)

func init() {

	cfg := facades.Config()
	cfg.Add("auth", map[string]any{
		"casbin": map[string]any{
			"table": cfg.Env("auth.casbin.table", "sys_casbin"),
		},
		"platforms": []uint16{auth.CodeOfPlatform, auth.CodeOfClique, auth.CodeOfStore},
		"permissions": func() []contractauth.Permission {
			return []contractauth.Permission{
				site(),
			}
		},
	})
}

func site() contractauth.Permission {
	return contractauth.Permission{
		Code: "site",
		Name: "站点",
		Children: []contractauth.Permission{
			{
				Code: "role",
				Name: "角色",
				Children: []contractauth.Permission{
					{
						Code:   "create",
						Name:   "创建",
						Common: true,
					},
					{
						Code:   "update",
						Name:   "修改",
						Common: true,
					},
					{
						Code:   "delete",
						Name:   "删除",
						Common: true,
					},
					{
						Code:   "paginate",
						Name:   "列表",
						Common: true,
					},
				},
			},
			{
				Code: "user",
				Name: "账号",
				Children: []contractauth.Permission{
					{
						Code:   "create",
						Name:   "创建",
						Common: true,
					},
					{
						Code:   "update",
						Name:   "修改",
						Common: true,
					},
					{
						Code:   "delete",
						Name:   "删除",
						Common: true,
					},
					{
						Code:   "enable",
						Name:   "启禁",
						Common: true,
					},
					{
						Code:   "paginate",
						Name:   "列表",
						Common: true,
					},
				},
			},
			crud("article", "热点资讯"),
			crud("page", "页面"),
			crud("page_builtin", "内置页面"),
			crud("nav", "导航"),
			crud("banner", "轮播"),
			crud("schedule", "日期安排"),
			crud("schedule_category", "日程分类"),
			crud("person", "人员", action("import", "导入")),
			crud("draw_category", "抽签类别"),
			{
				Code: "draw",
				Name: "人员抽签",
				Children: []contractauth.Permission{
					action("create", "抽签"),
					action("delete", "删除"),
					action("paginate", "列表"),
				},
			},
			crud("scene", "场景"),
			crud("media", "精彩媒体"),
			crud("manager", "管理人员", action("enable", "启禁")),
			crud("score", "成绩", action("import", "导入")),
		},
	}
}

func crud(code, name string, extras ...contractauth.Permission) contractauth.Permission {
	children := []contractauth.Permission{
		action("create", "创建"),
		action("update", "修改"),
		action("delete", "删除"),
		action("paginate", "列表"),
	}
	return contractauth.Permission{
		Code:     code,
		Name:     name,
		Children: append(children, extras...),
	}
}

func action(code, name string) contractauth.Permission {
	return contractauth.Permission{
		Code:   code,
		Name:   name,
		Common: true,
	}
}
