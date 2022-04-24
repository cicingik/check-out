package product

import (
	"context"
	"errors"
	"github.com/cicingik/check-out/models/database"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func (r *ProductRepository) Create(ctx context.Context, newProduct database.Product) (result database.Product, err error) {
	product := database.InitProduct(ctx, r.DB.G)

	result, err = product.Create(newProduct)
	if err != nil {
		log.Errorf("failed create product. Details %s", err.Error())
		return result, err
	}

	return result, nil
}

func (r *ProductRepository) Update(ctx context.Context, newProduct database.Product) (result database.Product, err error) {
	product := database.InitProduct(ctx, r.DB.G)

	result, err = product.Update(r.DB.G, newProduct)
	if err != nil {
		log.Errorf("failed update product. Details %s", err.Error())
		return result, err
	}

	return result, nil
}

func (r *ProductRepository) FindAllProduct(ctx context.Context, param database.RequestPaginatorParam) (result database.ProductPagination, err error) {
	product := database.InitProduct(ctx, r.DB.G)

	result, err = product.FindAll(param)
	if err != nil {
		log.Errorf("failed get all product. Details %s", err.Error())
		return result, err
	}

	return result, nil
}

func (r *ProductRepository) FindProduct(ctx context.Context, sku string) (result database.Product, err error) {
	product := database.InitProduct(ctx, r.DB.G)

	result, err = product.FindProduct("sku", sku)
	if err != nil {
		log.Errorf("failed find product. Details %s", err.Error())
		return result, err
	}

	return result, nil
}

func (r *ProductRepository) CheckProductStock(ctx context.Context, checkOutItem []database.Cart) (productStock []database.Product, trx *gorm.DB, err error) {
	var checkOutItemCode []string
	product := database.InitProduct(ctx, r.DB.G)
	trx = r.DB.G

	coItem := make(map[string]int)
	for i := 0; i < len(checkOutItem); i++ {
		coItem[checkOutItem[i].Sku] = checkOutItem[i].Quantity
		checkOutItemCode = append(checkOutItemCode, checkOutItem[i].Sku)
	}

	productStock, err = product.FindProductStock(checkOutItemCode)
	if err != nil {
		log.Errorf("failed find product. Details %s", err.Error())
		return productStock, trx, err
	}

	// check all item is available for purchase
	for _, prdct := range productStock {
		if coItem[prdct.Sku] > prdct.Quantity {
			return productStock, trx, errors.New("item out of stock")
		}
	}

	return productStock, trx, nil
}

func (r *ProductRepository) Purchase(ctx context.Context, trx *gorm.DB, checkOutItem []database.Cart, productStock []database.Product) (result float64, err error) {
	promo := database.InitPromo(ctx, r.DB.G)
	product := database.InitProduct(ctx, r.DB.G)

	coItem := make(map[string]int)
	for i := 0; i < len(checkOutItem); i++ {
		coItem[checkOutItem[i].Sku] = checkOutItem[i].Quantity
	}

	var totalPayment float64
	for _, item := range productStock {
		promoItem, err := promo.FindPromo("sku", item.Sku)
		if err != nil {
			return result, err
		}

		productItem, err := product.FindProduct("sku", item.Sku)
		if err != nil {
			return result, err
		}

		if len(promoItem) > 0 {
			promoType := promoItem[0].PromoType
			switch promoType {
			case database.Discount:
				qty, exist := coItem[promoItem[0].Sku]
				if exist && qty >= promoItem[0].MinimalPurchased {
					totalPayment += (float64(qty) * productItem.Price) * (1 - promoItem[0].Discount/100)
				} else {
					totalPayment += float64(qty) * productItem.Price
				}

				item.Quantity = item.Quantity - qty
				_, _ = product.Update(trx, item)

			case database.FreeItem:
				qty, exist := coItem[promoItem[0].Sku]
				promoProduct, _ := product.FindProduct(`sku`, promoItem[0].BonusProductSku)
				if exist && qty >= promoItem[0].MinimalPurchased {
					totalPayment += float64(qty)*productItem.Price - (promoProduct.Price)

					promoProduct.Quantity = promoProduct.Quantity - 1
					_, _ = product.Update(trx, promoProduct)

				} else {
					totalPayment += float64(qty) * productItem.Price
				}

				item.Quantity = item.Quantity - qty
				_, _ = product.Update(trx, item)

			default:
				continue
			}
		} else {
			qty, _ := coItem[item.Sku]
			totalPayment += float64(qty) * productItem.Price

			item.Quantity = item.Quantity - qty
			_, _ = product.Update(trx, item)
		}
	}

	return totalPayment, nil
}
