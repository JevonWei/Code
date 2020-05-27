package main

import (
	"fmt"
	"runtime"
)

func PrintChars(name int, notice chan int) {
	for ch := 'A'; ch <= 'Z'; ch++ {
		fmt.Printf("%d: %c\n", name, ch)

		runtime.Gosched()
	}
	notice <- name
}

func main() {
	var notice chan int = make(chan int)

	for i := 0; i < 10; i++ {
		go PrintChars(i, notice)
	}

	for i := 0; i < 10; i++ {
		<-notice
	}
	fmt.Println("over")
	// go PrintChars("2", notice)

	// go PrintChars("3", notice)
	// <-notice
	// <-notice
	// <-notice
	// 	fmt.Println(<-notice)
	// 	fmt.Println(<-notice)
	// 	fmt.Println(<-notice)
}
