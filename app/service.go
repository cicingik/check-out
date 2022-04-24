package app

import (
	"github.com/cicingik/check-out/delivery"
	"github.com/cicingik/check-out/repository/cart"
	"github.com/cicingik/check-out/repository/product"
	"github.com/cicingik/check-out/repository/promo"
	log "github.com/sirupsen/logrus"
)

func initService(
	httpServer *delivery.DeliveryHTTPEngine,
	cart *cart.CartRepository,
	promo *promo.PromoRepository,
	product *product.ProductRepository,
) error {

	aHandler, err := delivery.NewGraphQl(cart, promo, product)
	if err != nil {
		log.Errorf("could not initialize NewCheckout: %s", err)
		return err
	}

	httpServer.RegisterHandler(aHandler.Routes)

	return nil
}
