package main
import "github.com/astaxie/beego"

type TemplateController struct{
	beego.Controller
}

type Person struct{
	name string
	Friends []string
	Emails []string
        Introduce func() string 
}


func (this *Person)Welcome(welcome string) string{
	return welcome + " " + this.name
}

func (this *TemplateController) Handle(){
	this.Data["person"] = &Person{name: "astaxie", Friends: []string{"asta", "xie"}, Emails: []string{"a@b", "c@d"}, Introduce: func()string{return "I am astaxie"}} 
	this.TplName = "index.html"
}

func(this *TemplateController) TestLayout(){
	this.Layout = "layout.tpl"
	this.TplName = "tpl.tpl"
}

func(this *TemplateController) LayoutSection(){
	this.Layout = "layout_section.tpl"
	this.TplName = "tpl.tpl"
	
	this.LayoutSections = make(map[string]string)
	this.LayoutSections["Section1"] = "section1.tpl"
	this.LayoutSections["Section2"] = "section2.tpl"
	
	
}

type User struct { 
 		Id int `form:"-"` 
		Name interface{} `form:"username"` 
		Age int `form:"age,text,年龄："` 
		Sex string 
		Intro string `form:",textarea"` 
} 


func(this *TemplateController) Form(){
	this.Layout = "layout.tpl"
	this.Data["Form"] = &User{Id: 1, Name: "Kai", Age: 18, Sex: "男", Intro: "IT 男"}	
}



func main() {
	beego.Router("/", &TemplateController{}, "*:Handle")
	beego.AutoRouter(&TemplateController{})
	beego.SetStaticPath("/test","test") 
	//beego.SetStaticPath("/","static")
	beego.Run()
}
