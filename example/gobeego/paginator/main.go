package main
import "github.com/astaxie/beego"
import "fmt"
import "paginator/utils"

type User struct { 
 		Id int  
		Name interface{} 
		Age int  
		Sex string 
} 

const total = 95
const pers = 10
var users = [total]User{}
func initData(){
	for index, _ := range(users){
		user := &users[index]
		user.Name = "Name_" + fmt.Sprintf("%d", index)
		user.Age = 10 + index
		if index % 2 == 0{
			user.Sex = "男"
		}else{
			user.Sex = "女"
		}
	}
}

type PaginatorController struct{
	beego.Controller
}

func(this *PaginatorController) List(){
	
	this.Layout = "layout.tpl"
	this.TplName = "paginator.html"
	p := utils.NewPaginator(this.Ctx.Request, pers, total)
	this.Data["paginator"] = p
	if p.Offset() + pers > total{
		this.Data["Users"] = users[p.Offset(): ]
	}else{
		this.Data["Users"] = users[p.Offset(): p.Offset() + pers]
	}
}


func main() {
	initData()	
	beego.Router("/", &PaginatorController{}, "*:List")
	beego.Run()
}
