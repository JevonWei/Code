package main

import "fmt"

func main() {
	// defer在函数退出时执行
	// defer先进后出，按照声明的顺序，先声明后执行(堆栈)
	defer func() {
		fmt.Println("defer")
	}()

	defer func() {
		fmt.Println("defer01")
	}()

	fmt.Println("main defer")
}
