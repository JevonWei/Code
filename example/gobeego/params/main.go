package main

import (
	"github.com/astaxie/beego"
    "fmt"
    "encoding/json"
)

type ParamsController  struct{
    beego.Controller
}

func (this *ParamsController) HandleParams(){
    
    output := fmt.Sprintf("%T %v\n", this.GetString("string"), this.GetString("string"))
    output += fmt.Sprintf("%T %v\n", this.GetStrings("string"), this.GetStrings("string"))
    boolValue, err := this.GetBool("bool")
    if err == nil{
    	output += fmt.Sprintf("%T %v\n", boolValue, boolValue)
    }else{
        output += fmt.Sprintf("No bool value\n")
    }
    this.Ctx.Output.Body([]byte(output))
    

}



type user struct {
	Id    int        `form:"-"`
	Name  interface{} `form:"username"`
	Age   int         `form:"age"`
	Email string
}
func (this *ParamsController) HandleForm(){
    u := user{}
    this.ParseForm(&u)
    this.Ctx.Output.Body([]byte(fmt.Sprintf("%+v\n", u)))
    

}

func (this *ParamsController) HandleBody(){
    u := user{}
    err := json.Unmarshal(this.Ctx.Input.RequestBody, &u)
    if err != nil{
        this.Ctx.Output.Body([]byte(fmt.Sprintf("%v\n", err)))
    }else{
        this.Ctx.Output.Body([]byte(fmt.Sprintf("%+v\n", u)))
    }
}

func (this *ParamsController) HandleFile() {
	f, h, err := this.GetFile("uploadname")
	if err != nil {
        this.Ctx.Output.Body([]byte(fmt.Sprintf("%v\n", err)))
	} else {
	    defer f.Close()
		err := this.SaveToFile("uploadname", "/Users/uploads/"+h.Filename)
        if err == nil{
            this.Ctx.Output.Body([]byte("ok"))
        }else{
            this.Ctx.Output.Body([]byte(fmt.Sprintf("%v\n", err)))
        }
	}
}
func main() {
    beego.Router("/params", &ParamsController{}, "get,post:HandleParams") 
    beego.Router("/form", &ParamsController{}, "get,post:HandleForm")
    beego.Router("/body", &ParamsController{}, "get,post:HandleBody")
    beego.Router("/file", &ParamsController{}, "get,post:HandleFile")
	beego.Run()
}

