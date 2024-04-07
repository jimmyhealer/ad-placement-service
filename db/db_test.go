package db

import (
	"testing"

	testutils "github.com/jimmyhealer/ad-placement-service/test_utils"
)

func TestGetDB(t *testing.T) {
	testutils.SetUpTestEnv()

	db, err := NewDatabase()

	if err != nil {
		t.Errorf("Failed to create new database: %v", err)
	}

	if db == nil {
		t.Error("New database is nil")
	}

	if db.GetDB() == nil {
		t.Error("Database is nil")
	}
}
