package repositories

import (
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/jimmyhealer/ad-placement-service/db"
	"github.com/jimmyhealer/ad-placement-service/models"
	"gorm.io/gorm"
)

type ConcreteAdRepository struct {
	db db.Database
}

func NewAdRepository(db db.Database) AdRepository {
	return &ConcreteAdRepository{
		db: db,
	}
}

func isNonEmptySlice(value interface{}) bool {
	v := reflect.ValueOf(value)
	if v.Kind() == reflect.Slice || v.Kind() == reflect.Array {
		return v.Len() > 0
	}
	return true
}

func addCondition(query *gorm.DB, condition string, value interface{}, format string) *gorm.DB {
	if value != nil && value != "" && value != 0 && isNonEmptySlice(value) {
		return query.Where(condition, fmt.Sprintf(format, value))
	}
	return query
}

func removeBrackets(value []interface{}) string {
	strSlice := make([]string, len(value))
	for i, v := range value {
		strSlice[i] = fmt.Sprint(v)
	}
	return strings.Trim(strings.Join(strSlice, "\",\""), "[]")
}

func (repo *ConcreteAdRepository) CreateAd(ad *models.Advertisement) error {
	db := repo.db.GetDB()

	if err := db.Create(&ad).Error; err != nil {
		return err
	}

	return nil
}

func (repo *ConcreteAdRepository) FindActiveAds(
	nowTime time.Time,
	offset int,
	limit int,
	age int,
	genders []models.GenderType,
	countries []models.CountryCode,
	platforms []models.PlatformType,
) ([]models.Advertisement, error) {
	genderInterfaces := make([]interface{}, len(genders))
	platformInterfaces := make([]interface{}, len(platforms))
	countryInterfaces := make([]interface{}, len(countries))
	for i, v := range genders {
		genderInterfaces[i] = v
	}
	for i, v := range platforms {
		platformInterfaces[i] = v
	}
	for i, v := range countries {
		countryInterfaces[i] = v
	}

	db := repo.db.GetDB()

	var ads []models.Advertisement

	query := db.Debug().Joins("JOIN Conditions on Advertisements.ID = Conditions.advertisement_id")

	nowTimeStr := nowTime.UTC().Format("2006-01-02 15:04:05+00")
	query = addCondition(query, "start_at <= ?", nowTimeStr, "%s")
	query = addCondition(query, "end_at >= ?", nowTimeStr, "%s")
	query = addCondition(query, "age_start <= ?", age, "%d")
	query = addCondition(query, "age_end >= ?", age, "%d")
	query = addCondition(query, "gender @> ?", removeBrackets(genderInterfaces), `{"%s"}`)
	query = addCondition(query, "country @> ?", removeBrackets(countryInterfaces), `{"%s"}`)
	query = addCondition(query, "platform @> ?", removeBrackets(platformInterfaces), `{"%s"}`)

	if err := query.Order("end_at ASC").Offset(offset).Limit(limit).Find(&ads).Error; err != nil {
		return nil, err
	}

	return ads, nil
}
