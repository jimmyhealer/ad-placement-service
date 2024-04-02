package main

import (
	"log"

	"github.com/jimmyhealer/ad-placement-service/db"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Printf("Error loading .env file: %v", err)
	}

	if _, err := db.ConnectDatabase(); err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	r := setupRouter()
	r.Run(":8080")
}
