package promo

import (
	"context"
	"github.com/cicingik/check-out/models/database"
	log "github.com/sirupsen/logrus"
)

func (r *PromoRepository) Create(ctx context.Context, newPromo database.Promo) (result database.Promo, err error) {
	promo := database.InitPromo(ctx, r.DB.G)

	result, err = promo.Create(newPromo)
	if err != nil {
		log.Errorf("failed create promo. Details %s", err.Error())
		return result, err
	}

	return result, nil
}

func (r *PromoRepository) Update(ctx context.Context, newPromo database.Promo) (result database.Promo, err error) {
	promo := database.InitPromo(ctx, r.DB.G)

	result, err = promo.Update(r.DB.G, newPromo)
	if err != nil {
		log.Errorf("failed update promo. Details %s", err.Error())
		return result, err
	}

	return result, nil
}

func (r *PromoRepository) FindAllPromoWithPagination(ctx context.Context, param database.RequestPaginatorParam) (result database.PromoPagination, err error) {
	promo := database.InitPromo(ctx, r.DB.G)

	result, err = promo.FindAllWithPagination(param)
	if err != nil {
		log.Errorf("failed get all promo. Details %s", err.Error())
		return result, err
	}

	return result, nil
}

func (r *PromoRepository) FindPromo(ctx context.Context, sku string) (result database.Promo, err error) {
	promo := database.InitPromo(ctx, r.DB.G)

	resultPromo, err := promo.FindPromo("sku", sku)
	if err != nil {
		log.Errorf("failed find promo. Details %s", err.Error())
		return result, err
	}

	if len(resultPromo) > 0 {
		return resultPromo[0], nil
	}

	return result, nil
}

func (r *PromoRepository) FindAllActivePromo(ctx context.Context) (result []database.Promo, err error) {
	promo := database.InitPromo(ctx, r.DB.G)

	result, err = promo.FindAll()
	if err != nil {
		log.Errorf("failed get all promo. Details %s", err.Error())
		return result, err
	}

	return result, nil
}

func (r *PromoRepository) ApplyPromo(ctx context.Context, checkOutItem []database.Cart) (promoDeal []database.Promo, err error) {
	promo := database.InitPromo(ctx, r.DB.G)

	coItem := make(map[string]int)
	for i := 0; i < len(checkOutItem); i++ {
		coItem[checkOutItem[i].Sku] = checkOutItem[i].Quantity
	}

	allPromo, err := promo.FindAll()
	if err != nil {
		log.Errorf("failed get all promo. Details %s", err.Error())
		return promoDeal, err
	}

	for _, promoItem := range allPromo {
		switch promoItem.PromoType {
		case database.Discount:
			qty, exist := coItem[promoItem.Sku]
			if exist && qty >= promoItem.MinimalPurchased {
				promoDeal = append(promoDeal, promoItem)
			}

		case database.FreeItem:
			qty, exist := coItem[promoItem.Sku]
			if exist && qty >= promoItem.MinimalPurchased {
				promoDeal = append(promoDeal, promoItem)
			}

		default:
			continue
		}
	}

	return promoDeal, nil
}
