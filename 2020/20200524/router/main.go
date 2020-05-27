package main

import (
	"strings"

	"github.com/astaxie/beego"
)

type FixControlltr struct {
	beego.Controller
}

func (this *FixControlltr) Get() {
	this.Ctx.Output.Body([]byte("Hello World for Controller Get\n"))
}

func (this *FixControlltr) Put() {
	this.Ctx.Output.Body([]byte("Hello World for Controller put\n"))
}

func (this *FixControlltr) Any() {
	this.Ctx.Output.Body([]byte("Hello World for Controller Any\n"))
}

type RController struct {
	beego.Controller
}

func (this *RController) Get() {
	id := this.Ctx.Input.Param(":id")
	if id != "" {
		this.Ctx.Output.Body([]byte("id in regex: " + this.Ctx.Input.Param(":id") + "\n"))
	} else {
		this.Ctx.Output.Body([]byte("id in regex is empty" + "\n"))
	}
}

type RController1 struct {
	beego.Controller
}

func (this *RController1) Get() {
	output := "path in regex: " + this.Ctx.Input.Param(":path") + "\n"
	output += "ext in regex: " + this.Ctx.Input.Param(":ext") + "\n"
	output += "splat in regex: " + this.Ctx.Input.Param(":splat") + "\n"
	this.Ctx.Output.Body([]byte(output))
}

type SelfDefineController struct {
	beego.Controller
}

func (this *SelfDefineController) ForGet() {
	if strings.ToLower(this.Ctx.Input.Method()) == "get" {
		this.Ctx.Output.Body([]byte("Yes, it's get   "))
	}
}

func main() {
	beego.Router("/fix", &FixControlltr{})

	beego.Router("/api/?:id", &RController{})
	beego.Router("/api1/:id", &RController{})
	beego.Router("/api2/:id([0-9]+)", &RController{})
	beego.Router("/api3/:id:string", &RController{})
	beego.Router("/path_ext/*.*", &RController1{})
	beego.Router("/splat/*", &RController1{})

	// beego.Router("/selfdefine", &SelfDefineController{}, "get:ForGet")
	beego.AutoRouter(&SelfDefineController{})
	beego.Run()
}
