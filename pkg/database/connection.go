package database

import (
	"log"
	"media-server/internal/config"
	"media-server/pkg/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() *gorm.DB {
	cfg, err := config.LoadDBConfig()
	if err != nil {
		panic(err)
	}

	DB, err = gorm.Open(postgres.Open(cfg.DBUrl), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect to the database:", err)
	}

	DB.AutoMigrate(&model.Media{})
	return DB
}
