package main

import "fmt"

func main() {
	channel := make(chan string, 2)

	fmt.Println(len(channel)) // 获取chan中元素的数量，没读取一次chan中少一个元素，读入一次增加一个元素
	channel <- "x"
	fmt.Println(len(channel))
	channel <- "y"
	fmt.Println(len(channel))

	fmt.Println(<-channel)
	fmt.Println(len(channel))
	fmt.Println(<-channel)
	fmt.Println(len(channel))

	channel <- "z"
	channel <- "a"
	close(channel) // chan关闭之后，仅能读取chan，不能写入chan

	//fmt.Println(<-channel)
	//channel <- "a"

	for ch := range channel { // 需要有某个例程能够关闭管道，否则会发生死锁
		fmt.Println(ch)
	}
}
