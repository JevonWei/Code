package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
	"path/filepath"
	"strings"
	"time"
)

var (
	dirPath         string
	file404         string
	file500         string
	defaultfile     string
	defaultFileName string
)

func parseRequest(conn net.Conn) {

}

func handleResponse(conn net.Conn, path string) {
	fmt.Fprintf(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Server: JevonWei Server 1.10\r\n")

	fmt.Fprintf(conn, "\r\n")

	// 将文件中的数据写入客户端
	// bytes, _ := ioutil.ReadFile(path)
	// conn.Write(bytes)

	file, _ := os.Open(path)
	reader_path := bufio.NewReader(file)

	bytes := make([]byte, 1024)
	for {
		n, err := reader_path.Read(bytes)
		if err != nil {
			if err != io.EOF {
				fmt.Println(err)
			}
			break
		} else {
			conn.Write(bytes[:n])
		}
	}
}

func handle(conn net.Conn) {
	defer func() {
		conn.Close()
		fmt.Println("Client Closed: ", conn.RemoteAddr())
	}()

	fmt.Println("Client Connected: ", conn.RemoteAddr())

	//fmt.Println("sleep start")
	time.Sleep(time.Second * 3)
	//fmt.Println("sleep end")

	// 处理客户端数据
	reader := bufio.NewReader(conn)

	// 读取数据
	line, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)

	} else {
		//fmt.Printf("%q\n", strings.Fields(line))
		// 正常处理
		nodes := strings.Fields(line)
		path := filepath.Join(dirPath, nodes[1])

		if fileinfo, err := os.Stat(path); err != nil {
			if os.IsNotExist(err) {
				path = file404
			} else {
				path = file500
			}
		} else {
			// 目录
			if fileinfo.IsDir() {
				path = defaultfile
			}
			// 文件
		}

		// 再次对path进行检查是否存在
		if _, err := os.Stat(path); err == nil {
			//conn.Write([]byte(fmt.Sprintf("HTTP/1.1 200 OK\r\n")))
			// bufio
			// fmt.Fprint, fmt.Fpinft

			// 调用回复函数发送数据到客户端
			handleResponse(conn, path)

		} else {
			fmt.Fprintf(conn, "HTTP/1.1 404 NotFoound\r\n")
		}

	}

}

func init() {
	binPath, _ := filepath.Abs(os.Args[0])
	dirPath = filepath.Dir(binPath)
	defaultFileName = "index.html"
	file404 = filepath.Join(dirPath, "404.html")
	file500 = filepath.Join(dirPath, "500.html")
	defaultfile = filepath.Join(dirPath, defaultFileName)
}

func main() {

	addr := "0.0.0.0:9999"
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

		go handle(conn)

	}

}
