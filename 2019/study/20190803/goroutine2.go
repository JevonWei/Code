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
				fmt.Printf("%d:%d: %c\n", id, i, ch)
				runtime.Gosched() // 调度CPU资源
			}
			group.Done()
		}(i)
	}

	group.Wait()
}
