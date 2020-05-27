package main

import (
	"encoding/csv"
	"os"
)

func main() {
	file, err := os.Create("user.csv")

	if err == nil {
		defer file.Close()

		writer := csv.NewWriter(file)
		writer.Write([]string{"编号", "名字", "性别"})
		writer.Write([]string{"1", "aa", "男"})
		writer.Write([]string{"2", "bb", "男"})
		writer.Write([]string{"3", "cc", "女"})
		writer.Flush()
	}
}
