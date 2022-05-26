package services

import (
	"awesomeProject/models"
	"github.com/google/uuid"
)

func (h *Handler) GetResults(userId uuid.UUID) []models.Result {
	results := h.db.GetResults(userId)
	return results
}

func (h *Handler) GetResultById(id uuid.UUID) models.Result {
	result := h.db.GetResultById(id)
	return result
}

func (h *Handler) SaveResult(result models.Result) error {
	err := h.db.SaveResult(result)
	return err
}
