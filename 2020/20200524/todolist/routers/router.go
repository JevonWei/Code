package routers

import (
	"todolist/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.AutoRouter(&controllers.AuthController{})
	beego.AutoRouter(&controllers.TaskController{})
	beego.Router("/", &controllers.TaskController{}, "get:Index")
}
