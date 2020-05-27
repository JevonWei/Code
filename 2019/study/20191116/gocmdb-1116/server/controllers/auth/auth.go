package auth

import (
	"github.com/JevonWei/gocmdb/server/controllers/base"
	"github.com/JevonWei/gocmdb/server/models"
)

type LoginRequiredController struct{
	base.BaseController
	User *models.User
}

func (c *LoginRequiredController) Perpare() {
	c.BaseController.Prepare()

	// session认证，判断用户是否登录
	if user := DefaultManager.IsLogin(c); user == nil {
		DefaultManager.GoToLoginPage(c)
		c.StopRun()
	} else {
		c.User = user
		c.Data["user"] = user
	}
}

type AuthController struct {
	base.BaseController
}

func (c *AuthController) Login() {
	DefaultManager.Login(c)
	
}

func (c *AuthController) Logout() {
	DefaultManager.Logout(c)
}