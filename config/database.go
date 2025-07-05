package config

import (
	"log"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

type Database struct {
	DBUser     string
	DBPassword string
	DBName     string
	DBHost     string
	DBPort     string
}

func LoadDatabaseConfig() *Database {
	dir, _ := os.Getwd()
	err := godotenv.Load(filepath.Join(dir, ".env"))
	if err != nil {
		log.Fatal("Error loading .env file => ", err)
	}

	return &Database{
		DBUser:     os.Getenv("DB_USER"),
		DBPassword: os.Getenv("DB_PASSWORD"),
		DBName:     os.Getenv("DB_NAME"),
		DBHost:     os.Getenv("DB_HOST"),
		DBPort:     os.Getenv("DB_PORT"),
	}
}
