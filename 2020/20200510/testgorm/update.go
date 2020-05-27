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

	db.LogMode(true)
	db.AutoMigrate(&User{})

	// 查找对象进行更新
	var user User
	if db.First(&user, "name=?", "kk_2").Error == nil {
		user.Name = "Jevon"
		db.Save(user)
	}

	db.Model(&User{}).Where("id > ?", 10).UpdateColumn("sex", true)
	db.Model(&User{}).Where("id > ?", 15).UpdateColumns(map[string]interface{}{"telephone": "abc", "addr": "China"})
	db.Model(&User{}).Where("id > ?", 18).Updates(User{Tel: "xxx", Addr: "beijing", Desc: "test"})
}
