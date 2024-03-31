package main

import (
	v1 "github.com/jimmyhealer/ad-placement-service/api/v1"

	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/api/v1/ad", v1.GetAds)
	r.POST("/api/v1/ad", v1.CreateAd)

	return r
}
