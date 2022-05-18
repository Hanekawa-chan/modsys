package dao

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

type DB struct {
	*gorm.DB
}

func New() (*DB, error) {
	gormDB, err := gorm.Open(postgres.Open(os.Getenv("ELEPHANT")))
	if err != nil {
		return nil, err
	}
	return &DB{gormDB}, nil
}
