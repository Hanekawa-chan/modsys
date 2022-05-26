package dao

import (
	"awesomeProject/models"
	"github.com/google/uuid"
)

func (d *DB) GetResults(userId uuid.UUID) []models.Result {
	var results []models.Result
	d.First(&results, "student_id = ?", userId)
	return results
}

func (d *DB) GetResultById(id uuid.UUID) models.Result {
	var result models.Result
	d.First(&result, id)
	return result
}

func (d *DB) SaveResult(result models.Result) error {
	err := d.Create(result).Error
	return err
}
