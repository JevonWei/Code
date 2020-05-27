package main

import "fmt"

// 整数相加
func AddN(a, b int, args ...int) int {
	sum := a + b
	if len(args) > 0 {
		for _, arg := range args {
			sum += arg
		}
	}

	return sum
}

func calc(op string, a, b int, args ...int) int {
	switch op {
	case "add":
		return AddN(a, b, args...)
	}
	return -1
}

func main() {
	fmt.Println(AddN(1, 3, 3))
	fmt.Println(calc("add", 1, 3, 4, 5))

	args := []int{1, 2, 3, 4, 5}
	fmt.Println(AddN(1, 2, args...))

	args_slice := []int{1, 2, 3, 4, 5}
	args_slice = append(args_slice[:1], args_slice[2:]...)
	fmt.Println(args_slice)

}
