package main

import (
	"fmt"
	"io"
	"os"
)

func main() {

	path := "user.txt"
	file, err := os.Open(path)

	if err == nil {
		var bytes []byte = make([]byte, 20)

		for {
			n, err := file.Read(bytes)
			if err != nil {
				if err != io.EOF {
					fmt.Println(err)
				}
				break
			} else {
				fmt.Print(string(bytes[:n]))
			}
		}
		file.Close()
	} else {
		fmt.Println(err)
	}
}
