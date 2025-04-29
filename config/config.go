package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	POSTGRES_USER     string
	POSTGRES_DD_NAME  string
	POSTGRES_PASSWORD string
	POSTGRES_PORT     int64
	POSTGRES_HOST     string
	PORT              string
	DB_URL            string
)

func Init() {

	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file - ", err)
	}

	POSTGRES_USER = os.Getenv("POSTGRES_USER")
	POSTGRES_DD_NAME = os.Getenv("POSTGRES_DB_NAME")
	POSTGRES_PASSWORD = os.Getenv("POSTGRES_PASSWORD")
	POSTGRES_PORT, _ = strconv.ParseInt(os.Getenv("POSTGRES_PORT"), 10, 64)
	POSTGRES_HOST = os.Getenv("POSTGRES_HOST")
	PORT = ":" + os.Getenv("PORT")
	SSL_MODE := os.Getenv("SSL_MODE")

	os.Setenv("DB_URL", fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s",
		POSTGRES_HOST, POSTGRES_USER, POSTGRES_PASSWORD, POSTGRES_DD_NAME, POSTGRES_PORT, SSL_MODE))
		
	DB_URL = os.Getenv("DB_URL")
}
