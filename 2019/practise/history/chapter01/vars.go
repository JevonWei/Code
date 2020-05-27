package main

import "fmt"

var Version string = "1.0"

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

	var me string
	me = "JevonWei"
	fmt.Println(me)

	fmt.Println(Version)

	var name, user string = "Jevon", "Dan"
	fmt.Println(name, user)

	var (
		age    int     = 20
		height float64 = 2.1
	)

	fmt.Println(age, height)

	var (
		s = "DanRan"
		a = 20
	)
	fmt.Println(s, a)

	var ss, aa = "JevonWei", 22
	fmt.Println(ss, aa)

	// 简短声明，只能在函数内部使用
	isBoy := true
	fmt.Println(isBoy)

	ss, aa, isBoy = "Dan", 22, false
	fmt.Println(ss, aa, isBoy)

	fmt.Println(s, ss)
	s, ss = ss, s
	fmt.Println(s, ss)

	var a1 bool = true
	fmt.Println(a1)

}
