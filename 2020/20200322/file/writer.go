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

		writer.WriteString("abcddddd")
		writer.Write([]byte("1212121212"))
		writer.Flush()
	}
}
