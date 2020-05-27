package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {
	append_path := "user.log"

	file, err := os.OpenFile(append_path, os.O_CREATE|os.O_APPEND, os.ModePerm)
	if err != nil {
		fmt.Println(err)
	} else {
		file.Write([]byte("Danran\n"))
		file.WriteString(fmt.Sprintf("%d\n", time.Now().Unix()))
		file.WriteString(strconv.FormatInt(time.Now().Unix(), 10))
		file.WriteString("\n")
	}
	file.Close()
}
