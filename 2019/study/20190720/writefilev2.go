package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	err := ioutil.WriteFile("users.txt", []byte("XXXXXXXXXXX"), os.ModePerm)
	fmt.Println(err)

}
