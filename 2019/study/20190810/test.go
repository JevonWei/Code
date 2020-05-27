package main

import "fmt"

func copySlice(dest []int, src []int) {
	for i := 0; i < len(dest) && i < len(src); i++ {
		dest[i] = src[i]
	}
}

func main() {
	var slices []int = []int{1, 2, 3, 4}

	dest_slice := make([]int, 3)
	copySlice(dest_slice, slices)

	fmt.Println(dest_slice)
}
