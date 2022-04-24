package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"github.com/cicingik/check-out/graph/generated"
	"github.com/cicingik/check-out/graph/model"
	"github.com/cicingik/check-out/models/database"
)

func (r *mutationResolver) CreateProduct(ctx context.Context, input *model.NewProduct) (*model.Product, error) {
	result, err := r.product.Create(ctx, database.Product{
		Sku:      input.Sku,
		Name:     input.Name,
		Price:    input.Price,
		Quantity: input.Quantity,
	})

	if err != nil {
		return nil, err
	}

	resp := model.Product{
		ID:       result.ID,
		Sku:      result.Sku,
		Name:     result.Name,
		Price:    result.Price,
		Quantity: result.Quantity,
	}

	return &resp, nil
}

func (r *mutationResolver) CreatePromo(ctx context.Context, input *model.NewPromo) (*model.Promo, error) {
	result, err := r.promo.Create(ctx, database.Promo{
		Sku:              input.Sku,
		PromoType:        input.PromoType,
		MinimalPurchased: input.MinimalPurchased,
		BonusProductSku:  input.BonusProductSku,
		Discount:         input.Discount,
		IsActive:         true,
	})

	if err != nil {
		return nil, err
	}

	resp := model.Promo{
		ID:               result.ID,
		Sku:              result.Sku,
		PromoType:        result.PromoType,
		MinimalPurchased: result.MinimalPurchased,
		BonusProductSku:  result.BonusProductSku,
		Discount:         result.Discount,
		IsActive:         result.IsActive,
	}

	return &resp, nil
}

func (r *mutationResolver) UpdateProduct(ctx context.Context, input *model.NewProduct) (*model.Product, error) {
	result, err := r.product.Update(ctx, database.Product{
		Sku:      input.Sku,
		Name:     input.Name,
		Price:    input.Price,
		Quantity: input.Quantity,
	})

	if err != nil {
		return nil, err
	}

	resp := model.Product{
		ID:       result.ID,
		Sku:      result.Sku,
		Name:     result.Name,
		Price:    result.Price,
		Quantity: result.Quantity,
	}

	return &resp, nil
}

func (r *mutationResolver) UpdatePromo(ctx context.Context, input *model.NewPromo) (*model.Promo, error) {
	result, err := r.promo.Update(ctx, database.Promo{
		Sku:              input.Sku,
		PromoType:        input.PromoType,
		MinimalPurchased: input.MinimalPurchased,
		BonusProductSku:  input.BonusProductSku,
		Discount:         input.Discount,
		IsActive:         true,
	})

	if err != nil {
		return nil, err
	}

	resp := model.Promo{
		ID:               result.ID,
		Sku:              result.Sku,
		PromoType:        result.PromoType,
		MinimalPurchased: result.MinimalPurchased,
		BonusProductSku:  result.BonusProductSku,
		Discount:         result.Discount,
		IsActive:         result.IsActive,
	}

	return &resp, nil
}

func (r *queryResolver) Product(ctx context.Context, sku string) (*model.Product, error) {
	result, err := r.product.FindProduct(ctx, sku)
	if err != nil {
		return nil, err
	}

	resp := model.Product{
		ID:       result.ID,
		Sku:      result.Sku,
		Name:     result.Name,
		Price:    result.Price,
		Quantity: result.Quantity,
	}

	return &resp, nil
}

func (r *queryResolver) AllProducts(ctx context.Context, param model.PagingQuery) (*model.Products, error) {
	var (
		resp model.Products
	)

	result, err := r.product.FindAllProduct(ctx, database.RequestPaginatorParam{
		Page:  param.Page,
		Limit: param.Limit,
	})
	if err != nil {
		return nil, err
	}

	for _, prd := range result.Product {
		resp.Edges = append(resp.Edges, &model.Product{
			ID:       prd.ID,
			Sku:      prd.Sku,
			Name:     prd.Name,
			Price:    prd.Price,
			Quantity: prd.Quantity,
		})
	}

	resp.PageInfo = &model.PageInfo{
		After:     result.Pg.After,
		Before:    result.Pg.Before,
		TotalPage: result.Pg.TotalPage,
		Page:      result.Pg.Page,
	}

	return &resp, nil
}

func (r *queryResolver) Promo(ctx context.Context, sku string) (*model.Promo, error) {
	result, err := r.promo.FindPromo(ctx, sku)
	if err != nil {
		return nil, err
	}

	resp := model.Promo{
		ID:               result.ID,
		Sku:              result.Sku,
		PromoType:        result.PromoType,
		MinimalPurchased: result.MinimalPurchased,
		BonusProductSku:  result.BonusProductSku,
		Discount:         result.Discount,
		IsActive:         result.IsActive,
	}

	return &resp, nil
}

func (r *queryResolver) AllActivePromo(ctx context.Context, param model.PagingQuery) (*model.Promos, error) {
	var (
		resp model.Promos
	)

	result, err := r.promo.FindAllPromoWithPagination(ctx, database.RequestPaginatorParam{
		Page:  param.Page,
		Limit: param.Limit,
	})
	if err != nil {
		return nil, err
	}

	for _, promo := range result.Promo {
		resp.Edges = append(resp.Edges, &model.Promo{
			ID:               promo.ID,
			Sku:              promo.Sku,
			PromoType:        promo.PromoType,
			MinimalPurchased: promo.MinimalPurchased,
			BonusProductSku:  promo.BonusProductSku,
			Discount:         promo.Discount,
			IsActive:         promo.IsActive,
		})
	}

	resp.PageInfo = &model.PageInfo{
		After:     result.Pg.After,
		Before:    result.Pg.Before,
		TotalPage: result.Pg.TotalPage,
		Page:      result.Pg.Page,
	}

	return &resp, nil
}

func (r *queryResolver) CartList(ctx context.Context, param model.PagingQuery) (*model.Carts, error) {
	var (
		resp model.Carts
	)

	result, err := r.cart.FindAllCart(ctx, database.RequestPaginatorParam{
		Page:  param.Page,
		Limit: param.Limit,
	})
	if err != nil {
		return nil, err
	}

	for _, crt := range result.Cart {
		resp.Edges = append(resp.Edges, &model.Cart{
			ID:       crt.ID,
			Sku:      crt.Sku,
			Quantity: crt.Quantity,
		})
	}

	resp.PageInfo = &model.PageInfo{
		After:     result.Pg.After,
		Before:    result.Pg.Before,
		TotalPage: result.Pg.TotalPage,
		Page:      result.Pg.Page,
	}

	return &resp, nil
}

func (r *queryResolver) Checkout(ctx context.Context, input *model.CheckOutItem) (*model.ResponseCheckout, error) {
	var (
		checkOutItem []database.Cart
	)

	if input == nil || input.Contents == nil {
		return &model.ResponseCheckout{Total: 0}, nil
	}

	// transform to list of cart
	for _, item := range input.Contents {
		checkOutItem = append(checkOutItem, database.Cart{
			Sku:      item.Sku,
			Quantity: item.Quantity,
		})

	}

	productStock, trx, err := r.product.CheckProductStock(ctx, checkOutItem)
	if err != nil {
		return nil, err
	}

	// check for available purchase and calculate promo
	totalPayment, err := r.product.Purchase(ctx, trx, checkOutItem, productStock)
	if err != nil {
		return nil, err
	}

	return &model.ResponseCheckout{Total: totalPayment}, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
