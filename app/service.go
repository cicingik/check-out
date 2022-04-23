package app

import (
	"github.com/cicingik/check-out/config"
	"github.com/cicingik/check-out/http"
	"github.com/cicingik/check-out/repository/postgre"
)

func initService(
	cfg *config.AppConfig,
	db *postgre.DbEngine,
	httpServer *http.DeliveryHTTPEngine,
) error {

	//aHandler, err := rest.NewApiHandler(cfg, db, mCon, pConn)
	//if err != nil {
	//	log.Errorf("could not initialize QueryHttpHandler: %s", err)
	//	return err
	//}
	//
	//httpServer.RegisterHandler(aHandler.Routes)

	return nil
}
