package main

import "fmt"

func main() {
	var age = 30
	age = age + 1

	fmt.Println("明年： ", age)

	age = age + 1
	fmt.Println("后年： ", age)

	fmt.Println("第一行")

	// Print默认不打印换行
	fmt.Print("第2行")
	fmt.Print("第3行")

	fmt.Printf("\n%T, %s, %T, %d\n", "JevonWei", "JevonWei", 1, 1)

}
