package main

import (
	"github.com/astaxie/beego"
    "fmt"
)

type SessionController struct{
    beego.Controller
}
var session int = 1000
func (this *SessionController) HandleRequest(){
    if this.GetSession("sessionid") != nil{
       this.Ctx.Output.Body([]byte(fmt.Sprintf("Session id: %v\n", this.GetSession("sessionid")))) 
    }else{
       this.Ctx.Output.Body([]byte(fmt.Sprintf("New ession: %v\n", session)))
       this.SetSession("sessionid", session)
       session++
    }
}

func main() {
    //Invalid in latest version
    //beego.SessionOn = true
    beego.Router("/session", &SessionController{}, "get,post:HandleRequest") 
	beego.Run()
}

