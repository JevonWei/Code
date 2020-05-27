package main

import (
	"fmt"
)

// 冒泡排序

func main() {
	heights := []int{6, 7, 8, 9, 10, 5}

	// 先把最高的人移动到最后面
	for j := 0; j < len(heights)-1; j++ {
		for i := 0; i < len(heights) -1 -j; i++ {
			if heights[i] > heights[i+1] {
				heights[i], heights[i+1] = heights[i+1], heights[i]
			}
			fmt.Println(i, "交换", heights)
		}
		fmt.Println(j, "交换结果", heights)
	}
}