package models

import "github.com/google/uuid"

type ResultView struct {
	Name   string
	Author string
	Score  int16
	Id     uuid.UUID
}
