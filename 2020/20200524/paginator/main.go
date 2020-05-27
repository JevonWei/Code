package main

import (
	"flag"
	"fmt"
	"os"

	_ "paginator/controllers"
	"paginator/models"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

const (
	dbUser     string = "root"
	dbPassword string = "danran"
	dbHost     string = "127.0.0.1"
	dbPort     int    = 3306
	dbName     string = "paginator"
)

var dsn string = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&loc=Local&parseTime=true",
	dbUser, dbPassword, dbHost, dbPort, dbName)

func main() {
	help := flag.Bool("help", false, "print help")
	h := flag.Bool("h", false, "print help")
	gendata := flag.Int("gendata", -1, "产生随机数据")

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

	switch {
	case *gendata > 0:
		fmt.Println("Start to Gendata datas")
		orm.RunSyncdb("default", true, true)
		models.WriteRandomDataToDB(*gendata)
	default:
		fmt.Println("Run Server")

		beego.Run()
	}
}
