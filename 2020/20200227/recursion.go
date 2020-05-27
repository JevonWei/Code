package main

import "fmt"

func addN(n int) int {
	if n == 1 {
		return 1
	}
	return n + addN(n-1)
}

func main() {
	fmt.Println(addN(5))
}
