package main

import (
	"fmt"
)

func main() {
	var a int
	var b int32

	a = 15
	fmt.Printf("%T\n", a)

	a = a + a
	b = b + 5

	fmt.Printf("%d, %03d\n", a, b)

	// s := "Hello, 世界! Hello!"
	// ss := strings.Split(s, "")
	// fmt.Printf("%q\n", ss)

	// fmt.Println(time.Now().Year())

	// k := 6
	// switch k {
	// case 4:
	// 	fmt.Println("was <= 4")
	// 	fallthrough
	// case 5:
	// 	fmt.Println("was <= 5")
	// 	fallthrough
	// case 6:
	// 	fmt.Println("was <= 6")
	// 	fallthrough
	// case 7:
	// 	fmt.Println("was <= 7")
	// 	fallthrough
	// case 8:
	// 	fmt.Println("was <= 8")
	// 	fallthrough
	// default:
	// 	fmt.Println("default case")
	// }

	for i := 0; i < 5; i++ {
		var v int
		fmt.Printf("%d \n", v)
		v = 5
	}

	// for i := 0; ; i++ {
	// 	fmt.Println("Value of i is now:", i)
	// }

	// s := ""
	// for s != "aaaaa" {
	// 	fmt.Println("Value of s:", s)
	// 	s = s + "a"
	// }

	// for i, j, s := 0, 5, "a"; i < 3 && j < 100 && s != "aaaaa"; i, j,
	// 	s = i+1, j+1, s+"a" {
	// 	fmt.Println("Value of i, j, s:", i, j, s)
	// }

	// i := 0
	// for { //since there are no checks, this is an infinite loop
	// 	if i >= 3 {
	// 		break
	// 	}
	// 	//break out of this for loop when this condition is met
	// 	fmt.Println("Value of i is:", i)
	// 	i++
	// }
	// fmt.Println("A statement just after for loop.")

	// for i := 0; i < 7; i++ {
	// 	if i%2 == 0 {
	// 		continue
	// 	}
	// 	fmt.Println("Odd:", i)
	// }

}
