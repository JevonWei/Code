package main
import "github.com/astaxie/beego"


type ErrorController struct{
	beego.Controller
}

func (this* ErrorController) Error404(){
	this.Data["Content"] = "没找到页面"
	this.TplName = "404.html"
}

func (this* ErrorController) ErrorDb(){
	this.Data["Content"] = "数据库错误"
	this.TplName = "404.html"
}

type NormalController struct {
	beego.Controller
}

func (this* NormalController) Generate500(){
	this.Abort("500")
}

func (this* NormalController) GenerateDbError(){
	this.Abort("Db")

}

func (this* NormalController) GenerateUnknownError(){
	this.Abort("unknown")

}

func main() {
	beego.Router("/500", &NormalController{}, "get,post:Generate500")	
	beego.Router("/db", &NormalController{}, "get,post:GenerateDbError")	
	beego.Router("/unknown", &NormalController{}, "get,post:GenerateUnknownError")	
	beego.ErrorController(&ErrorController{})
	beego.Run()
}

