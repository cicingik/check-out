package app

import (
	"context"
	"github.com/cicingik/check-out/config"
	"github.com/cicingik/check-out/delivery"
	"github.com/cicingik/check-out/repository/cart"
	"github.com/cicingik/check-out/repository/postgre"
	"github.com/cicingik/check-out/repository/product"
	"github.com/cicingik/check-out/repository/promo"
	"go.uber.org/fx"
)

type WebApplication struct {
	*fx.App
	cfg *config.AppConfig
	db  *postgre.DbEngine
}

func NewWebApplication(cfg *config.AppConfig) (*WebApplication, error) {
	app := &WebApplication{}

	container := fx.New(
		coreServiceProviders(cfg),
		cart.CartRepositoryModule,
		promo.PromoRepositoryModule,
		product.ProductRepositoryModule,
		fx.Invoke(initService),
	)

	app.App = container
	return app, nil
}

func coreServiceProviders(cfg *config.AppConfig) fx.Option {
	return fx.Provide(
		func() *config.AppConfig {
			return cfg
		},
		NewDatabase,
		NewHttpServer,
	)
}

func NewDatabase(lifecycle fx.Lifecycle, cfg *config.AppConfig) *postgre.DbEngine {
	db := postgre.NewDbService(*cfg)

	lifecycle.Append(fx.Hook{
		OnStop: func(ctx context.Context) error {
			return nil
		},
	})

	return db
}

func NewHttpServer(
	lifecycle fx.Lifecycle,
	cfg *config.AppConfig,
) *delivery.DeliveryHTTPEngine {
	httpServer := delivery.NewHTTPServer(cfg)

	httpServer.InitMiddleware()

	lifecycle.Append(fx.Hook{
		OnStart: func(context.Context) error {
			go httpServer.Serve()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return nil
		},
	})

	return httpServer
}

func (w *WebApplication) Start(ctx context.Context) error {
	return w.App.Start(ctx)
}

// Stop perform gracefull stop
func (w *WebApplication) Stop(ctx context.Context) error {
	return w.App.Stop(ctx)
}
