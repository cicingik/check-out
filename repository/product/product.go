package product

import (
	"github.com/cicingik/check-out/config"
	"github.com/cicingik/check-out/repository/postgre"
	"go.uber.org/fx"
)

type (
	ProductRepository struct {
		DB  *postgre.DbEngine
		Cfg *config.AppConfig
	}
)

var ProductRepositoryModule = fx.Provide(NewProductRepository)

func NewProductRepository(cfg *config.AppConfig, db *postgre.DbEngine) (*ProductRepository, error) {
	return &ProductRepository{
		DB:  db,
		Cfg: cfg,
	}, nil
}
