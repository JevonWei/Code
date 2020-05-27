package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	//base64.RawStdEncoding
	//base64.RawURLEncoding
	//base64.URLEncoding

	x := base64.StdEncoding.EncodeToString([]byte("i am Jevon"))
	fmt.Println(x)

	bytes, err := base64.StdEncoding.DecodeString(x)
	fmt.Println(string(bytes), err)

}
