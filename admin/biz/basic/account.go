package basic

import (
	"context"
	"errors"
	"fmt"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/herhe-com/framework/auth"
	"github.com/herhe-com/framework/facades"
	"github.com/herhe-com/framework/http"
	req "github.com/herhe-com/template/admin/http/request/basic"
	res "github.com/herhe-com/template/admin/http/response/basic"
	"github.com/herhe-com/template/model"
	"github.com/samber/lo"
	"gorm.io/gorm"
)

func ToAccountOfInformation(c context.Context, ctx *app.RequestContext) {

	var user model.SysUser

	fu := facades.Gorm.First(&user, "`id`=?", auth.ID(ctx))

	if errors.Is(fu.Error, gorm.ErrRecordNotFound) {
		http.Unauthorized(ctx)
		return
	}

	responses := res.ToAccountOfInformation{
		Nickname: user.Nickname,
		Avatar:   user.Avatar,
	}

	if user.Username != nil {
		responses.Username = *user.Username
	}

	if user.Mobile != nil {
		responses.Mobile = *user.Mobile
	}

	if user.Email != nil {
		responses.Username = *user.Email
	}

	responses.Platform.Code = auth.Platform(ctx)
	responses.Platform.Name = facades.Cfg.GetString("app.title")

	http.Success(ctx, responses)
}

func ToAccountOfModules(c context.Context, ctx *app.RequestContext) {

	responses := make([]res.ToAccountOfModules, 0)

	modules := auth.Modules(auth.Platform(ctx))

	if ok, _ := facades.Casbin.HasRoleForUser(auth.NameOfUser(auth.ID(ctx)), auth.NameOfDeveloper()); ok {

		for _, item := range modules {
			responses = append(responses, res.ToAccountOfModules{
				Code: item.Code,
				Name: item.Name,
			})
		}

	} else {

		var moduleCodes []string

		facades.Gorm.
			Model(&model.SysRoleBindPermission{}).
			Distinct("module").
			Where("exists (?)", facades.Gorm.
				Model(&model.SysUserBindRole{}).
				Select("1").
				Where(fmt.Sprintf("`%s`.`role_id`=`%s`.`role_id` and `%s`.`user_id`=?", model.TableSysRoleBindPermission, model.TableSysUserBindRole, model.TableSysUserBindRole), auth.ID(ctx)),
			).
			Where("exists (?)", facades.Gorm.
				Model(&model.SysRole{}).
				Select("1").
				Where(fmt.Sprintf("`%s`.`role_id`=`%s`.`id`", model.TableSysRoleBindPermission, model.TableSysRole)),
			).
			Pluck("module", &moduleCodes)

		for _, item := range modules {

			if lo.Contains(moduleCodes, item.Code) {

				responses = append(responses, res.ToAccountOfModules{
					Code: item.Code,
					Name: item.Name,
				})
			}
		}
	}

	http.Success(ctx, responses)
}

func ToAccountOfPermissions(c context.Context, ctx *app.RequestContext) {

	var request req.ToAccountOfPermissions

	if err := ctx.BindAndValidate(&request); err != nil {
		http.BadRequest(ctx, err)
		return
	}

	responses := make([]string, 0)

	if ok, _ := facades.Casbin.HasRoleForUser(auth.NameOfUser(auth.ID(ctx)), auth.NameOfDeveloper()); ok {

		modules := auth.Modules(auth.Platform(ctx))

		for _, item := range modules {
			if item.Code == request.Module {
				responses = item.Permissions
				break
			}
		}

	} else {

		facades.Gorm.
			Model(&model.SysRoleBindPermission{}).
			Where("module = ?", request.Module).
			Where("exists (?)", facades.Gorm.
				Model(&model.SysUserBindRole{}).
				Select("1").
				Where(fmt.Sprintf("`%s`.`role_id`=`%s`.`role_id` and `%s`.`user_id`=?", model.TableSysRoleBindPermission, model.TableSysUserBindRole, model.TableSysUserBindRole), auth.ID(ctx)),
			).
			Where("exists (?)", facades.Gorm.
				Model(&model.SysRole{}).
				Select("1").
				Where(fmt.Sprintf("`%s`.`role_id`=`%s`.`id`", model.TableSysRoleBindPermission, model.TableSysRole)),
			).
			Pluck("permission", &responses)
	}

	http.Success(ctx, responses)
}

func DoAccount(c context.Context, ctx *app.RequestContext) {

	var request req.DoAccount

	if err := ctx.BindAndValidate(&request); err != nil {
		http.BadRequest(ctx, err)
		return
	}

	updates := make(map[string]any)

	if request.Mobile != "" {
		updates["mobile"] = request.Mobile
	}

	if request.Email != "" {
		updates["email"] = request.Email
	}

	if request.Password != "" {
		updates["password"] = auth.Password(request.Password)
	}

	if len(updates) > 0 {

		if result := facades.Gorm.Model(&model.SysUser{}).Where("`id` = ?", auth.ID(ctx)).Updates(updates); result.Error != nil {
			http.Fail(ctx, "修改失败：%v", result.Error)
			return
		}
	}

	http.Success[any](ctx)
}
