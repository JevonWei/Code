package controllers

import (
	"net/http"
	"strings"
	"time"
	"todolist/forms"
	"todolist/models"
	"todolist/utils"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/validation"
)

type UserController struct {
	LoginRequiredController
}

func (this *UserController) Prepare() {
	this.LoginRequiredController.Prepare()

	_, action := this.GetControllerAndAction()
	// 除user/password 外其他url只有超级管理员可操作
	if action != "Password" && !this.User.IsSuper {
		// 非管理员且并非请求user/password则跳转到主页面
		this.Redirect(beego.URLFor(beego.AppConfig.String("home")), http.StatusFound)
	}

	this.Layout = "layout/base.html"
	this.Data["nav"] = "user"
}

// 用户列表
func (this *UserController) Index() {
	q := strings.TrimSpace(this.GetString("q"))

	var users []models.User

	// 创建查询条件
	condition := orm.NewCondition()
	if q != "" {
		// 查询名称、电话、地址、描述中包含字符
		condition = condition.Or("name__icontains", q)
		condition = condition.Or("tel__icontains", q)
		condition = condition.Or("addr__icontains", q)
		condition = condition.Or("desc__icontains", q)
	}

	orm.NewOrm().QueryTable(&models.User{}).SetCond(condition).All(&users)

	this.Data["q"] = q
	this.Data["users"] = users
	this.TplName = "user/index.html"
}

// 用户创建
func (this *UserController) Create() {
	form := &forms.UserCreateForm{}   // 用户创建表单
	valid := &validation.Validation{} // 验证器

	if this.Ctx.Input.IsPost() {
		if this.ParseForm(form) == nil {
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

				// 设置密码
				user.SetPassword(form.Password)

				// 插入用户
				ormer := orm.NewOrm()
				ormer.Insert(user)

				// 通过flash提示用户操作结果
				flash := beego.NewFlash()
				flash.Success("添加用户成功")
				flash.Store(&this.Controller)
				this.Redirect(beego.URLFor("UserController.Index"), http.StatusFound)
			}
		}
	}

	this.TplName = "user/create.html"
	this.Data["xsrf_token"] = this.XSRFToken() // 生成csrftoken值
	this.Data["form"] = form
	this.Data["validation"] = valid
}

func (this *UserController) Modify() {
	form := &forms.UserModifyForm{}   // 修改表单
	valid := &validation.Validation{} // 验证器

	if this.Ctx.Input.IsGet() {
		// 获取用户信息显示
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
				// 更改用户信息
				birthday, _ := time.Parse("2006-01-02", form.Birthday)
				form.User.Name = form.Name
				form.User.Birthday = &birthday
				form.User.Sex = form.Sex == 1
				//form.User.Sex = true
				form.User.Tel = form.Tel
				form.User.Addr = form.Addr
				form.User.Desc = form.Desc

				// 更新数据库
				ormer := orm.NewOrm()
				ormer.Update(form.User)

				// 通过flash提示用户操作结果
				flash := beego.NewFlash()
				flash.Success("修改用户成功")
				flash.Store(&this.Controller)
				this.Redirect(beego.URLFor("UserController.Index"), http.StatusFound)
			}
		}
	}
	this.TplName = "user/modify.html"
	this.Data["xsrf_token"] = this.XSRFToken() // 生成csrftoken
	this.Data["form"] = form
	this.Data["validation"] = valid
}

// 用户删除
func (this *UserController) Delete() {
	if id, err := this.GetInt("id"); err == nil {
		// 删除用户
		orm.NewOrm().Delete(&models.User{Id: id})

		// 通过flash提示用户操作结果
		flash := beego.NewFlash()
		flash.Success("删除用户成功")
		flash.Store(&this.Controller)
	}
	this.Redirect(beego.URLFor("UserController.Index"), http.StatusFound)
}

// 修改密码
func (this *UserController) Password() {
	form := &forms.ModifyPasswordForm{User: this.User} //修改密码表单
	valid := &validation.Validation{}                  // 验证器
	if this.Ctx.Input.IsPost() {

		// 解析请求参数到form中(根据form标签)
		if this.ParseForm(form) == nil {

			// 表单验证
			if corret, err := valid.Valid(form); err == nil && corret {

				// 设置密码
				this.User.SetPassword(form.NewPassword)

				ormer := orm.NewOrm()

				// 只更新密码列
				ormer.Update(this.User, "Password")

				// 通过flash通知用户操作结果(注意，此处只使用flash对象，并未调用store到cookie中)
				flash := beego.NewFlash()
				flash.Success("修改密码成功")
				this.Data["flash"] = flash.Data
				this.Redirect(beego.URLFor("UserController.Index"), http.StatusFound)
			}
		}
	}
	this.TplName = "user/password.html"
	this.Data["xsrf_token"] = this.XSRFToken() //生成csrftoken
	this.Data["form"] = form
	this.Data["validation"] = valid

}

//重置密码
func (this *UserController) ResetPassword() {
	if id, err := this.GetInt("id"); err == nil && this.User.IsSuper {

		user := models.User{Id: id}
		ormer := orm.NewOrm()
		if err := ormer.Read(&user); err == nil {
			password := utils.RandomString(6) // 生成随机6位密码
			user.SetPassword(password)        // 设置密码
			ormer.Update(&user, "Password")   //更新数据库

			// 通过flash通知用户重置成功和重置后密码
			flash := beego.NewFlash()
			flash.Success("重置用户<%s>密码为<%s>", user.Name, password)
			flash.Store(&this.Controller)

		}
	}
	this.Redirect(beego.URLFor("UserController.Index"), http.StatusFound)
}
