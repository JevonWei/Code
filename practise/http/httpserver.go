package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"net"
	"os"
	"path/filepath"
	"strings"
)

var (
	dirPath     string
	defauleFile string
	file404     string
	file500     string
)

func handle(conn net.Conn) {
	defer func() {
		conn.Close()
		fmt.Println("Client Closed: ", conn.RemoteAddr())
	}()

	fmt.Println("Client Connerted: ", conn.RemoteAddr())
	// 处理客户端数据
	reader := bufio.NewReader(conn)
	// 读取数据
	line, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
	} else {
		nodes := strings.Fields(line)

		// 获取请求资源(本地)路径
		path := filepath.Join(dirPath, nodes[1])
		if info, err := os.Stat(path); err != nil {
			if os.IsNotExist(err) {
				path = file404
			} else {
				path = file500
			}
		} else {
			// 目录
			if info.IsDir() {
				path = defauleFile
			}
			// 文件
		}

		// 再次对path进行检查(404.html/500.html/index.html)
		if _, err := os.Stat(path); err == nil {
			// conn.Write([]byte(fmt.Sprintf("HTTP/1.1 200 OK\r\n")))

			fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
			fmt.Fprint(conn, "Server: JevonWei Server 1.0\r\n")
			fmt.Fprint(conn, "\r\n")

			bytes, _ := ioutil.ReadFile(path)
			conn.Write(bytes)

		} else {
			fmt.Fprint(conn, "HTTP/1.1 404 NotFound\r\n")

		}
	}
}

func init() {
	binPath, err := filepath.Abs(os.Args[0])
	if err != nil {
		fmt.Println(err)
	}
	dirPath = filepath.Dir(binPath)
	file404 = filepath.Join(dirPath, "404.html")
	file500 = filepath.Join(dirPath, "500.html")
	defauleFile = filepath.Join(dirPath, "index.html")
}

func main() {
	addr := ":9999"
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	defer listener.Close()

	fmt.Println("Listen on: ", listener.Addr())

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}

		// func() {
		// 	defer func() {
		// 		conn.Close()
		// 		fmt.Println("Client Closed: ", conn.RemoteAddr())
		// 	}()
		// }()

		go handle(conn)
	}

}
