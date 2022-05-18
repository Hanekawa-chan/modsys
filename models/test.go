package models

import "github.com/google/uuid"

type Test struct {
	Id        uuid.UUID
	TeacherId uuid.UUID
	Name      string
	Questions []Question
}

func NewTest(teacherId uuid.UUID, name string) *Test {
	return &Test{
		Id:        uuid.New(),
		TeacherId: teacherId,
		Name:      name,
	}
}
