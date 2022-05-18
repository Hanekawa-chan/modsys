package models

import "github.com/google/uuid"

type Question struct {
	Id       uuid.UUID
	TestId   uuid.UUID
	Question string
	Answer   string
}

func NewQuestion(testId uuid.UUID, question, answer string) *Question {
	return &Question{
		Id:       uuid.New(),
		TestId:   testId,
		Question: question,
		Answer:   answer,
	}
}
