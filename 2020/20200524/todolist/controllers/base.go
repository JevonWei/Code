package controllers

import (
	"fmt"
	"net/http"
	"todolist/models"

	"github.com/astaxie/beego"
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
		this.Redirect("/auth/login", http.StatusFound)
	}
}
