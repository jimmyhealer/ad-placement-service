package v1

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jimmyhealer/ad-placement-service/models"
	"github.com/jimmyhealer/ad-placement-service/repositories"
	"github.com/lib/pq"
)

type AdController struct {
	repo repositories.AdRepository
}

func NewAdController(repo repositories.AdRepository) *AdController {
	return &AdController{
		repo: repo,
	}
}

func (ctrl *AdController) CreateAd(c *gin.Context) {
	var req models.Advertisement
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Data validation and set default values
	if req.StartAt.After(req.EndAt) {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "startAt must be less than endAt",
		})
		return
	}
	if req.Conditions != nil {
		for i := range req.Conditions {
			if req.Conditions[i].AgeStart > req.Conditions[i].AgeEnd {
				c.JSON(http.StatusBadRequest, gin.H{
					"error": "ageStart must be less than or equal to ageEnd",
				})
				return
			}

			if req.Conditions[i].AgeStart == 0 {
				req.Conditions[i].AgeStart = 1
			}

			if req.Conditions[i].AgeEnd == 0 {
				req.Conditions[i].AgeEnd = 100
			}

			if req.Conditions[i].Gender == nil {
				req.Conditions[i].Gender = pq.StringArray{
					"M", "F",
				}
			}

			if req.Conditions[i].Country == nil {
				req.Conditions[i].Country = pq.StringArray{
					"TW", "JP",
				}
			}

			if req.Conditions[i].Platform == nil {
				req.Conditions[i].Platform = pq.StringArray{
					"android", "ios", "web",
				}
			}
		}
	}

	if err := ctrl.repo.CreateAd(&req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Ad created successfully",
	})
}

func (ctrl *AdController) GetAds(c *gin.Context) {
	var getAdsRequestErrorMessages = map[string]string{
		"Offset.min":     "offset must be greater than or equal to 1",
		"Offset.max":     "offset must be less than or equal to 100",
		"Limit.min":      "limit must be greater than or equal to 1",
		"Limit.max":      "limit must be less than or equal to 100",
		"Age.min":        "age must be greater than or equal to 1",
		"Age.max":        "age must be less than or equal to 100",
		"Gender.oneof":   "gender must be one of M, F",
		"Country.oneof":  "country must be one of TW, JP",
		"Platform.oneof": "platform must be one of android, ios, web",
	}

	var req models.GetAdsRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": errorMessageHandler(err, getAdsRequestErrorMessages),
		})
		return
	}

	if ads, err := ctrl.repo.FindActiveAds(
		time.Now(),
		req.Offset,
		req.Limit,
		req.Age,
		req.Gender,
		req.Country,
		req.Platform); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{
			"data": ads,
		})
	}
}
