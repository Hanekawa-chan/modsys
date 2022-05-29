package models

import (
	"github.com/google/uuid"
)

type User struct {
	Id       uuid.UUID `gorm:"primaryKey;type:uuid"`
	Name     string
	Surname  string
	Email    string
	Password string
	RoleId   int16
	Role     Role     `gorm:"ForeignKey:RoleId;References:Id"`
	Tests    []Test   `gorm:"foreignKey:TeacherId"`
	Results  []Result `gorm:"foreignKey:StudentId"`
	Teachers []User   `gorm:"many2many:students_teachers"`
}

func NewUser(name, surname, email, password string) *User {
	return &User{
		Id:       uuid.New(),
		Name:     name,
		Surname:  surname,
		Email:    email,
		Password: password,
		RoleId:   3,
	}
}
