package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	write_path := "user2.txt"
	read_path := "user.txt"
	var bytes []byte = make([]byte, 10)

	file, err := os.Create(write_path)
	if err != nil {
		fmt.Println(err)
	} else {
		read_file, err := os.Open(read_path)
		if err == nil {
			for {
				n, err := read_file.Read(bytes)
				if err != nil {
					if err != io.EOF {
						fmt.Println(err)
					}
					break
				} else {
					file.Write(bytes[:n])
				}
			}
		} else {
			fmt.Println(err)
		}
		file.Write([]byte("\nDanran"))
		file.WriteString("\nJevonWei")
	}
	file.Close()
}
