package routers

import (
	"github.com/astaxie/beego"

	"github.com/JevonWei/gocmdb/server/controllers"
	"github.com/JevonWei/gocmdb/server/controllers/auth"
)

func init() {
	beego.AutoRouter(&auth.AuthController{})
	beego.AutoRouter(&controllers.UserPageController{})
	beego.AutoRouter(&controllers.UserController{})
	beego.AutoRouter(&controllers.TokenController{})
}
