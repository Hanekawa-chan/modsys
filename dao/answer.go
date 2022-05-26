package dao

import "awesomeProject/models"

func (d *DB) SaveAnswers(answers []models.Answer) error {
	err := d.Create(answers).Error
	return err
}
