package cart

import (
	"github.com/cicingik/check-out/config"
	"github.com/cicingik/check-out/repository/postgre"
	"go.uber.org/fx"
)

type (
	CartRepository struct {
		DB  *postgre.DbEngine
		Cfg *config.AppConfig
	}
)

var CartRepositoryModule = fx.Provide(NewCartRepository)

func NewCartRepository(cfg *config.AppConfig, db *postgre.DbEngine) (*CartRepository, error) {
	return &CartRepository{
		DB:  db,
		Cfg: cfg,
	}, nil
}
