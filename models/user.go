package models

import (
	"github.com/google/uuid"
)

type User struct {
	//gorm.Model
	Id       uuid.UUID `gorm:"primaryKey;type:uuid"`
	Name     string
	Surname  string
	Email    string
	Password string
	Role     int16
	Tests    []Test   `gorm:"foreignKey:TeacherId"`
	Answers  []Answer `gorm:"foreignKey:StudentId"`
	Results  []Result `gorm:"foreignKey:StudentId"`
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
