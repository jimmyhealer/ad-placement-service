package v1

import (
	"github.com/gin-gonic/gin"
)

func CreateAd(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "createAd",
	})
}

func GetAds(c *gin.Context) {

	c.JSON(200, gin.H{
		"message": "getAds",
	})
}
