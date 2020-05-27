package main

import "log"

func main() {
	log.Printf("Printf日志: %s", "x") // 打印日志

	log.SetPrefix("prefix:")                   // 设置日志前缀
	log.SetFlags(log.Flags() | log.Lshortfile) //短文件
	log.SetFlags(log.Flags() | log.Llongfile)  //长文件,与短文件仅有一个生效
	log.Printf("Printf日志: %s", "x")

	//log.Panicf("Panic日志：%s", "y")	// 打印日志，并触发panic
	log.Fatalf("Fatalf日志: %s", "Z") // 打印日志，打印第一个日志之后退出
	log.Fatalf("Fatalf日志: %s", "Z")
}
