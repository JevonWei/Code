package main

import "fmt"

type SignalSender interface {
	Send(to, msg string) error
}

// 定义接口类型
type Sender interface {
	Send(to string, msg string) error
	SendAll(tos []string, msg string) error
}

type EmailSender struct {
	SmtpAddr string
}

func (s EmailSender) Send(to, msg string) error {
	fmt.Println("发送邮件给: ", to, ", 消息内容是：", msg)
	return nil
}

func (s EmailSender) SendAll(tos []string, msg string) error {
	for _, to := range tos {
		s.Send(to, msg)
	}
	return nil
}

type SmsSender struct {
	SmsAPI string
}

func (s *SmsSender) Send(to, msg string) error {
	fmt.Println("发送短信给: ", to, ", 消息内容是：", msg)
	return nil
}

func (s *SmsSender) SendAll(tos []string, msg string) error {
	for _, to := range tos {
		s.Send(to, msg)
	}
	return nil
}

type WechatSender struct {
	ID string
}

func (s WechatSender) Send(to, msg string) error {
	fmt.Println("发送微信给: ", to, ", 消息内容是：", msg)
	return nil
}

func (s *WechatSender) SendAll(tos []string, msg string) error {
	for _, to := range tos {
		s.Send(to, msg)
	}
	return nil
}

func do(sender Sender) {
	sender.Send("领导", "工作日志")
}

func main() {
	var sender Sender = EmailSender{"Jevon@danran.com"}
	// var sender1 Sender = SmsSender{}

	// fmt.Printf("%T, %#v\n", sender, sender)

	// sender.Send("Jevon", "XXXXX")
	// sender.SendAll([]string{"Jevon", "Dan"}, "AAAAA")
	// sender1.SendAll([]string{"Jevon", "Dan"}, "AAAAA")

	do(sender)

	//sender = SmsSender{}
	sender = &SmsSender{"Jevonwei"}
	do(sender)

	sender = &EmailSender{"Wei@danran.com"}
	do(sender)

	sender = &WechatSender{"Jevonwei"}
	do(sender)

	var ssender SignalSender = sender
	ssender.Send("DDD", "Hello")

	sender01, ok := ssender.(Sender)
	fmt.Printf("%T, %v\n", sender01, ok)
	sender01.SendAll([]string{"Jevon", "Danran"}, "Hello")

	wsender01, ok := ssender.(*WechatSender)
	fmt.Printf("%T, %v\n", wsender01, ok)
	fmt.Println(wsender01.ID)

	esender01, ok := ssender.(*EmailSender)
	fmt.Printf("%T, %v\n", esender01, ok)

	sender = EmailSender{"Testtest"}

	switch v := sender.(type) {
	case EmailSender:
		fmt.Println("emailsender", v.SmtpAddr)
	case *SmsSender:
		fmt.Println("smsender", v.SmsAPI)
	case *WechatSender:
		fmt.Println("*Wechatsender", v.ID)
	}

}
