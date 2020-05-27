package main

import (
	"flag"
	"fmt"
	"os"

	"todolist/models"
	_ "todolist/routers"
	"todolist/utils"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

const (
	dbUser     string = "root"
	dbPassword string = "danran"
	dbHost     string = "127.0.0.1"
	dbPort     int    = 3306
	dbName     string = "todolist2"
)

var dsn string = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&loc=Local&parseTime=true",
	dbUser, dbPassword, dbHost, dbPort, dbName)

func main() {
	init := flag.Bool("init", false, "Init Admin")
	force := flag.Bool("force", false, "force clear database")
	help := flag.Bool("help", false, "print help")
	h := flag.Bool("h", false, "print help")

	flag.Usage = func() {
		fmt.Println("Paginator usage")
		flag.PrintDefaults()
		os.Exit(0)
	}

	flag.Parse()

	if *h || *help {
		flag.Usage()
	}

	orm.Debug = true
	orm.RegisterDataBase("default", "mysql", dsn, 30)

	if *init {
		orm.RunSyncdb("default", *force, true)
		user := models.User{Name: "admin", IsSupper: true}
		password := utils.RandomString(6)
		user.SetPassword(password)

		if err := models.AddUser(&user); err != nil {
			fmt.Println(err)
		} else {
			fmt.Printf("Admin Password: %s\n", password)
		}
	} else {
		fmt.Println("Run Todolist App")
		beego.Run()
	}
}
