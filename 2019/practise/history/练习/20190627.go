package main

import "fmt"

func main() {
	const a = 3
	var b int
	b = 3
	fmt.Println(b)

	const (
		d int = 3
		e
		f
	)

	const (
		g int = iota + 1
		h
		i
	)
	fmt.Println(g, h, f)

	fmt.Println(d, e, f)
	var c int = 3
	fmt.Println(c)

	aa := 12
	fmt.Println(aa)

	var name, num = "Jevon", 12
	fmt.Println(name, num)

	var (
		bb string = "dan"
		cc int    = 111
	)
	fmt.Println(bb, cc)

	var a1 bool = true
	fmt.Println(a1)

	desc := "abcdefg"
	fmt.Printf("%T %c\n", desc[0], desc[0])
}
