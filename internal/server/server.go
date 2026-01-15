package server

import (
	"context"
	"embed"
	"errors"
	"net/http"
	"time"

	"github.com/ezh0v/pumpkin/internal/app"
	"github.com/ezh0v/pumpkin/internal/server/admin"
	"github.com/ezh0v/pumpkin/internal/server/api"
	"github.com/ezh0v/pumpkin/internal/server/web"
)

//go:embed static
var staticFS embed.FS

type Server struct {
	shutdownTimeout time.Duration
	httpServer      *http.Server
}

func New(ctx *app.Context, opts ...Option) (*Server, error) {
	options := &options{}

	for _, opt := range opts {
		opt(options)
	}

	optionsWithDefaults(options)

	handler := http.NewServeMux()
	handler.Handle("/", web.Handler(ctx))
	handler.Handle("/api/", api.Handler(ctx))
	handler.Handle("/admin/", admin.Handler(ctx))
	handler.Handle("/static/", http.FileServer(http.FS(staticFS)))

	return &Server{
		shutdownTimeout: options.shutdownTimeout,
		httpServer: &http.Server{
			Addr:    options.address,
			Handler: handler,
		},
	}, nil
}

func (s *Server) ListenAndServe() error {
	err := s.httpServer.ListenAndServe()
	if !errors.Is(err, http.ErrServerClosed) {
		return err
	}

	return nil
}

func (s *Server) Shutdown() error {
	ctx, cancel := context.WithTimeout(context.Background(), s.shutdownTimeout)
	defer cancel()

	return s.httpServer.Shutdown(ctx)
}
