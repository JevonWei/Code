package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	err := ioutil.WriteFile("user.txt", []byte("XXXXXXXXXXXXXXXXXXXXX"), os.ModePerm)
	fmt.Println(err)
}
