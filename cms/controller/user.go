package controller

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"order-go/service"
)

type UserController struct {
	Ctx     iris.Context
	Service service.UserService
}

func (uc UserController) GetLogin() mvc.Result {
	return mvc.View{
		Name: "user/login.html",
	}
}

func (uc UserController) PostLogin() mvc.Result {
	loginName := uc.Ctx.PostValueDefault("login_name", "")
	loginPwd := uc.Ctx.PostValueDefault("login_pwd", "")

	if len(loginName) < 1 {
		return mvc.Response{
			Object: map[string]interface{}{
				"code": -1,
				"msg":  "请输入正确的用户名！",
			},
		}
	}
	if len(loginPwd) < 1 {
		return mvc.Response{
			Object: map[string]interface{}{
				"code": -1,
				"msg":  "请输入正确的密码！",
			},
		}
	}
	_, err := uc.Service.GetByLoginName(loginName)
	if err != nil {
		return mvc.Response{
			Object: map[string]interface{}{
				"code": -1,
				"msg":  err,
			},
		}
	}

	return nil
}
