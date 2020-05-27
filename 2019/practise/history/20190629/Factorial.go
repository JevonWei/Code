package main

import "fmt"

func factorial1(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial1(n-1)
}
func main() {
	fmt.Println(factorial1(5))
}
