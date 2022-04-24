package promo

import (
	"github.com/cicingik/check-out/config"
	"github.com/cicingik/check-out/repository/postgre"
	"go.uber.org/fx"
)

type (
	PromoRepository struct {
		DB  *postgre.DbEngine
		Cfg *config.AppConfig
	}
)

var PromoRepositoryModule = fx.Provide(NewPromoRepository)

func NewPromoRepository(cfg *config.AppConfig, db *postgre.DbEngine) (*PromoRepository, error) {
	return &PromoRepository{
		DB:  db,
		Cfg: cfg,
	}, nil
}
