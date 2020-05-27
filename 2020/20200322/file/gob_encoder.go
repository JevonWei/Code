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
	Addr     string
	Tel      string
}

func main() {
	// users := map[int]string{1: "aa", 2: "bb", 3: "cc", 4: "dd"}

	// users := map[int]User{
	// 	1: User{1, "aa", time.Now(), "121212", "shanghai"},
	// 	2: User{2, "bb", time.Now(), "44444444", "hangzhou"},
	// 	3: User{3, "cc", time.Now(), "212121", "beijing"},
	// 	4: User{4, "dd", time.Now(), "111111", "henan"},
	// }

	stdu := User{1, "aa", time.Now(), "121212", "shanghai"},
	file, err := os.Create("user.gob")
	if err == nil {
		defer file.Close()

		encoder := gob.NewEncoder(file)
		encoder.Encode(stdu)
	}
}
