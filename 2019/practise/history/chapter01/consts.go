package main

import "fmt"

func main() {
	const NAME string = "JevonWei"

	fmt.Println(NAME)

	const a = "abc"
	fmt.Println(a)

	const d, b, c string = "dd", "bb", "cc"
	fmt.Println(d, b, c)

	const (
		q int    = 12
		w string = "ww"
	)
	fmt.Println(q, w)

	const z, x, v = "zz", "xx", "vv"
	fmt.Println(z, x, v)

	const (
		C1 int = 1
		C2
		C3
		C4 float64 = 3.14
		C5
		C6 string = "JevonWei"
	)
	fmt.Println(C1, C2, C3)
	fmt.Println(C4, C5, C6)

	//枚举，const + iota
	const (
		E1 int = iota
		E2
		E3
	)
	fmt.Println(E1, E2, E3)

	const (
		E4 int = (iota + 1) * 3
		E5
		E6
	)
	fmt.Println(E4, E5, E6)

}
