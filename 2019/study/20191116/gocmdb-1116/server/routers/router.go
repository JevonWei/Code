package routers

import (
	"github.com/JevonWei/gocmdb/server/controllers"
	"github.com/JevonWei/gocmdb/server/controllers/auth"
	"github.com/astaxie/beego"
)

func init() {
	beego.AutoRouter(&auth.AuthController{})
	beego.AutoRouter(&controllers.TestController{})
	beego.AutoRouter(&controllers.TestPageController{})
}
