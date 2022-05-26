package models

import "github.com/google/uuid"

type Result struct {
	Id        uuid.UUID `gorm:"primaryKey;type:uuid"`
	TestId    uuid.UUID `gorm:"type:uuid"`
	Score     int16
	StudentId uuid.UUID `gorm:"type:uuid"`
	Student   User      `gorm:"ForeignKey:StudentId;References:Id"`
	Test      Test      `gorm:"ForeignKey:TestId;References:Id"`
	Answers   []Answer  `gorm:"foreignKey:ResultId"`
}
