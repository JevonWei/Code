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
	Tel      string
	Addr     string
}

func main() {

	users := map[int]User{}
	file, err := os.Open("user.gob")

	if err == nil {
		defer file.Close()
		decode := gob.NewDecoder(file)
		decode.Decode(&users)

		fmt.Printf("%#v", users)
	}
}
