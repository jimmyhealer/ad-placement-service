package db

import (
	"log"
	"os"
	"testing"

	"github.com/joho/godotenv"
)

func TestMain(m *testing.M) {
	err := godotenv.Load("../.env.test")
	if err != nil {
		log.Fatalln("Error loading .env file")
	}

	os.Exit(m.Run())
}

func TestConnectDatabase(t *testing.T) {
	if testDB, err := ConnectDatabase(); err != nil {
		t.Errorf("Failed to connect to database: %v", err)
	} else {
		if testDB == nil {
			t.Error("Database connection is nil")
		}
	}
}
