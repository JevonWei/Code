package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"

	// _ "github.com/mattn/go-sqlite3"
	"github.com/astaxie/beego/orm"
)

const (
	dbUser     string = "root"
	dbPassword string = "danran"
	dbHost     string = "127.0.0.1"
	dbPort     int    = 3306
	dbName     string = "test"
)

var dsn string = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&loc=Local&parseTime=true",
	dbUser, dbPassword, dbHost, dbPort, dbName)

type User_orm struct {
	Id   int
	Name string `orm:size(100)`
}

func init() {

	//orm.RegisterDataBase("default", "mysql", "user=root password=danran dbname=test host=127.0.0.1 port=3306 sslmode=disable", 30)
	orm.RegisterDataBase("default", "mysql", dsn, 30)
	orm.RegisterModel(new(User_orm))

	orm.RunSyncdb("default", false, true)

}

func main() {
	o := orm.NewOrm()
	user := User_orm{Name: "jevon"}

	id, err := o.Insert(&user)
	fmt.Printf("ID: %d, ERR:%v\n", id, err)

	user.Name = "Danran"
	num, err := o.Update(&user)
	fmt.Printf("NUM: %d, ERR: %v\n", num, err)

	u := User_orm{Id: user.Id}
	err = o.Read(&u)
	fmt.Printf("ERR: %v\n", err)
}
