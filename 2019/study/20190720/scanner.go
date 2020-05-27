package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("users.txt")
	if err == nil {
		defer file.Close()

		i := 0
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			fmt.Println(i, scanner.Text())
			i++
		}

	}

}
