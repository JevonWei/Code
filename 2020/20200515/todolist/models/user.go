package models

import (
	"time"

	"todolist/utils"

	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Name       string     `gorm:"type:varchar(32); not null; default: ''"`
	Password   string     `gorm:"type:varchar(1024); not null; default: ''"`
	Birthday   time.Time  `gorm: type:date; not null`
	Sex        bool       `gorm:not null; default: false`
	Tel        string     `gorm:"type:varchar(16); not null; default: ''"`
	Addr       string     `gorm:"type:varchar(512); not null; default: ''"`
	Desc       string     `gorm:"column:description; type:varchar(1024); not null; default: ''"`
	CreateTime *time.Time `gorm:"column:create_time; type:datetime"`
}

type User_str struct {
	gorm.Model
	Name       string     `gorm:"type:varchar(32); not null; default:'' "`
	Birthday   string     `gorm:"type:date; not null"`
	Sex        bool       `gorm:"not null; default:false"`
	Addr       string     `gorm:"type:varchar(512); not null; default:'' "`
	Tel        string     `gorm:"type:varchar(16); not null; default:'' "`
	Desc       string     `gorm:"column:description; type:text; not null; default:'' "`
	Password   string     `gorm:"type:varchar(1024); not null; default:'' "`
	CreateTime *time.Time `gorm:"column:create_time; type:datetime "`
}

type Errors struct {
	Name           string
	Password       string
	PasswordVerify string
}

func (u User) ValidatePassword(password string) bool {
	return utils.Md5(password) == u.Password
}

// func loadUsers() (map[int]User, error) {
// 	if bytes, err := ioutil.ReadFile("datas/users.json"); err != nil {
// 		if os.IsNotExist(err) {
// 			return map[int]User{}, nil
// 		}
// 		return nil, err
// 	} else {
// 		var users map[int]User
// 		if err := json.Unmarshal(bytes, &users); err == nil {
// 			return users, nil
// 		} else {
// 			return nil, err
// 		}
// 	}
// }

// func storeUsers(users map[int]User) error {
// 	bytes, err := json.Marshal(users)
// 	if err != nil {
// 		return err
// 	}
// 	return ioutil.WriteFile("datas/users.json", bytes, 0X066)
// }

func GetUsers(q string) []User_str {
	var users []User
	if q == "" {
		db.Find(&users)
	} else {
		db.Where("name like ?", "%"+q+"%").Or("tel like ?", "%"+q+"%").Or("addr like ?", "%"+q+"%").Or("description like ?", "%"+q+"%").Find(&users)
	}
	users_str := make([]User_str, 0)
	var user_str User_str
	for _, user := range users {
		user_str.ID = user.ID
		user_str.Name = user.Name
		user_str.Sex = user.Sex
		user_str.Birthday = user.Birthday.Format("2006-01-02")
		user_str.Addr = user.Addr
		user_str.Tel = user.Tel
		user_str.Desc = user.Desc
		users_str = append(users_str, user_str)
	}

	return users_str
}

func GetUserByName(name string) (User, error) {
	var user User
	err := db.First(&user, "name=?", name).Error
	return user, err
}

func ValidateCreateUser(name, password, birthday, tel, addr, desc string) map[string]string {

	errors := map[string]string{}
	if len(name) > 12 || len(name) < 4 {
		errors["name"] = "名称长度必须在4~12之间"
	} else if _, err := GetUserByName(name); err == nil {
		errors["name"] = "名称重复"
	}
	if len(password) > 30 || len(password) < 6 {
		errors["password"] = "密码长度必须在6~30之间"
	}
	return errors
}

func CreateUser(name, password, birthday, tel, addr, desc string, sex bool) {
	now := time.Now()
	local, _ := time.LoadLocation("Local")
	t, _ := time.ParseInLocation("2006-01-02", birthday, local)
	var user User = User{
		Name:       name,
		Password:   password,
		Birthday:   t,
		Tel:        tel,
		Addr:       addr,
		Desc:       desc,
		Sex:        sex,
		CreateTime: &now,
	}

	if db.NewRecord(user) {
		db.Create(&user)
	}

}
