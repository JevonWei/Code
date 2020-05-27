package main

import "fmt"

type Sender interface {
	Send(to string, msg string) error
	SendAll(tos []string, msg string) error
}

type EmailSender struct {
}

func (s EmailSender) Send(to string, msg string) error {
	fmt.Println("发送邮件给:", to, "，消息内容是:", msg)
	return nil
}

func (s EmailSender) SendAll(tos []string, msg string) error {
	for _, to := range tos {
		s.Send(to, msg)
	}
	return nil
}

type SmsSender struct {
}

func (s SmsSender) Send(to string, msg string) error {
	fmt.Println("发送短信给:", to, "，消息内容是:", msg)
	return nil
}

func (s SmsSender) SendAll(tos []string, msg string) error {
	for _, to := range tos {
		s.Send(to, msg)
	}
	return nil
}

type WechatSender struct {
	ID string
}

func (s WechatSender) Send(to string, msg string) error {
	fmt.Println("发送微信给:", to, "，消息内容是:", msg)
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

func fargs(arg interface{}) {
	fmt.Println(arg)
}

func main() {
	// email := EmailSender{}
	// var sender Sender = EmailSender{}
	var sender Sender = SmsSender{}

	fmt.Printf("%T, %v\n", sender, sender)

	sender.Send("JevonWei", "Good")
	sender.SendAll([]string{"QQ", "AA"}, "Hello World")

	// email.Send("JevonWei", "Good")

	do(sender)

	sender = &EmailSender{}
	do(sender)

	sender = &WechatSender{"123"}
	do(sender)

	ssender, _ := sender.(*WechatSender)
	fmt.Println(ssender.ID)

	a := WechatSender{"1212"}
	fmt.Println(a.ID)

	fargs("1")

}
