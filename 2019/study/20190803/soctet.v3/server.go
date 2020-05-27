package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func Server(listener net.Listener) {
	defer listener.Close()
	input := bufio.NewScanner(os.Stdin)

	for {
		client, err := listener.Accept()
		if err != nil {
			fmt.Printf("client %s connect faild", client.RemoteAddr())
		} else {
			reader := bufio.NewReader(client)
			writer := bufio.NewWriter(client)

			fmt.Printf("客户端%s连接成功\n", client.RemoteAddr())

			for {
				fmt.Println("请输入(q 退出)...")
				input.Scan()
				text := strings.TrimSpace(input.Text())
				if text == "q" {
					break
				}

				for {
					if text == "" {
						fmt.Println("输入内容为空，请重新输入....")
						input.Scan()
						text = strings.TrimSpace(input.Text())
					} else {
						break
					}
				}

				_, err := writer.WriteString(strings.TrimSpace(text) + "\n")
				writer.Flush()
				if err != nil {
					break
				}

				//fmt.Println(n, err)
				line, err := reader.ReadString('\n')
				if err != nil {
					fmt.Printf("%s 客户端读取失败", client.RemoteAddr())
					break
				}
				fmt.Println("客户端:", strings.TrimSpace(line))
			}
			fmt.Printf("客户端%s关闭\n", client.RemoteAddr())

			client.Close()
		}

	}
}

func main() {
	addr := "0.0.0.0:9999"
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	fmt.Printf("Server监听地址: %s\n", addr)

	// 与客户端发送/接收信息
	Server(listener)

}
