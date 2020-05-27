package main

import "fmt"

func main() {
	// 多维数组
	// 长度为2的int类型数组 => [2]int

	var marrays [3][2]int
	fmt.Println(marrays)
	fmt.Println(marrays[0])
	marrays[0] = [2]int{1, 3}
	fmt.Println(marrays)

	marrays[1][1] = 100
	fmt.Println(marrays)

	var m3 [3][3][5]int
	fmt.Println(m3)

	m3[1][2] = [5]int{1, 3, 4, 2, 3}
	m3[1][1] = [5]int{1, 3, 4, 2}
	fmt.Println(m3)

	m3[1] = [3][5]int{{1, 2, 3, 4, 5}, {11, 11, 11, 11, 11}}
	fmt.Println(m3)
}
