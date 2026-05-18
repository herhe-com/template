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
		"permissions": []contractauth.Permission{
			site(),
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
		},
	}
}
