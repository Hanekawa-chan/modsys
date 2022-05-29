package models

type Role struct {
	Id    int16 `gorm:"primaryKey"`
	Role  string
	Users []User `gorm:"foreignKey:RoleId"`
}
