package dao

import "awesomeProject/models"

func (d *DB) Migrate() error {
	err := d.AutoMigrate(&models.User{}, &models.Test{}, &models.Question{},
		&models.Answer{}, &models.Result{}, &models.Record{})
	return err
}
