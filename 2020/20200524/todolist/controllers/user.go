package controllers

import (
	"fmt"
	"net/http"
	"strings"
	"time"
	"todolist/forms"
	"todolist/models"

	"todolist/utils"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/validation"
)

type UserController struct {
	LoginRequiredController
}

func (this *UserController) Prepare() {
	this.LoginRequiredController.Prepare()
	_, action := this.GetControllerAndAction()

	if action != "Password" && !this.User.IsSupper {
		this.Redirect(beego.URLFor(beego.AppConfig.String("home")), http.StatusFound)
	}

	this.Layout = "layout/base.html"
	this.Data["nav"] = "user"
}

func (this *UserController) Index() {
	q := strings.TrimSpace(this.GetString("q"))

	var users []models.User

	condition := orm.NewCondition()
	if q != "" {
		condition = condition.Or("name__icontains", q)
		condition = condition.Or("tel__icontains", q)
		condition = condition.Or("addr__icontains", q)
		condition = condition.Or("desc__icontains", q)
		condition = condition.AndCond(condition)
	}

	orm.NewOrm().QueryTable(&models.User{}).SetCond(condition).All(&users)

	this.TplName = "user/index.html"
	this.Data["users"] = users
	this.Data["q"] = q
}

func (c *UserController) Create() {
	form := &forms.UserCreateForm{}   // 用户创建表单
	valid := &validation.Validation{} //验证器

	if c.Ctx.Input.IsPost() {
		// 解析请求参数到form中(根据form标签)
		if c.ParseForm(form) == nil {

			// 表单验证
			if corret, err := valid.Valid(form); err == nil && corret {
				// 转换时间
				birthday, _ := time.Parse("2006-01-02", form.Birthday)

				// 创建结构体对象
				user := &models.User{
					Name:     form.Name,
					Birthday: &birthday,
					Sex:      form.Sex == 1,
					Tel:      form.Tel,
					Addr:     form.Addr,
					Desc:     form.Desc,
				}

				//设置密码
				user.SetPassword(form.Password)

				// 插入用户
				ormer := orm.NewOrm()
				ormer.Insert(user)

				// 通过flash提示用户操作结果
				flash := beego.NewFlash()
				flash.Success("添加用户成功")
				flash.Store(&c.Controller)
				c.Redirect(beego.URLFor("UserController.Index"), http.StatusFound)
			}
		}
	}

	c.TplName = "user/create.html"
	c.Data["form"] = form
	c.Data["xsrf_token"] = c.XSRFToken() //生成csrftoken值
	c.Data["validation"] = valid

}

func (this *UserController) Modify() {
	form := &forms.UserModifyForm{}
	valid := &validation.Validation{}

	if this.Ctx.Input.IsGet() {
		if id, err := this.GetInt("id"); err == nil {
			user := models.User{Id: id}
			if orm.NewOrm().Read(&user) == nil {
				form.Id = user.Id
				form.Name = user.Name
				if user.Sex {
					form.Sex = 1
				}

				if user.Birthday != nil {
					form.Birthday = user.Birthday.Format("2006-01-02")
				}

				form.Addr = user.Addr
				form.Tel = user.Tel
				form.Desc = user.Desc
			}
		}
	} else if this.Ctx.Input.IsPost() {
		if this.ParseForm(form) == nil {
			if corret, err := valid.Valid(form); err == nil && corret {
				birthday, _ := time.Parse("2006-01-02", form.Birthday)

				form.User.Name = form.Name
				form.User.Birthday = &birthday
				form.User.Sex = form.Sex == 1
				form.User.Tel = form.Tel
				form.User.Addr = form.Addr
				form.User.Desc = form.Desc

				o := orm.NewOrm()
				o.Update(form.User)

				flash := beego.NewFlash()
				flash.Success("修改用户成功")
				flash.Store(&this.Controller)
				this.Redirect(beego.URLFor("UserController.Index"), http.StatusFound)
			}
		}
	}

	this.TplName = "user/modify.html"
	this.Data["form"] = form
	this.Data["xsrf_token"] = this.XSRFToken() // 生成csrftoken

	this.Data["validation"] = valid
}

func (this *UserController) Delete() {
	if id, err := this.GetInt("id"); err == nil {

		orm.NewOrm().Delete(&models.User{Id: id})

		flash := beego.NewFlash()
		flash.Success("删除用户成功")
		flash.Store(&this.Controller)

	}

	this.Redirect(beego.URLFor("UserController.Index"), http.StatusFound)
}

func (this *UserController) Password() {
	form := &forms.ModifyPasswordForm{User: this.User}
	valid := &validation.Validation{}

	if this.Ctx.Input.IsPost() {
		if this.ParseForm(form) == nil {
			if corret, err := valid.Valid(form); err == nil && corret {
				this.User.SetPassword(form.NewPassword)

				o := orm.NewOrm()
				o.Update(this.User, "Password")

				flash := beego.NewFlash()
				flash.Success("修改密码成功")
				this.Data["flash"] = flash.Data
			}
		}
	}

	this.TplName = "user/password.html"
	this.Data["form"] = form
	this.Data["xsrf_token"] = this.XSRFToken() // 生成csrftoken

	this.Data["validation"] = valid

	// this.Redirect(beego.URLFor("UserController.Index"), http.StatusFound)
}

func (this *UserController) ResetPassword() {
	if id, err := this.GetInt("id"); err == nil && this.User.IsSupper {
		user := models.User{Id: id}
		o := orm.NewOrm()
		if err := o.Read(&user); err == nil {
			password := utils.RandomString(6)
			user.SetPassword(password)
			o.Update(&user, "Password")

			fmt.Printf("重置用户<%s>密码为<%s>\n", user.Name, password)
			logs.Info("重置用户<%s>密码为<%s>\n", user.Name, password)
			// 通过flash通知用户重置成功和重置后密码
			flash := beego.NewFlash()
			flash.Success("重置用户<%s>密码为<%s>", user.Name, password)
			flash.Store(&this.Controller)
		}
	}

	this.Redirect(beego.URLFor("UserController.Index"), http.StatusFound)
}

// func init() {
// 	// log := logs.NewLogger(1000)
// 	logs.SetLogger(logs.AdapterFile, `{"filename":"logs/todolist.log"}`)
// }
