package main

import (
	"fmt"
	"runtime"
	"time"
)

func PrintChars(name string) {
	for ch := 'A'; ch <= 'Z'; ch++ {
		fmt.Printf("%s: %c\n", name, ch)
		runtime.Gosched() // 调度CPU资源
	}
}

func main() {
	go PrintChars("1") // 工作例程
	go PrintChars("2")

	PrintChars("3") // 主例程
	time.Sleep(time.Second * 3)
}
