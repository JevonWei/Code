package users

type address struct {
	Region string
	street string
	no     string
}

type User struct {
	ID       int
	Name     string
	brithday string
	addr     address
}
