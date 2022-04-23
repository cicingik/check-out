package http

import (
	"fmt"
	"net/http"
	"os"

	"github.com/cicingik/check-out/config"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type (
	DeliveryHTTPEngine struct {
		mux      *chi.Mux
		HttpPort int
		HttpHost string
	}
)

func NewHTTPServer(cfg *config.AppConfig) *DeliveryHTTPEngine {
	return &DeliveryHTTPEngine{
		mux:      chi.NewMux(),
		HttpHost: cfg.HttpHost,
		HttpPort: cfg.HttpPort,
	}
}

func (h *DeliveryHTTPEngine) InitMiddleware(appMiddleware ...func(http.Handler) http.Handler) {
	c := h.mux

	//rateLimit := customMiddleware.RateLimit(1*time.Second, 3)

	c.Use(middleware.RequestID)
	c.Use(middleware.AllowContentType("application/json", "multipart/form-data"))
	c.Use(middleware.RealIP)
	c.Use(middleware.Logger)
	c.Use(middleware.Recoverer)

	// App-level middleware
	for _, m := range appMiddleware {
		c.Use(m)
	}

	c.Handle("/metrics", promhttp.HandlerFor(
		prometheus.DefaultGatherer,
		promhttp.HandlerOpts{
			// Opt into OpenMetrics to support exemplars.
			EnableOpenMetrics: false,
		},
	))
}

func (h *DeliveryHTTPEngine) RegisterHandler(registerFn func(*chi.Mux)) {
	registerFn(h.mux)
}

func (h *DeliveryHTTPEngine) Serve() error {
	err := h.initRoute()
	if err != nil {
		fmt.Printf("found error: %s", err)
		os.Exit(1)
	}
	binding := fmt.Sprintf("%s:%d", h.HttpHost, h.HttpPort)
	fmt.Printf("Running HTTP Server in %s", binding)
	return http.ListenAndServe(binding, h.mux)
}

func (h *DeliveryHTTPEngine) initRoute() error {
	h.mux.Get("/", indexHandler)
	h.mux.Get("/v", versionHandler)
	h.mux.Get("/healthzx", HealthZX)
	return nil
}
