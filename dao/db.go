package dao

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const dsn = os.Getenv("ELEPHANT")

type DB struct {
	*gorm.DB
}

func New() (*DB, error) {
	gormDB, err := gorm.Open(postgres.Open(dsn))
	if err != nil {
		return nil, err
	}
	return &DB{gormDB}, nil
}
