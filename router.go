package main

import (
	v1 "github.com/jimmyhealer/ad-placement-service/api/v1"

	"github.com/gin-gonic/gin"
)

func setupRouter(crtl *v1.AdController) *gin.Engine {
	r := gin.Default()

	r.GET("/api/v1/ads", crtl.GetAds)
	r.POST("/api/v1/ads", crtl.CreateAd)

	return r
}
