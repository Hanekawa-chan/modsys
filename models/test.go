package models

import "github.com/google/uuid"

type Test struct {
	Id        uuid.UUID `gorm:"primaryKey;type:uuid"`
	TeacherId uuid.UUID `gorm:"type:uuid"`
	Name      string
	Questions []Question `gorm:"foreignKey:TestId"`
	Results   []Result   `gorm:"foreignKey:TestId"`
	Teacher   User       `gorm:"ForeignKey:TeacherId;References:Id"`
}

func NewTest(teacherId uuid.UUID, name string) *Test {
	return &Test{
		Id:        uuid.New(),
		TeacherId: teacherId,
		Name:      name,
	}
}
