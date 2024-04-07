package v1

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/jimmyhealer/ad-placement-service/db"
	"github.com/jimmyhealer/ad-placement-service/repositories"
	testutils "github.com/jimmyhealer/ad-placement-service/test_utils"
	"github.com/stretchr/testify/assert"
)

var (
	DB   db.Database
	repo repositories.AdRepository
	ctrl *AdController
)

func TestMain(m *testing.M) {
	testutils.SetUpTestEnv("../../.env.test")

	DB, err := db.NewDatabase()
	if err != nil {
		panic(err)
	}

	repo = repositories.NewAdRepository(DB)
	ctrl = NewAdController(repo)

	m.Run()
}

func TestCreateAd(t *testing.T) {
	router := gin.Default()

	router.POST("/ad", ctrl.CreateAd)
	reqBody := map[string]interface{}{
		"title":   "Test Ad-01",
		"startAt": "2022-01-01T00:00:00Z",
		"endAt":   "2022-01-02T00:00:00Z",
		"conditions": []map[string]interface{}{
			{
				"ageStart": 18,
				"ageEnd":   30,
				"gender":   []string{"M", "F"},
				"country":  []string{"TW", "JP"},
				"platform": []string{"android", "ios", "web"},
			},
		},
	}

	reqBodyJSON, _ := json.Marshal(reqBody)
	req, _ := http.NewRequest("POST", "/ad", bytes.NewBuffer(reqBodyJSON))
	req.Header.Set("Content-Type", "application/json")

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, req)

	assert.Equal(t, http.StatusOK, recorder.Code)

	var responseBody map[string]interface{}
	json.Unmarshal(recorder.Body.Bytes(), &responseBody)

	assert.Equal(t, "Ad created successfully", responseBody["message"])
}
func TestGetAds(t *testing.T) {
	router := gin.Default()

	router.GET("/ads", ctrl.GetAds)

	t.Run("Should return bad request if query parameters are invalid", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/ads?offset=0&limit=101&age=0&gender=Invalid&country=Invalid&platform=Invalid", nil)

		recorder := httptest.NewRecorder()
		router.ServeHTTP(recorder, req)

		assert.Equal(t, http.StatusBadRequest, recorder.Code)

		var responseBody map[string]interface{}
		json.Unmarshal(recorder.Body.Bytes(), &responseBody)

		assert.Contains(t, responseBody["error"], "gender must be one of M, F")
		assert.Contains(t, responseBody["error"], "country must be one of TW, JP")
		assert.Contains(t, responseBody["error"], "platform must be one of android, ios, web")
	})

	t.Run("Should return ads if query parameters are valid", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/ads?offset=1&limit=10&age=18&gender=M&country=TW&platform=android", nil)

		recorder := httptest.NewRecorder()
		router.ServeHTTP(recorder, req)

		assert.Equal(t, http.StatusOK, recorder.Code)
	})
}
