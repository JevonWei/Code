package main

import (
	"bufio"
	"crypto/md5"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

func md5reader(reader *bufio.Reader) string {
	hasher := md5.New()
	bytes := make([]byte, 1024*1024*10)
	for {
		n, err := reader.Read(bytes)
		if err != nil {
			if err != io.EOF {
				return ""
				fmt.Println(err)
			}
			break
		} else {
			hasher.Write(bytes[:n])
		}
	}
	return fmt.Sprintf("%X", hasher.Sum(nil))
}
func md5file(path string) string {
	file, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
	} else {
		defer file.Close()
		return md5reader(bufio.NewReader(file))
	}
	return ""
}

func md5str(txt string) string {
	return md5reader(bufio.NewReader(strings.NewReader(txt)))
	// return fmt.Sprintf("%X", md5.Sum([]byte(txt)))
}

func main() {
	txt := flag.String("s", "", "md5 txt")
	path := flag.String("f", "", "file path")
	help := flag.Bool("h", false, "help")

	flag.Usage = func() {
		fmt.Println(`
Usage: md5.exe [-s 123abc] [-f filepath]

Options:
		`)
		flag.PrintDefaults()
	}

	flag.Parse()

	if *help || *txt == "" && *path == "" {
		flag.Usage()
	} else {
		var md5 string
		if *path != "" {
			md5 = md5file(*path)
		} else {
			md5 = md5str(*txt)
		}
		fmt.Println(md5)
	}
}
