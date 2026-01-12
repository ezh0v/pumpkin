package app

import (
	"embed"
	"fmt"
	"html/template"
	"io/fs"

	"github.com/ezh0v/pumpkin/internal/pkg/html"
	"github.com/ezh0v/pumpkin/internal/pkg/postgres"
)

var (
	//go:embed views
	viewsFS embed.FS

	//go:embed static
	staticFS embed.FS
)

type Context struct {
	Version  string
	StaticFS fs.FS
	database Database
	html.Renderer
}

func New(config *Config) (*Context, error) {
	if err := config.validate(); err != nil {
		return nil, err
	}

	views, err := template.ParseFS(viewsFS, "views/**/*.html")
	if err != nil {
		return nil, fmt.Errorf("templates parse failed %v", err)
	}

	database, err := postgres.New(config.DatabaseConnect)
	if err != nil {
		return nil, fmt.Errorf("postgres initialization failed %v", err)
	}

	return &Context{
		Version:  config.Version,
		StaticFS: staticFS,
		database: database,
		Renderer: html.Renderer{
			Views: views,
		},
	}, nil
}
