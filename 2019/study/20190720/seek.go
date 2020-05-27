package main

import (
	"fmt"
	"os"
)

func main() {
	file, _ := os.Open("users.txt")

	// 偏移量
	// 文件开始 0 os.SEEK_SET
	// 当前位置 1 os.SEEK_CUR
	// 文件末尾 2 os.SEEK_END

	fmt.Println(file.Seek(4, 0)) // 4 为偏移量，0 为文件开始位置
	bytes := make([]byte, 100)
	n, err := file.Read(bytes)

	fmt.Println(n, err, string(bytes[:n]))

	n, err = file.Read(bytes)
	fmt.Println(n, err, bytes[:n])

	file.Close()
}
