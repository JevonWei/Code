package users

import (
	"encoding/gob"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"time"
)

type Persistence1 interface {
	LoadUser() map[int]User
	StoreUser(users map[int]User)
}

type JSONFile struct {
	name string
}

func NewJSONFile(name string) JSONFile {
	return JSONFile{name + ".json"}
}

func (j JSONFile) LoadUser() map[int]User {
	users := map[int]User{}
	if bytes, err := ioutil.ReadFile(j.name); err == nil {
		// if file, err := os.Open(j.name); err == nil {
		// bytes, err := ioutil.ReadAll(file)
		if err == nil {
			err = json.Unmarshal(bytes, &users)
		} else {
			fmt.Println("[-]发生错误", err)
		}
	} else {
		if os.IsNotExist(err) {
			return make(map[int]User) //return map[int]User{}, nil
			// fmt.Println("[-]文件打开错误", err)
		}

	}
	return users
}

func (j JSONFile) StoreUser(users map[int]User) {
	// 将users.json文件重命名
	if _, err := os.Stat(j.name); err == nil {
		os.Rename(UserFile+".json", strconv.FormatInt(time.Now().Unix(), 10)+".users.json")
	}

	// 超过三份以上的文件删除
	if names, err := filepath.Glob("*.json"); err == nil {
		sort.Sort(sort.Reverse(sort.StringSlice(names)))
		// fmt.Println(names)
		if len(names) > 3 {
			for _, name := range names[3:] {
				os.Remove(name)
			}
		}

	}

	if file, err := os.Create(j.name); err == nil {
		defer file.Close()
		bytes, _ := json.Marshal(users)
		file.Write(bytes)
	}
}

type GobFile struct {
	name string
}

func NewGlobFile(name string) JSONFile {
	return JSONFile{name + ".gob"}
}

func (g GobFile) LoadUser() map[int]User {
	var users map[int]User

	file, err := os.Open(g.name)
	if err != nil {
		if os.IsNotExist(err) {
			return map[int]User{}
		}
		return nil
	}
	defer file.Close()
	decoder := gob.NewDecoder(file)
	err = decoder.Decode(&users)

	if err != nil {
		fmt.Println("[-]发生错误", err)
	}
	return users
}

func (g GobFile) StoreUser(users map[int]User) {
	file, err := os.Create(g.name)
	if err != nil {
		fmt.Println("文件创建错误", err)
	}
	defer file.Close()
	encoder := gob.NewEncoder(file)
	err = encoder.Encode(users)
	if err != nil {
		fmt.Println(err)
	}

}
