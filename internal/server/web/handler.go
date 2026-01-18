package web

import (
	"embed"
	"html/template"
	"net/http"

	"github.com/ezh0v/pumpkin/internal/app"
	"github.com/ezh0v/pumpkin/internal/pkg/html"
	"github.com/ezh0v/pumpkin/internal/server/pkg/middlewares"
)

var (
	//go:embed views
	viewsFS embed.FS

	renderer *html.Renderer
)

func init() {
	renderer = html.NewRenderer(
		template.Must(template.ParseFS(viewsFS, "views/**/*.html")),
	)
}

func Handler(app *app.Context) http.Handler {
	handler := http.NewServeMux()
	handler.HandleFunc("/", home(app))
	return middlewares.With(handler, middlewares.CSRF())
}
