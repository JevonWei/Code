package main

import "fmt"

func main() {
	var nums []int

	fmt.Printf("%T\n", nums)
	fmt.Printf("%#v\n", nums)
	fmt.Println(nums == nil)

	// 字面量
	nums = []int{1, 2, 3, 4}
	fmt.Printf("%#v\n", nums)
	nums = []int{1, 2, 3}
	fmt.Printf("%#v\n", nums)

	var arrays [10]int = [10]int{1, 3, 4, 5, 6, 7, 8}
	nums = arrays[1:10]
	fmt.Printf("%#v, %d, %d\n", nums, len(nums), cap(nums))

	// make函数
	nums = make([]int, 3)
	fmt.Printf("%#v, %d %d\n", nums, len(nums), cap(nums))

	nums = make([]int, 3, 5)
	fmt.Printf("%#v, %d %d\n", nums, len(nums), cap(nums))

	// 元素操作(增删改查)
	fmt.Println(nums[0])
	fmt.Println(nums[1])
	fmt.Println(nums[2])
	nums[2] = 2
	fmt.Println(nums[2])

	// 切片通过append添加元素
	nums = append(nums, 4, 5)
	fmt.Printf("%#v, %d %d\n", nums, len(nums), cap(nums))

	nums = append(nums, 4, 5)
	fmt.Printf("%#v, %d %d\n", nums, len(nums), cap(nums))

	// 遍历切片
	for i := 0; i < len(nums); i++ {
		fmt.Println(i, nums[i])
	}
	for i, v := range nums {
		fmt.Println(i, v)
	}

	nums = make([]int, 3, 10)

	// 切片操作
	n := nums[1:3:8]
	// 切片得我容量为n_cap - start
	fmt.Printf("%T, %#v %d %d\n", n, n, len(n), cap(n))
	n = nums[2:3]
	// 切片得我容量为src_cap - start
	fmt.Printf("%T, %#v %d %d\n", n, n, len(n), cap(n))

	nums = make([]int, 3, 5)
	nums02 := nums[1:3]
	fmt.Println(nums, nums02)
	nums02[0] = 1
	nums02 = append(nums02, 2)
	fmt.Println(nums, nums02)
	nums = append(nums, 10)
	fmt.Println(nums, nums02)
	nums = append(nums, 20)
	fmt.Println(nums, nums02)
	nums02 = append(nums02, 5)
	fmt.Println(nums, nums02)

	nums = arrays[:]
	fmt.Println(nums, arrays)
	arrays[0] = 100
	fmt.Println(nums, arrays)

	// 删除
	// copy
	nums04 := []int{1, 2, 3}
	nums05 := []int{10, 20, 30, 40}
	copy(nums05, nums04)
	fmt.Println(nums05)

	nums05 = []int{10, 20, 30, 40}
	copy(nums04, nums05)
	fmt.Println(nums04)

	// 删除索引为0的元素，删除最后一个元素
	nums06 := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(nums06[1:])
	fmt.Println(nums06[:len(nums06)-1])

	// 删除中间的元素，如索引为2的元素
	copy(nums06[2:], nums06[3:])
	fmt.Println(nums06[:len(nums06)-1])

	// 堆栈：每次添加在队尾，移除元素在队尾(先进后出)
	// 队列：每次添加在队尾，移除元素在队头(先进先出)
	queue := []int{}
	queue = append(queue, 1)
	queue = append(queue, 2)
	queue = append(queue, 3)
	queue = append(queue, 5)

	// 1， 2， 3， 5
	// 队头移出元素
	fmt.Println(queue)
	fmt.Println(queue[0])
	queue = queue[1:]
	//queue = queue[:len(queue)-1]
	fmt.Println(queue)
	queue = queue[1:]
	fmt.Println(queue)

	// 堆栈
	stack := []int{}
	stack = append(stack, 1)
	stack = append(stack, 2)
	stack = append(stack, 3)
	fmt.Println(stack[len(stack)-1])
	stack = stack[:len(stack)-1]
	fmt.Println(stack[len(stack)-1])
	stack = stack[:len(stack)-1]
	fmt.Println(stack[len(stack)-1])
	stack = stack[:len(stack)-1]

	// 多维切片
	points := [][]int{}
	points02 := make([][]int, 0)
	fmt.Printf("%T\n", points02)

	points = append(points, []int{1, 2, 3})
	points = append(points, []int{4, 5, 6})
	points = append(points, []int{2, 3, 4, 5, 6, 7})
	fmt.Println(points)
	fmt.Println(points[0])
	fmt.Println(points[2][1])

	// 数组是值类型
	// 切片
	slices01 := []int{1, 2, 3}
	slices02 := slices01

	slices02[0] = 10
	fmt.Println(slices01, slices02)

	// 数组
	array01 := [3]int{1, 2, 3}
	array02 := array01

	array02[0] = 10
	fmt.Println(array01, array02)

}

/*
func main() {
	var numbers []int
	fmt.Printf("%T\n", numbers)
	fmt.Printf("%#v\n", numbers)

	fmt.Println(numbers == nil)

	// 字面量
	numbers = []int{1, 2, 3}
	fmt.Printf("%#v\n", numbers)

	numbers = []int{1, 2, 3, 4, 5}
	fmt.Printf("%#v\n", numbers)

	// 数组切片赋值
	var arrays [10]int = [10]int{1, 2, 3, 4, 5, 6}
	numbers = arrays[1:10]
	fmt.Printf("%#v %d %d\n", numbers, len(numbers), cap(numbers))

	// make函数
	numbers = make([]int, 3)
	fmt.Printf("%#v %d %d\n", numbers, len(numbers), cap(numbers))

	numbers = make([]int, 3, 5)
	fmt.Printf("%#v %d %d\n", numbers, len(numbers), cap(numbers))

	// 元素操作(增， 删， 改， 查)
	fmt.Println(numbers[0])
	fmt.Println(numbers[2])
	// fmt.Println(numbers[3])
	numbers[2] = 10
	fmt.Println(numbers)


	// 删除
	// copy

	numbers04 := []int{1, 2, 3}
	numbers05 := []int{1, 2, 3, 4}

	copy(numbers05, numbers04)
	fmt.Println(numbers05)

	copy(numbers04, numbers05)
	fmt.Println(numbers04)

	// 删除索引为0， 最后一个索引的元素
	numbers06 := []int{1, 2, 3, 4, 5}
	fmt.Println(numbers06[1:])
	fmt.Println(numbers06[:len(numbers06)-1])

	// 删除中间(第二个)的元素 2
	copy(numbers06[2:], numbers06[3:])
	fmt.Println(numbers06[:len(numbers06)-1])

	// 堆栈：每次添加在队尾，移除元素在队尾(先进后出)
	// 堆列：每次添加在队尾，移除元素在队头(先进先出)
	queue := []int{}
	queue = append(queue, 1)
	queue = append(queue, 2)
	queue = append(queue, 3)
	queue = append(queue, 5)
	// 1, 2, 3, 5
	fmt.Println(queue[0])
	queue = queue[1:]
	// 2, 3, 5
	fmt.Println(queue)
	fmt.Println(queue[0])
	queue = queue[1:]
	// 3, 5
	fmt.Println(queue)

	// 堆栈
	stack := []int{}
	stack = append(stack, 1)
	stack= append(stack, 2)
	stack = append(stack, 3)

	fmt.Println(stack[len(stack)-1])
	stack = stack[:len(stack)-1]
	fmt.Println(stack[len(stack)-1])
	stack = stack[:len(stack)-1]
	fmt.Println(stack[len(stack)-1])
	stack = stack[:len(stack)-1]



	// 多维切片
	points := [][]int{}
	points02 := make([][]int, 0)
	fmt.Printf("%T\n", points02)

	points = append(points, []int{1, 2, 3})
	points = append(points, []int{3, 4, 0})
	points = append(points, []int{3, 4, 0, 2, 4, 5})

	fmt.Println(points)
	fmt.Println(points[0])
	fmt.Println(points[0][1])

	// 数组是值类型
	slice01 := []int{1, 2, 3}
	slice02 := slice01

	slice02[0] = 10
	fmt.Println(slice01, slice02)

	array01 := [3]int{1, 2, 3}
	array02 := array01
	array02[0] = 10
	fmt.Println(array01, array02)
}

*/
