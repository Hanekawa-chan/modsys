package services

import "awesomeProject/models"

func (h *Handler) SaveAnswers(answers []models.Answer) error {
	err := h.db.SaveAnswers(answers)
	return err
}
