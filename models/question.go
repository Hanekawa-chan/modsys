package models

import "github.com/google/uuid"

type Question struct {
	Id       uuid.UUID `gorm:"primaryKey;type:uuid"`
	TestId   uuid.UUID `gorm:"type:uuid"`
	Question string
	Answer   string
	Score    int
	Answers  []Answer `gorm:"foreignKey:QuestionId"`
	Student  Test     `gorm:"ForeignKey:TestId;References:Id"`
}

func NewQuestion(testId uuid.UUID, question, answer string, score int) *Question {
	return &Question{
		Id:       uuid.New(),
		TestId:   testId,
		Question: question,
		Answer:   answer,
		Score:    score,
	}
}
