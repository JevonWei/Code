package main

import (
	"gopkg.in/gomail.v2"
	// "github.com/astaxie/beego"
)

func main() {
	server := "smtp.qq.com"
	port := 465

	user := "1083101646@qq.com"
	password := "euiqrqteqqaufgjf"

	to := []string{"1083101646@qq.com"}

	m := gomail.NewMessage()
	m.SetHeader("From", user)
	m.SetHeader("To", to...)
	m.SetHeader("Subject", "测试邮件")
	m.SetBody("text/html", "<div style='color:red;'>测试邮件</div>测试邮件")
	m.Attach("./email.go")

	d := gomail.NewDialer(server, port, user, password)

	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}
}
