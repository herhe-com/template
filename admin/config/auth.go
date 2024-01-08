package config

import (
	"github.com/herhe-com/framework/contracts/auth"
	"github.com/herhe-com/framework/facades"
)

func init() {

	cfg := facades.Cfg
	cfg.Add("auth", map[string]any{
		"casbin": map[string]any{
			"table": cfg.Env("auth.casbin.table", "sys_casbin"),
		},
		"platforms": []uint16{auth.CodeOfPlatform, auth.CodeOfClique, auth.CodeOfStore},
		"permissions": []auth.Permission{
			site(),
		},
	})
}

func site() auth.Permission {
	return auth.Permission{
		Code: "site",
		Name: "站点",
		Children: []auth.Permission{
			{
				Code: "role",
				Name: "角色",
				Children: []auth.Permission{
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
				Children: []auth.Permission{
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
		},
	}
}
