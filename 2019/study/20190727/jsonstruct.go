package main

import (
	"encoding/json"
	"fmt"
)

const (
	Large = iota
	Medium
	Small
)

type Size int

func (s Size) MarshalText() ([]byte, error) {
	switch s {
	case Large:
		return []byte("large"), nil
	case Medium:
		return []byte("medium"), nil
	case Small:
		return []byte("small"), nil
	default:
		return []byte("unknown"), nil

	}
}

type Addr struct {
	Region string `json:"region"`
	Street string `json:"Street"`
	No     int    `json:"No"`
}

// 需要进行序列化的属性必须公开
type User struct {
	ID int `json:"id, string"`
	//ID   int
	Name string `json:"name"`
	Sex  int    `json:"Sex,int,omitempty"`
	Tel  string `json:"-"`
	Addr Addr   `json:"addr`
	Size Size
}

func main() {
	user := User{1, "danran", 1, "1222222", Addr{"hsanghai", "pudong", 12}, Medium}

	bytes, _ := json.MarshalIndent(user, "", "\t")
	fmt.Println(string(bytes))

	var user02 = user

	json.Unmarshal(bytes, &user02)
	fmt.Println(user02)
}
