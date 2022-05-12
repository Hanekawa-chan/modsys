package dao

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const dsn = "host=abul.db.elephantsql.com user=idjbfebr password=ONurXwA2nXGWBnhcMmYNiGk75zMBawJa port=5432 database=idjbfebr"

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
