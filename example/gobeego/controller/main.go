package main

import (
	"github.com/astaxie/beego"
    "controller/controllers"
)

func main(){
    beego.Router("/first", &controllers.FirstController{})
    beego.Run()
}
