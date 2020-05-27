package main
import "github.com/astaxie/beego"


func main(){
	//beego.SetLogger("file", `{"filename":"logs/test.log"}`)
	//beego.SetLogger("console", "")
	
	beego.SetLogFuncCall(true)

	beego.Emergency("this is emergency")
	beego.Alert("this is alert")
	beego.Critical("this is critical")
	beego.Error("this is error")
	beego.Warning("this is warning")
	beego.Notice("this is notice")
	beego.Informational("this is informational")
	beego.Debug("this is debug")
	beego.Run()
}
