package main

import "fmt"

func main() {
	a := 1
	b := true
	c := 1.10
	d := [3]int{1, 2, 3}
	e := []int{4, 5, 6}
	f := map[int]string{0: "AA", 1: "BB", 2: "CC"}
	h := &a

	a1, b1, c1, d1, e1, f1, h1 := a, b, c, d, e, f, h
	a1 = 2
	b1 = false
	c1 = 2.20
	d1[1] = 111
	e1[1] = 444
	f1[1] = "DD"
	h1 = &a1

	fmt.Println(a, b, c, d, e, f, h)
	fmt.Println(a1, b1, c1, d1, e1, f1, h1)
}
