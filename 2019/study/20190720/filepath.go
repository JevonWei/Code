package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	fmt.Println(filepath.Abs("."))
	fmt.Println(os.Args[0])
	fmt.Println(filepath.Abs(os.Args[0]))

	path, _ := filepath.Abs(os.Args[0])
	dirpath := filepath.Dir(path)

	fmt.Println(filepath.Base("c:/test/a.txt"))
	fmt.Println(filepath.Base("c:/test/xxx/"))
	fmt.Println(filepath.Base(path))

	fmt.Println(filepath.Dir("c:/test/a.txt"))
	fmt.Println(filepath.Dir("c:/test/xxx/"))
	fmt.Println(filepath.Dir(path))

	fmt.Println(filepath.Ext("c:/test/a.txt"))
	fmt.Println(filepath.Ext("c:/test/xxx/a"))
	fmt.Println(filepath.Ext(path))

	iniPath := dirpath + "/conf/ip.ini"
	fmt.Println(iniPath)

	// 将文件名和路径进行拼接
	fmt.Println(filepath.Join(dirpath, "conf", "ip.ini")

	// 获取指定路径下[cd]*开头的go文件
	fmt.Println(filepath.Glob("D:/Code/goang/code/20190720/[cd]*.go"))
	fmt.Println(filepath.Glob("././[cd]*.go"))

	// 变量当前目录下所有的文件
	filepath.Walk(".", func(path string, fileInfo os.FileInfo, err error) error {
		fmt.Println(path, fileInfo.Name())
		return nil
	})
}
