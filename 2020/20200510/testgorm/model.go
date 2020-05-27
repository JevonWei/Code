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
	charset    string = "utf8mb4"
)

var dsn string = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&loc=Local&parseTime=true",
	dbUser, dbPassword, dbHost, dbPort, dbName)

type User struct {
	Id       int    `gorm:"primary_key; auto_increment"`
	Name     string `gorm:"type:varchar(32); not null; default:''"`
	Password string
	Birthday time.Time
	Sex      bool
	Tel      string `gorm:"column:telephone"`
	Addr     string
	Desc     string `gorm:"type:varchar(256)"`
}

func (u *User) TableName() string {
	return "user"
}

func main() {

	db, err := gorm.Open("mysql", dsn)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	defer db.Close()

	db.AutoMigrate(&User{})

	user := User{
		Name:     "kk2",
		Password: "123456",
		Birthday: time.Date(1998, 11, 11, 0, 0, 0, 0, time.UTC),
		Sex:      false,
	}
	db.Create(&user)
	// 判断记录是否是新的
	if db.NewRecord(user) {
		db.Create(&user)
	}
}
