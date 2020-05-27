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

	db.AutoMigrate(&User{})

	user := User{
		Name:     "Jevon01",
		Password: "123456",
		Birthday: time.Date(1988, 11, 11, 0, 0, 0, 0, time.UTC),
		Sex:      false,
	}

	// for i := 0; i < 10; i++ {
	// 	user := User{
	// 		Name:     fmt.Sprintf("Jevon_%d", i),
	// 		Password: "123456",
	// 		Birthday: time.Date(1988, 11, 11, 0, 0, 0, 0, time.UTC),
	// 		Sex:      false,
	// 	}
	// 	db.Create(&user)
	// }

	fmt.Println(db.NewRecord(user))
	fmt.Println(user)

	db.Create(&user)
	fmt.Println(user)
	fmt.Println(db.NewRecord(user))

	if db.NewRecord(user) {
		db.Create(&user)
	}

	db.Close()
}
