package controllers

import (
	"net/http"
	"strings"
	"time"
	"todolist/forms"
	"todolist/models"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/validation"
)

type TaskController struct {
	LoginRequiredController
}

func (this *TaskController) Prepare() {
	this.LoginRequiredController.Prepare()
	this.Layout = "layout/base.html"
	this.Data["nav"] = "task"
}

func (this *TaskController) Index() {
	q := this.GetString("q")
	q = strings.TrimSpace(q)

	condition := orm.NewCondition()
	if q != "" {
		condition = condition.Or("name__icontains", q)
		condition = condition.Or("desc__icontains", q)
		condition = condition.AndCond(condition)
	}

	if !this.User.IsSupper {
		condition = condition.And("create_user__exact", this.User.Id)
	}

	var tasks []models.Task
	orm.NewOrm().QueryTable(&models.Task{}).SetCond(condition).All(&tasks)

	this.TplName = "task/index.html"
	this.Data["tasks"] = tasks
	this.Data["q"] = q
}

func (this *TaskController) Create() {
	form := &forms.TaskCreateForm{}
	valid := &validation.Validation{}

	if this.Ctx.Input.IsPost() {
		if this.ParseForm(form) == nil {
			task := &models.Task{
				Name:       form.Name,
				Worker:     form.Worker,
				CreateUser: this.User.Id,
				Desc:       form.Desc,
			}

			o := orm.NewOrm()
			o.Insert(task)

			flash := beego.NewFlash()
			flash.Success("创建任务成功")
			flash.Store(&this.Controller)

			this.Redirect(beego.URLFor("TaskController.Index"), http.StatusFound)
		}
	}
	this.TplName = "task/create.html"
	this.Data["form"] = form
	this.Data["validation"] = valid
}

func (this *TaskController) Modify() {
	form := &forms.TaskModifyForm{}
	valid := &validation.Validation{}

	if this.Ctx.Input.IsGet() {
		id, _ := this.GetInt("id")
		task := models.Task{Id: id}
		if orm.NewOrm().Read(&task) == nil {
			form.Id = task.Id
			form.Name = task.Name
			form.Status = task.Status
			form.Progress = task.Progress
			form.Worker = task.Worker
			form.Desc = task.Desc
		}
	} else {
		if this.ParseForm(form) == nil {
			form.User = this.User
			if correct, err := valid.Valid(form); err == nil && correct {
				form.Task.Name = form.Name
				form.Task.Status = form.Status
				form.Task.Progress = form.Progress
				form.Task.Worker = form.Worker
				form.Task.Desc = form.Desc

				if form.Status == models.TastStatusComplete {
					now := time.Now()
					form.Task.CompleteTime = &now
					form.Task.Progress = 100
				}
				orm.NewOrm().Update(form.Task)
			}
		}
		this.Redirect(beego.URLFor("TaskController.Index"), http.StatusFound)
	}
	this.TplName = "task/modify.html"
	this.Data["form"] = form
	this.Data["validation"] = valid
	this.Data["statusTexts"] = models.TaskStatusTexts
}

func (this *TaskController) Delete() {

}
