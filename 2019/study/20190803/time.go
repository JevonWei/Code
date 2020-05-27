package main

import (
	"fmt"
	"time"
)

func main() {

	fmt.Println(time.Now())

	fmt.Println("After")
	channel := time.After(3 * time.Second)
	fmt.Println(<-channel) // After只读一次    延时执行一次

	fmt.Println("Tick")
	ticker := time.Tick(3 * time.Second) // Tick可以循环读取
	fmt.Println(<-ticker)                // 每个ns产生一个管道信息  // 每隔ns执行动作
	fmt.Println(<-ticker)
	fmt.Println(<-ticker)

	for now := range ticker { // 循环读取ticker元素
		fmt.Println(now)
	}
}
