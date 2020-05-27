package main

import (
	"fmt"
	"bytes"
)

func main() {
	var bytes01 []byte = []byte{'a','b','c'}
	fmt.Printf("%T, %#v\n", bytes01, bytes01)

	s := string(bytes01)
	fmt.Printf("%T %v\n", s, s)

	bs := []byte(s)
	fmt.Printf("%T %#v\n", bs, bs)

	fmt.Println(bytes.Compare([]byte("abc"), []byte("abc")))    // 比较字节
	fmt.Println(bytes.Index([]byte("abcdefg"), []byte("abc")))
	fmt.Println(bytes.Contains([]byte("abcdefabc"), []byte("defa")))
}