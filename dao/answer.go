package dao

import (
	"awesomeProject/models"
	"github.com/google/uuid"
)

func (d *DB) SaveAnswers(answers []models.Answer) error {
	err := d.Create(answers).Error
	return err
}

func (d *DB) GetAnswersByResultId(resultId uuid.UUID) []models.Answer {
	var answers []models.Answer
	d.Debug().Find(&answers, "result_id = ?", resultId)
	return answers
}

func (d *DB) GetQuestionById(id uuid.UUID) models.Question {
	var question models.Question
	d.First(&question, id)
	return question
}
