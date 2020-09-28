package controller

import (
	"github.com/gorilla/securecookie"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"order-go/common/utils"
	"order-go/service"
)

var (
	hashKey  = securecookie.GenerateRandomKey(64)
	blockKey = securecookie.GenerateRandomKey(32)

	sc = securecookie.New(hashKey, blockKey)
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
	userInfo, err := uc.Service.GetByLoginName(loginName)
	if err != nil {
		return mvc.Response{
			Object: map[string]interface{}{
				"code": -1,
				"msg":  err,
			},
		}
	}
	if userInfo.LoginPwd != utils.Md5Encrypt(loginPwd, userInfo.LoginSalt) {
		return mvc.Response{
			Object: map[string]interface{}{
				"code": -1,
				"msg":  "请输入正确的用户名和密码！",
			},
		}
	}
	if userInfo.Status != 1 {
		return mvc.Response{
			Object: map[string]interface{}{
				"code": -1,
				"msg":  "账号已被禁用，请联系管理员！",
			},
		}
	}
	// 设置cookie
	uc.Ctx.SetCookieKV(loginName, utils.GenerateAuthCode(userInfo), iris.CookieEncode(sc.Encode))
	return mvc.Response{
		Object: map[string]interface{}{
			"code": 200,
			"msg":  "登录成功",
			"data": nil,
		},
	}
}
