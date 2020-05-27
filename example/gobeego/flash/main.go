package main

import (
	"github.com/astaxie/beego"
)

type GetController struct{
    	beego.Controller
}

type PostController struct{
    	beego.Controller
}

func (this *GetController) Get(){
	flash:=beego.ReadFromRequest(&this.Controller)
	if n,ok:=flash.Data["notice"];ok{
		this.Ctx.Output.Body([]byte(n))	
	}else{
		this.Ctx.Output.Body([]byte("Post not called"))
	}
}


func  (this *PostController) Post(){
	flash:=beego.NewFlash()
	flash.Notice("Post has been called")
	flash.Store(&this.Controller)
	this.Redirect("/get", 302)
}

func main() {
   	beego.Router("/get", &GetController{}, "get,post:Get") 
    	beego.Router("/post", &PostController{}, "get,post:Post") 
	beego.Run()
}

