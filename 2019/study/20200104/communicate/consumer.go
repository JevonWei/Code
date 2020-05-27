package main

import (
	"fmt"
	// "time"

	"github.com/gomodule/redigo/redis"
)

// 将数据从redis读取出来
func main() {
	conn, _ := redis.Dial("tcp", "localhost:6379")
	conn.Do("AUTH", "XnSRZj4B3XvMMJQX8PMzvBLLG5HsE5Ym")

	for {
		// 3s 秒钟从redis中读取一次
		// now, err := redis.String(conn.Do("RPOP", "test:logs"))
		// if err != nil {
		// 	fmt.Println(err)
		// 	time.Sleep(3 * time.Second)
		// } else {
		// 	fmt.Println(values[1])
		// }

		values, err := redis.Strings(conn.Do("BRPOP", "test:logs", 3))
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(values[1])
		}
	}
}
