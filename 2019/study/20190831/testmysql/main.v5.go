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

	sql := "delete from users where id = ?;"
	fmt.Println(sql)

	result, err := db.Exec(sql, 14)
	fmt.Println(err)
	fmt.Println(result.LastInsertId())
	fmt.Println(result.RowsAffected())

	db.Close()

}
