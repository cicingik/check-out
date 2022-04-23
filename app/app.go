package app

import (
	"context"
	"fmt"

	"github.com/cicingik/check-out/config"
	"github.com/cicingik/check-out/http"
	"github.com/cicingik/check-out/repository/postgre"
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
			err := db.G.Close()
			if err != nil {
				err = fmt.Errorf("error closing database connection: %s", err)
				return err
			}
			return nil
		},
	})

	return db
}

func NewHttpServer(
	lifecycle fx.Lifecycle,
	cfg *config.AppConfig,
	db *postgre.DbEngine,
) *http.DeliveryHTTPEngine {
	httpServer := http.NewHTTPServer(cfg)

	httpServer.InitMiddleware(
		ContextualizeDb(db),
	)

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
