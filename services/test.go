package services

import (
	"awesomeProject/models"
	"github.com/google/uuid"
)

func (h *Handler) GetTests() []models.Test {
	tests := h.db.GetTests()
	return tests
}

func (h *Handler) GetTestsByTeacherId(id uuid.UUID) []models.Test {
	tests := h.db.GetTestsByTeacherId(id)
	return tests
}

func (h *Handler) DeleteTestById(id uuid.UUID) error {
	err := h.db.DeleteTestById(id)
	return err
}

func (h *Handler) GetTestById(id uuid.UUID) (*models.Test, error) {
	var test *models.Test
	var err error
	test, err = h.db.GetTestById(id)
	if err != nil {
		return nil, err
	}
	return test, err
}

func (h *Handler) SaveTest(test *models.Test) error {
	err := h.db.SaveTest(test)
	return err
}
