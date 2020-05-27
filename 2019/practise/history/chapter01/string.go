package main

import "fmt"

func main() {
	// ""  = > 可解释字符串
	// `` => 原生字符串
	// 特殊字符 \n \f \t \r \b \v

	var name string = "JevonWei"
	var desc string = `来自\t china`

	fmt.Println("Jevon\\tWei") // \t 转义为 t -> \\t
	fmt.Println(name)
	fmt.Println(desc)

	// 操作
	// 算术运算符： + (连接)
	fmt.Println("名字为: " + name)

	// 关系运算
	fmt.Println("ab == ab:", "ab" == "ab")
	fmt.Println("ab != ab:", "ab" != "ab")
	fmt.Println("ab < ab:", "ab" < "bb")
	fmt.Println("ab >= ab:", "ab" >= "bb")
	fmt.Println("ab <= ab:", "ab" <= "ba")
	fmt.Println("ab < ab:", "ab" < "ab")

	// 赋值
	s := "Name: "
	s += "JevonWei"
	fmt.Println(s)

	// 字符串定义内容必须只能为ascii
	//索引 0 - n-1 (n 字符串长度)
	desc = "abcdefg"
	fmt.Printf("%T %c\n", desc[0], desc[0])

	// 切片[start:end]  start end -1
	fmt.Printf("%T %s\n", desc[0:2], desc[0:2])
	fmt.Printf("%T %s\n", desc[len(desc)-2:len(desc)-1], desc[len(desc)-2:len(desc)-1])
	fmt.Println(len(desc))

}
