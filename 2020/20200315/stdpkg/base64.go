package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	x := base64.StdEncoding.EncodeToString([]byte("i am jevon"))
	fmt.Println(x)

	bytes, err := base64.StdEncoding.DecodeString(x)
	fmt.Println(string(bytes), err)
}
