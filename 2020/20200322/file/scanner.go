package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("user.txt")
	if err == nil {
		defer file.Close()

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			fmt.Println(scanner.Text())
			fmt.Println(scanner.Bytes())
		}
	}
}
