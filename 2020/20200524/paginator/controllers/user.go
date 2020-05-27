package controllers

import (
	"paginator/models"
	"paginator/utils"

	"github.com/astaxie/beego"
)

type UserController struct {
	beego.Controller
}

func (this *UserController) List() {
	pers := 20
	count, users := models.GetAllUsers()
	p := utils.NewPaginator(this.Ctx.Request, pers, count)
	if p.Offset()+pers > int(count) {
		this.Data["Users"] = users[p.Offset()]
	} else {
		this.Data["Users"] = users[p.Offset() : p.Offset()+pers]
	}
	this.Data["paginator"] = p
	this.TplName = "index.html"
}

func init() {
	beego.Router("/user", &UserController{}, "get:List")
}
