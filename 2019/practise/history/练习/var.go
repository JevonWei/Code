package main

import "fmt"

// 变量只能被声明一次

// 函数外定义的变量可不适用
var Version string = "1.0"

// 函数内定义的变量必须使用
func main() {
	// 定义一个string类型的变量me
	/*
		变量名需要满足标识符命名规则
			1. 必须由非空的unicode字符串组成、数字、_
			2. 不能以数字开头
			3. 不能为go的关键字(25个)
			4. 避免和go预定义标识符冲突，true/false/nil/bool/string
			5. 驼峰命令法
			6. 标识符区分大小写
	*/

	var me1 string
	me1 = "JevonWei"
	fmt.Println("var me1 string 赋值为 ", me1)
	fmt.Println("var Version string = 1.0 函数外定义的变量：", Version)

	// 变量不赋值时打印为空
	//var name, user string
	//fmt.Println("var name, user string未赋值变量为空: ", name, user)

	// 给变量name、user 赋值同类型多个值
	var name, user string = "Jevon", "Wei"
	fmt.Println("var name, user string 赋值为: ", name, user)

	// 定义多个不同类型的值
	var (
		age    int     = 20
		height float64 = 1.78
	)
	fmt.Println("定义多个不同类型的值", age, height)

	// 定义变量时省略变量类型
	var (
		age1    int     = 22
		height1 float64 = 1.80
	)
	fmt.Println("定义多个变量时省略数据类型: ", age1, height1)

	var age2, name2 = 25, "DanRan"
	fmt.Println("定义多个变量时省略数据类型: ", age2, name2)

	// 简短声明，只能在函数内定义
	isBoy := true
	fmt.Println("简短声明只能在函数内使用: ", isBoy)

	// 更新已声明变量的值
	age1, name2, isBoy = 18, "JevonWei", false
	fmt.Println("重新给已声明变量赋值", age1, name2, isBoy)

	// 交换age1，age2的值
	fmt.Println("age1和age2交换之前的值：", age1, age2)
	age1, age2 = age2, age1
	fmt.Println("age1和age2交换之后的值：", age1, age2)
}
