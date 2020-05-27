package main

import "os"

func main() {
	os.Rename("user.txt", "user.v1.txt")
}
