package main

import "math/rand"
import "time"
import "fmt"

func main() {
	rand.Seed(time.Now().Unix())

	// 获取 0 - 100内整数
	fmt.Println(rand.Int() % 100)
	fmt.Println(rand.Intn(100))
}
