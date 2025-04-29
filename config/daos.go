package config

import (
	"embed"
	"log"

	_ "github.com/lib/pq"
	"github.com/pressly/goose/v3"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

//go:embed migrations/*.sql
var embedMigrations embed.FS

func ConnectDB() *gorm.DB {
	db, err := gorm.Open(postgres.Open(DB_URL), &gorm.Config{})
	if err != nil {
		log.Println("Found err while connecting to database", err)
		return nil
	}
	return db
}

func CloseDB(db *gorm.DB) {
	if db == nil {
		log.Println("Cannot close a nil database connection")
		return
	}
	dbSQL, err := db.DB()
	if err != nil {
		log.Println("Found err while closing the database", err)
	}
	dbSQL.Close()
}

func RunGooseMigration(db *gorm.DB) {
	goose.SetBaseFS(embedMigrations)
	if err := goose.SetDialect("postgres"); err != nil {
		log.Println("Setting Goose Postgres Dialect Failed")
		panic(err)
	}
	sql_db, err := db.DB()
	if err != nil {
		log.Println("Found err while getting sql db from gorm db", err)
		return
	}
	if err := goose.Up(sql_db, "migrations", goose.WithAllowMissing()); err != nil {
		log.Println("Goose Up Failed")
		panic(err)
	}
}
