package main

import (
	"time"

	"github.com/gomodule/redigo/redis"
)

// 将数据存入redis
func main() {
	conn, _ := redis.Dial("tcp", "localhost:6379")

	defer conn.Close()
	conn.Do("auth", "XnSRZj4B3XvMMJQX8PMzvBLLG5HsE5Ym")

	for {
		conn.Do("LPUSH", "test:logs", time.Now())
		time.Sleep(time.Second)
	}
}
