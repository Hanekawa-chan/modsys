package services

import (
	"awesomeProject/models"
	"github.com/google/uuid"
)

func (h *Handler) GetTestByID(id uuid.UUID) (*models.Test, error) {
	var test *models.Test
	var err error
	test, err = h.db.GetTestByID(id)
	if err != nil {
		return nil, err
	}
	return test, err
}
