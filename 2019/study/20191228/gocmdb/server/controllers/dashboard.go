package controllers

import (
	"github.com/JevonWei/gocmdb/server/controllers/auth"
	"github.com/JevonWei/gocmdb/server/models"
)

type DashboardPageController struct {
	LayoutController
}

func (c *DashboardPageController) Index() {
	c.Data["menu"] = "dashboard"

	c.TplName = "dashboard_page/index.html"
	c.LayoutSections["LayoutScript"] = "dashboard_page/index.script.html"
}

type DashboardController struct {
	auth.LoginRequiredController
}

func (c *DashboardController) Stat() {
	onlineCnt, offlineCnt := models.DefaultAgentManager.GetStatus()
	alarm_trend_days, alarm_trend_data := models.DefaultAlarmManager.GetLastestNStat(7)

	c.Data["json"] = map[string]interface{}{
		"code": 200,
		"text": "获取成功",
		"result": map[string]interface{}{
			"agent_offline_stat": offlineCnt,
			"agent_online_stat":  onlineCnt,
			"alarm_count":        models.DefaultAlarmManager.GetCountByNoComplate(),
			"alarm_dist":         models.DefaultAlarmManager.GetStatForNotComplete(),
			"alarm_trend": map[string]interface{}{
				"days": alarm_trend_days,
				"data": alarm_trend_data,
			},
		},
	}
	c.ServeJSON()
}
