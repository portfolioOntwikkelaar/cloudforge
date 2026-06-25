package database

import (
	"cloudforge/configs"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() error {

	cfg := configs.Load()

	db, err := gorm.Open(postgres.Open(cfg.DB), &gorm.Config{})
	if err != nil {
		return err
	}

	DB = db

	return nil
}
