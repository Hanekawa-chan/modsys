package services

import (
	"awesomeProject/models"
	"github.com/google/uuid"
)

func (h *Handler) SaveAnswers(answers []models.Answer) error {
	err := h.db.SaveAnswers(answers)
	return err
}

func (h *Handler) GetAnswersByResultId(resultId uuid.UUID) []models.Answer {
	answers := h.db.GetAnswersByResultId(resultId)
	return answers
}

func (h *Handler) GetQuestionById(id uuid.UUID) models.Question {
	question := h.db.GetQuestionById(id)
	return question
}
