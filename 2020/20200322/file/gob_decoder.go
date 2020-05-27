package main

import (
	"encoding/gob"
	"fmt"
	"os"
	"time"
)

type User struct {
	ID       int
	Name     string
	Birthday time.Time
	Addr     string
	Tel      string
}

func main() {
	users := map[int]User{}
	file, err := os.Open("user.gob")
	if err == nil {
		defer file.Close()

		decoder := gob.NewDecoder(file)
		decoder.Decode(&users)

		fmt.Println(users)
	}
}
