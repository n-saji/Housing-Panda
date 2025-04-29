package service

import (
	"housing_panda/db"

	"gorm.io/gorm"
)

type Service struct {
	db *db.DB
}

func NewService(db_conn *gorm.DB) *Service {
	return &Service{
		db: db.NewDB(db_conn),
	}
}
