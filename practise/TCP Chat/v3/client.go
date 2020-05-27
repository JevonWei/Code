package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

func Init() {
	logfile := "client.log"
	file, err := os.OpenFile(logfile, os.O_APPEND|os.O_CREATE, os.ModePerm)

	if err == nil {
		log.SetOutput(file)
		log.SetPrefix("client:")
		log.SetFlags(log.Flags() | log.Lshortfile)

	}
}

func main() {
	// 定义命令行参数
	host := flag.String("host", "127.0.0.1", "server host")
	port := flag.Int("port", 8080, "listen port")
	help := flag.Bool("help", false, "help")
	h := flag.Bool("h", false, "help")

	flag.Usage = func() {
		fmt.Println("Usage: client [--host=127.0.0.1] [--port=8080]")
		flag.PrintDefaults()
	}

	flag.Parse()

	if *help || *h {
		flag.Usage()
		os.Exit(0)
	}

	addr := fmt.Sprintf("%s:%d", *host, *port)
	client, err := net.Dial("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	// 获取带缓冲的服务器流读、写对象
	reader := bufio.NewReader(client)
	writer := bufio.NewWriter(client)

	// 获取带缓冲的命令行读对象
	input := bufio.NewReader(os.Stdin)

	// 提示用户输入用户名并发送给服务器端
	fmt.Print("请输入用户名:")
	line, _ := input.ReadString('\n')
	writer.WriteString(line)
	writer.Flush()

	// 使用goroutine处理服务器发送的消息
	go func() {
		for {
			line, err := reader.ReadString('\n')
			if err == nil {
				fmt.Println("\r*", line)
				fmt.Print("请输入信息:")
			} else {
				break
			}
		}
	}()

	// 循环让用户输入消息并发送到服务器端(输入q或Q退出)
	for {
		// 从控制台读取信息
		fmt.Print("请输入信息:")
		line, _ := input.ReadString('\n')
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		if line == "q" || line == "Q" {
			break
		}
		fmt.Println()

		_, err = writer.WriteString(line + "\n")
		writer.Flush()
		if err != nil {
			break
		}
	}

}
