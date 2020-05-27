package main

import (
	"fmt"
	"time"

	"github.com/gomodule/redigo/redis"
)

func main() {
	conn, err := redis.Dial("tcp", "localhost:6379")
	fmt.Println(conn, err)
	defer conn.Close()
	reply, err := conn.Do("auth", "XnSRZj4B3XvMMJQX8PMzvBLLG5HsE5Ym")
	fmt.Println(reply, err)

	// 字符串
	conn.Do("SET", "name2", "Jevon")
	svalue, err := redis.String(conn.Do("get", "name2"))
	fmt.Println(svalue, err)

	// 设置过期时间为30s
	conn.Do("EXPIRE", "name2", 30)

	time.Sleep(3 * time.Second)
	ivalue, err := redis.Int(conn.Do("TTL", "name2"))
	fmt.Println(ivalue, err)

	// 列表
	conn.Do("LPUSH", "testlist", 1, 2, 3, 4, 5)
	isvalue, err := redis.Ints(conn.Do("LRANGE", "testlist", 0, -1))
	fmt.Printf("%#v, %#v\n", isvalue, err)

	ssvalue, err := redis.Strings(conn.Do("LRANGE", "testlist", 0, -1))
	fmt.Printf("%#v, %#v\n", ssvalue, err)

	ivalue, err = redis.Int(conn.Do("LPOP", "testlist"))
	fmt.Printf("%#v, %#v\n", ivalue, err)

	ssvalue, err = redis.Strings(conn.Do("BLPOP", "testlist", 5))
	fmt.Printf("%#v, %#v\n", ssvalue, err)

	imvalue, err := redis.IntMap(conn.Do("BLPOP", "testlist", 5))
	fmt.Printf("%#v, %#v\n", imvalue, err)

	// 哈希
	conn.Do("HMSET", "user", "name", "jevon", "age", 32, "addr", "shanghai")
	smvalue, err := redis.StringMap(conn.Do("HGETALL", "user"))
	fmt.Printf("%#v, %#v\n", smvalue, err)

	// 发布
	for i := 0; i < 10; i++ {
		fmt.Println(conn.Do("PUBLISH", "testchannel", i))
		time.Sleep(time.Second)
	}

	//订阅
	pubSubConn := redis.PubSubConn{Conn: conn}
	pubSubConn.Subscribe("testchannel")
	for {
		switch v := pubSubConn.Receive().(type) {
		case redis.Subscription:
			fmt.Println("订阅成功:", v.Kind, v.Channel, v.Count)
		case redis.Message:
			fmt.Println(string(v.Data))
		}
		// fmt.Printf("%#v\n", v)
	}

}
