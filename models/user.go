package models

import (
	"github.com/google/uuid"
)

type User struct {
	//gorm.Model
	Id       uuid.UUID
	Name     string
	Surname  string
	Email    string
	Password string
	Role     int16
}

func NewUser(name, surname, email, password string) *User {
	return &User{
		Id:       uuid.New(),
		Name:     name,
		Surname:  surname,
		Email:    email,
		Password: password,
		Role:     0,
	}
}

func (u *User) GetID() uuid.UUID {
	return u.Id
}

func (u *User) ToString() string {
	return "name: " + u.Name +
		"\nsurname: " + u.Surname
}
