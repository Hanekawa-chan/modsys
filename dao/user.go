package dao

import (
	"awesomeProject/models"
	"github.com/google/uuid"
)

func (d *DB) GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	tx := d.First(&user, "email = ?", email)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &user, nil
}

func (d *DB) GetUserByID(id uuid.UUID) (*models.User, error) {
	var user models.User
	tx := d.First(&user, "id = ?", id)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &user, nil
}

func (d *DB) SaveUser(user *models.User) error {
	tx := d.Create(user)
	return tx.Error
}

func (d *DB) SetRole(user *models.User, role int16) error {
	tx := d.Model(user).Update("role", role)
	return tx.Error
}
