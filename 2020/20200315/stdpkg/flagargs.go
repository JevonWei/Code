package main

import (
	"flag"
	"fmt"
)

func main() {
	var port int
	var host string
	var help bool
	flag.IntVar(&port, "P", 22, "ssh port")
	flag.StringVar(&host, "H", "127.0.0.1", "ssh host")
	flag.BoolVar(&help, "h", false, "help")

	verbor := flag.Bool("v", false, "detail log")

	flag.Parse()

	flag.Usage = func() {
		fmt.Println("usage flagargs [-H 127.0.0.1] [-P 22] [-V] [args01 args02]")
		flag.PrintDefaults()
	}

	if help {
		flag.Usage()
	} else {
		fmt.Println(port)
		fmt.Println(host)
		fmt.Println(flag.Args())
	}

}
