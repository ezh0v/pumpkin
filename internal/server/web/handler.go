package web

import (
	"embed"
	"html/template"
	"net/http"

	"github.com/ezh0v/pumpkin/internal/app"
)

var (
	//go:embed views
	viewsFS embed.FS

	templates *template.Template
)

func init() {
	templates = template.Must(template.ParseFS(viewsFS, "views/**/*.html"))
}

func Handler(app *app.Context) http.Handler {
	handler := http.NewServeMux()
	handler.HandleFunc("GET /", home(app))
	return handler
}
