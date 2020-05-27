package main

import "fmt"

// 递归函数，需有终止的条件
func addN(n int) int {
	if n == 0 {
		return 1
	}
	fmt.Println("计算f(n): ", n)
	return n + addN(n-1)
}


}

func main() {
	fmt.Println(addN(10))

}
