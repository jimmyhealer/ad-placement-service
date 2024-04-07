package main

import (
	"log"

	v1 "github.com/jimmyhealer/ad-placement-service/api/v1"
	"github.com/jimmyhealer/ad-placement-service/db"
	"github.com/jimmyhealer/ad-placement-service/repositories"
	"github.com/joho/godotenv"
)

func main() {

	if err := godotenv.Load(); err != nil {
		log.Printf("Error loading .env file: %v", err)
	}

	newDB, err := db.NewDatabase()

	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	repo := repositories.NewAdRepository(newDB)
	crtl := v1.NewAdController(repo)

	r := setupRouter(crtl)
	r.Run(":8080")
}
