package view

import "awesomeProject/models"

type ResultFull struct {
	Name    string
	Author  string
	Score   int16
	Answers []models.Answer
}
