package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/gomodule/redigo/redis"
)

func main() {
	node := flag.String("node", "A", "node name")
	flag.Parse()

	masterKey := "test:worker:master"
	conn, err := redis.Dial("tcp", "localhost:6379")
	if err != nil {
		fmt.Println(err)
	}

	defer conn.Close()

	// 认证
	//reply, err := conn.Do("auth", "XnSRZj4B3XvMMJQX8PMzvBLLG5HsE5Ym")
	//fmt.Println(reply, err)

	conn.Do("auth", "XnSRZj4B3XvMMJQX8PMzvBLLG5HsE5Ym")

	for now := range time.Tick(10 * time.Second) {
		if value, err := redis.String(conn.Do("SET", masterKey, *node, "EX", 9, "NX")); err == nil && value == "OK" { // NX为当value值存在时，不在赋值
			fmt.Println("node is master:", *node, now)
		} else {
			value, _ := redis.String(conn.Do("GET", masterKey))
			fmt.Println("master:", value)
		}
	}
}
