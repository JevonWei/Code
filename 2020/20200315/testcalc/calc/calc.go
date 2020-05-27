package calc

func Add(x, y int) int {
	if x > 0 {
		return x + y
	} else {
		return x
	}

}

// func Multi(x, y int) int {
// 	return x * y
// }

func Fact(n int) int {
	if n == 0 {
		return 1
	}

	return n * Fact(n-1)
}
