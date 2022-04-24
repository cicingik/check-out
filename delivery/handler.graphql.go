package delivery

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/cicingik/check-out/graph"
	"github.com/cicingik/check-out/graph/generated"
	"github.com/cicingik/check-out/repository/cart"
	"github.com/cicingik/check-out/repository/product"
	"github.com/cicingik/check-out/repository/promo"
	"github.com/go-chi/chi"
)

type (
	GraphQl struct {
		srv *handler.Server
	}
)

func NewGraphQl(
	cart *cart.CartRepository,
	promo *promo.PromoRepository,
	product *product.ProductRepository,
) (*GraphQl, error) {

	resolver := graph.NewResolver(cart, promo, product)

	srv := handler.NewDefaultServer(
		generated.NewExecutableSchema(
			generated.Config{
				Resolvers: resolver,
			},
		),
	)

	checkout := &GraphQl{
		srv: srv,
	}

	return checkout, nil
}

func (p *GraphQl) Routes(c *chi.Mux) {
	c.Route("/v1/graphql", func(r chi.Router) {
		r.Group(func(r1 chi.Router) {
			r.Get("/", playground.Handler("GraphQL playground", "/v1/graphql/query"))
			r.Handle("/query", p.srv)
		})
	})

}
