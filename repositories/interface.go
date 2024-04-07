package repositories

import (
	"time"

	"github.com/jimmyhealer/ad-placement-service/models"
)

type AdRepository interface {
	CreateAd(ad *models.Advertisement) error
	FindActiveAds(
		nowTime time.Time,
		offset int,
		limit int,
		age int,
		gender []models.GenderType,
		country []models.CountryCode,
		platform []models.PlatformType,
	) ([]models.Advertisement, error)
}
