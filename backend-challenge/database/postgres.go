package database

import (
	"fmt"
	"github.com/amarsinghrathour/oolio_backend_challenge/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

type DBInstance struct {
	DB *gorm.DB
}

func NewDatabase(cfg *config.Config) (*DBInstance, error) {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		cfg.DBHost, cfg.DBUser, cfg.DBPassword, cfg.DBName, cfg.DBPort,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	log.Println("Connected to PostgreSQL!")

	return &DBInstance{DB: db}, nil
}
