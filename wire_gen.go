// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/jimmyhealer/ad-placement-service/api/v1"
	"github.com/jimmyhealer/ad-placement-service/db"
	"github.com/jimmyhealer/ad-placement-service/repositories"
)

// Injectors from wire.go:

func InitializeAdController() (*v1.AdController, error) {
	database, err := db.ProvideDatabase()
	if err != nil {
		return nil, err
	}
	adRepository := repositories.NewAdRepository(database)
	adController := v1.NewAdController(adRepository)
	return adController, nil
}
