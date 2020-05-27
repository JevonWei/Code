package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func Client(conn net.Conn) {
	defer conn.Close()
	reader := bufio.NewReader(conn)
	writer := bufio.NewWriter(conn)

	input := bufio.NewScanner(os.Stdin)

	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		//fmt.Print("服务器: ", strings.TrimSpace(line))
		fmt.Print("服务器: ", line)
		fmt.Println("请输入(q 退出)...")

		input.Scan()
		if input.Text() == "q" {
			break
		}

		for {
			if input.Text() == "" {
				fmt.Println("输入内容为空，请重新输入....")
				input.Scan()
				input.Text()
			} else {
				break
			}
		}
		_, err = writer.WriteString(strings.TrimSpace(input.Text()) + "\n")
		if err != nil {
			break
		}
		writer.Flush()

	}

}

func main() {
	addr := "127.0.0.1:9999"
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	Client(conn)
}
