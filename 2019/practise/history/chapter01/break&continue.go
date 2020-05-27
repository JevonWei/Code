package main

import "fmt"

func main() {
	// continue
	fmt.Println("continue:")
	for i := 0; i < 5; i++ {
		if i == 3 {
			continue  // 跳过本次循环
		}
		fmt.Println(i)
	}

	// break
	fmt.Println("break:")
	for i := 0; i < 5; i++ {
		if i == 3 {
			break  // 退出整个循环
		}
		fmt.Println(i)
	}
}
