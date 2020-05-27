package main

import "fmt"

func main() {

	add2 := func(n int) int {
		return n + 2
	}
	fmt.Println(add2(4))

	addBase := func(base int) func(int) int {
		// 返回函数
		return func(n int) int {
			return base + n
		}
	}

	add8 := addBase(8)
	fmt.Printf("%T\n", add8)
	fmt.Println(add8(10))

	add10 := addBase(5)(5)
	fmt.Println(add)
}
