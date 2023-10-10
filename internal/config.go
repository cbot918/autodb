package internal

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	IMAGE     string
	CONTAINER string
	SQL_FILE  string

	DB_DRIVER     string
	DB_USER       string
	DB_PASSWORD   string
	DB_HOST       string
	DB_PORT       string
	DB_NAME       string
	DB_TABLES     string
	TYPE_PKG_NAME string
}

func NewConfig() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}

	return &Config{
		IMAGE:         os.Getenv("IMAGE"),
		CONTAINER:     os.Getenv("CONTAINER"),
		SQL_FILE:      os.Getenv("SQL_FILE"),
		DB_DRIVER:     os.Getenv("DB_DRIVER"),
		DB_USER:       os.Getenv("DB_USER"),
		DB_PASSWORD:   os.Getenv("DB_PASSWORD"),
		DB_HOST:       os.Getenv("DB_HOST"),
		DB_PORT:       os.Getenv("DB_PORT"),
		DB_NAME:       os.Getenv("DB_NAME"),
		DB_TABLES:     os.Getenv("DB_TABLES"),
		TYPE_PKG_NAME: os.Getenv("TYPE_PKG_NAME"),
	}, nil
}
