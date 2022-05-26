package dao

import (
	"awesomeProject/models"
	"github.com/google/uuid"
)

func (d *DB) GetTests() []models.Test {
	var tests []models.Test
	d.First(&tests)
	return tests
}

func (d *DB) GetTestByID(id uuid.UUID) (*models.Test, error) {
	var test models.Test
	tx := d.Preload("Questions").First(&test, id)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &test, nil
}

func (d *DB) SaveTest(test *models.Test) error {
	err := d.Create(test).Error
	return err
}
