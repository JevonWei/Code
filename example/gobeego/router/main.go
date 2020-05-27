package main

import (
	"github.com/astaxie/beego"
    "github.com/astaxie/beego/context"
    "strings"
)

type FixController  struct{
    beego.Controller
}

func (this *FixController) Get(){
    this.Ctx.Output.Body([]byte("hello world for controller get"))
}


func (this *FixController) Put(){

    this.Ctx.Output.Body([]byte("hello world for controller put"))
}

func (this *FixController) Any(){
    //Can define but not invalid
    this.Ctx.Output.Body([]byte("hello world for controller any"))
}

type RController struct{
   beego.Controller
}

func (this *RController) Get(){
    id := this.Ctx.Input.Param( ":id")
    if id != ""{
        this.Ctx.Output.Body([]byte("id in regex: " + this.Ctx.Input.Param( ":id")))
    }else{
        this.Ctx.Output.Body([]byte("id in regex is empty "))
    }
}

type RController1 struct{
    beego.Controller
}

func (this *RController1) Get(){
    output := "path in regex: " + this.Ctx.Input.Param( ":path") + "\n"
    output += "ext in regex: " + this.Ctx.Input.Param( ":ext") + "\n"
    output += "splat in regex: " + this.Ctx.Input.Param( ":splat") + "\n"
    this.Ctx.Output.Body([]byte(output))
}

type SelfDefineController struct{
     beego.Controller
}

// @router /comment/:key [get]
func (this *SelfDefineController) ForGet(){
    if strings.ToLower(this.Ctx.Input.Method()) == "get"{
        this.Ctx.Output.Body([]byte("Yes, it's get"))
    }
}

 type BankAccount struct{
   beego.Controller
 }

// register the function
 func (b *BankAccount)Mapping(){
  b.Mapping()
}

 //@router /account/:id  [get]
 func (b *BankAccount) ShowAccount(){
    //logic
 }


 //@router /account/:id  [post]
 func (b *BankAccount) ModifyAccount(){
    //logic
 }

var ns = beego.NewNamespace("/v1",
	beego.NSNamespace("/shop",
		beego.NSGet("/:id", func(ctx *context.Context) {
			ctx.Output.Body([]byte("shopinfo"))
		}),
	),
	beego.NSNamespace("/order",
		beego.NSGet("/:id", func(ctx *context.Context) {
			ctx.Output.Body([]byte("orderinfo"))
		}),
	),
	beego.NSNamespace("/crm",
		beego.NSGet("/:id", func(ctx *context.Context) {
			ctx.Output.Body([]byte("crminfo"))
		}),
	),
)
func main() {
    //basic router
    beego.Get("/",func(ctx *context.Context){
     ctx.Output.Body([]byte("hello world for get"))
    })

    beego.Post("/",func(ctx *context.Context){
     ctx.Output.Body([]byte("hello world for post"))
    })

    beego.Put("/",func(ctx *context.Context){
     ctx.Output.Body([]byte("hello world for put"))
    })


    beego.Patch("/",func(ctx *context.Context){
     ctx.Output.Body([]byte("hello world for put"))
    })
    beego.Any("/",func(ctx *context.Context){
     ctx.Output.Body([]byte("hello world for " + ctx.Input.Method() + " in any"))
    })

    beego.Options("/",func(ctx *context.Context){
     ctx.Output.Body([]byte("hello world for options"))
    })

    //controller
    beego.Router("/fix", &FixController{}) 

    //regex， 
    //http://127.0.0.1:8080/api   http://127.0.0.1:8080/api/123
    beego.Router("/api/?:id", &RController{})
    //http://127.0.0.1:8080/api1/123
    beego.Router("/api1/:id", &RController{})
    //http://127.0.0.1:8080/api2/123  not http://127.0.0.1:8080/api2/abc
    beego.Router("/api2/:id([0-9]+)", &RController{})
    //http://127.0.0.1:8080/api2/abc   and http://127.0.0.1:8080/api2/123 
    beego.Router("/api3/:id([\\w]+)", &RController{})
    beego.Router("/api4/:id:int", &RController{})     
    beego.Router("/api5/:id:string", &RController{})     
    //http://127.0.0.1:8080/prefix/abc12  :id=12
    beego.Router("/prefix/abc:id:int", &RController{})
    
    //http://127.0.0.1:8080/path_ext/test.html  :path=test :ext=html
    beego.Router("/path_ext/*.*", &RController1{}) 
    //http://127.0.0.1:8080/splat/abc12  :id=abc12
    beego.Router("/splat/*", &RController1{}) 

    //self define  http://127.0.0.1:8080/selfdefine
    beego.Router("/selfdefine", &SelfDefineController{}, "get:ForGet")

    //auto  http://127.0.0.1:8080/selfdefine/forget
    beego.AutoRouter(&SelfDefineController{})

    //
    beego.Include(&SelfDefineController{})
    beego.Include(&BankAccount{})
    beego.AddNamespace(ns)
    beego.Run()
}

