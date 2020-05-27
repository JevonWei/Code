package main

import "fmt"

func test() (err error) {
	defer func() {
		if e := recover(); e != nil {
			err = fmt.Errorf("%v\n", e)
		}
	}()

	panic("error")
	return
}

func main() {
	// recover()只能用于defer函数中
	//defer func() {
	//	if err := recover(); err != nil {
	//		fmt.Println(err)
	//	}
	//}()

	//fmt.Println("main start")
	//panic("panic error")

	//fmt.Println("over")

	err := test()
	fmt.Println(err)
}
