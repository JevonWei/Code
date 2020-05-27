package main

import "fmt"

func main() {
	// const常量一定要设置默认值

	// 定义常量NAME
	const NAME string = "JevonWei"
	fmt.Println("常量NAME: ", NAME)

	// 省略类型
	const PI = 3.1415
	fmt.Println("省略类型常量PI: ", PI)

	// 定义多个常量(类型相同)
	const C1, C2 int = 1, 2
	fmt.Println("定义多个类型相同的常量: ", C1, C2)

	// 定义多个常量(类型不相同)
	const (
		C3 string = "Jevon"
		C4 int    = 3
	)
	fmt.Println("定义多个不同类型的常量: ", C3, C4)

	// 定义多个常量(省略类型)
	const (
		C5 = "DAN"
		C6 = 6
	)
	fmt.Println("定义多个不同类型的常量C5,C6: ", C5, C6)

	const C7, C8 = "RAN", 8
	fmt.Println("定义多个类型相同的常量C7,C8: ", C7, C8)

	// 后面的常量若跟前面定义的常量值相同，则可省略赋值(用来实现枚举类型)
	const (
		C10 int = 1
		C11
		C12
		C13 float64 = 3.14
		C14
		C15
		C16 string = "JevonWei"
		C17

	)
	fmt.Println("C10,C11,C12三个常量的值相同:", C10, C11, C12)
	fmt.Println("C13,C14,C15三个常量的值相同:", C13, C14, C15)
	fmt.Println("C16,C17两个常量的值相同:", C16, C17)

	// 实现枚举类型，const + iota
	const (
		E1 int = iota
		E2
		E3
	)

	fmt.Println("const + iota 枚举类型: ",E1, E2, E3)
	const (
		E4 int = (iota + 1) * 3
		E5
		E6
	)
	fmt.Println("const + iota 枚举类型((iota + 1) * 3): ",E4, E5, E6)

}
