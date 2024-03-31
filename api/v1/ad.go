package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jimmyhealer/ad-placement-service/models"
)

func CreateAd(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "createAd",
	})
}

func GetAds(c *gin.Context) {
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

	type getAdsRequest struct {
		Offset   int                   `form:"offset,default=0" binding:"omitempty,min=1,max=100"`
		Limit    int                   `form:"limit,default=5" binding:"omitempty,min=1,max=100"`
		Age      int                   `form:"age" binding:"omitempty,min=1,max=100"`
		Gender   []models.GenderType   `form:"gender" binding:"omitempty,dive,oneof=M F"`
		Country  []models.CountryCode  `form:"country" binding:"omitempty,dive,oneof=TW JP"`
		Platform []models.PlatformType `form:"platform" binding:"omitempty,dive,oneof=android ios web"`
	}

	var req getAdsRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": errorMessageHandler(err, getAdsRequestErrorMessages),
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "getAds",
		"req":     req,
	})
}
