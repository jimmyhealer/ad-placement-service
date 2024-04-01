package db

import (
	"fmt"
	"os"

	"github.com/jimmyhealer/ad-placement-service/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var postgreDB *gorm.DB

func ConnectDatabase() error {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s TimeZone=Asia/Taipei",
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_DB"),
		os.Getenv("POSTGRES_PORT"),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	postgreDB = db

	db.AutoMigrate(&models.Advertisement{})
	db.AutoMigrate(&models.Conditions{})

	return nil
}

func GetDB() *gorm.DB {
	return postgreDB
}
