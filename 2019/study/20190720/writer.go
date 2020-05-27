package main

import (
	"bufio"
	"os"
)

func main() {
	file, err := os.Create("user.txt")
	if err == nil {
		defer file.Close()

		writer := bufio.NewWriter(file)

		writer.WriteString("abcdefg")
		writer.Write([]byte("123456"))
		writer.Flush()
	}
}
