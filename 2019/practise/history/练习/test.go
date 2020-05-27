package main

import "fmt"

// 当未发生panic，则recover函数得到的结果为nil
func success() {
	defer func() {
		fmt.Println(recover())
	}()
	fmt.Println("success")
}

// 当发生panic，则recover函数得到的结果为panic传递的参数
func failure() {
	defer func() {
		fmt.Println(recover())
	}()
	fmt.Println("failure")
	panic("error")
}

// recover只能获取最后一次的panic信息
func failure2() {
	defer func() {
		fmt.Println(recover())
	}()
	defer func() {
		fmt.Println("failure02")
		panic("error 02")
	}()
	fmt.Println("failure")
	panic("error")
}

func main() {
	a := "1"
	b := "1"
	if a == b {
		fmt.Println("True")
	}

	printArgs(1, 2, []string{3, 4, 5, 5, 6}...)
	args := []string{3, 4, 5, 6, 7, 7, 8}
	printArgs(1, 3, args)

}
