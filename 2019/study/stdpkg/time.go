package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Printf("%T\n", now)
	fmt.Printf("%v\n", now)

	fmt.Println(now.Format("2006/01/02 15:04:05"))
	fmt.Println(now.Format("2006-01-02"))
	fmt.Println(now.Format("03:04:05"))

	fmt.Println(now.Unix())     // 时间戳
	fmt.Println(now.UnixNano()) // 纳秒
	fmt.Println(now.Date())
	fmt.Println(now.Hour())

	t, err := time.Parse("2006-01-02 15:04:05", "2006/01/02 03:04:05")
	fmt.Println(err, t)

	t = time.Unix(0, 0) // 第一个0代表0秒，第二个0秒代表纳秒
	fmt.Println(t)

	d := now.Sub(t)
	fmt.Printf("%T, %v\n", d, d)

	fmt.Println(now.Hour())
	fmt.Println(now.Minute())
	fmt.Println(now.Second())

	fmt.Println(time.Now())
	time.Sleep(time.Second * 5) // 延时5秒
	fmt.Println(time.Now())     // 显示延时5秒之后的时间

	t = now.Add(3 * time.Hour) // 在当前时间上加3个小时
	fmt.Println(t)

	// 时间段
	d, err = time.ParseDuration("3h2m4s")  // 将3h2m4s转换为区间格式
	fmt.Println(err, d)
	fmt.Println(d.Hours()) // 时间转为hours显示
	fmt.Println(d.Minutes())
	fmt.Println(d.Seconds())

}
