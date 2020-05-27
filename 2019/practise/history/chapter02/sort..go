package main

import "sort"
import "fmt"

func main() {
	nums := []int{1, 4, 1, 7, 6}
	sort.Ints(nums)
	fmt.Println(nums)

	names := []string{"wei", "zi", "Jevon", "dan"}
	sort.Strings(names)
	fmt.Println(names)

	heights := []float64{1.10, 1.23, 4.5, 3.14}
	sort.Float64s(heights)
	fmt.Println(heights)

}  