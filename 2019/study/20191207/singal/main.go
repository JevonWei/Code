package main

import (
	"fmt"
	"time"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	go func() {
		for now := range time.Tick(time.Second) {
			fmt.Println(now)
		}
	}()

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	fmt.Println("walt")
	<-ch
	fmt.Println("over")
}
