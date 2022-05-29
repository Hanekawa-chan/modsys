package dao

import (
	"awesomeProject/models"
	"github.com/google/uuid"
)

func (d *DB) GetTestsByTeacherId(teacherId uuid.UUID) []models.Test {
	var tests []models.Test
	d.Find(&tests, d.Where("teacher_id = ?", teacherId))
	return tests
}

func (d *DB) GetTests() []models.Test {
	var tests []models.Test
	d.Find(&tests)
	return tests
}

func (d *DB) GetTestById(id uuid.UUID) (*models.Test, error) {
	var test models.Test
	tx := d.Preload("Questions").First(&test, id)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &test, nil
}

func (d *DB) DeleteTestById(id uuid.UUID) error {
	var test models.Test
	tx := d.Delete(&test, id)
	return tx.Error
}

func (d *DB) SaveTest(test *models.Test) error {
	err := d.Create(test).Error
	return err
}
