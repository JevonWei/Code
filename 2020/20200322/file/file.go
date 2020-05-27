package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	//os.Rename("user.log", "user2.log")
	os.Remove("user.txt")
	t := time.Now()
	fmt.Println(t.Unix())
}
