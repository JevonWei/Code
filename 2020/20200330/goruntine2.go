package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	group := &sync.WaitGroup{}

	n := 2
	group.Add(n)

	for i := 1; i <= n; i++ {
		go func(id int) {
			for ch := 'A'; ch <= 'Z'; ch++ {
				fmt.Printf("%d: %c\n", id, ch)
				runtime.Gosched()
			}
			group.Done()
		}(i)
	}

	group.Wait()
}
