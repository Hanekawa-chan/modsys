package models

import "strconv"

type User struct {
	name    string
	surname string
	number  int
}

func NewUser(name, surname string, number int) *User {
	return &User{
		name:    name,
		surname: surname,
		number:  number,
	}
}

func (u *User) ToString() string {
	return "name: " + u.name +
		"\nsurname: " + u.surname +
		"\nnumber: " + strconv.Itoa(u.number)
}