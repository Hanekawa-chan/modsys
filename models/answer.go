package models

import "github.com/google/uuid"

type Answer struct {
	QuestionId uuid.UUID `gorm:"type:uuid"`
	StudentId  uuid.UUID `gorm:"type:uuid"`
	Answer     string
	Question   Question `gorm:"ForeignKey:QuestionId;References:Id"`
	Student    User     `gorm:"ForeignKey:StudentId;References:Id"`
}
