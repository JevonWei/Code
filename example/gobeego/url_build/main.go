package main
import "github.com/astaxie/beego"
import "fmt"

type TestController struct {
	beego.Controller
}

func (this *TestController) Get() {
	this.Data["Username"] = "astaxie"
	this.Ctx.Output.Body([]byte("ok"))
}

func (this *TestController) List() {
	this.Ctx.Output.Body([]byte("i am list"))
}


func (this *TestController) Myext() {
	this.Ctx.Output.Body([]byte(this.Ctx.Input.Param(":ext")))
}

func (this *TestController) GetUrl() {
	this.Ctx.Output.Body([]byte(this.URLFor(".Myext")))
}



func main(){
    	beego.Router("/api/list", &TestController{}, "*:List")
	beego.Router("/person/:last/:first", &TestController{})
	beego.AutoRouter(&TestController{})

        fmt.Println(beego.URLFor("TestController.List"))//  /api/list
        fmt.Println(beego.URLFor("TestController.Get", ":last", "xie", ":first", "asta"))//  /person/xie/asta
        fmt.Println(beego.URLFor("TestController.Myext"))//  /test/myext  
        fmt.Println(beego.URLFor("TestController.GetUrl"))//  /test/getUrl

}
