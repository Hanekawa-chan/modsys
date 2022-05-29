package models

import "time"

type Record struct {
	Id   uint `gorm:"primaryKey;AUTO_INCREMENT"`
	Text string
	Date time.Time
}
