package view

import "github.com/google/uuid"

type Teacher struct {
	Id      uuid.UUID
	Name    string
	Surname string
	Email   string
	IsAdded bool
}
