package main

import (
	"fmt"
	"sync"
)

func main() {
	var users sync.Map

	users.Store(10, "Jevon")
	users.Store(20, "Dan")

	if value, ok := users.Load(10); ok {
		//fmt.Println(value.(string))
		fmt.Println(value)
	}

	if value, ok := users.Load(30); ok {
		fmt.Println(value)
	}

	users.Delete(10)

	if value, ok := users.Load(10); ok {
		fmt.Println(value)
	}

}
