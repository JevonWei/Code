package users

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"time"
)

type Persistence interface {
	SaveUser(users map[int]Users)
	LoadUser() map[int]Users
}

type JSONFile struct {
	name string
}

func (j JSONFile) LoadUser() map[int]Users {
	users := map[int]Users{}
	json_file := filepath.Join(json_dir, j.name)
	file, err := os.Open(json_file)
	if err != nil {
		//fmt.Println(err)
		return map[int]Users{}
	} else {
		defer file.Close()
		bytes, err := ioutil.ReadAll(file)
		if err != nil {
			fmt.Println(err)
			fmt.Println("反序列化文件读取失败")
			log.Println("反序列化文件读取失败")
			return map[int]Users{}
		} else {
			//err = json.Unmarshal(bytes, &users)
			err = json.Unmarshal(bytes, &users)
			if err == nil {
				return users
			} else {
				log.Println(err)
				return map[int]Users{}
			}

		}

	}
}

func (j JSONFile) SaveUser(users map[int]Users) {
	os.Mkdir(json_dir, 0644)
	//json_file := filepath.Join(json_dir, (strconv.FormatInt(time.Now().Unix(), 10) + ".user.json"))
	json_file := filepath.Join(json_dir, j.name)

	// 将user文件重命名
	if _, err := os.Stat(json_file); err == nil {
		os.Rename(json_file, filepath.Join(json_dir, (strconv.FormatInt(time.Now().Unix(), 10)+j.name)))
		log.Printf("%s文件已备份\n", json_file)
	}

	// 仅保留三个用户历史文件
	if names, err := filepath.Glob(json_dir + "/*" + j.name); err == nil {
		sort.Sort(sort.Reverse(sort.StringSlice(names)))
		if len(names) > 3 {
			for _, name := range names[3:] {
				os.Remove(name)
				log.Printf("%s文件已删除\n", name)
			}
		}

	}

	//bytes, err := json.Marshal(users)
	bytes, err := json.MarshalIndent(users, "", "\t")
	err = ioutil.WriteFile(json_file, bytes, os.ModePerm)
	//file.Seek(0, 1))
	if err != nil {
		fmt.Println(err)
		fmt.Println("user序列化错误")
		log.Println("user序列化错误")

	}
}
