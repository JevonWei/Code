package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	bytes, err := ioutil.ReadFile("users.txt")
	if err == nil {
		fmt.Println(string(bytes))
	}
}
