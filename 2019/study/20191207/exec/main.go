package main

import (
	"os/exec"
	"fmt"
)

func main() {
	cmd := exec.Command("ping", "-n", "2","www.baidu.com")
	cxt, _ := cmd.Output()
	fmt.Println(string(cxt))
}