package controllers

import (
	"net/http"

	"github.com/astaxie/beego"
)

type HomeController struct {
	BaseController
}

func (this *HomeController) Get() {
	this.Redirect(beego.URLFor(beego.AppConfig.String("home")), http.StatusFound)
}
