package main

import "fmt"

// 定义无参、无返回值函数
func sayHello() {
	fmt.Println("Hello World!!")
}

// 定义有参数，无返回值的函数
func sayhi(name string) {
	fmt.Println("你好:", name)
}

// 定义有返回值函数
func add(a int, b int) int {
	return a + b
}

func main() {
	// 打印标识符 sapHell类型
	fmt.Printf("%T\n", sayHello)
	sayHello()

	// 调用有参，无返回值函数
	sayhi("Jevon")

	// 定义返回值函数add
	rt := add(1, 2)
	fmt.Println(rt)
}
