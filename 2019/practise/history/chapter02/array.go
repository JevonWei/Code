package main

import "fmt"

func main() {
	var nums [10]int
	var t2 [5]bool
	var t3 [3]string

	fmt.Printf("%T\n", nums)
	fmt.Println(nums)
	fmt.Println(t2)
	fmt.Printf("%q\n", t3)

	// 字面量
	nums = [10]int{1, 2, 3, 4}
	fmt.Println(nums)

	// 通过索引给特定位置的元素赋值
	nums = [10]int{0: 10, 9: 20}
	fmt.Println(nums)

	// [...] 根据数组的赋值自动匹配数组的长度
	nums01 := [...]int{1, 2, 3, 4}
	fmt.Println(nums01)
	fmt.Printf("%T\n", nums01)

	nums02 := [10]int{10, 20, 30}
	fmt.Printf("%T, %#v\n", nums02, nums02)

	nums03 := [...]string{"a", "b", "c"}
	fmt.Printf("%T %#v\n", nums03, nums03)

	nums04 := [10]int{2: 10, 5: 20, 6: 30}
	fmt.Printf("%T, %#v\n", nums04, nums04)

	// 操作
	nums05 := [3]int{1, 3, 4}
	nums06 := [3]int{2, 3, 5}
	fmt.Println(nums05 == nums06)
	fmt.Println(nums05 != nums06)

	// 获取数组的长度(len())
	fmt.Println(len(nums03))

	// 通过索引访问数组  0， 1， 2, ... len(array)-1
	fmt.Println(nums04[2], nums04[1])

	// 通过索引修改数组
	nums04[1] = 200
	fmt.Println(nums04)

	for i := 0; i < len(nums04); i++ {
		fmt.Println(i, ":", nums04[i])
	}

	for index, value := range nums04 {
		fmt.Println(index, ":", value)
	}

	for _, value := range nums04 {
		fmt.Println(value)
	}

	for index := range nums04 {
		fmt.Println(index)
	}

	// 切片
	// 数组切片的数据类型不是数组，而是切片类型
	fmt.Printf("%T,%T\n", nums04, nums04[0:5])
	fmt.Printf("%T, %v\n", nums04[0:5:8], nums04[0:5:6])
	fmt.Printf("%T, %#v\n", nums04[0:5:8], nums04[0:5:6])

}

/*
func main() {
	var nums [10]int
	var t2 [5]bool
	var t3 [3]string

	fmt.Printf("%T\n", nums)
	fmt.Println(nums)
	fmt.Println(t2)
	fmt.Printf("%q\n", t3)

	// 字面量
	nums = [10]int{10,20,30}
	fmt.Println(nums)

	nums =  [10]int{0: 5, 9: 18}
	fmt.Println(nums)

	nums = [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(nums)

	nums02 := [5]int{10,20,30}
	fmt.Printf("%T %#v\n", nums02, nums02)

	nums03 := [...]int{1, 2}
	fmt.Printf("%T, %#v\n", nums03, nums03)

	nums04 := [15]int{1:2, 2: 3, 3: 5}
	fmt.Printf("%T, %#v\n", nums04, nums04)

	// 操作
	nums05 := [3]int{1,3,4}
	nums06 := [3]int{1,3,4}
	fmt.Println(nums05 == nums06)
	fmt.Println(nums05 != nums06)

	// 获取数组的长度
	//fmt.Println(len(nums04))

}

*/
