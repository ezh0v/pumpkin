package app

import (
	"embed"
	"fmt"
	"html/template"
	"io/fs"

	"github.com/ezh0v/pumpkin/internal/pkg/html"
)

var (
	//go:embed views
	viewsFS embed.FS

	//go:embed static
	staticFS embed.FS
)

type Context struct {
	AppVersion string
	StaticFS   fs.FS
	html.Renderer
}

func New(config *Config) (*Context, error) {
	views, err := template.ParseFS(viewsFS, "views/**/*.html")
	if err != nil {
		return nil, fmt.Errorf("templates parse failed %v", err)
	}

	return &Context{
		AppVersion: config.AppVersion,
		StaticFS:   staticFS,
		Renderer: html.Renderer{
			Views: views,
		},
	}, nil
}
