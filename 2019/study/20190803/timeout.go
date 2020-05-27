package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// 一个操作3s未执行完成即任务超时

	rand.Seed(time.Now().Unix())
	result := make(chan int)
	//timeout := make(chan int)

	//任务例程
	go func() {
		interval := time.Duration(rand.Int()%10) * time.Second
		fmt.Println(interval)
		time.Sleep(interval)
		result <- 0
	}()

	// go func() {
	// 	time.Sleep(3 * time.Second)
	// 	timeout <- 0
	// }()

	select {
	case <-result:
		fmt.Println("执行完成")
	case <-time.After(3 * time.Second):
		fmt.Println("执行超时")
		// case <-timeout:
		// 	fmt.Println("执行超时")
	}
}
