package main

import (
	"encoding/gob"
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

	users := map[int]User{
		1: User{1, "Dan", time.Now(), "12345", "DANRAN"},
		2: User{2, "Ran", time.Now(), "12345", "Ran"},
		3: User{3, "Jevon", time.Now(), "12345", "Jevon"},
	}

	file, err := os.Create("user.gob")

	if err == nil {
		defer file.Close()
		decode := gob.NewEncoder(file)
		decode.Encode(users)
	}
}
