package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

func main() {
	x := fmt.Sprintf("%X", md5.Sum([]byte("i am jevon")))

	bytes := md5.Sum([]byte("i am jevon"))
	fmt.Println(hex.EncodeToString(bytes[:]))
	fmt.Println(x)

	m := md5.New()
	m.Write([]byte("i am "))
	m.Write([]byte("jevon"))
	fmt.Printf("%x", m.Sum(nil))
}
