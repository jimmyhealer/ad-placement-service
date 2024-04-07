package main

import (
	"log"

	"github.com/joho/godotenv"
)

func main() {

	if err := godotenv.Load(); err != nil {
		log.Printf("Error loading .env file: %v", err)
	}

	crtl, err := InitializeAdController()
	if err != nil {
		log.Fatalf("failed to initialize ad controller: %v", err)
	}

	r := setupRouter(crtl)
	r.Run(":8080")
}
