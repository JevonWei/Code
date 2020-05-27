package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	var counter int
	group := &sync.WaitGroup{}
	lock := &sync.Mutex{}

	incr := func() {
		defer group.Done()
		for i := 0; i <= 100; i++ {
			lock.Lock()
			counter++
			lock.Unlock()
			runtime.Gosched()
		}
	}

	decr := func() {
		defer group.Done()
		for i := 0; i <= 100; i++ {
			lock.Lock()
			counter--
			lock.Unlock()
			runtime.Gosched()
		}
	}

	for i := 0; i < 10; i++ {
		group.Add(2)
		go incr()
		go decr()
	}

	group.Wait()
	fmt.Println(counter)
}
