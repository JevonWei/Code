package controllers

import (
	"strings"

	"github.com/JevonWei/gocmdb/server/cloud"
	"github.com/JevonWei/gocmdb/server/controllers/auth"
	"github.com/JevonWei/gocmdb/server/forms"
	"github.com/JevonWei/gocmdb/server/models"
	"github.com/astaxie/beego/validation"
)

type CloudPlatformPageController struct {
	LayoutController
}

func (c *CloudPlatformPageController) Index() {
	c.Data["menu"] = "cloud_management"
	c.Data["expand"] = "cloud_plateform_management"
	c.TplName = "cloud_plateform_page/index.html"
	c.LayoutSections["LayoutScript"] = "cloud_plateform_page/index.script.html"
}

type CloudPlatformController struct {
	auth.LoginRequiredController
}

func (c *CloudPlatformController) List() {
	//draw,start, length, q
	draw, _ := c.GetInt("draw")
	start, _ := c.GetInt64("start")
	length, _ := c.GetInt("length")
	q := strings.TrimSpace(c.GetString("q"))
	result, total, queryTotal := models.DefaultCloudPlatformManager.Query(q, start, length)
	// []*User, total, queryTotall

	c.Data["json"] = map[string]interface{}{
		"code":            200,
		"text":            "获取成功",
		"result":          result,
		"draw":            draw,
		"recordsTotal":    total,
		"recordsFiltered": queryTotal,
	}
	c.ServeJSON()
}

func (c *CloudPlatformController) Create() {
	if c.Ctx.Input.IsPost() {
		form := &forms.CloudPlatformCreateForm{}
		valid := &validation.Validation{}
		json := map[string]interface{}{
			"code":   400,
			"text":   "提交数据错误",
			"result": nil,
		}

		if err := c.ParseForm(form); err != nil {
			valid.SetError("error", err.Error())
			json["result"] = valid.Errors
		} else {
			if ok, err := valid.Valid(form); err != nil {
				valid.SetError("error", err.Error())
				json["result"] = valid.Errors
			} else if ok {
				result, err := models.DefaultCloudPlatformManager.Create(
					form.Name,
					form.Type,
					form.Addr,
					form.Region,
					form.AccessKey,
					form.SecrectKey,
					form.Remark,
					c.User,
				)
				if err == nil {
					json = map[string]interface{}{
						"code":   200,
						"text":   "创建成功",
						"result": result,
					}
				} else {
					json = map[string]interface{}{
						"code":   500,
						"text":   "创建失败,请重试",
						"result": nil,
					}
				}
			} else {
				json["result"] = valid.Errors
			}
		}
		c.Data["json"] = json
		c.ServeJSON()
	} else {
		c.TplName = "cloud_platform/create.html"
		c.Data["types"] = cloud.DefaultManager.Plugins
	}
}

func (c *CloudPlatformController) Modify() {
	if c.Ctx.Input.IsPost() {
		json := map[string]interface{}{
			"code": 400,
			"text": "提交数据错误",
		}
		form := &forms.CloudPlatformModifyForm{}
		valid := &validation.Validation{}
		if err := c.ParseForm(form); err == nil {
			if ok, err := valid.Valid(form); err != nil {
				valid.SetError("error", err.Error())
				json["result"] = valid.Errors
			} else if ok {
				result, err := models.DefaultCloudPlatformManager.Modify(
					form.Id,
					form.Name,
					form.Type,
					form.Addr,
					form.Region,
					form.AccessKey,
					form.SecrectKey,
					form.Remark,
				)

				if err == nil {
					json = map[string]interface{}{
						"code":   200,
						"text":   "更新成功",
						"result": result,
					}
				} else {
					json = map[string]interface{}{
						"code": 500,
						"text": "服务器错误",
					}
				}
			} else {
				json["result"] = valid.Errors
			}
		} else {
			valid.SetError("error", err.Error())
			json["result"] = valid.Errors
		}
		c.Data["json"] = json
		c.ServeJSON()

	} else {
		//get
		pk, _ := c.GetInt("pk")
		c.TplName = "cloud_platform/modify.html"
		c.Data["object"] = models.DefaultCloudPlatformManager.GetById(pk)
		c.Data["types"] = cloud.DefaultManager.Plugins
	}
}

func (c *CloudPlatformController) Delete() {
	c.Data["json"] = map[string]interface{}{
		"code":   400,
		"text":   "提交数据错误",
		"result": nil,
	}
	if c.Ctx.Input.IsPost() {
		pk, _ := c.GetInt("pk")
		models.DefaultCloudPlatformManager.DeleteById(pk)

		c.Data["json"] = map[string]interface{}{
			"code":   200,
			"text":   "删除成功",
			"result": nil,
		}
	}

	c.ServeJSON()
}

func (c *CloudPlatformController) Disable() {
	c.Data["json"] = map[string]interface{}{
		"code":   400,
		"text":   "提交数据错误",
		"result": nil,
	}

	if c.Ctx.Input.IsPost() {
		pk, _ := c.GetInt("pk")
		models.DefaultCloudPlatformManager.DisableById(pk)
	}

	c.Data["json"] = map[string]interface{}{
		"code":   200,
		"text":   "Disable成功",
		"result": nil,
	}
	c.ServeJSON()
}

func (c *CloudPlatformController) Enable() {
	c.Data["json"] = map[string]interface{}{
		"code":   400,
		"text":   "提交数据错误",
		"result": nil,
	}

	if c.Ctx.Input.IsPost() {
		pk, _ := c.GetInt("pk")
		models.DefaultCloudPlatformManager.EnableById(pk)
	}

	c.Data["json"] = map[string]interface{}{
		"code":   200,
		"text":   "Enable成功",
		"result": nil,
	}
	c.ServeJSON()
}

type VirtualMachinePageController struct {
	LayoutController
}

func (c *VirtualMachinePageController) Index() {
	c.Data["menu"] = "cloud_management"
	c.Data["expand"] = "virtual_machine_management"
	c.TplName = "virtual_machine_page/index.html"
	c.LayoutSections["LayoutScript"] = "virtual_machine_page/index.script.html"
}

type VirtualMachineController struct {
	auth.LoginRequiredController
}

func (c *VirtualMachineController) List() {
	//draw,start, length, q
	draw, _ := c.GetInt("draw")
	start, _ := c.GetInt64("start")
	length, _ := c.GetInt("length")
	q := strings.TrimSpace(c.GetString("q"))
	platform, _ := c.GetInt("platform")

	result, total, queryTotal := models.DefaultVirtualMachineManager.Query(q, platform, start, length)
	// []*User, total, queryTotall

	c.Data["json"] = map[string]interface{}{
		"code":            200,
		"text":            "获取成功",
		"result":          result,
		"draw":            draw,
		"recordsTotal":    total,
		"recordsFiltered": queryTotal,
	}
	c.ServeJSON()
}

func (c *VirtualMachineController) Start() {
	pk, _ := c.GetInt("pk")
	if vm := models.DefaultVirtualMachineManager.GetById(pk); vm != nil {
		if sdk, ok := cloud.DefaultManager.Cloud(vm.Platform.Type); ok {
			sdk.Init(vm.Platform.Addr, vm.Platform.Region, vm.Platform.AccessKey, vm.Platform.SecrectKey)
			if sdk.StartInstance(vm.UUID) == nil {
				c.Data["json"] = map[string]interface{}{
					"code":   200,
					"text":   "启动成功",
					"result": nil,
				}
				c.ServeJSON()
			}
		}
	}
	c.Data["json"] = map[string]interface{}{
		"code":   400,
		"text":   "启动失败",
		"result": nil,
	}
	c.ServeJSON()
}

func (c *VirtualMachineController) Stop() {
	pk, _ := c.GetInt("pk")
	if vm := models.DefaultVirtualMachineManager.GetById(pk); vm != nil {
		if sdk, ok := cloud.DefaultManager.Cloud(vm.Platform.Type); ok {
			sdk.Init(vm.Platform.Addr, vm.Platform.Region, vm.Platform.AccessKey, vm.Platform.SecrectKey)
			if sdk.StopInstance(vm.UUID) == nil {
				c.Data["json"] = map[string]interface{}{
					"code":   200,
					"text":   "停止成功",
					"result": nil,
				}
				c.ServeJSON()
			}
		}
	}
	c.Data["json"] = map[string]interface{}{
		"code":   400,
		"text":   "停止失败",
		"result": nil,
	}
	c.ServeJSON()
}

func (c *VirtualMachineController) Reboot() {
	pk, _ := c.GetInt("pk")
	if vm := models.DefaultVirtualMachineManager.GetById(pk); vm != nil {
		if sdk, ok := cloud.DefaultManager.Cloud(vm.Platform.Type); ok {
			sdk.Init(vm.Platform.Addr, vm.Platform.Region, vm.Platform.AccessKey, vm.Platform.SecrectKey)
			if sdk.RebootInstance(vm.UUID) == nil {
				c.Data["json"] = map[string]interface{}{
					"code":   200,
					"text":   "重启成功",
					"result": nil,
				}
				c.ServeJSON()
			}
		}
	}
	c.Data["json"] = map[string]interface{}{
		"code":   400,
		"text":   "重启失败",
		"result": nil,
	}
	c.ServeJSON()
}
