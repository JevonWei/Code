package models

import (
	"crypto/md5"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

type User struct {
	Id       int
	Name     string
	Password string
	Desc     string
	Tel      string
	Addr     string
	Super    bool
}

func LoadUser() ([]User, error) {
	if bytes, err := ioutil.ReadFile("datas/users.json"); err != nil {
		if os.IsNotExist(err) {
			log.Println("users.json 文件不存在")
			return []User{}, nil
		}
		log.Fatalf("users.json read error: %s\n", err)
		return nil, err
	} else {
		if len(bytes) == 0 {
			log.Println("users.json数据为空")
			return []User{}, nil
		}

		var users []User
		if err := json.Unmarshal(bytes, &users); err == nil {
			log.Println("User Load Success")
			return users, nil
		} else {
			log.Fatalln("User Load error:", err)
			return nil, err
		}
	}
}

func StoreUser(users []User) error {
	bytes, err := json.Marshal(users)
	if err != nil {
		log.Fatalln("User store error:", err)
		return err
	}

	log.Println("User store Success")
	return ioutil.WriteFile("datas/users.json", bytes, 0X066)
}

func PasswdMd5(password string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(password)))
}

func GetUserByName(name string) (User, error) {
	var user User
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return user, err
	}

	if err := db.Ping(); err != nil {
		return user, err
	}

	defer db.Close()

	row := db.QueryRow("select id, name, password, `desc`, tel, addr, super from user where name=?", name)
	err = row.Scan(&user.Id, &user.Name, &user.Password, &user.Desc, &user.Tel, &user.Addr, &user.Super)
	// fmt.Println(user, err)
	return user, err

	// users, err := LoadUser()
	// if err != nil {
	// 	return User{}, err
	// }

	// for _, user := range users {
	// 	if user.Name == name {
	// 		return user, nil
	// 	}
	// }
	// return User{}, errors.New("Not Found")
}

func GetUserId() (int, error) {
	users, err := LoadUser()
	if err != nil {
		log.Fatalln("GetUserId LoadUser error: ", err)
		return -1, err
	}

	var id int
	for _, user := range users {
		if id < user.Id {
			id = user.Id
		}
	}

	return id + 1, nil
}

func GetSearchUser(q string) []User {
	users, err := LoadUser()
	if err != nil {
		log.Panicln("User Read Faild", err)
		panic(err)
		return []User{}
	}

	usernew := make([]User, 0)

	for _, user := range users {
		if q == "" || strings.Contains(user.Name, q) || strings.Contains(user.Tel, q) ||
			strings.Contains(user.Desc, q) || strings.Contains(user.Addr, q) {
			usernew = append(usernew, user)
		}
	}
	return usernew
}

func Register(name, password, desc, tel, addr string, super bool) {
	id, err := GetUserId()
	if err != nil {
		log.Fatalln(err)
	}

	user := User{
		Id:       id,
		Name:     name,
		Password: password,
		Desc:     desc,
		Tel:      tel,
		Addr:     addr,
		Super:    super,
	}

	if users, err := LoadUser(); err == nil {
		users = append(users, user)
		log.Printf("%s User Register successful\n", user.Name)
		StoreUser(users)
	}
}

func CreateUser(name, password, tel, addr, desc string, super bool) {
	Register(name, password, desc, tel, addr, super)
}

func CheckUserName(name string) bool {
	users, err := LoadUser()
	if err != nil {
		log.Panicln(err)
	}

	for _, user := range users {
		if name == user.Name {
			log.Printf("%s 用户已存在", name)
			return false
		}
	}
	return true
}

func SuperUser(name string) bool {
	users, _ := LoadUser()

	for _, user := range users {
		if user.Super {
			if user.Name == name {
				log.Printf("%s为管理员用户\n", name)
				return true
			}
		}
	}
	return false
}

func LoginAuth(name, password string) (bool, error) {
	users, err := LoadUser()
	if err != nil {
		log.Panicln(err)
	}

	for _, user := range users {
		if name == user.Name {
			if user.Password == PasswdMd5(password) {
				log.Printf("用户%s登录成功", name)
				return true, nil
			} else {
				log.Printf("用户%s验证失败", name)
				return false, errors.New("用户密码验证失败")
			}
		}
	}
	log.Printf("用户%s不存在，请先注册账号", name)
	return false, errors.New("输入的用户不存在")

}
