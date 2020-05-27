package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
    "fmt"
)

type FilterController struct{
    beego.Controller
}
var user int = 1000


func (this *FilterController) Login(){
    this.SetSession("user", user)
    user++
}

func  (this *FilterController) HandleRequest(){
    this.Ctx.Output.Body([]byte(fmt.Sprintf("Handle request for: %v\n", this.GetSession("user"))))
}

var ValidateLogin = func(ctx *context.Context) {
    
    user := ctx.Input.Session("user")
    if user == nil && ctx.Request.RequestURI != "/login" {
        ctx.Redirect(302, "/login")
    }
}
func main() {
    beego.Router("/login", &FilterController{}, "get,post:Login") 
    beego.Router("/req", &FilterController{}, "get,post:HandleRequest") 
    beego.InsertFilter("/*",beego.BeforeRouter, ValidateLogin)
	beego.Run()
}

