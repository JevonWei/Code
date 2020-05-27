package main

import (
	"fmt"
	"sync"
)

func main() {
	var users sync.Map

	users.Store(10, "a")
	users.Store(20, "c")

	value, ok := users.Load(10)
	fmt.Println(value.(string), ok)

	users.Delete(10)
	value, ok = users.Load(10)
	fmt.Println(value, ok)
}
