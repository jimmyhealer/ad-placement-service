//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	v1 "github.com/jimmyhealer/ad-placement-service/api/v1"
	"github.com/jimmyhealer/ad-placement-service/db"
	"github.com/jimmyhealer/ad-placement-service/repositories"
)

func InitializeAdController() (*v1.AdController, error) {
	wire.Build(db.ProvideDatabase, repositories.NewAdRepository, v1.NewAdController)
	return nil, nil
}
