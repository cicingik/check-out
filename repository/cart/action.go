package cart

import (
	"context"
	"github.com/cicingik/check-out/models/database"
	log "github.com/sirupsen/logrus"
)

func (r *CartRepository) FindAllCart(ctx context.Context, param database.RequestPaginatorParam) (result database.CartPagination, err error) {
	cart := database.InitCart(ctx, r.DB.G)

	result, err = cart.FindAll(param)
	if err != nil {
		log.Errorf("failed get all cart. Details %s", err.Error())
		return result, err
	}

	return result, nil
}

func (r *CartRepository) FindCart(ctx context.Context, id string) (result database.Cart, err error) {
	cart := database.InitCart(ctx, r.DB.G)

	result, err = cart.FindCart("id", id)
	if err != nil {
		log.Errorf("failed find cart. Details %s", err.Error())
		return result, err
	}

	return result, nil
}
