package db

import (
	"fmt"

	"github.com/dhofa/gofiber-clean-arch/config"
	"github.com/dhofa/gofiber-clean-arch/internal/entity"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect(cfg *config.Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		cfg.DBHost, cfg.DBUser, cfg.DBPassword, cfg.DBName, cfg.DBPort)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// Auto migrate
	db.AutoMigrate(&entity.User{})

	return db, nil
}
