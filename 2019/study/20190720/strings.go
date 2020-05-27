package main

import (
	"bufio"
	"fmt"
	"strings"
)

func main() {
	reader := strings.NewReader("asdfgzxc\nvwer")

	// 没三个字节输出异常
	// bytes := make([]byte, 3)

	// for {
	// 	n, err := reader.Read(bytes)
	// 	if err != nil {
	// 		if err != io.EOF {
	// 			fmt.Println(err)
	// 		}
	// 		break
	// 	} else {
	// 		fmt.Println(n, string(bytes[:n]))
	// 	}
	// }

	// 使用bufio按行处理
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	var builder strings.Builder
	builder.Write([]byte("abc"))
	builder.WriteString("zxcvbasdfq")
	fmt.Println(builder.String())
}
