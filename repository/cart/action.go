package cart

import (
	"context"
	"github.com/cicingik/check-out/models/database"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func (r *CartRepository) Create(ctx context.Context, trx *gorm.DB, newCart database.Cart) (result database.Cart, err error) {
	cart := database.InitCart(ctx, r.DB.G)

	cartItem, err := cart.FindCart("sku", newCart.Sku)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			result, err = cart.Create(newCart)
			if err != nil {
				log.Errorf("failed create cart. Details %s", err.Error())
				return result, err
			}
		} else {
			return result, err
		}
	}

	cartItem.Quantity += newCart.Quantity

	result, err = cart.Update(trx, cartItem)
	if err != nil {
		log.Errorf("failed create cart. Details %s", err.Error())
		return result, err
	}

	trx.Commit()

	return result, nil
}

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

func (r *CartRepository) FindCartById(ctx context.Context, id int) (result database.Cart, err error) {
	cart := database.InitCart(ctx, r.DB.G)

	result, err = cart.FindCartById(id)
	if err != nil {
		log.Errorf("failed find cart. Details %s", err.Error())
		return result, err
	}

	return result, nil
}
