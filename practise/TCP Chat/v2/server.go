package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	addr := ":9999"
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	defer listener.Close()
	fmt.Println("listen:", addr)

	input := bufio.NewScanner(os.Stdin)

	for {
		client, err := listener.Accept()
		if err == nil {
			reader := bufio.NewReader(client)
			writer := bufio.NewWriter(client)

			fmt.Printf("客户端%s连接成功\n", client.RemoteAddr())

			for {
				// _, err := writer.WriteString(time.Now().Format("2006-01-02 15:04:05") + "\n")
				fmt.Print("请输入(q退出):")
				input.Scan()
				if input.Text() == "q" {
					break
				}
				_, err := writer.WriteString(input.Text() + "\n")
				writer.Flush()
				if err != nil {
					break
				}
				line, err := reader.ReadString('\n')
				if err != nil {
					break
				}
				fmt.Printf("客户端(%s): %s\n", client.RemoteAddr(), strings.TrimSpace(line))
			}
		} else {
			fmt.Println(err)
		}

		client.Close()
		fmt.Printf("客户端%s关闭", client.RemoteAddr())
	}

}
