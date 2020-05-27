package controllers

import (
	"strings"
	"time"

	"todolist/forms"
	"todolist/models"

	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/validation"
)

// 任务控制器
type TaskController struct {
	LoginRequiredController
}

// 任务列表
func (c *TaskController) Index() {

	c.Layout = "layout/base.html" // 设置layout
	c.Data["nav"] = "task"        // 设置菜单
	c.LayoutSections = map[string]string{}
	c.LayoutSections["LayoutScripts"] = "task/index_scripts.html"

	c.TplName = "task/index.html"
	c.Data["xsrf_token"] = c.XSRFToken() // csrftoken
	c.Data["statusTexts"] = models.TaskStatusTexts
}

func (c *TaskController) List_Task() {
	orderByColumns := map[string]bool{
		"id":          true,
		"name":        true,
		"status":      true,
		"progress":    true,
		"worker":      true,
		"create_user": true,
	}

	draw := c.GetString("draw")
	start, err := c.GetInt("start")
	if err != nil {
		start = 0
	}
	length, err := c.GetInt("length")
	if err != nil {
		length = 10
	}

	q := strings.TrimSpace(c.GetString("search[value]", ""))

	orderBy := c.GetString("orderBy")
	if _, ok := orderByColumns[orderBy]; !ok {
		orderBy = "id"
	}

	orderDir := c.GetString("orderDir")
	if orderDir == "desc" {
		orderBy = "-" + orderBy
	}

	var tasks []*models.Task

	// 创建查询条件
	condition := orm.NewCondition()

	// 若非超级管理员只查看当前用户任务
	if !c.User.IsSuper {
		condition = condition.And("create_user__exact", c.User.Id)
	}

	ormer := orm.NewOrm()
	queryset := ormer.QueryTable("task")
	total, _ := queryset.SetCond(condition).Count()

	totalFilter := total

	if q != "" {
		// 创建查询条件
		qcondition := orm.NewCondition()
		// 查询名称和描述中包含字符
		qcondition = qcondition.Or("name__icontains", q)
		qcondition = qcondition.Or("desc__icontains", q)
		qcondition = qcondition.Or("worker__icontains", q)
		condition = condition.AndCond(qcondition)

		// 过滤查询之后的数据数量
		totalFilter, _ = queryset.SetCond(condition).Count()
	}

	// 查询数据
	queryset.SetCond(condition).OrderBy(orderBy).Limit(length).Offset(start).All(&tasks)
	for _, task := range tasks {
		task.Patch()
	}
	c.Data["json"] = map[string]interface{}{
		"code":            200,
		"text":            "获取任务成功",
		"result":          tasks,
		"draw":            draw,
		"recordsTotal":    total,
		"recordsFiltered": totalFilter,
	}
	c.ServeJSON()
}

// 任务创建
func (c *TaskController) Create() {
	json := map[string]interface{}{
		"code":   405,
		"text":   "请求方式错误",
		"result": nil,
	}

	if c.Ctx.Input.IsPost() {
		json = map[string]interface{}{
			"code":   400,
			"text":   "提交数据错误",
			"result": nil,
		}

		form := &forms.TaskCreateForm{}   // 任务创建表单
		valid := &validation.Validation{} // 验证器

		// 解析请求参数到form中(根据form标签)
		if err := c.ParseForm(form); err != nil {
			json["text"] = err.Error()
		} else {
			// 表单验证
			if corret, err := valid.Valid(form); err != nil {
				json["text"] = err.Error()
			} else if !corret {
				json["result"] = valid.Errors
			} else {
				// 创建任务
				task := &models.Task{
					Name:       form.Name,
					Worker:     form.Worker,
					Desc:       form.Desc,
					CreateUser: c.User.Id,
				}

				ormer := orm.NewOrm()
				if _, err := ormer.Insert(task); err == nil {
					json = map[string]interface{}{
						"code":   200,
						"text":   "创建成功",
						"result": task,
					}
				} else {
					json = map[string]interface{}{
						"code":   500,
						"text":   "服务端错误",
						"result": nil,
					}
				}
			}
		}
	}
	c.Data["json"] = json
	c.ServeJSON()
}

func (c *TaskController) Modify() {
	json := map[string]interface{}{
		"code":   405,
		"text":   "请求方式错误",
		"result": nil,
	}

	form := &forms.TaskModifyForm{User: c.User} // 任务修改表单
	valid := &validation.Validation{}

	if c.Ctx.Input.IsGet() {
		// 获取任务信息（超级管理员可查看任意任务信息，普通管理员只能查看自己创建的任务信息）
		if id, err := c.GetInt("id"); err == nil {
			task := models.Task{Id: id}
			if orm.NewOrm().Read(&task) == nil && (c.User.IsSuper || task.CreateUser == c.User.Id) {
				json["code"] = 200
				json["text"] = "获取数据成功"
				json["result"] = task
			}
		}
		c.Data["json"] = json
		c.ServeJSON()
	} else if c.Ctx.Input.IsPost() {
		// 任务修改
		json = map[string]interface{}{
			"code":   400,
			"text":   "请求数据错误",
			"result": nil,
		}
		// 解析请求参数到form中(根据form标签)
		if err := c.ParseForm(form); err != nil {
			json["text"] = err.Error()
		} else {
			if corret, err := valid.Valid(form); err != nil {
				json["text"] = err.Error()
			} else if !corret {

				json["result"] = valid.Errors

			} else {
				// 验证任务信息（超级管理员可修改任意任务信息，普通管理员只能修改自己创建的任务信息）
				form.Task.Name = form.Name
				form.Task.Progress = form.Progress
				form.Task.Status = form.Status
				form.Task.Worker = form.Worker
				form.Task.Desc = form.Desc

				// 任务完成修改完成时间和进度
				if form.Status == models.TastStatusComplete {
					now := time.Now()
					form.Task.CompleteTime = &now
					form.Task.Progress = 100
				}

				// 任务修改
				if _, err := orm.NewOrm().Update(form.Task); err == nil {
					json["code"] = 200
					json["text"] = "修改成功"
					json["result"] = form.Task

				} else {

					json["code"] = 500
					json["text"] = "服务器错误"
					json["result"] = nil
				}
			}
		}
	}
	c.Data["json"] = json
	c.ServeJSON()

}

// 任务删除逻辑
func (c *TaskController) Delete() {
	if id, err := c.GetInt("id"); err == nil {
		ormer := orm.NewOrm()
		task := models.Task{Id: id}
		if ormer.Read(&task) == nil && (c.User.IsSuper || c.User.Id == task.CreateUser) {
			ormer.Delete(&task)
		}
	}

	c.Data["json"] = map[string]interface{}{
		"code": 200,
		"text": "success",
	}
	c.ServeJSON()
}
