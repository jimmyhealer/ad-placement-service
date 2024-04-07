package testutils

import (
	"fmt"
	"log"
	"time"

	"github.com/jimmyhealer/ad-placement-service/models"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

func SetUpTestEnv() {
	if err := godotenv.Load("../.env.test"); err != nil {
		log.Fatalf("Failed to load .env file: %v", err)
	}
}

func ClearTestDB(db *gorm.DB) {
	db.Migrator().DropTable(&models.Advertisement{}, &models.Conditions{})
}

func CreateTestAds(db *gorm.DB) {
	ad := models.Advertisement{
		Title:   "Test Ad-01",
		StartAt: time.Now().Add(-time.Hour * 24),
		EndAt:   time.Now().Add(time.Hour * 24),
		Conditions: []models.Conditions{
			{
				AgeStart: 18,
				AgeEnd:   65,
				Gender:   []string{"M", "F"},
				Country:  []string{"TW", "US"},
				Platform: []string{"ios", "android", "web"},
			},
		},
	}

	db.Create(&ad)

	ad = models.Advertisement{
		Title:   "Test Ad-02",
		StartAt: time.Now().Add(-time.Hour * 48),
		EndAt:   time.Now().Add(time.Hour * 48),
		Conditions: []models.Conditions{
			{
				AgeStart: 0,
				AgeEnd:   100,
				Gender:   []string{"M"},
				Country:  []string{"TW", "JP", "US"},
				Platform: []string{"ios", "android", "web"},
			},
		},
	}

	db.Create(&ad)

	for i := 0; i < 10; i++ {
		ad = models.Advertisement{
			Title:   fmt.Sprintf("Test loop-Ad-%02d", i),
			StartAt: time.Now().Add(time.Hour * 4),
			EndAt:   time.Now().Add(time.Hour * 8),
			Conditions: []models.Conditions{
				{
					AgeStart: 77,
					AgeEnd:   77,
					Gender:   []string{"M"},
					Country:  []string{"TW"},
					Platform: []string{"ios"},
				},
			},
		}

		db.Create(&ad)
	}
}
