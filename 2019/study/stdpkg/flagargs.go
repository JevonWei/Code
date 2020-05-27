package main

import (
	"flag"
	"fmt"
)

func main() {
	var port int
	var host string
	var verbor bool
	var help bool

	// 绑定命令行参数与变量关系
	flag.IntVar(&port, "P", 22, "ssh port")
	flag.StringVar(&host, "H", "10.127.0.1", "ssh host")
	flag.BoolVar(&verbor, "v", false, "detail log")
	flag.BoolVar(&help, "h", false, "help")

	flag.Usage = func() {
		fmt.Println("usage: flagargs [-H 127.0.0.1] [-P 22] [-v] [args01 args02]")
		flag.PrintDefaults()
	}

	// 解析命令行参数
	flag.Parse()

	if help {
		flag.Usage()
	} else {
		fmt.Printf("%s:%d %v\n", host, port, verbor)
		fmt.Println(flag.Args())
	}

}
