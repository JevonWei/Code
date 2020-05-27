package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

const (
	dbUser     string = "root"
	dbPassword string = "danran"
	dbHost     string = "127.0.0.1"
	dbPort     int    = 3306
	dbName     string = "todolist"
)

func main() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&loc=Local&parseTime=true",
		dbUser, dbPassword, dbHost, dbPort, dbName)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	if err := db.Ping(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	//userName, userPassword := "danran", "abc"
	//rows, err := db.Query("select id, name, password, sex, birthday, tel, addr, desc, createTime from users")
	//rows := db.QueryRow(fmt.Sprintf("select name,password from users where name = '%s' and password = '%s'", userName, userPassword))
	//rows := db.QueryRow("select name,password from users where name = 'dan' and password = '123'")

	// var (
	// 	name     string
	// 	password string
	// )
	//sql := "select name,password from users where name = ? and password = md5(?)"
	//rows := db.QueryRow(sql, userName, userPassword)

	//sql := "insert into users(name, password, sex, birthday, tel, addr, `desc`, ceate_time) values('wei', md5('abc'), 1, '2000-03-01', '121212', '121212', now())"
	sql := "insert into users(name, password, sex, birthday, tel, addr, `desc`, create_time) values(?, md5(?), ?, ?, ?, ?, ?, ?)"
	result, err := db.Exec(sql, "wei", "123", 1, "2000-01-01", "123", "123", "123", "2019-01-01")
	fmt.Println(err)
	fmt.Println(result.LastInsertId())
	fmt.Println(result.RowsAffected())

	// if rows.Next() {
	// 	rows.Scan(&name, &password)
	// 	fmt.Println(name, password)
	// }

	// rows.Scan(&name, &password)
	// fmt.Println(name, password)
	db.Close()

}
