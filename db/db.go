package db

import (
	"gorm.io/gorm"
)

type DB struct {
	db_conn *gorm.DB
}

func NewDB(db_conn *gorm.DB) *DB {
	return &DB{
		db_conn: db_conn,
	}
}
