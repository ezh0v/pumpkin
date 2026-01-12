package server

import (
	"net/http"

	"github.com/ezh0v/pumpkin/internal/app"
	"github.com/ezh0v/pumpkin/internal/server/api"
	"github.com/ezh0v/pumpkin/internal/server/web"
)

func New(app *app.Context, opts ...Option) (*http.Server, error) {
	options := &options{}

	for _, opt := range opts {
		opt(options)
	}

	optionsWithDefaults(options)

	handler := http.NewServeMux()
	handler.Handle("/", web.Route(app))
	handler.Handle("/api/", api.Route(app))
	handler.Handle("/static/", http.FileServer(http.FS(app.StaticFS)))

	return &http.Server{
		Addr:    options.address,
		Handler: handler,
	}, nil
}
