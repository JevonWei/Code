package main

import (
	"fmt"
	"os"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

const (
	dbUser     string = "root"
	dbPassword string = "danran"
	dbHost     string = "127.0.0.1"
	dbPort     int    = 3306
	dbName     string = "testgorm"
)

var dsn string = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&loc=Local&parseTime=true", dbUser, dbPassword, dbHost, dbPort, dbName)

type User struct {
	gorm.Model
	Name     string `gorm:"type:varchar(32); not null; default:' '"`
	Password string
	Birthday time.Time
	Sex      bool
	Addr     string
	Tel      string `gorm:column:telephone`
	Desc     string `gorm:"type:text`
}

func (u *User) TableName() string {
	return "user3"
}

func main() {
	db, err := gorm.Open("mysql", dsn)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	//db.LogMode(true)
	db.AutoMigrate(&User{})

	// for i := 0; i < 10; i++ {
	// 	user := User{
	// 		Name:     fmt.Sprintf("Jevon_%d", i),
	// 		Password: fmt.Sprintf("password_%d", i),
	// 		Birthday: time.Date(1988, 11, 11, 0, 0, 0, 0, time.UTC),
	// 	}
	// 	db.Create(&user)
	// }

	var users []User
	db.Raw("select name from user3 where name = ?", "Jevon_2").Scan(&users)
	fmt.Println(users)

	var name string
	db.Raw("select name from user3 where id = ?", 2).Row().Scan(&name)
	fmt.Println(name)
	fmt.Println(db.Exec("insert into user3(name) value(?)", "Jevon_6"))
	//fmt.Println(db.Exec("insert into user3(name) value(?)", "Jevon_6").RowsAffected)
	//fmt.Println(db.Exec("delete from user3 where name = ?", "Jevon_7").RowsAffected)

	//fmt.Println(db.Exec("update user3 set name=? where id = ?", "danran", 9).RowsAffected)
	fmt.Println(db.Debug().Exec("update user3 set name=? where id = ?", "danran", 9).RowsAffected)
	db.Close()
}
