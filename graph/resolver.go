package graph

import (
	"github.com/cicingik/check-out/repository/cart"
	"github.com/cicingik/check-out/repository/product"
	"github.com/cicingik/check-out/repository/promo"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	cart    *cart.CartRepository
	promo   *promo.PromoRepository
	product *product.ProductRepository
}

func NewResolver(cart *cart.CartRepository, promo *promo.PromoRepository, product *product.ProductRepository) *Resolver {
	return &Resolver{
		cart:    cart,
		promo:   promo,
		product: product,
	}
}
