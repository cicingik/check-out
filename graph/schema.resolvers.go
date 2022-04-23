package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/cicingik/check-out/graph/generated"
	"github.com/cicingik/check-out/graph/model"
)

func (r *mutationResolver) CreateProduct(ctx context.Context, input *model.NewProduct) (*model.ResponseCreated, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateProduct(ctx context.Context, input *model.NewProduct) (*model.ResponseUpdated, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreatePromo(ctx context.Context, input *model.NewPromo) (*model.ResponseCreated, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdatePromo(ctx context.Context, input *model.NewPromo) (*model.ResponseUpdated, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) AddCart(ctx context.Context, input *model.NewCart) (*model.Cart, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Product(ctx context.Context, sku string) (*model.Product, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Products(ctx context.Context) ([]*model.Product, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Promo(ctx context.Context, sku string) (*model.Promo, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Promolist(ctx context.Context) ([]*model.Promo, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Cartlist(ctx context.Context) ([]*model.Cart, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Checkout(ctx context.Context, input *model.Carts) (*model.ResponseCheckout, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
