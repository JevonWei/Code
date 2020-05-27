package main

import "fmt"

func main() {
	// 布尔类型 表示真假
	// 标识符bool
	// 字面量： true/false
	// 零值： false

	// 零值
	var zero bool

	isBoy := true
	isGirl := false

	fmt.Println("零值zero，isBoy，isGirl: ", zero, isBoy, isGirl)

	// 操作
	// 逻辑运算(与 &&， 或 ||， 非 ！)
	// aBool bBool
	// aBool && b Bool 当两个都为True时，结果为True
	// 关系运算 (==, !=)

	// && 运算
	fmt.Println("true && true:", true && true)
	fmt.Println("true && false:", true && false)
	fmt.Println("false && false:", false && false)

	// || 运算
	fmt.Println("true || true:", true || true)
	fmt.Println("true || false:", true || false)
	fmt.Println("false || false:", false || false)

	// 取反
	fmt.Println("!:", "!")
	fmt.Println("!true:", !true)
	fmt.Println("isBoy - !isBoy:", isBoy, !isBoy)

	// 关系运算
	fmt.Println("a == a:", "a" == "a")
	fmt.Println("a != b:", "a" != "b")
	fmt.Println("isBoy == isGirl:", isBoy == isGirl)
	fmt.Println("isBoy != zero:", isBoy != zero)

	// %t 占位符
	fmt.Printf("\n%T, %t, %t", isBoy, isBoy, !isBoy)


}
