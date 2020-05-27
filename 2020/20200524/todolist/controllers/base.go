package controllers

import (
	"fmt"
	"net/http"
	"todolist/models"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

type BaseController struct {
	beego.Controller
	User *models.User
}

func (this *BaseController) ValidateSession() {
	if session := this.GetSession("user"); session != nil {
		if id, ok := session.(int); ok {
			user := models.User{Id: id}
			if err := models.GetUser(&user); err == nil {
				this.User = &user
				this.Data["user"] = &user
			} else {
				fmt.Printf("Can not get user for id %s, error: %s\n", id, err.Error())
			}
		} else {
			fmt.Printf("Invalid session id %v\n", session)
		}
	} else {
		fmt.Printf("Not Login\n")
	}
}

type LoginRequiredController struct {
	BaseController
}

func (this *LoginRequiredController) Prepare() {
	this.ValidateSession()
	if this.User == nil {
		this.Redirect(beego.URLFor(beego.AppConfig.String("login")), http.StatusFound)
	}
}

func init() {
	// log := logs.NewLogger(1000)
	logs.SetLogger(logs.AdapterFile, `{"filename":"logs/todolist.log"}`)
}
