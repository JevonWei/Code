package controllers

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/validation"
	"time"
	"strings"
	"todolist/forms"
	"todolist/models"
)

type TaskController struct {
	LoginRequiredController
}

// 访问前验证，session存在则继续，否则跳转至登录界面
func (this *TaskController) Prepare() {
	this.LoginRequiredController.Prepare()
}

func (this *TaskController) Index() {
	// 将数据返回
	this.Layout = "layout/base.html"
	this.Data["nav"] = "task"
	this.LayoutSections = map[string]string{}
	this.LayoutSections["LayoutScripts"] = "task/index_scripts.html"
	this.TplName = "task/index.html"
	this.Data["xsrf_token"] = this.XSRFToken()
	this.Data["statusTexts"] = models.TaskStatusTexts
}

func (this *TaskController) List() {
	orderByColumns := map[string]bool{
		"id": true,
		"name": true,
		"status": true,
		"progress": true,
		"worker": true,
	} 
	draw := this.GetString("draw")
	start, err := this.GetInt("start")
	if err != nil {
		start = 0
	}
	length, err := this.GetInt("length")
	if err != nil {
		length = 10
	}

	

	q := strings.TrimSpace(c.GetString("search[value]", ""))

	orderBy := this.GetString("orderBy")
	if _, ok := orderByColumns[orderBy]; !ok {
		orderBy = "id"
	}
	orderDir := this.GetString("orderDir")
	
	if orderDir == "desc" {
		orderBy = "-" + orderBy
	}


	var tasks []*models.Task

	condition := orm.NewCondition()

	// 如果不是管理员用户，则增加一条，只查询属于自己的数据的条件
	if !this.User.IsSuper{
		condition = condition.And("create_user__exact", this.User.Id)
	}

	ormer := orm.NewOrm()
	queryset := ormer.QueryTable("task")
	total, _ := queryset.SetCond(condition).Count()

	totalFilter := total

	if q != "" {
		qcondition := orm.NewCondition()
		qcondition = qcondition.Or("name__icontains", q)
		qcondition = qcondition.Or("desc__icontains", q)
		qcondition = qcondition.Or("worker__icontains", q)
		condition = condition.AndCond(qcondition)

		totalFilter, _ = queryset.SetCond(condition).Count()
	}

	

	// 根据条件查询
	queryset.SetCond(condition).OrderBy(orderBy).Limit(length).Offset(start).All(&tasks)
	for _, task := range tasks {
		task.Patch()
	}
	this.Data["json"] = map[string]interface{} {
		"code": 200,
		"text": "获取任务成功",
		"result": tasks,
		"draw": draw,
		"recordsTotal": total,
		"recordsFiltered": totalFilter,
	}
	//this.Data["json"] = json
	this.ServeJSON()
}

func (this *TaskController) Create() {
	json := map[string]interface{}{
		"code":   405,
		"text":   "请求方式错误",
		"result": nil,
	}

	form := &forms.TaskCreateForm{}
	valid := &validation.Validation{}

	// 验证请求方法
	if this.Ctx.Input.IsPost() {
		json["code"], json["text"] = 400, "请求数据错误"
		// 验证数据输入是否正确
		if err := this.ParseForm(form); err != nil {
			json["text"] = err.Error()
		} else {
			fmt.Println(form)
			if corret, err := valid.Valid(form); err != nil {
				json["text"] = err.Error()
			} else if !corret {
				json["result"] = valid.Errors
			} else {
				task := &models.Task{
					Name:       form.Name,
					Worker:     form.Worker,
					CreateUser: this.User.Id,
					Desc:       form.Desc,
				}
				// 验证通过则插入用户数据
				o := orm.NewOrm()
				if _, err := o.Insert(task); err == nil {
					json["code"], json["text"], json["result"] = 200, "创建成功", task
				} else {
					json["code"], json["text"], json["result"] = 500, "服务器端错误", nil
				}
			}
		}
	}

	//  失败则返回错误信息至创建页面
	this.Data["json"] = json
	this.ServeJSON()
}

func (this *TaskController) Detail() {
	json := map[string]interface{}{
		"code":   400,
		"text":   "请求数据错误",
		"result": nil,
	}

	if id, err := this.GetInt("id"); err == nil {
		task := &models.Task{Id: id}
		if orm.NewOrm().Read(task) == nil && (this.User.IsSuper || task.CreateUser == this.User.Id) {
			json["code"], json["text"], json["result"] = 200, "获取数据成功", task
		}
	}
	this.Data["json"] = json
	this.ServeJSON()
}

func (this *TaskController) Modify() {
	json := map[string]interface{}{
		"code":   405,
		"text":   "请求方式错误",
		"result": nil,
	}

	form := &forms.TaskModifyForm{User: this.User}
	valid := &validation.Validation{}

	// 验证请求方法
	if this.Ctx.Input.IsPost() {
		json["code"], json["text"] = 400, "请求数据错误"
		if err := this.ParseForm(form); err != nil {
			json["text"] = err.Error()
		} else {
			// 如果校验通过
			if correct, err := valid.Valid(form); err != nil {
				json["text"] = err.Error()
			} else if !correct {
				json["result"] = valid.Errors
			} else {
				// 给form加上用户属性
				form.User = this.User
				// 给任务结构体赋值
				form.Task.Name = form.Name
				form.Task.Progress = form.Progress
				form.Task.Status = form.Status
				form.Task.Worker = form.Worker
				form.Task.Desc = form.Desc

				// 如果进度为完成，则将当前时间写入完成时间
				if form.Status == models.TaskStatusComplete {
					now := time.Now()
					form.Task.CompleteTime = &now
					form.Task.Progress = 100
				}

				// 更新数据
				if _, err := orm.NewOrm().Update(form.Task); err == nil {
					json["code"], json["text"], json["result"] = 200, "修改成功", form.Task
				} else {
					json["code"], json["text"], json["result"] = 500, "服务器端错误", nil
				}
			}
		}
	}

	// 将json数据返回
	this.Data["json"] = json
	this.ServeJSON()
}

func (this *TaskController) Delete() {
	// if id, err := this.GetInt("id"); err == nil {
	// 	ormer := orm.NewOrm()
	// 	task := models.Task{Id: id}
	// 	if ormer.Read(&task) == nil && (this.User.IsSuper || this.User.Id == task.CreateUser ) {
	// 		ormer.Delete(&task)
	// 	}
	// }

	// c.Data["json"] = map[string]interface{} {
	// 	"code": 200,
	// 	"text": "删除成功"，
	// }
	// this.ServeJSON()

	json := map[string]interface{}{
		"code":   405,
		"text":   "请求方式错误",
		"result": nil,
	}
	// 判断请求方法
	if this.Ctx.Input.IsPost() {
		json["code"], json["text"] = 400, "请求数据错误"
		// 获取url 传入的id
		id, _ := this.GetInt("id")
		task := models.Task{Id: id}
		// 根据id删除用户
		if this.User.IsSuper || this.User.Id == task.CreateUser {
			if result, err := orm.NewOrm().Delete(&task); err == nil {
				json["code"], json["text"], json["result"] = 200, "删除成功", result
			} else {
				json["code"], json["text"], json["result"] = 500, "服务器端错误", err.Error()
			}
		}
	}
	this.Data["json"] = json
	this.ServeJSON()
}
