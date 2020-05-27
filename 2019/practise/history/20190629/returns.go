package main

import "fmt"

// 定义多个返回值
func calc(a, b int) (int, int, int, int) {
	return a + b, a - b, a * b, a / b
}

// 指定返回值变量
func calc2(a, b int) (sum int, diff int, product int, merchant int) {
	sum = a + b
	diff = a - b
	product = a * b
	merchant = a / b
	return
}

func main() {
	q, w, e, _ := calc(1, 2)
	fmt.Println(q, w, e)

	fmt.Println(calc2(6, 3))
}
