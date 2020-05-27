package main

import "os"

func main() {
	//os.Mkdir("Test", 0644)
	//os.Rename("Test", "Test1")
	//os.Remove("test1")
	//os.MkdirAll("Test/xx.txt", 0644)
	os.RemoveAll("Test")
}
