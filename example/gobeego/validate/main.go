package main
import "github.com/astaxie/beego"
import "github.com/astaxie/beego/validation"
import "fmt"

type User struct{
	Name string `valid:"AlphaNumeric"`
	AccessIp string `valid:"IP"`
	Phone string `valid:"Mobile"`
	Email string `valid:"Email"`
	Age int `valid:"Min(0);Max(120)`
}

type ValidateController struct{
	beego.Controller
}


func (this *ValidateController) Post(){
	user := User{}
	this.ParseForm(&user)
	output := ""
	output += fmt.Sprintf("%+v\n", user)	
	valid := validation.Validation{}
	b, err := valid.Valid(&user)
    	if err != nil {
		output += fmt.Sprintf("Error: %v",  err)
        	this.Ctx.Output.Body([]byte(output))
    	}
    	if !b {
        	for _, err := range valid.Errors {
			output += fmt.Sprintln(err.Key, err.Message)
            		this.Ctx.Output.Body([]byte(output))
        	}
   	 }else{
		output += "OK, Passed!\n"
		this.Ctx.Output.Body([]byte(output))
	}
	
}

func main() {
	beego.Router("/", &ValidateController{}, "get,post:Post")		

	beego.Run()
}

