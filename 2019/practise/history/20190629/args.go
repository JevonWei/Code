package main

import "fmt"

func add(a, b int) int {
	return a + b
}

// 传入可变参数
func addN(a, b int, args ...int) int {
	//fmt.Println(a, b, args)
	//fmt.Printf("%T\n", args)
	total := a + b
	for _, v := range args {
		total += v
	}
	return total
}

func cacl(op string, a, b int, args ...int) int {
	switch op {
	case "add":
		fmt.Println(args)
		return addN(a, b, args...)
	}
	return -1
}

// 删除传入切片中n位置的元素
func del(n int, args ...int) int {
	args = append(args[:n], args[n+1:]...)
	fmt.Println(args)
	return 1
}

func main() {
	fmt.Println(add(1, 2))

	// 调用可变参数函数
	fmt.Println(addN(1, 2))
	fmt.Println(addN(1, 2, 3, 4))
	fmt.Println(addN(1, 2, 4, 6, 7))

	fmt.Println(cacl("add", 1, 2))
	fmt.Println(cacl("add", 1, 2, 3, 4))
	fmt.Println(cacl("add", 1, 2, 4, 6, 7))

	nums := []int{1, 3, 4, 6}
	nums = append(nums[:1], nums[2:]...)
	fmt.Println(nums)

	// 调用删除元素del函数，传入可变参数为1, 2, 3, 4，删除位置为1的元素
	del(1, 1, 2, 3, 4)
}
