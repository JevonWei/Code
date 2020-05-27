package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	path := "users.txt"
	file, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
	} else {
		var bytes []byte = make([]byte, 20)
		for {
			n, err := file.Read(bytes)
			if err != nil {
				if err != io.EOF {
					fmt.Println(err)
				}
				break
			} else {
				fmt.Printf(string(bytes[:n]))
			}
		}
		file.Close()
	}
	//fmt.Println(err)
	//fmt.Println(file)
}
