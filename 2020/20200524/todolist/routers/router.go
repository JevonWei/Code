package routers

import (
	"todolist/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.HomeController{})
	beego.AutoRouter(&controllers.AuthController{})
	beego.AutoRouter(&controllers.TaskController{})
	beego.AutoRouter(&controllers.UserController{})
	// beego.Router("/", &controllers.TaskController{}, "get:Index")
}
