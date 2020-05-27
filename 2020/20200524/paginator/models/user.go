package models

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/astaxie/beego/orm"
)

type User struct {
	Id       int
	Name     string `orm:"unique"`
	Age      uint32
	Birthday time.Time `orm:"type(date)"`
}

func generateRandomTime() time.Time {
	max := time.Now().Unix()
	second := rand.Int63() % int64(max)

	return time.Unix(second, 0)
}

func generateRandomData(count int) []User {
	users := make([]User, count)

	for index, _ := range users {
		users[index].Name = fmt.Sprintf("Name_%d", index)
		users[index].Age = rand.Uint32() % 90
		users[index].Birthday = generateRandomTime()
	}

	return users
}

func WriteRandomDataToDB(count int) int {
	users := generateRandomData(count)
	o := orm.NewOrm()
	o.InsertMulti(100, users)

	return count
}

func GetAllUsers() (int64, []User) {
	o := orm.NewOrm()
	qs := o.QueryTable("user")
	users := []User{}
	if count, err := qs.OrderBy("Id").All(&users); err == nil {
		return count, users
	} else {
		return 0, users
	}

}

func init() {
	orm.RegisterModel(new(User))
}
