package main

import (
	"crypto/rand"
	"fmt"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())
	// 一个操作3s未执行完成就任务超时
	result := make(chan int)
	timeout := make(chan int)

	go func() {
		interval := time.Duration(rand.Int()%10) * time.Second
		time.Sleep(interval)
		result <- 0
	}()

	// go func() {
	// 	time.Sleep(3 * time.Second)
	// 	result <- 0
	// }()

	select {
	case <-result:
		fmt.Println("执行完成")
	case <-time.After(3 * time.Second):
		fmt.Println("执行超时")
	}

}
