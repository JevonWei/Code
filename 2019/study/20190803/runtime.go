package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	fmt.Println(runtime.GOROOT())
	fmt.Println(runtime.NumCPU())      // 获取CPU 核数
	fmt.Println(runtime.GOMAXPROCS(1)) // 获取使用CPU数量为1

	go func() {
		time.Sleep(time.Second * 3)
		runtime.Gosched()
	}()

	fmt.Println(runtime.NumGoroutine()) // 获取例程的数量
}
