package main

import (
	"fmt"
	"time"
)

func main() {

	local, _ := time.LoadLocation("Local")
	//fmt.Println(local)
	t, _ := time.ParseInLocation("2006-01-02 15:04:05", "2019-07-16 20:22:15", local)
	fmt.Println(t)
	fmt.Println(t.Format("2006/01/02"))

	t1, err := time.Parse("2006-01-02", "2019-07-16")
	fmt.Printf("%T,%T\n", t1, err)

}
