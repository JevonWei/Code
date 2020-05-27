package controllers

import (
	"net/http"
	"todolist/forms"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/validation"
)

type AuthController struct {
	BaseController
}

func (this *AuthController) Login() {
	this.ValidateSession()
	if this.User != nil {
		this.Redirect(beego.URLFor(beego.AppConfig.String("home")), http.StatusFound)
	} else {
		this.TplName = "auth/index.html"
		this.Data["form"] = forms.AuthForm{}
		this.Data["validation"] = &validation.Validation{}
	}
}

func (this *AuthController) Auth() {
	form := &forms.AuthForm{}
	valid := &validation.Validation{}

	if err := this.ParseForm(form); err == nil {
		if correct, err := valid.Valid(form); err == nil && correct {
			this.SetSession("user", form.User.Id)
			this.Redirect(beego.URLFor(beego.AppConfig.String("home")), http.StatusFound)
		}
	}

	this.TplName = "auth/index.html"
	this.Data["form"] = form
	this.Data["validation"] = valid
}

func (this *AuthController) Logout() {
	this.DestroySession()
	this.Redirect(beego.URLFor(beego.AppConfig.String("logout")), http.StatusFound)
}
