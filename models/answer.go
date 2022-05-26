package models

import "github.com/google/uuid"

type Answer struct {
	QuestionId uuid.UUID `gorm:"type:uuid"`
	ResultId   uuid.UUID `gorm:"type:uuid"`
	Answer     string
	Question   Question `gorm:"ForeignKey:QuestionId;References:Id"`
	Result     Result   `gorm:"ForeignKey:ResultId;References:Id"`
}
