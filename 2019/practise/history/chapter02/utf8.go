package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s := "淡然归心"
	fmt.Println(len(s))
	fmt.Println(utf8.RuneCountInString(s))  // 获取unicode字符的数量
}