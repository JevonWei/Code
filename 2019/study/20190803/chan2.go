package main

import (
	"fmt"
	"runtime"
)

func PrintChars(name int, channel chan int) {
	for ch := 'A'; ch <= 'Z'; ch++ {
		fmt.Printf("%d: %c\n", name, ch)
		runtime.Gosched() // 调度CPU资源
	}
	channel <- name
	//fmt.Println("写入:", name)
}

func main() {
	var channel chan int = make(chan int)

	for i := 0; i <= 10; i++ {
		go PrintChars(i, channel)
	}

	for i := 0; i < 10; i++ {
		//fmt.Println("读取:", <-channel)
		<-channel
	}

	fmt.Println("over")

}
