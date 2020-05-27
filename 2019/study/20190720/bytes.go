package main

import (
	"bytes"
	"fmt"
)

func main() {
	buffer := bytes.NewBufferString("asdzxca")

	buffer.Write([]byte("1234"))
	buffer.WriteString("!@#$")

	fmt.Println(buffer.String())

	bytes := make([]byte, 2)
	buffer.Read(bytes)
	fmt.Println(string(bytes))

	line, _ := buffer.ReadString('!')
	fmt.Println(line)
}
