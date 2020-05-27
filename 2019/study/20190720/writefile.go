package main

import (
	"fmt"
	"os"
)

func main() {
	path := "users2.txt"
	file, err := os.Create(path)
	if err != nil {
		fmt.Println(err)
	} else {
		// 写文件
		file.Write([]byte("JevonWei\n"))
		file.WriteString("Danran\n")

		// 关闭文件
		file.Close()
	}
}
