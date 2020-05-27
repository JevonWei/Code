package main

import "fmt"

// a 是源，b 借助， c 目的长度
func tower(a, b, c string, layer int) {
	if layer == 1 {
		fmt.Println(a, "111->", c)
		return
	}

	// n-1 个 a 借助 c 到 b
	tower(a, c, b, layer-1)
	fmt.Println(a, "11->", c)
	// b n-1 借助a移动到c
	tower(b, a, c, layer-1)

}
func main() {
	tower("a", "b", "c", 3)
}
