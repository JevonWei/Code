package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strings"
	"time"
)

// 定义客户端结构体
type Client struct {
	Name   string
	Conn   net.Conn
	Input  *bufio.Reader
	Output *bufio.Writer
}

// 定义消息结构体
type Message struct {
	From *Client
	Msg  string
	Time time.Time
}

// 为消息结构体定义String方法
func (m Message) String() string {
	return fmt.Sprintf("[%s](%s):%s", m.Time.Format("2006-01-02 15:04:05"), m.From.Name, m.Msg)
}

func init() {
	logfile := "server.log"
	file, err := os.OpenFile(logfile, os.O_APPEND|os.O_CREATE, os.ModePerm)

	if err == nil {
		log.SetOutput(file)
		log.SetPrefix("server:")
		log.SetFlags(log.Flags() | log.Lshortfile)
	}
}

func main() {
	// 定义命令行参数
	host := flag.String("host", "0.0.0.0", "host")
	port := flag.Int("port", 8080, "port")
	help := flag.Bool("help", false, "help")
	h := flag.Bool("h", false, "help")

	flag.Usage = func() {
		fmt.Println("Usage: server [--host 0.0.0.0 ] [--port=8080]")
		flag.PrintDefaults()
	}
	flag.Parse()

	if *h || *help {
		flag.Usage()
		os.Exit(0)
	}

	addr := fmt.Sprintf("%s:%d", *host, *port)
	// 监听服务
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalln("server error: ", err)
	}

	defer func() {
		listener.Close()
		log.Println("Server is Closed")
	}()

	fmt.Println("Listen:", addr)
	log.Printf("Server is Listen on: [%s]\n", addr)

	// 定义客户端容器，用于保存当前所有用户的客户端信息
	clients := make(map[*Client]bool)

	// 定义广播管道
	boardcast := make(chan Message, 10)
	timeout, _ := time.ParseDuration("1s")

	index := 1

	// 使用goroutine定义将当前连接数量输出到日志
	go func() {
		// 定义定时器
		ticker := time.Tick(60 * time.Second)
		for _ = range ticker {
			log.Println("online client: ", len(clients))
		}
	}()

	// 使用goroutine处理消息广播
	go func() {
		for {
			// 从管道读取数据，发送给所有客户端
			msg := <-boardcast
			for client := range clients {
				// 设置写入消息到客户端的超时时间
				client.Conn.SetWriteDeadline(time.Now().Add(timeout))
				// 写入消息到客户端
				client.Output.WriteString(msg.String())
				client.Output.Flush()
			}
		}
	}()

	// 循环监听客户端连接
	for {
		// 接收客户端连接
		conn, err := listener.Accept()
		if err != nil {
			log.Println("client connection error:", err)
			continue
		}
		client := &Client{
			Name:   fmt.Sprintf("Jevon%3d", index),
			Conn:   conn,
			Input:  bufio.NewReader(conn),
			Output: bufio.NewWriter(conn),
		}
		index++

		clients[client] = true

		// 处理客户端消息
		go func(client *Client) {
			log.Println("Client is Connected: ", client.Conn.RemoteAddr())

			// 延时处理客户端离线
			defer func() {
				delete(clients, client)
				boardcast <- Message{From: client, Msg: "我要走了\n", Time: time.Now()}

				client.Conn.Close()
				log.Println("Client is Closed: ", client.Conn.RemoteAddr())
			}()

			// 获取客户端消息(名称)
			line, err := client.Input.ReadString('\n')
			if err == io.EOF {
				return
			}
			if name := strings.TrimSpace(line); name != "" {
				client.Name = name
			}

			// 将上线消息发送到广播管道
			boardcast <- Message{From: client, Msg: "I Come In\n", Time: time.Now()}

			// 循环处理客户端发送的消息
			for {
				// 设置读取客户端消息的超时时间
				client.Conn.SetWriteDeadline(time.Now().Add(timeout))
				// 读取客户端信息
				line, err := client.Input.ReadString('\n')
				if err == nil {
					boardcast <- Message{From: client, Msg: line, Time: time.Now()}
				} else if err == io.EOF {
					break
				}
			}
		}(client)
	}
}
