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

func (d *DB) GetTeachers() []models.User {
	var users []models.User
	d.Find(&users, "role_id = ?", 1)
	return users
}

func (d *DB) GetTeachersByStudentId(studentId uuid.UUID) ([]models.User, error) {
	var teachers []models.User
	var student models.User
	err := d.First(&student, studentId).Error
	if err != nil {
		return teachers, err
	}
	err = d.Model(&student).Association("Teachers").Find(&teachers)
	return teachers, err
}

func (d *DB) SaveUser(user *models.User) error {
	tx := d.Create(user)
	return tx.Error
}

func (d *DB) AddTeacher(studentId, teacherId uuid.UUID) error {
	var student models.User
	err := d.First(&student, studentId).Error
	if err != nil {
		return err
	}
	var teacher models.User
	err = d.First(&teacher, teacherId).Error
	if err != nil {
		return err
	}
	err = d.Model(&student).Association("Teachers").Append(&teacher)
	return err
}

func (d *DB) DeleteTeacher(studentId, teacherId uuid.UUID) error {
	var student models.User
	err := d.First(&student, studentId).Error
	if err != nil {
		return err
	}
	var teacher models.User
	err = d.First(&teacher, teacherId).Error
	if err != nil {
		return err
	}
	err = d.Model(&student).Association("Teachers").Delete(&teacher)
	return err
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
