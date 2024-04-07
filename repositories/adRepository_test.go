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
	tests := []struct {
		name     string
		input    interface{}
		expected bool
	}{
		{
			name:     "Empty slice",
			input:    []int{},
			expected: false,
		},
		{
			name:     "Non-empty slice",
			input:    []int{1, 2, 3},
			expected: true,
		},
		{
			name:     "Non-slice value",
			input:    10,
			expected: true,
		},
		{
			name:     "Empty array",
			input:    [0]int{},
			expected: false,
		},
		{
			name:     "Non-empty array",
			input:    [3]int{1, 2, 3},
			expected: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := isNonEmptySlice(test.input)
			if result != test.expected {
				t.Errorf("Expected %v, got %v", test.expected, result)
			}
		})
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

	nowTime := time.Now()
	offset := 0
	limit := 5
	t.Run("Test case 1: Find active ads successfully", func(t *testing.T) {
		age := 18

		ads, err := repo.FindActiveAds(nowTime, offset, limit, age, nil, nil, nil)

		if err != nil {
			t.Errorf("Expected nil, got %v", err)
		}

		if len(ads) != 2 {
			t.Errorf("Expected 2 ads, got %d", len(ads))
		}

		for i := range ads {
			if i > 0 && ads[i].EndAt.Before(ads[i-1].EndAt) {
				t.Errorf("Ads are not sorted by end time")
			}
			if ads[i].StartAt.After(nowTime) || ads[i].EndAt.Before(nowTime) {
				t.Errorf("Ad %d is not active", ads[i].ID)
			}
		}
	})

	t.Run("Test case 2: Find active ads with filter", func(t *testing.T) {
		genders := []models.GenderType{"M"}
		contries := []models.CountryCode{}
		platforms := []models.PlatformType{}
		ads, err := repo.FindActiveAds(nowTime, offset, limit, 0, genders, contries, platforms)

		if err != nil {
			t.Errorf("Expected nil, got %v", err)
		}

		if len(ads) != 2 {
			t.Errorf("Expected 2 ads, got %d", len(ads))
		}
	})

	t.Run("Test case 3: Find no active ads with filter", func(t *testing.T) {
		age := 18
		genders := []models.GenderType{"F"}
		contries := []models.CountryCode{"JP"}
		platforms := []models.PlatformType{"web"}
		ads, err := repo.FindActiveAds(nowTime, offset, limit, age, genders, contries, platforms)

		if err != nil {
			t.Errorf("Expected nil, got %v", err)
		}

		if len(ads) != 0 {
			t.Errorf("Expected 0 ads, got %d", len(ads))
		}
	})

	t.Run("Test case 4: Find no active ads with time filter", func(t *testing.T) {
		nowTime = time.Now().Add(-time.Hour * 96)
		ads, err := repo.FindActiveAds(nowTime, offset, limit, 18, nil, nil, nil)

		if err != nil {
			t.Errorf("Expected nil, got %v", err)
		}

		if len(ads) != 0 {
			t.Errorf("Expected 0 ads, got %d", len(ads))
		}
	})

	t.Run("Test case 5: Check limit and offset", func(t *testing.T) {
		nowTime = time.Now().Add(time.Hour * 5)
		offset = 3
		limit = 3
		ads, err := repo.FindActiveAds(nowTime, offset, limit, 77, nil, nil, nil)

		if err != nil {
			t.Errorf("Expected nil, got %v", err)
		}

		if len(ads) != limit {
			t.Errorf("Expected 3 ads, got %d", len(ads))
		}

		for i := range ads {
			if ads[i].Title != fmt.Sprintf("Test loop-Ad-%02d", i+offset) {
				t.Errorf("Ad %s is not correct", ads[i].Title)
			}
		}
	})
}
