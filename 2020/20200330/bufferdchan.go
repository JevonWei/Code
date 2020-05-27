package main

import "fmt"

func main() {
	channel := make(chan string, 2)

	channel <- "x"
	channel <- "y"

	fmt.Println(<-channel)
	fmt.Println(<-channel)

	channel <- "z"
	channel <- "y"
	close(channel)

	for ch := range channel {
		fmt.Println(ch)
	}
}
