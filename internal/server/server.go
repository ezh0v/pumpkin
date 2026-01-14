package server

import (
	"context"
	"embed"
	"errors"
	"io"
	"net/http"
	"time"

	"github.com/ezh0v/pumpkin/internal/pkg/postgres"
	"github.com/ezh0v/pumpkin/internal/server/admin"
	"github.com/ezh0v/pumpkin/internal/server/api"
	"github.com/ezh0v/pumpkin/internal/server/web"
)

//go:embed static
var staticFS embed.FS

type Server struct {
	shutdownTimeout time.Duration
	httpServer      *http.Server
	resources       []io.Closer
}

func New(opts ...Option) (*Server, error) {
	options := &options{}

	for _, opt := range opts {
		opt(options)
	}

	optionsWithDefaults(options)

	database, err := postgres.New("")
	if err != nil {
		return nil, err
	}

	handler := http.NewServeMux()
	handler.Handle("/", web.Handler())
	handler.Handle("/api/", api.Handler())
	handler.Handle("/admin/", admin.Handler())
	handler.Handle("/static/", http.FileServer(http.FS(staticFS)))

	return &Server{
		shutdownTimeout: options.shutdownTimeout,
		httpServer: &http.Server{
			Addr:    options.address,
			Handler: handler,
		},
		resources: []io.Closer{
			database,
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

	shutdownErr := s.httpServer.Shutdown(ctx)

	for _, resource := range s.resources {
		if err := resource.Close(); err != nil {
			shutdownErr = errors.Join(shutdownErr, err)
		}
	}

	return shutdownErr
}
