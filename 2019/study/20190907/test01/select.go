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
	Id       int    `gorm:"primary key"`
	Name     string `gorm:"type:varchar(32); not null; default:' '"`
	Password string
	Birthday time.Time
	Sex      bool
	Addr     string
	Tel      string `gorm:column:telephone`
	Desc     string `gorm:"type:text`
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

	db.LogMode(true)
	db.AutoMigrate(&User{})

	// for i := 0; i < 10; i++ {
	// 	user := User{
	// 		Name:     fmt.Sprintf("Jevon_%d", i),
	// 		Password: "123456",
	// 		Birthday: time.Date(1988, 11, 11, 0, 0, 0, 0, time.UTC),
	// 		Sex:      false,
	// 	}
	// 	db.Create(&user)
	// }

	var user User
	db.First(&user, "name = ?", "Jevon_2")
	fmt.Println(user)

	var user2 User
	db.Last(&user2, "name = ?", "Jevon_3")
	fmt.Println(user2)

	var users []User
	//db.Where("name=?", "Jevon_5").Find(&users)
	//db.Where("name like ?", "%5").Find(&users)
	//db.Where("name in (?)", []string{"Jevon_2", "Jevon_6"}).Find(&users)
	//db.Where("name = ? and password=?", "Jevon_3", "123456").Find(&users)
	//db.Where("name = ?", "Jevon_3").Where("password = ?", "123456").Find(&users)
	//db.Where("name = ?","Jevon_3").Not("password = ?","123456").Find(&users)
	//db.Where("name = ?", "Jevon_11").Or("name = ?", "Jevon_4").Find(&users)

	db.Select("name, password").Find(&users)
	db.Select([]string{"name", "password"}).Find(&users)

	//db.Order("id desc, name asc").Find(&users)
	// db.Order("id desc, name asc").Limit(3).Offset(5).Find(&users)
	// fmt.Println(users)

	// var count int
	// db.Model(&User{}).Count(&count)
	// fmt.Println(count)

	// var count1 int
	// db.Model(&User{}).Where("name=?","Jevon_2").Count(&count1)
	// fmt.Println(count1)

	rows, _ := db.Model(&User{}).Select("name, password").Rows()
	for rows.Next() {
		var name, password string
		rows.Scan(&name, &password)
		fmt.Println(name, password)
	}
	rows, _ = db.Model(&User{}).Select("name, count(*) as cnt").Group("name").Having("count(*) = ?", 1).Rows()
	for rows.Next() {
		var name string
		var count int
		rows.Scan(&name, &count)
		fmt.Println(name, count)
	}

	db.Close()
}
