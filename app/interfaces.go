package app

import "github.com/cicingik/check-out/repository/postgre"

type (
	IWebApplication interface {
		//ManageDb
		ManageHttpServer
		Start() error
		Stop() error
	}

	ManageDb interface {
		SetDb(*postgre.DbEngine)
		GetDb() *postgre.DbEngine
	}

	DeliveryHTTPEngine interface {
		Serve() error
	}

	ManageHttpServer interface {
		SetHttpServer(*DeliveryHTTPEngine)
		GetHttpServer() *DeliveryHTTPEngine
	}
)
