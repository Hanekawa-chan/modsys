package services

import (
	"awesomeProject/models"
	"time"
)

func (h *Handler) SaveRecord(record models.Record) error {
	record.Date = time.Now()
	err := h.db.AddRecord(record)
	return err
}

func (h *Handler) GetRecords() []models.Record {
	var records []models.Record
	records = h.db.GetRecords()
	return records
}
