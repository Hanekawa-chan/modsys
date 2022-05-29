package dao

import (
	"awesomeProject/models"
	"github.com/google/uuid"
)

func (d *DB) GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	tx := d.Preload("Role").First(&user, "email = ?", email)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &user, nil
}

func (d *DB) GetUserByID(id uuid.UUID) (*models.User, error) {
	var user models.User
	tx := d.Preload("Role").First(&user, "id = ?", id)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &user, nil
}

func (d *DB) GetUsers() []models.User {
	var users []models.User
	d.Preload("Role").Find(&users)
	return users
}

func (d *DB) SaveUser(user *models.User) error {
	tx := d.Create(user)
	return tx.Error
}

func (d *DB) GetRoleById(roleId int16) *models.Role {
	var role models.Role
	d.First(&role, roleId)
	return &role
}

func (d *DB) SetRole(user *models.User) error {
	tx := d.Save(user)
	return tx.Error
}
