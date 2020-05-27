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
	dbName     string = "test"
	charset    string = "utf8mb4"
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

	rows, err := db.Query("select * from user")
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	var (
		id       int
		name     string
		password string
		desc     string
		tel      string
		addr     string
		super    bool
	)

	for rows.Next() {
		rows.Scan(&id, &name, &password, &desc, &tel, &addr, &super)
		fmt.Println(id, name, password, desc, tel, addr, super)
	}

	db.Close()
}
