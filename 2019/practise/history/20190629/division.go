package main

import (
	"errors"
	"fmt"
)

// 返回值 怎么定义错误类型  error
// 怎么创建错误类型的值， errors.New(), fmt.Errorf()
// 无错误，返回nil

// errors.New()创建错误类型
func division(a, b int) (int, error) {
	if b == 0 {
		return -1, errors.New("division by zero")
	}
	return a / b, nil
}

// fmt.Errorf()创建错误类型
func division1(a, b int) (int, error) {
	if b == 0 {
		return -1, fmt.Errorf("Error: %s", "division by zero")
	}
	return a / b, nil
}

func main() {
	// fmt.Println(division(6, 3))
	if v, err := division(6, 3); err == nil {
		fmt.Println(v)
	} else {
		fmt.Println(err)
	}

	if v, err := division1(6, 0); err == nil {
		fmt.Println(v)
	} else {
		fmt.Println(err)
	}

	//e := fmt.Errorf("Error: %s", "division by zero")
	//fmt.Printf("%T, %v", e, e)

}
