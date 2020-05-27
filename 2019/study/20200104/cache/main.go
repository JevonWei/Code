package main

import (
	"fmt"
	"time"

	"github.com/gomodule/redigo/redis"
)

func fact(n int) int64 {
	if n == 0 {
		return 1
	}
	return int64(n) * fact(n-1)
}

func main() {
	conn, err := redis.Dial("tcp", "localhost:6379")
	if err != nil {
		return
	}

	defer conn.Close()
	conn.Do("auth", "XnSRZj4B3XvMMJQX8PMzvBLLG5HsE5Ym")

	s := time.Now()
	for n := 0; n < 50; n++ {
		//n := 10
		key := fmt.Sprintf("fact:%d", n)

		// if ok, err := redis.Bool(conn.Do("EXISTS", key)); err == nil && ok {
		// 	value, _ := redis.Int64(conn.Do("GET", key))
		// 	fmt.Println("cache", n, ":", value)
		// } else {
		// 	result := fact(n)
		// 	conn.Do("SET", key, result)
		// 	fmt.Println(n, ":", result)
		// }

		if value, err := redis.Int64(conn.Do("EXISTS", key)); err == nil {
			fmt.Println("cache", n, ":", value)
		} else {
			result := fact(n)
			conn.Do("SET", key, result)
			fmt.Println(n, ":", result)
		}
	}
	fmt.Println(time.Now().Sub(s))

}
