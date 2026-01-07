package server

import (
	"fmt"
	"html/template"
	"io/fs"
	"net/http"

	"github.com/ezh0v/pumpkin/internal/app"
	"github.com/ezh0v/pumpkin/internal/app/api"
	"github.com/ezh0v/pumpkin/internal/app/web"
	"github.com/ezh0v/pumpkin/internal/pkg/html"
)

func New(viewsFS, staticFS fs.FS, opts ...Option) (*http.Server, error) {
	options := &options{}

	for _, opt := range opts {
		opt(options)
	}

	optionsWithDefaults(options)

	views, err := template.ParseFS(viewsFS, options.viewsPatterns)
	if err != nil {
		return nil, fmt.Errorf("templates parse failed %v", err)
	}

	c := &app.Context{
		Renderer: html.Renderer{
			Views: views,
		},
	}

	handler := http.NewServeMux()
	handler.Handle("/", web.Route(c))
	handler.Handle("/api/", api.Route(c))
	handler.Handle("/static/", http.FileServer(http.FS(staticFS)))

	return &http.Server{
		Addr:    options.address,
		Handler: handler,
	}, nil
}
