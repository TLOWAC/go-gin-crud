package db

import (
	"example/web-service-gin/api/models"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init() *gorm.DB {
	dbURL := "postgresql://postgres:root12#@localhost:5432/go-local"

	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	db.Migrator().AutoMigrate(&models.Book{})

	return db
}
