package dao

import (
	"awesomeProject/models"
	"github.com/google/uuid"
)

func (d *DB) GetTestByID(id uuid.UUID) (*models.Test, error) {
	var test models.Test
	tx := d.First(&test, "id = ?", id)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &test, nil
}

func (d *DB) SaveTest(test *models.Test) error {
	tx := d.Create(test)
	return tx.Error
}
