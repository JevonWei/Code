package controllers
import  "github.com/astaxie/beego"

type FirstController struct{
    beego.Controller
    output string
}

func (this *FirstController) Prepare(){
    //called before method
    this.output = "Prepare\n"
}
func outputMethod(controller *FirstController){
    controller.output += (controller.Ctx.Input.Method() + "\n")
}
func (this *FirstController)Get(){
    outputMethod(this)
}
func (this *FirstController)Post(){
    outputMethod(this)
}
func (this *FirstController)Put(){
    outputMethod(this)
}
func (this *FirstController)Delete(){
    outputMethod(this)
}
func (this *FirstController)Head(){
    outputMethod(this)
}
func (this *FirstController)Patch(){
    outputMethod(this)
}
func (this *FirstController)Options(){
    outputMethod(this)
}
func (this *FirstController)Finish(){
    //not output, Finish is called after Render
    this.output += "Finish\n"
}
func (this *FirstController)Render() error{
   this.Ctx.Output.Body([]byte(this.output))
   return nil
}
