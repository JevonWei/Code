package main

import (
	"encoding/json"
	"fmt"

	"github.com/astaxie/beego"
)

type ParamsControlltr struct {
	beego.Controller
}

func (this *ParamsControlltr) HandleParams() {
	output := fmt.Sprintf("%T %v\n", this.GetString("string"), this.GetString("string"))
	output += fmt.Sprintf("%T %v\n", this.GetStrings("string"), this.GetStrings("string"))
	boolValue, err := this.GetBool("bool")
	if err == nil {
		output += fmt.Sprintf("%T %v\n", boolValue, boolValue)
	} else {
		output += fmt.Sprintf("No bool Value\n")
	}

	this.Ctx.Output.Body([]byte(output))
}

type user struct {
	Id    int         `form:"_"`
	Name  interface{} `form:username`
	Age   int         `form:age`
	Email string
}

func (this *ParamsControlltr) HandleForm() {
	u := user{}
	this.ParseForm(&u)
	this.Ctx.Output.Body([]byte(fmt.Sprintf("%+v\n", u)))
}

func (this *ParamsControlltr) HandleBody() {
	u := user{}
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &u)
	if err != nil {
		this.Ctx.Output.Body([]byte(fmt.Sprintf("%v\n", err)))
	} else {
		this.Ctx.Output.Body([]byte(fmt.Sprintf("%v\n", u)))
	}
}

func (this *ParamsControlltr) HandleFile() {
	f, h, err := this.GetFile("uploadname")
	if err != nil {
		this.Ctx.Output.Body([]byte(fmt.Sprintf("%v\n", err)))
	} else {
		defer f.Close()
		err := this.SaveToFile("uploadname", "D:\\Code\\goang\\practise\\20200524"+h.Filename)
		if err == nil {
			this.Ctx.Output.Body([]byte("ok"))
		} else {
			this.Ctx.Output.Body([]byte(fmt.Sprintf("%V\n", err)))
		}
	}
}

func main() {
	beego.Router("/params", &ParamsControlltr{}, "get,post:HandleParams")
	beego.Router("/form", &ParamsControlltr{}, "get,post:HandleForm")
	beego.Router("/body", &ParamsControlltr{}, "get,post:HandleBody")
	beego.Router("/file", &ParamsControlltr{}, "get,post:HandleFile")
	beego.Run()
}
