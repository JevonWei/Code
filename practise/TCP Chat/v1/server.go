package main

import (
	"fmt"
	"net"
	"os"
	"time"
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

	for {
		client, err := listener.Accept()
		if err == nil {
			client.Write([]byte(time.Now().Format("2006-01-02 15:04:05")))
		} else {
			fmt.Println(err)
		}

		client.Close()
	}

}
