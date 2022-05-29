package dao

import "awesomeProject/models"

func (d *DB) GetRecords() []models.Record {
	var records []models.Record
	d.Find(&records)
	return records
}

func (d *DB) AddRecord(record models.Record) error {
	err := d.Create(&record).Error
	return err
}
