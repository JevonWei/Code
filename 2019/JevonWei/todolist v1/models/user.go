package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
)

type User struct {
	ID       int    `json:id`
	Name     string `json:name`
	Birthday string `json:birthday`
	Addr     string `json:addr`
	Tel      string `json:tel`
	Desc     string `json:desc`
}

type Account struct {
	AccountName string
	Passwd      string
}

func loadAccount() ([]Account, error) {
	if bytes, err := ioutil.ReadFile("datas/accounts.json"); err != nil {
		if os.IsNotExist(err) {
			return []Account{}, nil
		}
		return nil, err
	} else {
		var accounts []Account
		if string(bytes) == "" {
			return []Account{}, nil
		}
		if err := json.Unmarshal(bytes, &accounts); err == nil {
			return accounts, nil
		} else {
			return nil, err
		}
	}
}

func Storepasswd(accounts []Account) error {
	bytes, err := json.Marshal(accounts)
	if err != nil {
		return err
	}
	return ioutil.WriteFile("datas/accounts.json", bytes, 0X066)
}

func GetAccountByname(name string) (Account, error) {
	accounts, err := loadAccount()
	if err != nil {
		panic(err)
	}
	for _, account := range accounts {
		if name == account.AccountName {
			return account, nil
		}
	}
	return Account{}, errors.New("Account Not Found")
}

func ModifyPasswd(name, passwd string) {
	accounts, err := loadAccount()
	if err != nil {
		panic(err)
	}

	newAccounts := make([]Account, len(accounts))
	for i, account := range accounts {
		if name == account.AccountName {
			account.AccountName = name
			account.Passwd = passwd
		}
		newAccounts[i] = account
	}
	Storepasswd(newAccounts)
}

func AccountCreate(name, passwd string) error {
	accounts, err := loadAccount()
	if err != nil {
		fmt.Println(err)
	}
	account := Account{
		AccountName: name,
		Passwd:      passwd,
	}

	for _, account := range accounts {
		if account.AccountName == name {
			return errors.New("Account Exist")
		}
	}

	accounts = append(accounts, account)
	Storepasswd(accounts)
	return nil

}

func GetAccount() []Account {
	accounts, err := loadAccount()
	fmt.Println(accounts, err)
	if err == nil {
		fmt.Printf("%T\n", accounts)
		return accounts
	}
	panic(err)
}

func AccountVerify(name, passwd string) (bool, error) {
	accounts := GetAccount()

	for _, account := range accounts {
		if account.AccountName == name {
			if account.Passwd == passwd {
				return true, nil
			} else {
				return false, errors.New("Verity Failed")
			}
		}
	}

	return false, errors.New("Account not Exist")
}

func loadUsers() ([]User, error) {
	if bytes, err := ioutil.ReadFile("datas/user.json"); err != nil {
		if os.IsNotExist(err) {
			return []User{}, nil
		}
		return nil, err
	} else {
		var users []User
		if err := json.Unmarshal(bytes, &users); err == nil {
			return users, nil
		} else {
			return nil, err
		}
	}
}

func storeUsers(users []User) error {
	bytes, err := json.Marshal(users)
	if err != nil {
		return err
	}
	return ioutil.WriteFile("datas/user.json", bytes, 0X066)
}

func GetUsers() []User {
	users, err := loadUsers()
	fmt.Println(users, err)
	if err == nil {
		return users
	}
	panic(err)
}

func GetUserId() (int, error) {
	users, err := loadUsers()
	if err != nil {
		return -1, err
	}
	var id int
	for _, user := range users {
		if id < user.ID {
			id = user.ID
		}
	}
	return id + 1, nil
}

func CreateUser(name, birthday, addr, tel, desc string) {
	id, err := GetUserId()
	if err != nil {
		panic(err)
	}
	user := User{
		ID:       id,
		Name:     name,
		Birthday: birthday,
		Desc:     desc,
		Addr:     addr,
		Tel:      tel,
	}
	users, err := loadUsers()
	if err != nil {
		panic(err)
	}
	users = append(users, user)
	storeUsers(users)
}

func GetUserById(id int) (User, error) {
	users, err := loadUsers()
	if err != nil {
		panic(err)
	}
	for _, user := range users {
		if id == user.ID {
			return user, nil
		}
	}

	return User{}, errors.New("Not Found")
}

func ModifyUser(id int, name, birthday, addr, tel, desc string) {
	users, err := loadUsers()
	if err != nil {
		panic(err)
	}

	newUsers := make([]User, len(users))
	for i, user := range users {
		if id == user.ID {
			user.Name = name
			user.Desc = desc
			user.Birthday = birthday
			user.Addr = addr
			user.Tel = tel
		}
		newUsers[i] = user
	}
	storeUsers(newUsers)
}

func DeleteUser(id int) {
	users, err := loadUsers()
	if err != nil {
		panic(err)
	}

	newUsers := make([]User, 0)
	for _, user := range users {
		if id != user.ID {
			newUsers = append(newUsers, user)
		} else {
			fmt.Println(user)
		}
	}
	fmt.Println(storeUsers(newUsers))
}
