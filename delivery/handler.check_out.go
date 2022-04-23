package delivery

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/cicingik/check-out/config"
	"github.com/cicingik/check-out/graph"
	"github.com/cicingik/check-out/graph/generated"
	"github.com/go-chi/chi"
)

type (
	Checkout struct {
		cfg *config.AppConfig
		srv *handler.Server
	}
)

func NewCheckout(cfg *config.AppConfig) (*Checkout, error) {
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

	checkout := &Checkout{
		cfg: cfg,
		srv: srv,
	}

	return checkout, nil
}

func (p *Checkout) Routes(c *chi.Mux) {
	c.Route("/v1/graphql", func(r chi.Router) {
		r.Group(func(r1 chi.Router) {
			r.Get("/", playground.Handler("GraphQL playground", "/query"))
			r.Handle("/query", p.srv)
		})
	})

}
