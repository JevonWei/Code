package main

import (
	"fmt"
	"sort"
)

type User struct {
	ID    int
	Name  string
	Score float64
}

func main() {
	list := [][2]int{{1, 2}, {3, 2}, {5, 7}, {8, 4}}

	// 排序， 使用数组的第二个(索引为1)元素比较大小排序 []Type
	sort.Slice(list, func(i, j int) bool {
		return list[i][1] > list[j][1]
	})

	fmt.Println(list)

	users := []User{{1, "kk", 6}, {2, "danran", 5}, {3, "JevonWei", 7}}

	sort.Slice(users, func(i, j int) bool {
		return users[i].Score < users[j].Score
	})

	fmt.Println(users)
}
