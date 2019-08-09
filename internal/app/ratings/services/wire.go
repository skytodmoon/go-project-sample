// +build wireinject

package services

import (
	"github.com/google/wire"
	"github.com/sdgmf/go-project-sample/internal/pkg/config"
	"github.com/sdgmf/go-project-sample/internal/pkg/database"
	"github.com/sdgmf/go-project-sample/internal/pkg/log"
	"github.com/sdgmf/go-project-sample/internal/app/ratings/repositorys"
)

var testProviderSet = wire.NewSet(
	log.ProviderSet,
	config.ProviderSet,
	database.ProviderSet,
	ProviderSet,
)

func CreateRatingsService(cf string, sto repositorys.RatingsRepository) (RatingsService, error) {
	panic(wire.Build(testProviderSet))
}
