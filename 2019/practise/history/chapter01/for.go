package main

import "fmt"

func main() {
	// 统计100以内数字相加
	// 索引 => 记录已经加到的n
	// 记录结果

	result := 0

	/*
		result += 1
		result += 2
		....
		ersult += 100
	*/

	/*
		i => 1 .... 100
		result += i
	*/

	// 初始化子语句；条件子语句；后置子语句
	for i := 1; i <= 100; i++ {
		result += i
	}

	fmt.Println(result)

	fmt.Println("------------")
	result = 0
	i := 1

	for i <= 100 {
		result += i
		i++
	}
	fmt.Println(result)

	// 死循环
	/*fmt.Println("------------")
	i = 0
	for {
		fmt.Println(i)
		i++
	}
	*/

	//遍历(字符串， 数组， 切片， 映射， 管道)
	desc := "我在地球"
	for i, ch := range desc {
		//fmt.Println(i, ch)
		fmt.Printf("索引：%d; 数据类型：%T; 字符字面值：%q\n", i, ch, ch)
	}

}
