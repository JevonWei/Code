package main

import (
	"os"

	"github.com/sirupsen/logrus"
)

func main() {
	// 设置日志级别
	logrus.SetLevel(logrus.DebugLevel)
	logrus.Info("Info信息")
	logrus.Debug("Debug信息")
	logrus.Warn("Warn信息")
	logrus.Error("Error信息")

	// Field
	logrus.WithFields(logrus.Fields{
		"module": "main",
		"text":   "Jevon",
	}).Info("带field的Info日志")

	// 调用关系
	logrus.SetReportCaller(true)
	logrus.Error("记录调用关系")

	// JSON格式
	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.WithFields(logrus.Fields{
		"module": "main",
		"text":   "Jevon",
	}).Info("JSON格式的Info日志")

	// 日志输出到文件
	logfile, _ := os.OpenFile("main.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0777)
	defer func() {
		logfile.Close()
	}()
	logrus.SetOutput(logfile)
	logrus.Info("记录日志到文件")

	// 定义局部log环境
	log := logrus.New()
	log.SetLevel(logrus.DebugLevel)
	log.Debug("text")
	log.Debugf("占位输出_%s", "Jevon")
}
