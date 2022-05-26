package dao

import "awesomeProject/models"

func (d *DB) Migrate() error {
	err := d.AutoMigrate(&models.User{})
	if err != nil {
		return err
	}
	err = d.AutoMigrate(&models.Test{})
	if err != nil {
		return err
	}
	err = d.AutoMigrate(&models.Question{})
	if err != nil {
		return err
	}
	err = d.AutoMigrate(&models.Answer{})
	if err != nil {
		return err
	}
	err = d.AutoMigrate(&models.Result{})
	if err != nil {
		return err
	}
	return nil
}
