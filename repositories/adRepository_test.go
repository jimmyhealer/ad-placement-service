package repositories

import (
	"fmt"
	"testing"
	"time"

	"github.com/jimmyhealer/ad-placement-service/db"
	"github.com/jimmyhealer/ad-placement-service/models"
	testutils "github.com/jimmyhealer/ad-placement-service/test_utils"
)

func TestMain(m *testing.M) {
	testutils.SetUpTestEnv()
	m.Run()
}

func TestIsNonEmptySlice(t *testing.T) {
	// Test case 1: Empty slice
	if isNonEmptySlice([]int{}) {
		t.Errorf("Test case 1 Empty slice: Expected false, got true")
	}

	// Test case 2: Non-empty slice
	if !isNonEmptySlice([]int{1, 2, 3}) {
		t.Errorf("Test case 2 Non-empty slice: Expected true, got false")
	}

	// Test case 3: Non-slice value
	if !isNonEmptySlice(10) {
		t.Errorf("Test case 3 Non-slice value: Expected true, got false")
	}

	// Test case 4: Empty array
	if isNonEmptySlice([0]int{}) {
		t.Errorf("Test case 4 Empty array: Expected false, got true")
	}

	// Test case 5: Non-empty array
	if !isNonEmptySlice([3]int{1, 2, 3}) {
		t.Errorf("Test case 5 Non-empty array: Expected true, got false")
	}
}

func TestCreateAd(t *testing.T) {
	db, err := db.NewDatabase()
	if err != nil {
		t.Errorf("Failed to create new database: %v", err)
	}
	defer testutils.ClearTestDB(db.GetDB())

	repo := NewAdRepository(db)

	// Create ad successfully
	ad := &models.Advertisement{
		Title:   "Test Ad-01",
		StartAt: time.Now(),
		EndAt:   time.Now().Add(time.Hour * 24),
		Conditions: []models.Conditions{
			{
				AgeStart: 18,
				AgeEnd:   65,
			},
		},
	}

	err = repo.CreateAd(ad)
	if err != nil {
		t.Errorf("Expected nil, got %v", err)
	}

	if ad.ID == 0 {
		t.Error("Ad ID not set")
	}
}

func TestFindActiveAds(t *testing.T) {
	db, err := db.NewDatabase()
	if err != nil {
		t.Errorf("Failed to create new database: %v", err)
	}
	defer testutils.ClearTestDB(db.GetDB())
	testutils.CreateTestAds(db.GetDB())

	repo := NewAdRepository(db)

	// Test case 1: Find active ads successfully
	nowTime := time.Now()
	offset := 0
	limit := 5
	age := 18

	ads, err := repo.FindActiveAds(nowTime, offset, limit, age, nil, nil, nil)

	if err != nil {
		t.Errorf("Test case 1: Expected nil, got %v", err)
	}

	if len(ads) != 2 {
		t.Errorf("Test case 1: Expected 2 ads, got %d", len(ads))
	}

	for i := range ads {
		if i > 0 && ads[i].EndAt.Before(ads[i-1].EndAt) {
			t.Errorf("Test case 1: Ads are not sorted by end time")
		}
		if ads[i].StartAt.After(nowTime) || ads[i].EndAt.Before(nowTime) {
			t.Errorf("Test case 1: Ad %d is not active", ads[i].ID)
		}
	}

	// Test case 2: Find active ads with filter
	age = 0
	genders := []models.GenderType{"M"}
	contries := []models.CountryCode{}
	platforms := []models.PlatformType{}

	ads, err = repo.FindActiveAds(nowTime, offset, limit, age, genders, contries, platforms)

	if err != nil {
		t.Errorf("Test case 2: Expected nil, got %v", err)
	}

	if len(ads) != 2 {
		t.Errorf("Test case 2: Expected 2 ad, got %d", len(ads))
	}

	// Test case 3: Find no active ads with filter
	genders = []models.GenderType{"F"}
	contries = []models.CountryCode{"JP"}
	platforms = []models.PlatformType{"web"}

	ads, err = repo.FindActiveAds(nowTime, offset, limit, age, genders, contries, platforms)

	if err != nil {
		t.Errorf("Test case 3: Expected nil, got %v", err)
	}

	if len(ads) != 0 {
		t.Errorf("Test case 3: Expected 0 ad, got %d", len(ads))
	}

	// Test case 4: Find no active ads with time filter
	nowTime = time.Now().Add(-time.Hour * 96)
	age = 18

	ads, err = repo.FindActiveAds(nowTime, offset, limit, age, nil, nil, nil)

	if err != nil {
		t.Errorf("Test case 4: Expected nil, got %v", err)
	}

	if len(ads) != 0 {
		t.Errorf("Test case 4: Expected 0 ad, got %d", len(ads))
	}

	// Test case 5: Check limit and offset
	nowTime = time.Now().Add(time.Hour * 5)
	offset = 6
	limit = 3
	age = 77

	ads, err = repo.FindActiveAds(nowTime, offset, limit, age, nil, nil, nil)

	if err != nil {
		t.Errorf("Test case 5: Expected nil, got %v", err)
	}

	if len(ads) != 3 {
		t.Errorf("Test case 5: Expected 3 ad, got %d", len(ads))
	}

	for i := range ads {
		if ads[i].Title != fmt.Sprintf("Test loop-Ad-%02d", i+offset) {
			t.Errorf("Test case 4: Ad %s is not correct", ads[i].Title)
		}
	}
}
