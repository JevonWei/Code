package main

import (
	"github.com/astaxie/beego"
)

type ConfController struct {
	beego.Controller
}

func (this *ConfController) Get() {
	this.Ctx.Output.Body([]byte(beego.AppConfig.String("test") + "\n" + "Hello World" + "\n"))
}

func main() {
	beego.Router("/", &ConfController{})
	beego.Run()
}
