package main

import (
	"fmt"
	"os"
)

func main() {
	// Seek
	// 偏移量，相对位置
	// 文件开始 0 os.SEEK_SET
	// 当前位置 1 os.SEEK_CUR
	// 文件末尾 2 os.SEEK_END

	file, _ := os.Open("user.txt")
	bytes := make([]byte, 100)

	file.Seek(0, 0)
	n, err := file.Read(bytes)
	fmt.Println(string(bytes[:n]), err)

	// 偏移量， 相对位置
	file.Seek(-5, 2)
	n, _ = file.Read(bytes)
	fmt.Println(string(bytes[:n]))

	file.Close()
}
