package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"

	"github.com/JevonWei/gocmdb/server/models"
	_ "github.com/JevonWei/gocmdb/server/routers"
	"github.com/JevonWei/gocmdb/server/utils"
)

func main() {
	// 初始化命令行参数
	h := flag.Bool("h", false, "help")
	help := flag.Bool("help", false, "help")
	verbose := flag.Bool("v", false, "verbose")

	flag.Usage = func() {
		fmt.Println("usage: alarm -h")
		flag.PrintDefaults()
	}
	// 解析命令行参数
	flag.Parse()

	if *h || *help {
		flag.Usage()
		os.Exit(0)
	}

	// 设置日志到文件
	beego.SetLogger("file", `{
		"filename" : "logs/alarm.log",
		"level" : 7}`,
	)
	if !*verbose {
		//删除控制台日志
		beego.BeeLogger.DelLogger("console")
	} else {
		orm.Debug = true
	}

	// 初始化orm
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", beego.AppConfig.String("dsn"))

	// 测试数据库连接是否正常
	if db, err := orm.GetDB(); err != nil || db.Ping() != nil {
		beego.Error("数据库连接错误")
		os.Exit(-1)
	}

	host := beego.AppConfig.String("smtp::host")
	port, _ := beego.AppConfig.Int("smtp::port")
	user := beego.AppConfig.String("smtp::user")
	password := beego.AppConfig.String("smtp::password")
	emailSender := utils.NewEmail(host, port, user, password)
	to := beego.AppConfig.Strings("smtp::to")

	smsSender := utils.NewSms(
		beego.AppConfig.String("sms::endpoint"),
		beego.AppConfig.String("sms::secretId"),
		beego.AppConfig.String("sms::secretKey"),
		beego.AppConfig.String("sms::appid"),
		beego.AppConfig.String("sms::sign"),
	)

	templateOfflineId := beego.AppConfig.String("sms::templateOfflineId")
	templateCPUId := beego.AppConfig.String("sms::templateCPUId")
	templateRamId := beego.AppConfig.String("sms::templateRamId")
	phones := beego.AppConfig.Strings("sms::phones")

	go func() {
		// 离线告警
		offlineTime := 5
		noticeWindowTime := 60
		noticeCounter := int64(2)

		for now := range time.Tick(time.Minute) {
			// for now := range time.Tick(time.Second * 3) {
			//fmt.Println("离线告警", now)
			beego.Debug("离线告警", now)
			endTime := now.Add(-1 * time.Duration(offlineTime) * time.Minute) // 5 根据配置

			// 告警开始时间
			noticeStartTime := now.Add(-1 * time.Duration(noticeWindowTime) * time.Minute)
			var result []orm.Params
			orm.NewOrm().Raw("SELECT uuid, heartbeat_time from agent where deleted_time is null and heartbeat_time < ?", endTime).Values(&result)
			//fmt.Println(result)
			for _, line := range result {
				uuid := line["uuid"].(string)
				heartbeat_time := line["heartbeat_time"].(string)
				content := fmt.Sprintf("终端[%s]最后一次发送心跳时间为%s, 已超过离线时间%d分钟", uuid, heartbeat_time, offlineTime)

				alarmCnt := models.DefaultAlarmManager.GetCountByUuidAndType(uuid, models.AlarmTypeOffline, noticeStartTime)
				if alarmCnt >= noticeCounter {
					beego.Info(fmt.Sprintf("通知次数(%d)超过限制(%d), %s", alarmCnt, noticeCounter, content))
					continue
				}

				emailerr := emailSender.Send(to, "[CMDB]终端离线告警", content, []string{})

				params := []string{uuid, heartbeat_time, strconv.Itoa(offlineTime)}
				smsErr := smsSender.Send(templateOfflineId, phones, params)

				beego.Info("终端离线告警: ", content, ", email通知:", emailerr, ", sms通知:", smsErr)
				models.DefaultAlarmManager.Create(uuid, models.AlarmTypeOffline, content, now)
			}
		}
	}()

	go func() {
		windowTime := 5
		cpuThreshold := 5
		cpuCounter := 3
		noticeWindowTime := 60
		noticeCounter := int64(2)

		// CPU使用率
		for now := range time.Tick(time.Minute) {
			//fmt.Println("CPU使用率告警", now)
			beego.Debug("CPU使用率告警", now)
			startTime := now.Add(-1 * time.Duration(windowTime) * time.Minute) // 5 根据配置
			noticeStartTime := now.Add(-1 * time.Duration(noticeWindowTime) * time.Minute)
			var result []orm.Params
			orm.NewOrm().Raw("SELECT uuid, count(*) as cnt from resource where deleted_time is null and created_time >= ? and cpu_precent >= ? group by uuid having count(*) >= ?", startTime, cpuThreshold, cpuCounter).Values(&result)
			// fmt.Println(result)
			for _, line := range result {
				uuid := line["uuid"].(string)
				cntString, _ := line["cnt"].(string)
				cnt, _ := strconv.Atoi(cntString)
				content := fmt.Sprintf("终端[%s]在最近%d分钟内CPU使用率大于%d%%的次数为%d, 已超过%d", uuid, windowTime, cpuThreshold, cnt, cpuCounter)

				alarmCnt := models.DefaultAlarmManager.GetCountByUuidAndType(uuid, models.AlarmTypeCPU, noticeStartTime)
				if alarmCnt >= noticeCounter {
					beego.Info(fmt.Sprintf("通知次数(%d)超过限制(%d), %s", alarmCnt, noticeCounter, content))
					continue
				}

				emailerr := emailSender.Send(to, "[CMDB]终端CPU告警", content, []string{})
				params := []string{uuid, strconv.Itoa(windowTime), strconv.Itoa(cpuThreshold), strconv.Itoa(cnt), strconv.Itoa(cpuCounter)}
				smsErr := smsSender.Send(templateCPUId, phones, params)

				beego.Info("终端CPU告警: ", content, ", email通知:", emailerr, ", sms通知:", smsErr)

				models.DefaultAlarmManager.Create(uuid, models.AlarmTypeCPU, content, now)
			}
		}
	}()

	// 内存使用率
	windowTime := 5
	ramThreshold := 20
	ramCounter := 3
	noticeWindowTime := 60
	noticeCounter := int64(2)

	for now := range time.Tick(time.Minute) {
		//fmt.Println("内存使用率告警", now)
		beego.Debug("内存使用率告警", now)
		startTime := now.Add(-1 * time.Duration(windowTime) * time.Minute) // 5 根据配置
		noticeStartTime := now.Add(-1 * time.Duration(noticeWindowTime) * time.Minute)
		var result []orm.Params
		orm.NewOrm().Raw("SELECT uuid, count(*) as cnt from resource where deleted_time is null and created_time >= ? and ram_precent >= ? group by uuid having count(*) >= ?", startTime, ramThreshold, ramCounter).Values(&result)
		// fmt.Println(result)
		for _, line := range result {
			uuid := line["uuid"].(string)
			cntString, _ := line["cnt"].(string)
			cnt, _ := strconv.Atoi(cntString)
			content := fmt.Sprintf("终端[%s]在最近%d分钟内内存使用率大于%d%%的次数为%d, 已超过%d", uuid, windowTime, ramThreshold, cnt, ramCounter)

			alarmCnt := models.DefaultAlarmManager.GetCountByUuidAndType(uuid, models.AlarmTypeRAM, noticeStartTime)
			if alarmCnt >= noticeCounter {
				beego.Info(fmt.Sprintf("通知次数(%d)超过限制(%d), %s", alarmCnt, noticeCounter, content))
				continue
			}

			emailerr := emailSender.Send(to, "[CMDB]终端内存告警", content, []string{})
			params := []string{uuid, strconv.Itoa(windowTime), strconv.Itoa(ramThreshold), strconv.Itoa(cnt), strconv.Itoa(ramCounter)}
			smsErr := smsSender.Send(templateRamId, phones, params)

			beego.Info("终端内存告警: ", content, ", email通知:", emailerr, ", sms通知:", smsErr)

			models.DefaultAlarmManager.Create(uuid, models.AlarmTypeRAM, content, now)
		}
	}

}
