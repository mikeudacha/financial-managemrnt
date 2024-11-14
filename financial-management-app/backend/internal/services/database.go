package services

import (
	"financial-management-app/backend/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

func ConnectDatabase() {
	dsn := "host=localhost user=db_user password=db_password dbname=db_name port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Error connecting to the database:", err)

	}
	db.AutoMigrate(&models.User{})
	DB = db
}
