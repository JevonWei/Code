package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	fmt.Println(filepath.Abs("."))
	fmt.Println(filepath.Abs(os.Args[0]))

	fmt.Println(filepath.Base("c:/test/a.txt"))
	fmt.Println(filepath.Dir(os.Args[0]))
	fmt.Println(filepath.Ext("c:/test/a.txt"))

	path, _ := filepath.Abs(os.Args[0])
	dirpath := filepath.Dir(path)

	iniPath := dirpath + "/conf/ip.ini"
	fmt.Println(iniPath)

	fmt.Println(filepath.Join(dirpath, "conf", "ip.ini"))

	fmt.Println(filepath.Glob("./*.go"))
	fmt.Println(filepath.Glob("./[ab]*.go"))

	filepath.Walk(".", func(path string, fileInfo os.FileInfo, err error) error {
		fmt.Println(path, fileInfo.Name())
		return nil
	})
}
