package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println(time.Now())

	fmt.Println("after")
	channel := time.After(3 * time.Second)
	fmt.Println(<-channel)

	fmt.Println("ticker")
	ticker := time.Tick(3 * time.Second)
	fmt.Println(<-ticker)
	fmt.Println(<-ticker)
	fmt.Println(<-ticker)

	for now := range ticker {
		fmt.Println(now)
	}
}
