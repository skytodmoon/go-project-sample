// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package controllers

import (
	"github.com/google/wire"
	"github.com/sdgmf/go-project-sample/internal/app/reviews/repositorys"
	"github.com/sdgmf/go-project-sample/internal/app/reviews/services"
	"github.com/sdgmf/go-project-sample/internal/pkg/config"
	"github.com/sdgmf/go-project-sample/internal/pkg/database"
	"github.com/sdgmf/go-project-sample/internal/pkg/log"
)

// Injectors from wire.go:

func CreateReviewsController(cf string, sto repositorys.ReviewsRepository) (*ReviewsController, error) {
	viper, err := config.New(cf)
	if err != nil {
		return nil, err
	}
	options, err := log.NewOptions(viper)
	if err != nil {
		return nil, err
	}
	logger, err := log.New(options)
	if err != nil {
		return nil, err
	}
	reviewsService := services.NewReviewService(logger, sto)
	reviewsController := NewReviewsController(logger, reviewsService)
	return reviewsController, nil
}

// wire.go:

var testProviderSet = wire.NewSet(log.ProviderSet, config.ProviderSet, database.ProviderSet, services.ProviderSet, ProviderSet)
