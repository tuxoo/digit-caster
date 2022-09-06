package server

import (
	"context"
	"digit-caster/internal/config"
	"fmt"
	"net/http"
)

type HTTPServer struct {
	httpServer *http.Server
}

func NewHTTPServer(cfg *config.Config, handler http.Handler) *HTTPServer {
	return &HTTPServer{
		httpServer: &http.Server{
			Addr:           fmt.Sprintf(":%s", cfg.HTTPConfig.Port),
			Handler:        handler,
			MaxHeaderBytes: cfg.HTTPConfig.MaxHeaderMegabytes << 28,
		},
	}
}

func (s *HTTPServer) Run() error {
	return s.httpServer.ListenAndServe()
}

func (s *HTTPServer) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
