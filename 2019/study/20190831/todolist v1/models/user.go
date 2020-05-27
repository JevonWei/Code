package models

import (
	"database/sql"
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"time"
)

type User struct {
	ID         int       `json:id`
	Name       string    `json:name`
	Birthday   time.Time `json:birthday`
	Sex        bool
	Addr       string `json:addr`
	Tel        string `json:tel`
	Desc       string `json:desc`
	Password   string `json:password`
	CreateTime time.Time
}

func (u User) ValidatePassword(passwd string) bool {
	log.Printf("%s Account Verify Success", u.Name)
	return passwd == u.Password
}

func Init() {
	logfile := "user.log"
	file, err := os.OpenFile(logfile, os.O_APPEND|os.O_CREATE, os.ModePerm)

	if err == nil {
		log.SetOutput(file)
		log.SetFlags(log.Flags() | log.Lshortfile)
	}
}

func loadUsers() (map[int]User, error) {
	if bytes, err := ioutil.ReadFile("datas/user.json"); err != nil {
		if os.IsNotExist(err) {
			return map[int]User{}, nil
		}
		return nil, err
	} else {
		var users map[int]User
		if err := json.Unmarshal(bytes, &users); err == nil {
			return users, nil
		} else {
			return nil, err
		}
	}
}

func storeUsers(users map[int]User) error {
	bytes, err := json.Marshal(users)
	if err != nil {
		return err
	}
	return ioutil.WriteFile("datas/user.json", bytes, 0X066)
}

func GetUsers(q string) []User {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	if err := db.Ping(); err != nil {
		panic(err)
	}
	defer db.Close()

	rows, err := db.Query("select id, name, sex, birthday, addr, tel, `desc` from users")
	if err != nil {
		panic(err)
	}

	users := make([]User, 0)
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.ID, &user.Name, &user.Sex, &user.Birthday, &user.Addr, &user.Tel, &user.Desc); err == nil {
			if q == "" || strings.Contains(user.Name, q) ||
				strings.Contains(user.Tel, q) || strings.Contains(user.Addr, q) ||
				strings.Contains(user.Desc, q) {
				users = append(users, user)
			}
		}
	}
	return users
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
	sql := "select id,name,password,sex,birthday,tel,addr,create_time from users where name=?"
	row := db.QueryRow(sql, name)
	err = row.Scan(&user.ID, &user.Name, &user.Password, &user.Sex, &user.Birthday, &user.Tel, &user.Addr, &user.CreateTime)
	return user, err

}

func ValidateCreateUser(name, password, birthday, tel, addr, desc string) map[string]string {
	errors := map[string]string{}
	if len(name) > 12 || len(name) < 4 {
		errors["name"] = "Name长度需在4-12位之间"
	} else if _, err := GetUserByName(name); err == nil {
		errors["name"] = "Name已存在"
	}

	if len(password) > 30 || len(password) < 6 {
		errors["password"] = "password长度需在6-30位之间"
	}

	return errors
}

func CreateUser(name, password, birthday, tel, addr, desc string, sex bool) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}

	if err := db.Ping(); err != nil {
		panic(err)
	}
	defer db.Close()

	_, err = db.Exec("insert into users(name, password, sex, birthday, tel, addr, `desc`, create_time) values(?,?,?,?,?,?,?,?)", name, password, sex, birthday, tel, addr, desc, time.Now())

	if err != nil {
		panic(err)
	}
}

func GetUserById(id int) (User, error) {
	var user User
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return user, nil
	}

	if err := db.Ping(); err != nil {
		return user, nil
	}

	defer db.Close()

	row := db.QueryRow("select id, name, birthday, tel, addr, `desc` from users where id=?", id)
	err = row.Scan(&user.ID, &user.Name, &user.Birthday, &user.Tel, &user.Addr, &user.Desc)

	return user, err

}

func ModifyUser(id int, name, birthday, addr, tel, desc string) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}

	if err := db.Ping(); err != nil {
		panic(err)
	}

	defer db.Close()

	_, err = db.Exec("update users set name=?, birthday=?, addr=?, tel=?, `desc`=? where id = ?", name, birthday, addr, tel, desc, id)

	if err != nil {
		panic(err)
	}

	log.Printf("%s user Create Success", name)

}

func DeleteUser(id int) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}

	if err := db.Ping(); err != nil {
		panic(err)
	}

	defer db.Close()

	_, err = db.Exec("delete from users where id = ?", id)

	if err != nil {
		panic(err)
	}

	log.Printf("ID=%d user Delete Success", id)

}

func ModifyPasswd(name, passwd string) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}

	if err := db.Ping(); err != nil {
		panic(err)
	}

	defer db.Close()

	_, err = db.Exec("update users set password=? where name = ?", passwd, name)

	if err != nil {
		panic(err)
	}

	log.Printf("Name为 %s 的Passwd Modidy Success", name)

}
