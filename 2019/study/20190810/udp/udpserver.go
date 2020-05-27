package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	addr := "localhost:9999"
	server, err := net.ListenPacket("udp", addr)

	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	for {
		bytes := make([]byte, 1024)
		n, addr, err := server.ReadFrom(bytes)

		go func(addr net.Addr, bytes []byte) {
			fmt.Println(n, addr, err, string(bytes[:n]))
			server.WriteTo([]byte("JevonWei"), addr)
		}(addr, bytes)

	}
}
