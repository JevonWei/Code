package main

import "fmt"

func changeInt(a int) {
	a = 100
}

func changeSlice(s []int) {
	s[1] = 100
}

func changeIntByoint(p *int) {
	*p = 100
}

func main() {
	num := 1
	changeInt(num)
	fmt.Println(num)

	nums := []int{1, 2, 3}
	changeSlice(nums)
	fmt.Println(nums)

	changeIntByoint(&num)
	fmt.Println(num)

}
