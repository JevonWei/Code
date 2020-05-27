package main

import (
	"fmt"
	"math/rand"
	"time"
)

// 创建一个随机种子并随机生成10个100以内的int类型元素，进行冒泡排序
func main() {
	rand.Seed(time.Now().UnixNano())
	var arr [10]int

	for i := 0; i < 10; i++ {
		f := rand.Intn(100)
		arr[i] = f
	}

	fmt.Println(arr)

	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr)-1; j++ {
			// fmt.Println(i)
			tmp := arr[j]
			if arr[j] < arr[j+1] {
				arr[j] = arr[j+1]
				arr[j+1] = tmp
			}
		}
	}
	fmt.Println(arr)
}
