package models

import (
	"fmt"
	"time"

	"github.com/JevonWei/gocmdb/server/utils"
	"github.com/astaxie/beego/orm"
)

type User struct {
	Id          int        `orm:"column(id)"`
	Name        string     `orm:"column(name);size(32);default();"`      //用户名
	Password    string     `orm:"column(password);size(1024);default();` // 密码
	Birthday    *time.Time `orm:"column(birthday);null; default(null)"`  //出生日期，允许为null
	Gender      int        `orm:"column(gender);"default(0)"`            //性别，true：男，false： 女
	Tel         string     `orm:"column(tel);size(1024);default()"`      //电话号码
	Email       string     `orm:"column(email);size(1024);default()"`    // 邮箱
	Addr        string     `orm:"column(addr);size(1024);default()"`     // 住址
	Remark      string     `orm:"column(remark);size(1024);default()"`   // 备注
	IsSuperman  bool       `orm:"column(is_superman);"default(false)"`   //是否为超级管理员, true:是，false：非
	Status      int        `orm:"column(status);default(0);"`            //状态
	CreatedTime *time.Time `orm:"column(created_time);auto_now_add;"`    // 创建时间，在创建时自动设置（auto_now_add）
	Updatedtime *time.Time `orm:"column(updated_time);auto_now;"`
	Deletedtime *time.Time `orm:"column(deleted_time);null;default(null);"`
	Token       *Token     `orm:"reverse(one);"`
}

func (u *User) SetPassword(password string) {
	u.Password = utils.Md5Salt(password, "")
}

func (u *User) ValidatePassword(password string) bool {
	salt, _ := utils.SplitMd5Salt(u.Password)
	return utils.Md5Salt(password, salt) == u.Password
}

func (u *User) IsLock() bool {
	return u.Status == StatusLock
}

type UserManager struct{}

func NewUserManager() *UserManager {
	return &UserManager{}
}

func (m *UserManager) GetById(id int) *User {
	user := &User{}
	err := orm.NewOrm().QueryTable(user).Filter("Id__exact", id).Filter("DeletedTime__isnull", true).One(user)

	if err == nil {
		return user
	}
	return nil
}

func (m *UserManager) GetByName(name string) *User {
	user := &User{}
	err := orm.NewOrm().QueryTable(user).Filter("Name__exact", name).Filter("DeletedTime__isnull", true).One(user)
	if err == nil {
		return user
	}
	return nil
}

type Token struct {
	Id          int        `orm:"column(id);"`
	User        *User      `orm:"column(user);rel(one);"`
	AccessKey   string     `orm:"column(access_key);size(1024);"`
	SecrectKey  string     `orm:"column(secrect_key);size(1024);"`
	CreatedTime *time.Time `orm:"column(created_time);auto_now_add;"`
	UpdatedTimd *time.Time `orm:"column(updated_time);auto_now;"`
}

type TokenManager struct {
}

func NewTokenManager() *TokenManager {
	return &TokenManager{}
}

func (m *TokenManager) GetByKey(accessKey, secrectKey string) *Token {
	fmt.Println(accessKey, secrectKey)
	token := &Token{AccessKey: accessKey, SecrectKey: secrectKey}
	ormer := orm.NewOrm()
	if err := ormer.Read(token, "AccessKey", "SecrectKey"); err == nil {
		ormer.LoadRelated(token, "User")
		return token
	}
	return nil
}

var DefaultUserManager = NewUserManager()
var DefaultTokenManager = NewTokenManager()

func init() {
	orm.RegisterModel(new(Token), new(User))
}
