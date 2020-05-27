package main

import "fmt"

/*
	定义无参无返回值的函数
*/
func sayHello() {
	fmt.Println("Hello World")
}

/*
定义两个参数且有一个返回值的函数
*/
func add(n1 int, n2 int) int {
	return n1 + n2
}

/*
定义有一个参数无返回值的函数
*/

func sayhi(name string) {
	fmt.Printf("Hi, %s\n", name)
}

/*
	合并相同类型参数类型名
*/
func mergeFuncArgsType(n1, n2 int, s1, s2, s3 string, b1 bool) {
	fmt.Printf("%T, %T, %T, %T, %T, %T\n", n1, n2, s1, s2, s3, b1)
	fmt.Println(n1, n2, s1, s2, s3, b1)
}

/*
	定义可变参数列表函数，至少两个参数,调用函数后，可变参数则被初始化为对应类型的切片
	打印所有参数到控制台
*/
func printArgs(n1, n2 int, args ...string) {
	fmt.Printf("%T, %T, %T\n", n1, n2, args)
	fmt.Println(n1, n2, args)
}

/*
	定义多个返回值函数
	计算两个参数的四则运算并返回
*/

func calc(n1, n2 int) (int, int, int, int) {
	return n1 + n2, n1 - n2, n1 * n2, n1 / n2
}

/*
	定义命名返回值函数
*/
func calcReturnNamecalc(n1, n2 int) (sum, difference, product, quotient int) {
	sum, difference, product, quotient = n1+n2, n1-n2, n1*n2, n1/n2
	return
}

/*
	n 阶乘
	n < 0: 错误
	n = 0: 1
	n >0 : n! = n * n-1
*/
func factorial(n int) int {
	if n < 0 {
		return -1
	} else if n == 0 {
		return 1
	} else {
		return n * factorial(n-1)
	}
}

/*
	汉罗塔游戏
	将所有a柱上的圆盘借助b柱移动到c柱， 在移动过程中保证每个柱子的上面圆盘比下面圆盘小
	n: a -> c(b)
	n=1: a -> c
	n>1: n-1 (a -> b(c));a -> c; n-1(b -> c(a))
*/
func tower(a, b, c string, layer int) {
	if layer <= 0 {
		return
	}
	if layer == 1 {
		fmt.Printf("%s -> %s\n", a, c)
		return
	}
	tower(a, c, b, layer-1)
	fmt.Printf("%s -> %s\n", a, c)
	tower(b, a, c, layer-1)
}

/*
	定义接收函数类型作为参数的函数
*/
func printResult(pf func(...string), list ...string) {
	pf(list...)
}

func line(list ...string) {
	fmt.Print("!")
	for _, e := range list {
		fmt.Print(e)
		fmt.Printf("\t!")
	}
	fmt.Println()
}

func column(list ...string) {
	for _, e := range list {
		fmt.Println(e)
	}
	fmt.Println("")
}

/*
	定义闭包函数，返回一个匿名函数用于计算与base元素的和
*/
func addBase(base int) func(int) int {
	return func(num int) int {
		return base + num
	}
}

func main() {
	sayHello()

	num := add(2, 4)
	println(num)

	sayhi("Jevon")

	mergeFuncArgsType(1, 2, "A", "B", "C", true)

	printArgs(1, 12, "aa", "cc", "FF")

	// 通过切片解包并调用可变参数函数
	printArgs(1, 22, []string{"a", "b", "c", "d"}...)
	args := []string{"a", "b", "c", "d", "f"}
	printArgs(1, 2, args...)
	printArgs(1, 2, args[:3]...)

	c1, c2, c3, c4 := calc(4, 2)
	fmt.Println(c1, c2, c3, c4)

	fmt.Println(calcReturnNamecalc(6, 3))

	fmt.Println(factorial(4))

	tower("a", "b", "c", 4)

	// 定义函数类型变量， 并使用零值nil进行初始化
	var callback func(n1, n2 int) (r1, r2, r3, r4 int)
	fmt.Printf("%T, %v\n", callback, callback)

	// 赋值为函数calc
	callback = calc
	// 调用calc函数
	fmt.Println(callback(5, 2))

	// 赋值为函数calcReturnNamecalc
	callback = calcReturnNamecalc
	fmt.Println(calcReturnNamecalc(8, 2))

	// 调用参数类型为函数的函数
	names := []string{"wei", "Jevon", "Dan"}
	printResult(line, names...)
	printResult(column, names...)

	// 定义匿名函数并赋值给hi
	hi := func(name string) {
		fmt.Printf("Hi, %s\n", name)
	}
	// 调用匿名函数hi
	hi("Jevon")

	// 定义匿名函数并进行调用
	func() {
		fmt.Println("我是匿名函数")
	}()

	// 使用匿名函数作为printResult的参数
	printResult(func(list ...string) {
		for i, v := range list {
			fmt.Printf("%d: %s\n", i, v)
		}
	}, names...)

	// 使用闭包函数
	base2 := addBase(2)
	base10 := addBase(10)

	fmt.Println(base2(1), base10(1))
	fmt.Println(base2(5), base10(5))
	fmt.Println(base2(10), base10(10))

	/*
		判断数据类型是引用类型还是值类型(映射，切片，接口为引用类型)
	*/
	// 定义字符串，数值，布尔类型
	name, age, height, isBoy := "Jevon", 20, 1.68, false
	// 定义指针类型
	pointer := new(int)
	// 定义数组类型
	scores := [...]int{1, 2, 3}
	// 定义切片类型
	names1 := make([]string, 1, 3)
	// 定义映射类型
	user := make(map[int]string)

	name2, age2, height2, isBoy2, pointer2, scores2, names2, user2 := name, age, height, isBoy, pointer, scores, names1, user

	name2 = "Dan"
	age2 = 10
	height2 = 2.0
	isBoy2 = true
	scores2[0] = 100
	pointer2 = &age
	names2[0] = "Ran"
	user2[1] = "ran"

	fmt.Println(name, age, height, isBoy, pointer, scores, names1, user)
	fmt.Println(name2, age2, height2, isBoy2, pointer2, scores2, names2, user2)

	/*
		函数：值传递
	*/
	e1, e2 := "Jevon", []string{"JevonWei", "Dan"}

	// 值传递
	// 在函数内修改值类型
	fmt.Printf("e1: %p %v\n", &e1, e1)
	func(e string) {
		fmt.Printf("e: %p %v\n", &e, e)
		e = "dan"
	}(e1)

	// 在函数内修改引用类型
	fmt.Printf("e2: %p %v\n", &e2, e2)
	func(e []string) {
		fmt.Printf("e: %p %v\n", &e, e)
		e[1] = "DanRan"
	}(e2)

	fmt.Println(e1, e2)

}
