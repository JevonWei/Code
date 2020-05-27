package main

import "fmt"

func add(a, b int) int {
	return a + b
}

// 格式化函数
// callback 将传递的数据按照每行打印还是一行按| 分割
func print(callback func(...string), args ...string) {
	fmt.Println("print函数输出")
	callback(args...)
}

func list(args ...string) {
	for i, v := range args {
		fmt.Println(i, ":", v)
	}

	for _, v := range args {
		fmt.Printf("!%v ", v)
	}
	fmt.Println("")
}

func main() {
	fmt.Printf("%T\n", add)

	var f func(int, int) int = add
	fmt.Println(f(1, 2))

	print(list, "a", "b", "c")

	// 匿名函数
	say := func(name string) {
		fmt.Println("Hello ", name)
	}
	say("Jevon")

	func(name1 string) {
		fmt.Println("hi", name1)
	}("Jevon")

	values := func(args ...string) {
		for _, v := range args {
			fmt.Println(v)
		}
	}

	print(values, "a", "b", "c")

	print(func(args ...string) {
		for _, v := range args {
			fmt.Print(v, "\t")
		}
		fmt.Println()
	}, "a0", "bb", "cc")

}
