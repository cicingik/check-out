package app

import (
	"github.com/cicingik/check-out/config"
	"github.com/cicingik/check-out/delivery"
	"github.com/cicingik/check-out/http"
	"github.com/cicingik/check-out/repository/postgre"
	log "github.com/sirupsen/logrus"
)

func initService(
	cfg *config.AppConfig,
	db *postgre.DbEngine,
	httpServer *http.DeliveryHTTPEngine,
) error {

	aHandler, err := delivery.NewCheckout(cfg)
	if err != nil {
		log.Errorf("could not initialize NewCheckout: %s", err)
		return err
	}

	httpServer.RegisterHandler(aHandler.Routes)

	return nil
}
