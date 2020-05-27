package main

import (
	"fmt"
	"os"
	"strings"
)

// 获取目录的子文件信息
func main() {
	file, err := os.Open("xxx")
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("文件不存在")
		}
	} else {
		file.Close()
	}

	for _, path := range []string{"xxx", "reader.go", "111"} {
		fileinfo, err := os.Stat(path)
		if err != nil {
			if os.IsNotExist(err) {
				fmt.Println("文件不存在")
			}
		} else {
			fmt.Println(strings.Repeat("*", 20))
			fmt.Println(fileinfo.Name())
			fmt.Println(fileinfo.IsDir())
			fmt.Println(fileinfo.Size())
			fmt.Println(fileinfo.ModTime())

			if fileinfo.IsDir() {
				dirfile, err := os.Open(path)
				if err == nil {
					defer dirfile.Close()
					// 	childrens, _ := dirfile.Readdir(-1)
					// 	for _, children := range childrens {
					// 		fmt.Println(children.Name(), children.IsDir(), children.Size(), children.ModTime())
					// 	}

					names, _ := dirfile.Readdirnames(-1)
					for _, name := range names {
						fmt.Println(name)
					}
				}
			}
		}
	}
}
